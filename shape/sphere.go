package shape

import (
	"goray/matrix"
	"goray/ray"
	"goray/tuple"
	"math"
)

type Sphere struct {
	Transformation *matrix.Matrix
}

func NewSphere() *Sphere {
	return &Sphere{Transformation: matrix.NewIdentityMatrix4x4()}
}

func (s *Sphere) Intersect(r *ray.Ray) ray.Intersections {
	r2 := r.Transform(s.Transformation.Invert())

	sphereToRay := r2.Origin.Sub(tuple.NewPoint(0, 0, 0))

	a := r2.Direction.Dot(r2.Direction)
	b := 2 * r2.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b*b) - 4 * a * c

	if discriminant < 0 {
		return *ray.NewIntersections()
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	xs := ray.NewIntersections()
	xs.Add(ray.NewIntersection(math.Min(t1, t2), s))
	xs.Add(ray.NewIntersection(math.Max(t1, t2), s))

	return *xs
}
