package shape

import (
	"github.com/stretchr/testify/assert"
	"goray/ray"
	"goray/tuple"
	"testing"
)

func TestNormalOfPlaneIsConstantEverywhere(t *testing.T) {
	p := Plane{}

	n1 := p.calculateNormalAt(tuple.NewPoint(0, 0, 0))
	n2 := p.calculateNormalAt(tuple.NewPoint(10, 0, -10))
	n3 := p.calculateNormalAt(tuple.NewPoint(-5, 0, 150))

	assert.True(t, n1.Equals(tuple.NewVector(0, 1, 0)))
	assert.True(t, n2.Equals(tuple.NewVector(0, 1, 0)))
	assert.True(t, n3.Equals(tuple.NewVector(0, 1, 0)))
}

func TestIntersectWithRayParallelToPlane(t *testing.T) {
	p := Plane{}
	r := ray.NewRay(tuple.NewPoint(0, 10, 0), tuple.NewVector(0, 0, 1))

	xs := p.calculateIntersections(r, &Shape{})

	assert.Zero(t, xs.Len())
}

func TestIntersectWithCoplanarRay(t *testing.T) {
	p := Plane{}
	r := ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))

	xs := p.calculateIntersections(r, &Shape{})

	assert.Zero(t, xs.Len())
}

func TestRayIntersectingPlaneFromAbove(t *testing.T) {
	p := Plane{}
	r := ray.NewRay(tuple.NewPoint(0, 1, 0), tuple.NewVector(0, -1, 0))
	s := &Shape{}

	xs := p.calculateIntersections(r, s)

	assert.Equal(t, 1, xs.Len())
	assert.Equal(t, float64(1), xs.Get(0).T)
	assert.Equal(t, s, xs.Get(0).Object)
}

func TestRayIntersectingPlaneFromBelow(t *testing.T) {
	p := Plane{}
	r := ray.NewRay(tuple.NewPoint(0, -1, 0), tuple.NewVector(0, 1, 0))
	s := &Shape{}

	xs := p.calculateIntersections(r, s)

	assert.Equal(t, 1, xs.Len())
	assert.Equal(t, float64(1), xs.Get(0).T)
	assert.Equal(t, s, xs.Get(0).Object)
}
