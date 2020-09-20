package material

import (
	"github.com/stretchr/testify/assert"
	"goray/color"
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
