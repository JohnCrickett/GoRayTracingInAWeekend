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

func ray_color(r *tracer.Ray) tracer.Colour {
	unitDirection := tracer.UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)
	return tracer.Colour{1.0, 1.0, 1.0}.Scale(1.0 - a).Plus(tracer.Colour{0.5, 0.7, 1.0}.Scale(a))
}

func main() {
	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400

	// Calculate the image height, and ensure that it's at least 1.
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := tracer.Point{0, 0, 0}

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := tracer.Vec{viewportWidth, 0, 0}
	viewportV := tracer.Vec{0, -viewportHeight, 0}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Scale(1 / float64(imageWidth))
	pixelDeltaV := viewportV.Scale(1 / float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Minus(tracer.Vec{0, 0, focalLength}).Minus(viewportU.Scale(1 / 2)).Minus(viewportV.Scale(1 / 2))
	pixel00Loc := viewportUpperLeft.Plus((pixelDeltaU.Plus(pixelDeltaV)).Scale(0.5))

	// File to render to
	f, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Rendering
	fmt.Fprintf(f, "P3\n%d %d 255\n", width, height)

	for row := 0; row < height; row++ {
		fmt.Printf("\rScanlines remaining: %d", (height - row))
		for col := 0; col < width; col++ {
			pixelCenter := pixel00Loc.Plus((pixelDeltaU.Scale(float64(col)))).Plus(pixelDeltaV.Scale(float64(row)))
			rayDirection := pixelCenter.Minus(cameraCenter)
			r := tracer.NewRay(cameraCenter, rayDirection)

			pixel_color := ray_color(r)
			pixel_color.Write(f)
		}
	}
	fmt.Println("\rDone.                           ")
}
