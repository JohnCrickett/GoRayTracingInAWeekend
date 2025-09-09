package tracer

type Material interface {
	scatter(r *Ray, h *HitRecord) (bool, *Ray, Colour)
}

type Lambertian struct {
	Albedo Colour
}

func (l Lambertian) scatter(r *Ray, h *HitRecord) (bool, *Ray, Colour) {
	scatterDirection := h.Normal.Plus(RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = h.Normal
	}

	scattered := NewRay(h.P, scatterDirection)
	at := l.Albedo
	return true, scattered, at
}

type Metal struct {
	Albedo Colour
	Fuzz   float64
}

func (l Metal) scatter(r *Ray, h *HitRecord) (bool, *Ray, Colour) {
	reflected := reflect(r.Direction(), h.Normal)
	reflected = UnitVector(reflected).Plus(RandomUnitVector().Scale(l.Fuzz))
	scattered := NewRay(h.P, reflected)
	at := l.Albedo
	var res bool
	if Dot(scattered.Direction(), h.Normal) > 0 {
		res = true
	}
	return res, scattered, at
}
