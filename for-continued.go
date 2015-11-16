package main

import "fmt"

func main() {
    sum := 1
    // 前処理をスキップできる
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}
