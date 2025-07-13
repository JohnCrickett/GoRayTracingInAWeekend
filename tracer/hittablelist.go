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

func (h HittableList) Hit(ray *Ray, tmin float64, tmax float64) (bool, *HitRecord) {
	var hitAnything bool
	var rec *HitRecord

	closest := tmax

	for _, obj := range h.Objects {
		hit, hitRecord := obj.Hit(ray, tmin, closest)
		if hit {
			closest = hitRecord.T
			rec = hitRecord
			hitAnything = true
		}
	}

	return hitAnything, rec
}
