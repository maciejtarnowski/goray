package shape

import (
	"goray/material"
	"goray/matrix"
	"goray/ray"
	"goray/tuple"
	"math"
)

type shapeType interface {
	calculateIntersections(r *ray.Ray) (float64, float64, bool)
	calculateNormalAt(point *tuple.Tuple) *tuple.Tuple
}

type Shape struct {
	transformation *matrix.Matrix
	material       *material.Material
	shapeType      shapeType
}

func NewShape(shapeType shapeType) *Shape {
	return &Shape{
		transformation: matrix.NewIdentityMatrix4x4(),
		material: material.NewMaterial(),
		shapeType: shapeType,
	}
}

func (s *Shape) SetMaterial(m *material.Material) {
	s.material = m
}

func (s *Shape) GetMaterial() *material.Material {
	return s.material
}

func (s *Shape) SetTransformation(m *matrix.Matrix) {
	s.transformation = m
}

func (s *Shape) GetTransformation() *matrix.Matrix {
	return s.transformation
}

func (s *Shape) Intersect(r *ray.Ray) ray.Intersections {
	objectRay := r.Transform(s.transformation.Invert())

	t1, t2, found := s.shapeType.calculateIntersections(objectRay)

	if !found {
		return ray.Intersections{}
	}

	xs := ray.NewIntersections()
	xs.Add(ray.NewIntersection(math.Min(t1, t2), s))
	xs.Add(ray.NewIntersection(math.Max(t1, t2), s))

	return *xs
}

func (s *Shape) NormalAt(point *tuple.Tuple) *tuple.Tuple {
	objectPoint := s.transformation.Invert().MultiplyTuple(point)
	objectNormal := s.shapeType.calculateNormalAt(objectPoint)

	worldNormal := s.transformation.Invert().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
