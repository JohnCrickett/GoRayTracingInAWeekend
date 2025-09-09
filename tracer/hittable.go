package tracer

type HitRecord struct {
	P         Point
	Normal    Vec
	Material  Material
	T         float64
	FrontFace bool
}

type Hittable interface {
	Hit(ray *Ray, rayT Interval) (bool, *HitRecord)
}

func (h *HitRecord) SetFaceNormal(ray *Ray, outwardNormal Vec) {
	dot := Dot(ray.Direction(), outwardNormal)

	if dot < 0 {
		h.FrontFace = true
	} else {
		h.FrontFace = false
	}

	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Scale(-1)
	}
}
