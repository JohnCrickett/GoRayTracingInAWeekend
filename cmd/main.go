package main

import (
	"fmt"
	"github.com/JohnCrickett/GoRayTracingInAWeekend/tracer"
)
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

			p := tracer.Colour{float64(col) / (width - 1), float64(row) / (height - 1), 0.0}
			p.Write(f)
		}
	}
	fmt.Println("\rDone.                           ")
}
