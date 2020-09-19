package shape

import (
	"github.com/stretchr/testify/assert"
	"goray/ray"
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
