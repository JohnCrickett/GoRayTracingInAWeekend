package main

import "fmt"
import "os"

const (
	width      = 256
	height     = 256
	targetFile = "test.ppm"
)

func main() {
	f, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "P3\n%d %d 255\n", width, height)

	for row := 0; row < height; row++ {
		fmt.Printf("\rScanlines remaining: %d", (height - row))
		for col := 0; col < width; col++ {
			red := float64(col) / (width - 1)
			green := float64(row) / (height - 1)
			blue := 0.0

			r := int(255.999 * red)
			g := int(255.999 * green)
			b := int(255.999 * blue)

			fmt.Fprintf(f, "%d %d %d\n", r, g, b)
		}
	}
	fmt.Println("\rDone.                           ")
}
