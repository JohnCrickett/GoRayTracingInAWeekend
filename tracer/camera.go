package tracer

import (
	"fmt"
	"math"
	"os"
)

type Camera struct {
	imageWidth        int
	imageHeight       int
	aspectRatio       float64
	cameraCenter      Point
	pixel00Loc        Point
	pixelDeltaU       Vec
	pixelDeltaV       Vec
	samplesPerPixel   int
	pixelSamplesScale float64
	maxDepth          int
	vFoV              float64
	LookFrom          Vec
	LookAt            Vec
	VUp               Vec
}

func NewCamera(imageWidth int, aspectRatio float64, samplesPerPixel int, maxDepth int, vFoV float64, LookFrom, LookAt, VUp Vec) *Camera {
	// Calculate the image height, and ensure that it's at least 1.
	imageHeight := int(float64(imageWidth) / aspectRatio)
	if imageHeight < 1 {
		imageHeight = 1
	}
	
	cameraCenter := LookFrom

	// Camera
	focalLength := LookFrom.Minus(LookAt).Length()
	theta := DegreesToRadians(vFoV)
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h * focalLength
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))

	// Calculate the u,v,w unit basis vectors for the camera coordinate frame.
	w := UnitVector(LookFrom.Minus(LookAt))
	u := UnitVector(Cross(VUp, w))
	v := Cross(w, u)

	// Calculate the vectors across the horizontal and down the vertical viewport edges.
	viewportU := u.Scale(viewportWidth)
	viewportV := v.Scale(-viewportHeight)

	// Calculate the horizontal and vertical delta vectors from pixel to pixel.
	pixelDeltaU := viewportU.Divide(float64(imageWidth))
	pixelDeltaV := viewportV.Divide(float64(imageHeight))

	// Calculate the location of the upper left pixel.
	viewportUpperLeft := cameraCenter.Minus(w.Scale(focalLength)).Minus(viewportU.Scale(0.5)).Minus(viewportV.Scale(0.5))
	pixel100Loc := viewportUpperLeft.Plus((pixelDeltaU.Plus(pixelDeltaV)).Scale(0.5))

	return &Camera{
		imageWidth:        imageWidth,
		imageHeight:       imageHeight,
		cameraCenter:      cameraCenter,
		pixel00Loc:        pixel100Loc,
		pixelDeltaV:       pixelDeltaV,
		pixelDeltaU:       pixelDeltaU,
		samplesPerPixel:   samplesPerPixel,
		pixelSamplesScale: 1.0 / float64(samplesPerPixel),
		maxDepth:          maxDepth,
		vFoV:              vFoV,
		LookFrom:          LookFrom,
		LookAt:            LookAt,
		VUp:               VUp,
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
		fmt.Printf("\rScanlines remaining: %4d", c.imageHeight-row)
		for col := 0; col < c.imageWidth; col++ {
			pixelColour := Colour{0, 0, 0}
			for sample := 0; sample < c.samplesPerPixel; sample++ {
				r := c.getRay(col, row)
				pixelColour = pixelColour.Add(c.rayColor(&r, c.maxDepth, world))
			}
			pixelColour = pixelColour.Scale(c.pixelSamplesScale)
			pixelColour.Write(f)
		}
	}
	fmt.Println("\rDone.                           ")
}

func (c *Camera) rayColor(r *Ray, depth int, world Hittable) Colour {
	if depth < 0 {
		return Colour{0, 0, 0}
	}
	var hitRecord *HitRecord

	hit, hitRecord := world.Hit(r, NewInterval(0.001, math.Inf(1)))
	if hit {
		//	direction := hitRecord.Normal.Plus(RandomUnitVector())
		//	return c.rayColor(&Ray{hitRecord.P, direction}, depth-1, world).Scale(0.5)
		scatters, scatteredRay, atColour := hitRecord.Material.scatter(r, hitRecord)
		if scatters {
			return atColour.Multiply(c.rayColor(scatteredRay, depth-1, world))
		}
		return Colour{}
	}

	unitDirection := UnitVector(r.Direction())
	a := 0.5 * (unitDirection.Y() + 1.0)
	return Colour{1.0, 1.0, 1.0}.Scale(1.0 - a).Plus(Colour{0.5, 0.7, 1.0}.Scale(a))
}

func (c *Camera) getRay(i, j int) Ray {
	// Construct a camera ray originating from the origin and directed at randomly sampled
	// point around the pixel location i, j.
	offset := sampleSquare()
	pixelSample := c.pixel00Loc.Plus(c.pixelDeltaU.Scale(float64(i) + offset.X()).Plus(c.pixelDeltaV.Scale(float64(j) + offset.Y())))

	rayOrigin := c.cameraCenter
	rayDirection := pixelSample.Minus(rayOrigin)

	return Ray{rayOrigin, rayDirection}
}

func sampleSquare() Vec {
	// Returns the vector to a random point in the [-.5,-.5]-[+.5,+.5] unit square.
	return Vec{RandomDouble() - 0.5, RandomDouble() - 0.5, 0}
}
