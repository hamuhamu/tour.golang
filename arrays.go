package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    // 型が違うからエラーになる
    // a[1] = 100
    fmt.Println(a[0], a[1])
    fmt.Println(a)
}
