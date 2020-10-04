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