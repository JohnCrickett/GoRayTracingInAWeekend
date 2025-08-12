package main

import (
	"github.com/JohnCrickett/GoRayTracingInAWeekend/tracer"
)

const (
	targetFile = "test.ppm"
)

func main() {
	// World
	var world tracer.HittableList
	world.Add(tracer.NewSphere(tracer.Vec{0, 0, -1}, 0.5))
	world.Add(tracer.NewSphere(tracer.Vec{0, -100.5, -1}, 100))

	// Camera
	c := tracer.NewCamera(400, 16.0/9.0, 100, 50)
	c.Render(world, targetFile)
}
