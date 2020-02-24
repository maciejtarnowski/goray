package canvas

import (
	"fmt"
	"goray/color"
	"strings"
)

type Canvas struct {
	Width  int
	Height int
	Pixels [][]*color.Color
}

func NewCanvas(width, height int) *Canvas {
	c := Canvas{Width: width, Height: height}
	c.FillWith(color.NewColor(0, 0, 0))

	return &c
}

func (c *Canvas) FillWith(col *color.Color) {
	c.Pixels = make([][]*color.Color, c.Height)
	for i := range c.Pixels {
		c.Pixels[i] = make([]*color.Color, c.Width)

		for j := range c.Pixels[i] {
			c.Pixels[i][j] = col
		}
	}
}

func (c *Canvas) PixelAt(x, y int) *color.Color {
	return c.Pixels[y][x]
}

func (c *Canvas) WriteAt(x, y int, color *color.Color) {
	c.Pixels[y][x] = color.Clone()
}

func (c *Canvas) ToPPM() string {
	var out strings.Builder

	out.WriteString("P3\n")
	out.WriteString(fmt.Sprintf("%d %d\n", c.Width, c.Height))
	out.WriteString("255\n")

	var line strings.Builder
	for y := range c.Pixels {
		for x := range c.Pixels[y] {
			pixel := c.PixelAt(x, y)
			rgb := pixel.ToRGB()

			for i := range rgb {
				colorStr := fmt.Sprintf("%d", rgb[i])
				if line.Len()+1+len(colorStr) > 69 {
					out.WriteString(line.String())
					out.WriteString("\n")
					line.Reset()
				}
				if line.Len() != 0 {
					line.WriteByte(' ')
				}
				line.WriteString(colorStr)
			}
		}
		out.WriteString(line.String())
		out.WriteString("\n")
		line.Reset()
	}

	return out.String()
}
