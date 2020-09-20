package light

import (
	"github.com/stretchr/testify/assert"
	"goray/color"
	"goray/tuple"
	"testing"
)

func TestPointLightHasPositionAndIntensity(t *testing.T) {
	intensity := color.NewColor(1, 1, 1)
	position := tuple.NewPoint(0, 0, 0)

	light := NewPointLight(position, intensity)

	assert.Equal(t, position, light.Position)
	assert.Equal(t, intensity, light.Intensity)
}
