package main

import (
    "fmt"
    "syscall/js"
    "encoding/base64"
    "maps"
)

func exportWorld(inlet [][]maps.Block) [][]map[string]interface{} {
    limitX := len(inlet)
    limitY := len(inlet[0])
    outlet := make([][]map[string]interface{}, limitX)
    for i := 0; i < limitX; i++ {
        outlet[i] = make([]map[string]interface{}, limitY)
        for j := 0; j < limitY; j++ {
            outlet[i][j] = inlet[i][j].ToMap()
        }
    }
    return outlet
}

func loadMap(this js.Value, i []js.Value) interface{} {
    encodedPngFromBrowser := i[0].String()
    pngFromBrowser, oops := base64.StdEncoding.DecodeString(encodedPngFromBrowser)
    if oops != nil {
        return oops
    }
    binaryWorld := maps.Png2Map(pngFromBrowser)
    gomapWorld := exportWorld(binaryWorld)
    return ValueOf(gomapWorld)
}

func tick(this js.Value, i[]js.Value) interface{} {
    // converting map from JS
    mJs := i[0]
    limitI := mJs.Length()
    m := make([][]maps.Block, limitI)
    for i := 0; i < limitI; i++ {
        mi := mJs.Index(i)
        limitJ := mi.Length()
        m[i] = make([]maps.Block, limitJ)

        for j := 0; j < limitJ; j++ {
            mj := mi.Index(j)

            kind := mj.Get("kind").String()

            color := make([]uint32, 3)
            colorJs := mj.Get("color")
            for k := 0; k < 3; k++ {
                color[k] = uint32(colorJs.Index(k).Int())
            }

            m[i][j] = maps.Block{
                Kind: kind,
                Color: color,
            }
        }
    }

    // converting society from JS
    sJs := i[1]
    sLen := sJs.Length()
    s := make([]maps.Institution, sLen)

    for i := 0; i < sLen; i++ {
        si := sJs.Index(i)
        s[i] = maps.Institution{
            What: si.Get("what").String(),
            Where: []int{
                si.Get("where").Get("x").Int(),
                si.Get("where").Get("y").Int(),
            },
            Memory: si.Get("memory").String(),
            Script: si.Get("script").String(),
        }
    }

    return ValueOf(maps.Tick(m, s))
}

func registerCallbacks() {
    js.Global().Set("loadMap", js.FuncOf(loadMap))
    js.Global().Set("tick", js.FuncOf(tick))
}

func main() {
    c := make(chan struct{}, 0)
    fmt.Println("WASM Go Initialized")
    registerCallbacks()
    <-c
}
