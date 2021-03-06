package world

import (
	"goray/color"
	"goray/light"
	"goray/material"
	"goray/ray"
	"goray/shape"
	"goray/transformation"
	"goray/tuple"
	"sort"
)

type World struct {
	Light   *light.Light
	Objects []ray.Object
}

func NewWorld() *World {
	return &World{}
}

func NewDefaultWorld() *World {
	s1 := shape.NewSphere()
	m := material.NewMaterial()
	m.Color = color.NewColor(0.8, 1.0, 0.6)
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.SetMaterial(m)

	s2 := shape.NewSphere()
	s2.SetTransformation(transformation.NewScaling(0.5, 0.5, 0.5))

	return &World{
		Light:   light.NewPointLight(tuple.NewPoint(-10, 10, -10), color.NewColor(1, 1, 1)),
		Objects: []ray.Object{s1, s2},
	}
}

func (w *World) Intersect(r *ray.Ray) *ray.Intersections {
	var worldXs []*ray.Intersection
	for _, obj := range w.Objects {
		objXs := obj.Intersect(r)
		for _, objX := range objXs.GetAll() {
			worldXs = append(worldXs, objX)
		}
	}

	xs := ray.NewIntersections(worldXs...)
	sort.Sort(xs)

	return xs
}

func (w *World) ShadeHit(comps *ray.Computation) *color.Color {
	return comps.Object.GetMaterial().Lighting(w.Light, comps.OverPoint, comps.EyeV, comps.NormalV, w.IsShadowed(comps.OverPoint))
}

func (w *World) ColorAt(r *ray.Ray) *color.Color {
	xs := w.Intersect(r)

	if xs.Hit() == nil {
		return color.NewColor(0, 0, 0)
	}

	comps := xs.Hit().PrepareComputations(r)

	return w.ShadeHit(comps)
}

func (w *World) IsShadowed(p *tuple.Tuple) bool {
	v := w.Light.Position.Sub(p)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := ray.NewRay(p, direction)
	xs := w.Intersect(r)

	hit := xs.Hit()

	return hit != nil && hit.T < distance
}
