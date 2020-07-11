package main

import (
    "fmt"
    "syscall/js"
    "github.com/ishiikurisu/reign-library/maps"
)

func registerCallbacks() {
    js.Global().Set("loadMap", js.FuncOf(maps.LoadMap))
    js.Global().Set("tick", js.FuncOf(maps.Tick))
}

func main() {
    c := make(chan struct{}, 0)
    fmt.Println("WASM Go Initialized")
    registerCallbacks()
    <-c
}
