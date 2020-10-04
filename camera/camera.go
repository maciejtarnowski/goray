package camera

import (
	"goray/canvas"
	"goray/matrix"
	"goray/ray"
	"goray/tuple"
	"goray/world"
	"math"
)

type Camera struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   *matrix.Matrix

	PixelSize  float64
	HalfWidth  float64
	HalfHeight float64
}

func NewCamera(hsize, vsize int, fov float64) *Camera {
	c := &Camera{HSize: hsize, VSize: vsize, FieldOfView: fov, Transform: matrix.NewIdentityMatrix4x4()}

	halfView := math.Tan(c.FieldOfView / 2)
	aspect := float64(c.HSize) / float64(c.VSize)

	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}

	c.PixelSize = (c.HalfWidth * 2) / float64(c.HSize)

	return c
}

func (c *Camera) RayForPixel(x, y int) *ray.Ray {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize

	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	transformInvert := c.Transform.Invert()
	pixel := transformInvert.MultiplyTuple(tuple.NewPoint(worldX, worldY, -1))
	origin := transformInvert.MultiplyTuple(tuple.NewPoint(0, 0, 0))
	direction := pixel.Sub(origin).Normalize()

	return ray.NewRay(origin, direction)
}

func (c *Camera) Render(w *world.World) *canvas.Canvas {
	im := canvas.NewCanvas(c.HSize, c.VSize)

	for y := 0; y < c.VSize; y++ {
		for x := 0; x < c.HSize; x++ {
			r := c.RayForPixel(x, y)
			c := w.ColorAt(r)

			im.WriteAt(x, y, c)
		}
	}

	return im
}
