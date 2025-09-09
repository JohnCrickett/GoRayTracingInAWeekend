package tracer

import "math"

type Sphere struct {
	Hittable
	Center   Point
	Material Material
	Radius   float64
}

func NewSphere(center Point, radius float64, material Material) *Sphere {
	radius = math.Max(0, radius)
	return &Sphere{
		Center:   center,
		Material: material,
		Radius:   radius,
	}
}

func (s Sphere) Hit(ray *Ray, rayT Interval) (bool, *HitRecord) {
	oc := s.Center.Minus(ray.Origin())
	a := ray.Direction().LengthSquared()
	h := Dot(ray.Direction(), oc)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := h*h - a*c

	if discriminant < 0 {
		return false, nil
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (h - sqrtd) / a
	//if root <= rayT.Min || rayT.Max <= root {
	if !rayT.Surrounds(root) {
		root = (h + sqrtd) / a
		if !rayT.Surrounds(root) {

			return false, nil
		}
	}

	var rec HitRecord
	rec.T = root
	rec.P = ray.At(rec.T)
	rec.Material = s.Material
	outwardNormal := rec.P.Minus(s.Center).Divide(s.Radius)
	rec.SetFaceNormal(ray, outwardNormal)

	return true, &rec
}
