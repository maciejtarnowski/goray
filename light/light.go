package light

import (
	"goray/color"
	"goray/tuple"
)

type Light struct {
	Position  *tuple.Tuple
	Intensity *color.Color
}

func NewPointLight(position *tuple.Tuple, intensity *color.Color) *Light {
	return &Light{Position: position, Intensity: intensity}
}
