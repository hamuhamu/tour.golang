package main

import "fmt"

func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        // deferは、呼びだされた順にstackされていく
        // last-in first-out
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
