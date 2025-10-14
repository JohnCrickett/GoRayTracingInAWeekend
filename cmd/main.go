package main

import (
	"github.com/JohnCrickett/GoRayTracingInAWeekend/tracer"
)

const (
	targetFile = "test.ppm"
)

func main() {
	// Materials
	materialGround := tracer.Lambertian{tracer.Colour{0.8, 0.8, 0.0}}
	materialCenter := tracer.Lambertian{tracer.Colour{0.1, 0.2, 0.5}}
	materialLeft := tracer.Dielectric{1.5}
	materialBubble := tracer.Dielectric{1.00 / 1.50}
	materialRight := tracer.Metal{tracer.Colour{0.8, 0.6, 0.2}, 1.0}

	// World
	var world tracer.HittableList
	world.Add(tracer.NewSphere(tracer.Vec{0, 0, -1.2}, 0.5, materialCenter))
	world.Add(tracer.NewSphere(tracer.Vec{0, -100.5, -1}, 100, materialGround))
	world.Add(tracer.NewSphere(tracer.Vec{-1.0, 0.0, -1.}, 0.5, materialLeft))
	world.Add(tracer.NewSphere(tracer.Vec{-1.0, 0.0, -1.0}, 0.4, materialBubble))
	world.Add(tracer.NewSphere(tracer.Vec{1.0, 0.0, -1.0}, 0.5, materialRight))

	//R := math.Cos(math.Pi / 4)
	//
	//materialLeft := tracer.Lambertian{tracer.Colour{0, 0, 1}}
	//materialRight := tracer.Lambertian{tracer.Colour{1, 0, 0}}
	//
	//var world tracer.HittableList
	//world.Add(tracer.NewSphere(tracer.Vec{-R, 0, -1}, R, materialLeft))
	//world.Add(tracer.NewSphere(tracer.Vec{R, 0, -1}, R, materialRight))

	// Camera
	c := tracer.NewCamera(400, 16.0/9.0, 100, 50, 90.0, tracer.Vec{-2, 2, 1}, tracer.Vec{0, 0, -1}, tracer.Vec{0, 1, 0})
	c.Render(world, targetFile)
}
