package main

import (
	"fmt"
	"goray/canvas"
	"goray/color"
	"goray/ray"
	"goray/shape"
	"goray/transformation"
	"goray/tuple"
)

const (
	WALL_Z    = 10
	WALL_SIZE = 7.0
	HALF_WALL = WALL_SIZE / 2

	CANVAS_PIXELS = 100

	PIXEL_SIZE = WALL_SIZE / CANVAS_PIXELS
)

func main() {
	rayOrigin := tuple.NewPoint(0, 0, -5)

	canvs := canvas.NewCanvas(CANVAS_PIXELS, CANVAS_PIXELS)
	red := color.NewColor(1, 0, 0)
	s := shape.NewSphere()

	// s.SetTransformation(transformation.NewScaling(1, 0.5, 1)) // shrink along Y axis
	// s.SetTransformation(transformation.NewScaling(0.5, 1, 1)) // shrink along X axis
	// s.SetTransformation(transformation.NewRotationZ(math.Pi / 4).MultiplyMatrix(transformation.NewScaling(0.5, 1, 1))) // shrink and rotate
	s.SetTransformation(transformation.NewShearing(1, 0, 0, 0, 0, 0).MultiplyMatrix(transformation.NewScaling(0.5, 1, 1))) // shrink and skew

	for y := 0; y < CANVAS_PIXELS; y++ {
		worldY := HALF_WALL - PIXEL_SIZE*float64(y)

		for x := 0; x < CANVAS_PIXELS; x++ {
			worldX := -HALF_WALL + PIXEL_SIZE*float64(x)

			pos := tuple.NewPoint(worldX, worldY, WALL_Z)

			r := ray.NewRay(rayOrigin, pos.Sub(rayOrigin).Normalize())

			xs := s.Intersect(r)

			if xs.Hit() != nil {
				canvs.WriteAt(x, y, red)
			}
		}
	}

	fmt.Print(canvs.ToPPM())
}
