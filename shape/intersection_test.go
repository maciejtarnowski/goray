package shape

import (
	"github.com/stretchr/testify/assert"
	"goray/ray"
	"goray/tuple"
	"testing"
)

func TestIntersectionHasTAndObject(t *testing.T) {
	s := NewSphere()

	i := ray.NewIntersection(3.5, s)

	assert.Equal(t, 3.5, i.T)
	assert.Equal(t, s, i.Object)
}

func TestAggregatingIntersections(t *testing.T) {
	s := NewSphere()

	i1 := ray.NewIntersection(1, s)
	i2 := ray.NewIntersection(2, s)

	xs := ray.NewIntersections(i1, i2)

	assert.Equal(t, 2, xs.Len())

	assert.Equal(t, float64(1), xs.ValueAt(0))
	assert.Equal(t, s, xs.ObjectAt(0))

	assert.Equal(t, float64(2), xs.ValueAt(1))
	assert.Equal(t, s, xs.ObjectAt(1))
}

func TestHitWithAllPositiveT(t *testing.T) {
	s := NewSphere()

	i1 := ray.NewIntersection(1, s)
	i2 := ray.NewIntersection(2, s)

	xs := ray.NewIntersections(i1, i2)

	assert.Equal(t, i1, xs.Hit())
}

func TestHitWithSomeNegativeT(t *testing.T) {
	s := NewSphere()

	i1 := ray.NewIntersection(-1, s)
	i2 := ray.NewIntersection(1, s)

	xs := ray.NewIntersections(i1, i2)

	assert.Equal(t, i2, xs.Hit())
}

func TestHitWithAllNegativeT(t *testing.T) {
	s := NewSphere()

	i1 := ray.NewIntersection(-2, s)
	i2 := ray.NewIntersection(-1, s)

	xs := ray.NewIntersections(i1, i2)

	assert.Nil(t, xs.Hit())
}

func TestHitIsLowestNonNegativeT(t *testing.T) {
	s := NewSphere()

	i1 := ray.NewIntersection(5, s)
	i2 := ray.NewIntersection(7, s)
	i3 := ray.NewIntersection(-3, s)
	i4 := ray.NewIntersection(2, s)

	xs := ray.NewIntersections(i1, i2, i3, i4)

	assert.Equal(t, i4, xs.Hit())
}

func TestPrecomputingStateOfIntersection(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	i := ray.NewIntersection(4, s)

	comps := i.PrepareComputations(r)

	assert.Equal(t, i.T, comps.T)
	assert.Equal(t, i.Object, comps.Object)
	assert.True(t, tuple.NewPoint(0, 0, -1).Equals(comps.Point))
	assert.True(t, tuple.NewVector(0, 0, -1).Equals(comps.EyeV))
	assert.True(t, tuple.NewVector(0, 0, -1).Equals(comps.NormalV))
}

func TestHitWithIntersectionOnTheOutside(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	i := ray.NewIntersection(4, s)

	comps := i.PrepareComputations(r)

	assert.False(t, comps.Inside)
}

func TestHitWithIntersectionOnTheInside(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	i := ray.NewIntersection(1, s)

	comps := i.PrepareComputations(r)

	assert.True(t, tuple.NewPoint(0, 0, 1).Equals(comps.Point))
	assert.True(t, tuple.NewVector(0, 0, -1).Equals(comps.EyeV))
	assert.True(t, comps.Inside)
	assert.True(t, tuple.NewVector(0, 0, -1).Equals(comps.NormalV))
}
