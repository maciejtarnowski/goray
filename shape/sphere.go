package shape

import (
	"goray/ray"
	"goray/tuple"
	"math"
)

type Sphere struct {}

func NewSphere() *Shape {
	return NewShape(Sphere{})
}

func (s Sphere) calculateIntersections(r *ray.Ray) (float64, float64, bool) {
	sphereToRay := r.Origin.Sub(tuple.NewPoint(0, 0, 0))

	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return 0, 0, false
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return t1, t2, true
}

func (s Sphere) calculateNormalAt(point *tuple.Tuple) *tuple.Tuple {
	return point.Sub(tuple.NewPoint(0, 0, 0))
}
