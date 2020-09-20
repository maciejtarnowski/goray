package main

import (
	"fmt"
	"goray/canvas"
	"goray/color"
	"goray/light"
	"goray/ray"
	"goray/shape"
	"goray/transformation"
	"goray/tuple"
	"math"
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

	lightPosition := tuple.NewPoint(-10, 10, -10)
	lightColor := color.NewColor(1, 1, 1)
	l := light.NewPointLight(lightPosition, lightColor)

	s := shape.NewSphere()
	s.Material.Color = color.NewColor(1, 0.2, 1)

	// s.Transformation = transformation.NewScaling(1, 0.5, 1) // shrink along Y axis
	// s.Transformation = transformation.NewScaling(0.5, 1, 1) // shrink along X axis
	s.Transformation = transformation.NewRotationZ(math.Pi / 4).MultiplyMatrix(transformation.NewScaling(0.5, 1, 1)) // shrink and rotate
	// s.Transformation = transformation.NewShearing(1, 0, 0, 0, 0, 0).MultiplyMatrix(transformation.NewScaling(0.5, 1, 1)) // shrink and skew

	for y := 0; y < CANVAS_PIXELS; y++ {
		worldY := HALF_WALL - PIXEL_SIZE*float64(y)

		for x := 0; x < CANVAS_PIXELS; x++ {
			worldX := -HALF_WALL + PIXEL_SIZE*float64(x)

			pos := tuple.NewPoint(worldX, worldY, WALL_Z)

			r := ray.NewRay(rayOrigin, pos.Sub(rayOrigin).Normalize())

			xs := s.Intersect(r)

			if xs.Hit() != nil {
				hit := xs.Hit()

				p := r.Position(hit.T)
				normal := hit.Object.(*shape.Sphere).NormalAt(p)
				eye := r.Direction.Negate()

				c := hit.Object.(*shape.Sphere).Material.Lighting(l, p, eye, normal)

				canvs.WriteAt(x, y, c)
			}
		}
	}

	fmt.Print(canvs.ToPPM())
}
