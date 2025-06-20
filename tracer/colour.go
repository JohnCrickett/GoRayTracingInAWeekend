package tracer

import (
	"fmt"
	"io"
)

type Colour [3]float64

func (c Colour) R() float64 {
	return c[0]
}

func (c Colour) Rbyte() uint8 {
	return uint8(255.999 * c.R())
}

func (c Colour) G() float64 {
	return c[1]
}

func (c Colour) Gbyte() uint8 {
	return uint8(255.999 * c.G())
}

func (c Colour) B() float64 {
	return c[2]
}

func (c Colour) Bbyte() uint8 {
	return uint8(255.999 * c.B())
}

func (c Colour) Write(f io.Writer) {
	fmt.Fprintf(f, "%d %d %d\n", c.Rbyte(), c.Gbyte(), c.Bbyte())
}
