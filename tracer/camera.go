package tracer

import (
	"fmt"
	"math"
	"os"
)

type Camera struct {
	imageWidth   int
	imageHeight  int
	aspectRatio  float64
	cameraCenter Point
	pixel100Loc  Point
	pixelDeltaU  Vec
	pixelDeltaV  Vec
}

func NewCamera(imageWidth int, aspectRatio float64) *Camera {
	// Calculate the image height, and ensure that it's at least 1.
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}

	// Camera
	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCenter := Point{0, 0, 0}

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := Vec{viewportWidth, 0, 0}
	viewportV := Vec{0, -viewportHeight, 0}

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Divide(float64(imageWidth))
	pixelDeltaV := viewportV.Divide(float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Minus(Vec{0, 0, focalLength}).Minus(viewportU.Scale(0.5)).Minus(viewportV.Scale(0.5))
	pixel100Loc := viewportUpperLeft.Plus((pixelDeltaU.Plus(pixelDeltaV)).Scale(0.5))

	return &Camera{
		imageWidth:   imageWidth,
		imageHeight:  imageHeight,
		cameraCenter: cameraCenter,
		pixel100Loc:  pixel100Loc,
		pixelDeltaV:  pixelDeltaV,
		pixelDeltaU:  pixelDeltaU,
	}
}

func (c *Camera) Render(world HittableList, targetFile string) {
	// File to render to
	f, err := os.OpenFile(targetFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Rendering
	fmt.Fprintf(f, "P3\n%d %d 255\n", c.imageWidth, c.imageHeight)

	for row := 0; row < c.imageHeight; row++ {
		fmt.Printf("\rScanlines remaining: %d", (c.imageHeight - row))
		for col := 0; col < c.imageWidth; col++ {
			pixelCenter := c.pixel100Loc.Plus(c.pixelDeltaU.Scale(float64(col))).Plus(c.pixelDeltaV.Scale(float64(row)))
			rayDirection := pixelCenter.Minus(c.cameraCenter)
			r := NewRay(c.cameraCenter, rayDirection)

			pixel_color := c.rayColor(r, world)
			pixel_color.Write(f)
		}
	}
	fmt.Println("\rDone.                           ")
}

func (c *Camera) rayColor(r *Ray, world Hittable) Colour {
	var hitRecord *HitRecord

	hit, hitRecord := world.Hit(r, NewInterval(0, math.Inf(1)))
	if hit {
		c := hitRecord.Normal.Plus(Vec{1.0, 1.0, 1.0}).Scale(0.5)
		return Colour{c.X(), c.Y(), c.Z()}
	}

	unitDirection := UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)
	return Colour{1.0, 1.0, 1.0}.Scale(1.0 - a).Plus(Colour{0.5, 0.7, 1.0}.Scale(a))
}
