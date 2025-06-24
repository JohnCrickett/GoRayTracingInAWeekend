package tracer

type Ray struct {
	origin    Point
	direction Vec
}

func NewRay(origin Point, direction Vec) *Ray {
	return &Ray{origin, direction}
}

func (r Ray) Origin() Point {
	return r.origin
}

func (r Ray) Direction() Vec {
	return r.direction
}

func (r Ray) At(t float64) Point {
	return r.origin.Plus(r.direction.Scale(t))
}
