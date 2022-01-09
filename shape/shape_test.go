package shape

import (
	"github.com/stretchr/testify/assert"
	"goray/material"
	"goray/matrix"
	"goray/ray"
	"goray/transformation"
	"goray/tuple"
	"testing"
)

type TestShape struct {
	Shape
}

func NewTestShape() *Shape {
	return NewShape(TestShape{})
}

func (ts TestShape) calculateIntersections(r *ray.Ray, s *Shape) ray.Intersections {
	return ray.Intersections{}
}

func (ts TestShape) calculateNormalAt(point *tuple.Tuple) *tuple.Tuple {
	return tuple.NewVector(point.X, point.Y, point.Z)
}

func TestDefaultTransformation(t *testing.T) {
	s := NewTestShape()

	assert.True(t, s.transformation.Equals(matrix.NewIdentityMatrix4x4()))
}

func TestAssigningTransformation(t *testing.T) {
	s := NewTestShape()

	s.SetTransformation(transformation.NewTranslation(2, 3, 4))

	assert.True(t, s.transformation.Equals(transformation.NewTranslation(2, 3, 4)))
}

func TestDefaultMaterial(t *testing.T) {
	s := NewTestShape()

	assert.Equal(t, material.NewMaterial(), s.material)
}

func TestAssigningMaterial(t *testing.T) {
	s := NewTestShape()

	m := material.NewMaterial()
	m.Ambient = 1

	s.SetMaterial(m)

	assert.Equal(t, m, s.material)
}
