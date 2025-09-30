package tracer

import "math"

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

type Dielectric struct {
	RefractionIndex float64
}

func (d Dielectric) scatter(r *Ray, h *HitRecord) (bool, *Ray, Colour) {
	at := Colour{1.0, 1.0, 1.0}
	var ri float64
	if h.FrontFace {
		ri = 1.0 / d.RefractionIndex
	} else {
		ri = d.RefractionIndex
	}

	unitDirection := UnitVector(r.Direction())

	cosTheta := math.Min(Dot(unitDirection.Scale(-1.0), h.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := ri*sinTheta > 1.0
	var direction Vec

	if cannotRefract || reflectance(cosTheta, ri) > RandomDouble() {
		direction = reflect(unitDirection, h.Normal)
	} else {
		direction = refract(unitDirection, h.Normal, ri)
	}

	scattered := NewRay(h.P, direction)
	return true, scattered, at
}

func reflectance(cosine, refractionIndex float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
