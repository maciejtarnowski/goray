package shape

import (
	"goray/material"
	"goray/matrix"
	"goray/ray"
	"goray/tuple"
	"math"
)

type Sphere struct {
	transformation *matrix.Matrix
	material       *material.Material
}

func NewSphere() *Sphere {
	return &Sphere{transformation: matrix.NewIdentityMatrix4x4(), material: material.NewMaterial()}
}

func (s *Sphere) SetMaterial(m *material.Material) {
	s.material = m
}

func (s *Sphere) GetMaterial() *material.Material {
	return s.material
}

func (s *Sphere) SetTransformation(m *matrix.Matrix) {
	s.transformation = m
}

func (s *Sphere) Intersect(r *ray.Ray) ray.Intersections {
	r2 := r.Transform(s.transformation.Invert())

	sphereToRay := r2.Origin.Sub(tuple.NewPoint(0, 0, 0))

	a := r2.Direction.Dot(r2.Direction)
	b := 2 * r2.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - 4*a*c

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

func (s *Sphere) NormalAt(point *tuple.Tuple) *tuple.Tuple {
	objectPoint := s.transformation.Invert().MultiplyTuple(point)

	objectNormal := objectPoint.Sub(tuple.NewPoint(0, 0, 0))

	worldNormal := s.transformation.Invert().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
