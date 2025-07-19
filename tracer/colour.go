package tracer

import (
	"fmt"
	"io"
)

var intensity = NewInterval(0.000, 0.999)

type Colour [3]float64

func (c Colour) R() float64 {
	return c[0]
}

func (c Colour) Rbyte() uint8 {
	return uint8(256 * intensity.Clamp(c.R()))
}

func (c Colour) G() float64 {
	return c[1]
}

func (c Colour) Gbyte() uint8 {
	return uint8(256 * intensity.Clamp(c.G()))
}

func (c Colour) B() float64 {
	return c[2]
}

func (c Colour) Bbyte() uint8 {
	return uint8(256 * intensity.Clamp(c.B()))
}

func (c Colour) Write(f io.Writer) {
	fmt.Fprintf(f, "%d %d %d\n", c.Rbyte(), c.Gbyte(), c.Bbyte())
}

func (c Colour) Scale(n float64) Colour {
	return Colour{c[0] * n, c[1] * n, c[2] * n}
}

func (c Colour) Plus(oc Colour) Colour {
	return Colour{c[0] + oc[0], c[1] + oc[1], c[2] + oc[2]}
}

func (c Colour) Add(colour Colour) Colour {
	return Colour{c.R() + colour.R(), c.G() + colour.G(), c.B() + colour.B()}
}
