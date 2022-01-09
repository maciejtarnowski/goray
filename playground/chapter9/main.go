package main

import (
	"fmt"
	"goray/camera"
	"goray/color"
	"goray/light"
	"goray/material"
	"goray/ray"
	"goray/shape"
	tr "goray/transformation"
	"goray/tuple"
	"goray/world"
	"math"
)

func main() {
	bgMaterial := material.NewMaterial()
	bgMaterial.Color = color.NewColor(1, 0.9, 0.9)
	bgMaterial.Specular = 0

	w := world.NewWorld()
	w.Objects = append(w.Objects, getFloor(bgMaterial), getMiddleSphere(), getRightSphere(), getLeftSphere())
	w.Light = light.NewPointLight(tuple.NewPoint(-10, 10, -10), color.NewColor(1, 1, 1))

	c := camera.NewCamera(600, 300, math.Pi/3)
	c.Transform = tr.ViewTransform(tuple.NewPoint(0, 1.5, -5), tuple.NewPoint(0, 1, 0), tuple.NewVector(0, 1, 0))

	im := c.Render(w)

	fmt.Print(im.ToPPM())
}

func getFloor(m *material.Material) ray.Object {
	floor := shape.NewPlane()
	floor.SetTransformation(tr.NewScaling(10, 1, 10))
	floor.SetMaterial(m)

	return floor
}

func getMiddleSphere() ray.Object {
	middle := shape.NewSphere()
	middle.SetTransformation(tr.NewTranslation(-0.5, 1, 0.5))
	mat := material.NewMaterial()
	mat.Color = color.NewColor(0.1, 1, 0.5)
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	middle.SetMaterial(mat)

	return middle
}

func getRightSphere() ray.Object {
	right := shape.NewSphere()
	right.SetTransformation(tr.NewTranslation(1.5, 0.5, -0.5).MultiplyMatrix(tr.NewScaling(0.5, 0.5, 0.5)))
	mat := material.NewMaterial()
	mat.Color = color.NewColor(0.5, 1, 0.1)
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	right.SetMaterial(mat)

	return right
}

func getLeftSphere() ray.Object {
	left := shape.NewSphere()
	left.SetTransformation(tr.NewTranslation(-1.5, 0.33, -0.75).MultiplyMatrix(tr.NewScaling(0.33, 0.33, 0.33)))
	mat := material.NewMaterial()
	mat.Color = color.NewColor(1, 0.8, 0.1)
	mat.Diffuse = 0.7
	mat.Specular = 0.3
	left.SetMaterial(mat)

	return left
}
