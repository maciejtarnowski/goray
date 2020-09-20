package material

import (
	"github.com/stretchr/testify/assert"
	"goray/color"
	"goray/light"
	"goray/tuple"
	"math"
	"testing"
)

func TestDefaultMaterial(t *testing.T) {
	m := NewMaterial()

	assert.Equal(t, color.NewColor(1, 1, 1), m.Color)
	assert.Equal(t, 0.1, m.Ambient)
	assert.Equal(t, 0.9, m.Diffuse)
	assert.Equal(t, 0.9, m.Specular)
	assert.Equal(t, 200.0, m.Shininess)
}

func TestLightingWithEyeBetweenLightAndSurface(t *testing.T) {
	m := NewMaterial()
	position := tuple.NewPoint(0, 0, 0)

	eyeV := tuple.NewVector(0, 0, -1)
	normalV := tuple.NewVector(0, 0, -1)
	l := light.NewPointLight(tuple.NewPoint(0, 0, -10), color.NewColor(1, 1, 1))

	result := m.Lighting(l, position, eyeV, normalV)

	assert.Equal(t, color.NewColor(1.9, 1.9, 1.9), result)
}

func TestLightingWithEyeBetweenLightAndSurfaceWithEyeOffset45Degrees(t *testing.T) {
	m := NewMaterial()
	position := tuple.NewPoint(0, 0, 0)

	eyeV := tuple.NewVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV := tuple.NewVector(0, 0, -1)
	l := light.NewPointLight(tuple.NewPoint(0, 0, -10), color.NewColor(1, 1, 1))

	result := m.Lighting(l, position, eyeV, normalV)

	assert.Equal(t, color.NewColor(1.0, 1.0, 1.0), result)
}

func TestLightingWithEyeOppositeSurfaceWithLightOffset45Degrees(t *testing.T) {
	m := NewMaterial()
	position := tuple.NewPoint(0, 0, 0)

	eyeV := tuple.NewVector(0, 0, -1)
	normalV := tuple.NewVector(0, 0, -1)
	l := light.NewPointLight(tuple.NewPoint(0, 10, -10), color.NewColor(1, 1, 1))

	result := m.Lighting(l, position, eyeV, normalV)

	assert.True(t, result.Equals(color.NewColor(0.7364, 0.7364, 0.7364)))
}

func TestLightingWithEyeInPathOfReflectionVector(t *testing.T) {
	m := NewMaterial()
	position := tuple.NewPoint(0, 0, 0)

	eyeV := tuple.NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalV := tuple.NewVector(0, 0, -1)
	l := light.NewPointLight(tuple.NewPoint(0, 10, -10), color.NewColor(1, 1, 1))

	result := m.Lighting(l, position, eyeV, normalV)

	assert.True(t, result.Equals(color.NewColor(1.6364, 1.6364, 1.6364)))
}

func TestLightingWithLightBehindSurface(t *testing.T) {
	m := NewMaterial()
	position := tuple.NewPoint(0, 0, 0)

	eyeV := tuple.NewVector(0, 0, -1)
	normalV := tuple.NewVector(0, 0, -1)
	l := light.NewPointLight(tuple.NewPoint(0, 0, 10), color.NewColor(1, 1, 1))

	result := m.Lighting(l, position, eyeV, normalV)

	assert.True(t, result.Equals(color.NewColor(0.1, 0.1, 0.1)))
}
