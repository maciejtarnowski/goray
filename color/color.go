package color

import "goray/utils"

type Color struct {
	Red, Green, Blue float64
}

func (c *Color) Equals(other *Color) bool {
	return utils.Compare(c.Red, other.Red) &&
		utils.Compare(c.Green, other.Green) &&
		utils.Compare(c.Blue, other.Blue)
}

func (c *Color) Add(other *Color) *Color {
	return NewColor(c.Red+other.Red, c.Green+other.Green, c.Blue+other.Blue)
}

func (c *Color) Sub(other *Color) *Color {
	return NewColor(c.Red-other.Red, c.Green-other.Green, c.Blue-other.Blue)
}

func (c *Color) MultiplyScalar(factor float64) *Color {
	return NewColor(c.Red*factor, c.Green*factor, c.Blue*factor)
}

func (c *Color) Multiply(other *Color) *Color {
	return NewColor(c.Red*other.Red, c.Green*other.Green, c.Blue*other.Blue)
}

func NewColor(red, green, blue float64) *Color {
	return &Color{Red: red, Green: green, Blue: blue}
}
