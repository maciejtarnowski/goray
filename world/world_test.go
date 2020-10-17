package world

import (
	"github.com/stretchr/testify/assert"
	"goray/color"
	"goray/light"
	"goray/ray"
	"goray/shape"
	"goray/transformation"
	"goray/tuple"
	"testing"
)

func TestEmptyWorld(t *testing.T) {
	w := NewWorld()

	assert.Nil(t, w.Light)
	assert.Len(t, w.Objects, 0)
}

func TestDefaultWorld(t *testing.T) {
	w := NewDefaultWorld()

	l := light.NewPointLight(tuple.NewPoint(-10, 10, -10), color.NewColor(1, 1, 1))

	s1 := shape.NewSphere()
	s1.GetMaterial().Color = color.NewColor(0.8, 1.0, 0.6)
	s1.GetMaterial().Diffuse = 0.7
	s1.GetMaterial().Specular = 0.2

	s2 := shape.NewSphere()
	s2.SetTransformation(transformation.NewScaling(0.5, 0.5, 0.5))

	assert.Equal(t, l, w.Light)
	assert.Contains(t, w.Objects, s1)
	assert.Contains(t, w.Objects, s2)
}

func TestIntersectWorldWithRay(t *testing.T) {
	w := NewDefaultWorld()
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))

	xs := w.Intersect(r)

	assert.Equal(t, 4, xs.Len())
	assert.Equal(t, float64(4), xs.ValueAt(0))
	assert.Equal(t, 4.5, xs.ValueAt(1))
	assert.Equal(t, 5.5, xs.ValueAt(2))
	assert.Equal(t, float64(6), xs.ValueAt(3))
}

func TestShadingIntersection(t *testing.T) {
	w := NewDefaultWorld()
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := w.Objects[0]
	i := ray.NewIntersection(4, s)

	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)

	assert.True(t, color.NewColor(0.38066, 0.47583, 0.2855).Equals(c))
}

func TestShadingIntersectionFromTheInside(t *testing.T) {
	w := NewDefaultWorld()
	w.Light = light.NewPointLight(tuple.NewPoint(0, 0.25, 0), color.NewColor(1, 1, 1))
	r := ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	s := w.Objects[1]
	i := ray.NewIntersection(0.5, s)

	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)

	assert.True(t, color.NewColor(0.90498, 0.90498, 0.90498).Equals(c))
}

func TestColorWhenRayMisses(t *testing.T) {
	w := NewDefaultWorld()
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 1, 0))

	c := w.ColorAt(r)

	assert.True(t, color.NewColor(0, 0, 0).Equals(c))
}

func TestColorWhenRayHits(t *testing.T) {
	w := NewDefaultWorld()
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))

	c := w.ColorAt(r)

	assert.True(t, color.NewColor(0.38066, 0.47583, 0.2855).Equals(c))
}

func TestColorWithIntersectionBehindRay(t *testing.T) {
	w := NewDefaultWorld()
	outer := w.Objects[0]
	outer.GetMaterial().Ambient = 1
	inner := w.Objects[1]
	inner.GetMaterial().Ambient = 1
	r := ray.NewRay(tuple.NewPoint(0, 0, 0.75), tuple.NewVector(0, 0, -1))

	c := w.ColorAt(r)

	assert.True(t, inner.GetMaterial().Color.Equals(c))
}

func TestNoShadowWhenNothingIsCollinearWithPointAndLight(t *testing.T) {
	w := NewDefaultWorld()
	p := tuple.NewPoint(0, 10, 0)

	assert.False(t, w.IsShadowed(p))
}

func TestShadowWhenObjectIsBetweenPointAndLight(t *testing.T) {
	w := NewDefaultWorld()
	p := tuple.NewPoint(10, -10, 10)

	assert.True(t, w.IsShadowed(p))
}

func TestNoShadowWhenObjectIsBehindLight(t *testing.T) {
	w := NewDefaultWorld()
	p := tuple.NewPoint(-20, 20, -20)

	assert.False(t, w.IsShadowed(p))
}

func TestNoShadowWhenObjectIsBehindPoint(t *testing.T) {
	w := NewDefaultWorld()
	p := tuple.NewPoint(-2, 2, -2)

	assert.False(t, w.IsShadowed(p))
}

func TestShadingIntersectionInShadow(t *testing.T) {
	w := NewWorld()
	w.Light = light.NewPointLight(tuple.NewPoint(0, 0, -10), color.NewColor(1, 1, 1))
	s1 := shape.NewSphere()
	s2 := shape.NewSphere()
	s2.SetTransformation(transformation.NewTranslation(0, 0, 10))
	w.Objects = []ray.Object{s1, s2}
	r := ray.NewRay(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
	i := ray.NewIntersection(4, s2)

	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)

	assert.True(t, c.Equals(color.NewColor(0.1, 0.1, 0.1)))
}
