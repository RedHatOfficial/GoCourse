package main

import "fmt"
import "os"

func onFinish() {
    fmt.Println("onFinish() called")
}

func main() {
    defer onFinish()
    os.Exit(0)
}
