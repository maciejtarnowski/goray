package material

import (
	"goray/color"
	"goray/light"
	"goray/tuple"
	"math"
)

type Material struct {
	Color     *color.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() *Material {
	return &Material{Color: color.NewColor(1, 1, 1), Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200.0}
}

func (m *Material) Lighting(l *light.Light, point *tuple.Tuple, eyeV *tuple.Tuple, normalV *tuple.Tuple) *color.Color {
	effectiveColor := m.Color.Multiply(l.Intensity)

	lightV := l.Position.Sub(point).Normalize()

	ambient := effectiveColor.MultiplyScalar(m.Ambient)

	lightDotNormal := lightV.Dot(normalV)

	var diffuse *color.Color
	var specular *color.Color
	if lightDotNormal < 0 {
		diffuse = color.NewColor(0, 0, 0)
		specular = color.NewColor(0, 0, 0)
	} else {
		diffuse = effectiveColor.MultiplyScalar(m.Diffuse).MultiplyScalar(lightDotNormal)

		reflectV := lightV.Negate().Reflect(normalV)
		reflectDotEye := reflectV.Dot(eyeV)

		if reflectDotEye <= 0 {
			specular = color.NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.MultiplyScalar(m.Specular).MultiplyScalar(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
