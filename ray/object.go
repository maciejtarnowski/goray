package ray

import (
	"goray/material"
	"goray/tuple"
)

type Object interface {
	Intersect(r *Ray) Intersections
	NormalAt(point *tuple.Tuple) *tuple.Tuple

	GetMaterial() *material.Material
}
