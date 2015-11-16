package main

import "fmt"

func main() {
    // 引数の評価はすぐ行われるが、deferの呼び出しは周囲の関数が戻るまで実行されない
    defer fmt.Println("world")

    fmt.Println("hello")
}
