package main

import "fmt"

const (
	width  = 256
	height = 256
)

func main() {
	fmt.Printf("P3\n%d %d 255\n", width, height)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			red := float64(col) / (width - 1)
			green := float64(row) / (height - 1)
			blue := 0.0

			r := int(255.999 * red)
			g := int(255.999 * green)
			b := int(255.999 * blue)

			fmt.Printf("%d %d %d\n", r, g, b)
		}
	}
}
