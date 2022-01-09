package shape

import (
	"goray/material"
	"goray/matrix"
	"goray/ray"
	"goray/tuple"
)

type shapeType interface {
	calculateIntersections(r *ray.Ray, s *Shape) ray.Intersections
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

	return s.shapeType.calculateIntersections(objectRay, s)
}

func (s *Shape) NormalAt(point *tuple.Tuple) *tuple.Tuple {
	objectPoint := s.transformation.Invert().MultiplyTuple(point)
	objectNormal := s.shapeType.calculateNormalAt(objectPoint)

	worldNormal := s.transformation.Invert().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
