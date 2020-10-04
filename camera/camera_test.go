package camera

import (
	"github.com/stretchr/testify/assert"
	"goray/color"
	"goray/matrix"
	"goray/transformation"
	"goray/tuple"
	"goray/world"
	"math"
	"testing"
)

func TestConstructingCamera(t *testing.T) {
	hsize := 160
	vsize := 120
	fieldOfView := math.Pi / 2

	c := NewCamera(hsize, vsize, fieldOfView)

	assert.Equal(t, hsize, c.HSize)
	assert.Equal(t, vsize, c.VSize)
	assert.Equal(t, fieldOfView, c.FieldOfView)
	assert.True(t, matrix.NewIdentityMatrix4x4().Equals(c.Transform))
}

func TestPixelSizeForHorizontalCanvas(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)

	assert.Equal(t, 0.01, c.PixelSize)
}

func TestPixelSizeForVerticalCanvas(t *testing.T) {
	c := NewCamera(125, 200, math.Pi/2)

	assert.Equal(t, 0.01, c.PixelSize)
}

func TestConstructingRayThroughTheCenterOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)

	r := c.RayForPixel(100, 50)

	assert.True(t, tuple.NewPoint(0, 0, 0).Equals(r.Origin))
	assert.True(t, tuple.NewVector(0, 0, -1).Equals(r.Direction))
}

func TestConstructingRayThroughCornerOfTheCanvas(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)

	r := c.RayForPixel(0, 0)

	assert.True(t, tuple.NewPoint(0, 0, 0).Equals(r.Origin))
	assert.True(t, tuple.NewVector(0.66519, 0.33259, -0.66851).Equals(r.Direction))
}

func TestConstructingRayWithTransformedCamera(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	c.Transform = transformation.NewRotationY(math.Pi / 4).MultiplyMatrix(transformation.NewTranslation(0, -2, 5))

	r := c.RayForPixel(100, 50)

	assert.True(t, tuple.NewPoint(0, 2, -5).Equals(r.Origin))
	assert.True(t, tuple.NewVector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2).Equals(r.Direction))
}

func TestRenderingWorldWithCamera(t *testing.T) {
	w := world.NewDefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := tuple.NewPoint(0, 0, -5)
	to := tuple.NewPoint(0, 0, 0)
	up := tuple.NewVector(0, 1, 0)
	c.Transform = transformation.ViewTransform(from, to, up)

	im := c.Render(w)

	assert.True(t, color.NewColor(0.38066, 0.47583, 0.2855).Equals(im.PixelAt(5, 5)))
}
