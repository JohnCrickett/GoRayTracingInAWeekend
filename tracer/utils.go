package tracer

import "math/rand/v2"

const Pi = 3.1415926535897932385

func DegreesToRadians(degrees float64) float64 {
	return degrees * Pi / 180.0
}

func RandomDouble() float64 {
	return rand.Float64()
}

func RandomDoubleInRange(min, max float64) float64 {
	return min + (max-min)*RandomDouble()
}
