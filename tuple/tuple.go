package tuple

import (
	"goray/utils"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t *Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t *Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t *Tuple) Equals(other *Tuple) bool {
	return utils.Compare(t.X, other.X) &&
		utils.Compare(t.Y, other.Y) &&
		utils.Compare(t.Z, other.Z) &&
		utils.Compare(t.W, other.W)
}

func (t *Tuple) Add(other *Tuple) *Tuple {
	return NewTuple(
		t.X+other.X,
		t.Y+other.Y,
		t.Z+other.Z,
		t.W+other.W,
	)
}

func (t *Tuple) Sub(other *Tuple) *Tuple {
	return NewTuple(
		t.X-other.X,
		t.Y-other.Y,
		t.Z-other.Z,
		t.W-other.W,
	)
}

func (t *Tuple) Negate() *Tuple {
	return NewTuple(
		-t.X,
		-t.Y,
		-t.Z,
		-t.W,
	)
}

func (t *Tuple) Multiply(factor float64) *Tuple {
	return NewTuple(
		t.X*factor,
		t.Y*factor,
		t.Z*factor,
		t.W*factor,
	)
}

func (t *Tuple) Divide(factor float64) *Tuple {
	return NewTuple(
		t.X/factor,
		t.Y/factor,
		t.Z/factor,
		t.W/factor,
	)
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W + t.W)
}

func (t *Tuple) Normalize() *Tuple {
	mag := t.Magnitude()
	return NewTuple(
		t.X/mag,
		t.Y/mag,
		t.Z/mag,
		t.W/mag,
	)
}

func (t *Tuple) Dot(other *Tuple) float64 {
	return t.X*other.X + t.Y*other.Y + t.Z*other.Z + t.W*other.W
}

func (t *Tuple) Cross(other *Tuple) *Tuple {
	if t.IsPoint() || other.IsPoint() {
		panic("cannot cross points")
	}

	return NewVector(
		t.Y*other.Z-t.Z*other.Y,
		t.Z*other.X-t.X*other.Z,
		t.X*other.Y-t.Y*other.X,
	)
}

func (t *Tuple) Reflect(normal *Tuple) *Tuple {
	return t.Sub(normal.Multiply(2).Multiply(t.Dot(normal)))
}

func NewTuple(x, y, z, w float64) *Tuple {
	return &Tuple{X: x, Y: y, Z: z, W: w}
}

func NewPoint(x, y, z float64) *Tuple {
	return NewTuple(x, y, z, 1.0)
}

func NewVector(x, y, z float64) *Tuple {
	return NewTuple(x, y, z, 0.0)
}
