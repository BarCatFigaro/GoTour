package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var x, y int
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			x += 2
			y += 4
			a[i][j] = uint8((x^y + x^y)/2)
		}
		
	}
	return a
}

func main() {
	pic.Show(Pic)
}
