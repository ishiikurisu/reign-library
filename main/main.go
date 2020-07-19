package main

import (
    "fmt"
    "syscall/js"
)

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
