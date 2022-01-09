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

func (sp Sphere) calculateIntersections(r *ray.Ray, s *Shape) ray.Intersections {
	sphereToRay := r.Origin.Sub(tuple.NewPoint(0, 0, 0))

	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

	if discriminant < 0 {
		return ray.Intersections{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	xs := ray.NewIntersections()
	xs.Add(ray.NewIntersection(math.Min(t1, t2), s))
	xs.Add(ray.NewIntersection(math.Max(t1, t2), s))

	return *xs
}

func (sp Sphere) calculateNormalAt(point *tuple.Tuple) *tuple.Tuple {
	return point.Sub(tuple.NewPoint(0, 0, 0))
}
