package shape

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goray/material"
	"goray/matrix"
	"goray/ray"
	"goray/transformation"
	"goray/tuple"
	"math"
	"testing"
)

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	xs := s.Intersect(r)

	assert.Equal(t, 2, xs.Len())
	assert.InDelta(t, 4.0, xs.ValueAt(0), 0.00001)
	assert.InDelta(t, 6.0, xs.ValueAt(1), 0.00001)
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	xs := s.Intersect(r)

	assert.Equal(t, 2, xs.Len())
	assert.InDelta(t, 5.0, xs.ValueAt(0), 0.00001)
	assert.InDelta(t, 5.0, xs.ValueAt(1), 0.00001)
}

func TestRayMissesSphere(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	xs := s.Intersect(r)

	assert.Equal(t, 0, xs.Len())
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	xs := s.Intersect(r)

	assert.Equal(t, 2, xs.Len())
	assert.InDelta(t, -1.0, xs.ValueAt(0), 0.00001)
	assert.InDelta(t, 1.0, xs.ValueAt(1), 0.00001)
}

func TestSphereBehindRay(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	xs := s.Intersect(r)

	assert.Equal(t, 2, xs.Len())
	assert.InDelta(t, -6.0, xs.ValueAt(0), 0.00001)
	assert.InDelta(t, -4.0, xs.ValueAt(1), 0.00001)
}

func TestIntersectSetsObjectInIntersection(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	xs := s.Intersect(r)

	assert.Equal(t, 2, xs.Len())
	assert.Equal(t, s, xs.ObjectAt(0))
	assert.Equal(t, s, xs.ObjectAt(1))
}

func TestSphereHasDefaultTransformation(t *testing.T) {
	s := NewSphere()

	assert.Equal(t, matrix.NewIdentityMatrix4x4(), s.transformation)
}

func TestChangingSphereTransformation(t *testing.T) {
	s := NewSphere()
	tr := transformation.NewTranslation(2, 3, 4)

	s.transformation = tr

	assert.Equal(t, tr, s.transformation)
}

func TestIntersectingScaledSphere(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.transformation = transformation.NewScaling(2, 2, 2)

	xs := s.Intersect(r)

	require.Equal(t, 2, xs.Len())

	assert.Equal(t, float64(3), xs.ValueAt(0))
	assert.Equal(t, float64(7), xs.ValueAt(1))
}

func TestIntersectingTranslatedSphere(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.transformation = transformation.NewTranslation(5, 0, 0)

	xs := s.Intersect(r)

	assert.Equal(t, 0, xs.Len())
}

func TestNormalAtPointOnXAxis(t *testing.T) {
	s := NewSphere()

	n := s.NormalAt(tuple.NewPoint(1, 0, 0))

	assert.Equal(t, tuple.NewVector(1, 0, 0), n)
}

func TestNormalAtPointOnYAxis(t *testing.T) {
	s := NewSphere()

	n := s.NormalAt(tuple.NewPoint(0, 1, 0))

	assert.Equal(t, tuple.NewVector(0, 1, 0), n)
}

func TestNormalAtPointOnZAxis(t *testing.T) {
	s := NewSphere()

	n := s.NormalAt(tuple.NewPoint(0, 0, 1))

	assert.Equal(t, tuple.NewVector(0, 0, 1), n)
}

func TestNormalAtPointOnNonAxialPoint(t *testing.T) {
	s := NewSphere()

	n := s.NormalAt(tuple.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	assert.Equal(t, tuple.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), n)
}

func TestNormalIsNormalized(t *testing.T) {
	s := NewSphere()

	n := s.NormalAt(tuple.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	assert.Equal(t, n.Normalize(), n)
}

func TestNormalOnTranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.transformation = transformation.NewTranslation(0, 1, 0)

	n := s.NormalAt(tuple.NewPoint(0, 1.70711, -0.70711))

	assert.True(t, tuple.NewVector(0, 0.70711, -0.70711).Equals(n))
}

func TestNormalOnTransformedSphere(t *testing.T) {
	s := NewSphere()
	s.transformation = transformation.NewScaling(1, 0.5, 1).MultiplyMatrix(transformation.NewRotationZ(math.Pi / 5))

	n := s.NormalAt(tuple.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))

	assert.True(t, tuple.NewVector(0, 0.97014, -0.24254).Equals(n))
}

func TestSphereHasDefaultMaterial(t *testing.T) {
	s := NewSphere()

	assert.Equal(t, material.NewMaterial(), s.material)
}

func TestChangingSphereMaterial(t *testing.T) {
	s := NewSphere()
	m := material.NewMaterial()
	m.Ambient = 1

	s.material = m

	assert.Equal(t, m, s.material)
}
