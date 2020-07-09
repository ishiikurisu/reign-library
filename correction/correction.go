package main

import (
    "fmt"
    "os"
    "maps"
    "io/ioutil"
    "math"
    "image"
    "image/color"
    "image/png"
)

func main() {
    args := os.Args[1:]
    inputFile := args[0]
    outputFile := args[1]

    pngBytes, oops := ioutil.ReadFile(inputFile)
    if oops != nil {
        fmt.Printf("ERROR LOADING FILE (%s)\n", inputFile)
        return
    }
    m := maps.Png2Map(pngBytes)

    allowedColors := [][]uint8{
        []uint8{224, 236, 137},
        []uint8{31, 168, 36},
        []uint8{0, 36, 156},
        []uint8{0, 87, 63},
    }

    height := len(m)
    width := len(m[0])
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}
    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            block := m[y][x]

            originalColor := block.Color
            minDiff := 2570.0
            correctedColorIndex := -1

            for c, color := range allowedColors {
                diff := 0.0

                for d := 0; d < len(color); d++ {
                    diff += math.Abs(float64(color[d]) - float64(originalColor[d]))
                }

                if diff < minDiff {
                    correctedColorIndex = c
                    minDiff = diff
                }
            }

            c := color.RGBA{
                allowedColors[correctedColorIndex][0],
                allowedColors[correctedColorIndex][1],
                allowedColors[correctedColorIndex][2],
                0xff,
            }
            img.Set(x, y, c)
        }
    }

    fp, _ := os.Create(outputFile)
    png.Encode(fp, img)
}
