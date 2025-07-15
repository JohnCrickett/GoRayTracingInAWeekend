package tracer

import "math"

type Interval struct {
	Min float64
	Max float64
}

func (i Interval) Size() float64 {
	return i.Max - i.Min
}

func (i Interval) Contains(v float64) bool {
	return v >= i.Min && v <= i.Max
}

func (i Interval) Surrounds(v float64) bool {
	return v > i.Min && v < i.Max
}

func NewInterval(min, max float64) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

var EmptyInterval = NewInterval(math.MaxInt, math.MinInt)
var UniversalInterval = NewInterval(math.MinInt, math.MaxInt)
