package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	peepee := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pee := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pee[x] = uint8(x ^ y)
		}
		peepee[y] = pee
	}

	return peepee
}

func main() {
	pic.Show(Pic)
}
