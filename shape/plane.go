package shape

import (
	"goray/ray"
	"goray/tuple"
	"goray/utils"
	"math"
)

type Plane struct {}

func NewPlane() *Shape {
	return NewShape(Plane{})
}

func (p Plane) calculateIntersections(r *ray.Ray, s *Shape) ray.Intersections {
	if math.Abs(r.Direction.Y) < utils.EPSILON {
		return ray.Intersections{}
	}

	t := -r.Origin.Y / r.Direction.Y
	return *ray.NewIntersections(ray.NewIntersection(t, s))
}

func (p Plane) calculateNormalAt(point *tuple.Tuple) *tuple.Tuple {
	return tuple.NewVector(0, 1, 0)
}
