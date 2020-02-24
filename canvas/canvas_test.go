package canvas

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goray/color"
	"strings"
	"testing"
)

func TestCreatingCanvas(t *testing.T) {
	canvas := NewCanvas(10, 20)

	require.Equal(t, 10, canvas.Width)
	require.Equal(t, 20, canvas.Height)

	expectedColor := color.NewColor(0, 0, 0)

	for x := 0; x < 10; x++ {
		for y := 0; y < 20; y++ {
			assert.True(t, canvas.PixelAt(x, y).Equals(expectedColor))
		}
	}
}

func TestWritingPixels(t *testing.T) {
	canvas := NewCanvas(10, 20)
	red := color.NewColor(1, 0, 0)

	canvas.WriteAt(2, 3, red)

	require.True(t, canvas.PixelAt(2, 3).Equals(red))
}

func TestConstructingPPMHeader(t *testing.T) {
	canvas := NewCanvas(5, 3)

	expected := []string{"P3", "5 3", "255"}

	ppm := canvas.ToPPM()
	ppmLines := strings.Split(ppm, "\n")

	for i, expectedLine := range expected {
		assert.Equal(t, expectedLine, ppmLines[i])
	}
}

func TestConstructingPPM(t *testing.T) {
	canvas := NewCanvas(5, 3)
	canvas.WriteAt(0, 0, color.NewColor(1.5, 0, 0))
	canvas.WriteAt(2, 1, color.NewColor(0, 0.5, 0))
	canvas.WriteAt(4, 2, color.NewColor(-0.5, 0, 1))

	expected := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	require.Equal(t, expected, canvas.ToPPM())
}

func TestConstructingPPMWithLongLines(t *testing.T) {
	canvas := NewCanvas(10, 2)
	canvas.FillWith(color.NewColor(1, 0.8, 0.6))

	expected := `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`

	require.Equal(t, expected, canvas.ToPPM())
}

func TestPPMEndsWithLineFeed(t *testing.T) {
	canvas := NewCanvas(5, 3)

	ppm := canvas.ToPPM()

	require.Equal(t, "\n", ppm[len(ppm)-1:])
}
