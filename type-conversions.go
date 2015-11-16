package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z int = int(f)
    fmt.Println(x, y, z)
    main2()
}

func main2() {
    i := 42
    f := float64(i)
    u := uint(f)
    fmt.Println(i, f, u)
}
