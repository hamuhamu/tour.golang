package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	m := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		m[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			m[i][j] = uint8(i * j)
		}
	}
	return m
}

func main() {
	pic.Show(Pic)
}
