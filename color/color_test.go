package color

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColor(t *testing.T) {
	color := Color{Red: -0.5, Green: 0.4, Blue: 1.7}

	require.InDelta(t, -0.5, color.Red, 0.00001)
	require.InDelta(t, 0.4, color.Green, 0.00001)
	require.InDelta(t, 1.7, color.Blue, 0.00001)
}

func TestAddingColors(t *testing.T) {
	color1 := NewColor(0.9, 0.6, 0.75)
	color2 := NewColor(0.7, 0.1, 0.25)

	expected := NewColor(1.6, 0.7, 1)

	require.True(t, color1.Add(color2).Equals(expected))
}

func TestSubtractingColors(t *testing.T) {
	color1 := NewColor(0.9, 0.6, 0.75)
	color2 := NewColor(0.7, 0.1, 0.25)

	expected := NewColor(0.2, 0.5, 0.5)

	require.True(t, color1.Sub(color2).Equals(expected))
}

func TestMultiplyingColorsByScalar(t *testing.T) {
	color := NewColor(0.9, 0.6, 0.75)

	expected := NewColor(1.8, 1.2, 1.5)

	require.True(t, color.MultiplyScalar(2).Equals(expected))
}

func TestMultiplyingColorByColor(t *testing.T) {
	color1 := NewColor(1, 0.2, 0.4)
	color2 := NewColor(0.9, 1, 0.1)

	expected := NewColor(0.9, 0.2, 0.04)

	require.True(t, color1.Multiply(color2).Equals(expected))
}

func TestColorClone(t *testing.T) {
	color := NewColor(1, 0.5, 0)

	colorClone := color.Clone()

	require.True(t, color.Equals(colorClone))
	require.NotSame(t, color, colorClone)
}

func TestConvertingColorToRGB(t *testing.T) {
	color := NewColor(1, 0.5, 0)

	rgb := color.ToRGB()

	assert.Equal(t, 255, rgb[0])
	assert.Equal(t, 128, rgb[1])
	assert.Equal(t, 0, rgb[2])
}

func TestConvertingColorToRGBWithClamping(t *testing.T) {
	color := NewColor(1.5, -0.5, 0.5)

	rgb := color.ToRGB()

	assert.Equal(t, 255, rgb[0])
	assert.Equal(t, 0, rgb[1])
	assert.Equal(t, 128, rgb[2])
}
