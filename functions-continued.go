package main

import "fmt"

// ２つ以上同じ方が続く場合は、引数を省略できるみたい
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}

