package main

import "fmt"

// "naked" return.
// 裸のretuenというらしい
// returnの引数を省略できる
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
