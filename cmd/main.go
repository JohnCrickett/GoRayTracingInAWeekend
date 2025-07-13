package main

import (
	"fmt"
	"github.com/JohnCrickett/GoRayTracingInAWeekend/tracer"
	"math"
)
import "os"

const (
	targetFile = "test.ppm"
)

func rayColor(r *tracer.Ray, world tracer.Hittable) tracer.Colour {
	var hitRecord *tracer.HitRecord

	hit, hitRecord := world.Hit(r, 0, math.Inf(1))
	if hit {
		c := hitRecord.Normal.Plus(tracer.Vec{1.0, 1.0, 1.0}).Scale(0.5)
		return tracer.Colour{c.X(), c.Y(), c.Z()}
	}

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

	// World
	var world tracer.HittableList
	world.Add(tracer.NewSphere(tracer.Vec{0, 0, -1}, 0.5))
	world.Add(tracer.NewSphere(tracer.Vec{0, -100.5, -1}, 100))

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := tracer.Point{0, 0, 0}

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := tracer.Vec{viewportWidth, 0, 0}
	viewportV := tracer.Vec{0, -viewportHeight, 0}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Divide(float64(imageWidth))
	pixelDeltaV := viewportV.Divide(float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Minus(tracer.Vec{0, 0, focalLength}).Minus(viewportU.Scale(0.5)).Minus(viewportV.Scale(0.5))
	pixel00Loc := viewportUpperLeft.Plus((pixelDeltaU.Plus(pixelDeltaV)).Scale(0.5))

	// File to render to
	f, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Rendering
	fmt.Fprintf(f, "P3\n%d %d 255\n", imageWidth, imageHeight)

	for row := 0; row < imageHeight; row++ {
		fmt.Printf("\rScanlines remaining: %d", (imageHeight - row))
		for col := 0; col < imageWidth; col++ {
			pixelCenter := pixel00Loc.Plus(pixelDeltaU.Scale(float64(col))).Plus(pixelDeltaV.Scale(float64(row)))
			rayDirection := pixelCenter.Minus(cameraCenter)
			r := tracer.NewRay(cameraCenter, rayDirection)

			pixel_color := rayColor(r, world)
			pixel_color.Write(f)
		}
	}
	fmt.Println("\rDone.                           ")
}
