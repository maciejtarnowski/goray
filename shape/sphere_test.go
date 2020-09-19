package shape

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goray/matrix"
	"goray/ray"
	"goray/transformation"
	"goray/tuple"
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

	assert.Equal(t, matrix.NewIdentityMatrix4x4(), s.Transformation)
}

func TestChangingSphereTransformation(t *testing.T) {
	s := NewSphere()
	tr := transformation.NewTranslation(2, 3, 4)

	s.Transformation = tr

	assert.Equal(t, tr, s.Transformation)
}

func TestIntersectingScaledSphere(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.Transformation = transformation.NewScaling(2, 2, 2)

	xs := s.Intersect(r)

	require.Equal(t, 2, xs.Len())

	assert.Equal(t, float64(3), xs.ValueAt(0))
	assert.Equal(t, float64(7), xs.ValueAt(1))
}

func TestIntersectingTranslatedSphere(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.Transformation = transformation.NewTranslation(5, 0, 0)

	xs := s.Intersect(r)

	assert.Equal(t, 0, xs.Len())
}
