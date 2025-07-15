package tracer

type HittableList struct {
	Hittable
	Objects []Hittable
}

func (h *HittableList) Add(obj Hittable) {
	h.Objects = append(h.Objects, obj)
}

func (h *HittableList) Clear() {
	h.Objects = nil
}

func (h HittableList) Hit(ray *Ray, rayT Interval) (bool, *HitRecord) {
	var hitAnything bool
	var rec *HitRecord

	closest := rayT.Max

	for _, obj := range h.Objects {
		hit, hitRecord := obj.Hit(ray, NewInterval(rayT.Min, closest))
		if hit {
			closest = hitRecord.T
			rec = hitRecord
			hitAnything = true
		}
	}

	return hitAnything, rec
}
