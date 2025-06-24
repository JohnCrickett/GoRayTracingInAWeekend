package tracer

import (
	"fmt"
	"io"
	"math"
)

type Vec [3]float64

type Point = Vec

func (v Vec) X() float64 {
	return v[0]
}

func (v Vec) Y() float64 {
	return v[1]
}

func (v Vec) Z() float64 {
	return v[2]
}

func (v Vec) Minus(ov Vec) Vec {
	return Vec{v[0] - ov[0], v[1] - ov[1], v[2] - ov[2]}
}

func (v Vec) Plus(ov Vec) Vec {
	return Vec{v[0] + ov[0], v[1] + ov[1], v[2] + ov[2]}
}

func (v Vec) Multiply(ov Vec) Vec {
	return Vec{v[0] * ov[0], v[1] * ov[1], v[2] * ov[2]}
}

func (v Vec) Scale(n float64) Vec {
	return Vec{v[0] * n, v[1] * n, v[2] * n}
}

func (v Vec) LengthSquared() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v Vec) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec) Write(f io.Writer) {
	fmt.Fprintf(f, "%d %d %d\n", v.X(), v.Y(), v.Z())
}

func Dot(v, ov Vec) float64 {
	return v[0]*ov[0] + v[1]*ov[1] + v[2]*ov[2]
}

func Cross(v, ov Vec) Vec {
	return Vec{
		v[1]*ov[2] - v[2]*ov[1],
		v[2]*ov[0] - v[0]*ov[2],
		v[0]*ov[1] - v[1]*ov[0],
	}
}

func UnitVector(v Vec) Vec {
	return v.Scale(1 / v.Length())
}
