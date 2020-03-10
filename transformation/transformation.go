package transformation

import (
	"goray/matrix"
	"math"
)

func NewTranslation(x, y, z float64) *matrix.Matrix {
	m := matrix.NewIdentityMatrix4x4()

	m.Set(0, 3, x)
	m.Set(1, 3, y)
	m.Set(2, 3, z)

	return m
}

func NewScaling(x, y, z float64) *matrix.Matrix {
	m := matrix.NewIdentityMatrix4x4()

	m.Set(0, 0, x)
	m.Set(1, 1, y)
	m.Set(2, 2, z)

	return m
}

func NewRotationX(rad float64) *matrix.Matrix {
	m := matrix.NewIdentityMatrix4x4()

	m.Set(1, 1, math.Cos(rad))
	m.Set(2, 2, math.Cos(rad))
	m.Set(1, 2, -math.Sin(rad))
	m.Set(2, 1, math.Sin(rad))

	return m
}

func NewRotationY(rad float64) *matrix.Matrix {
	m := matrix.NewIdentityMatrix4x4()

	m.Set(0, 0, math.Cos(rad))
	m.Set(2, 2, math.Cos(rad))
	m.Set(0, 2, math.Sin(rad))
	m.Set(2, 0, -math.Sin(rad))

	return m
}

func NewRotationZ(rad float64) *matrix.Matrix {
	m := matrix.NewIdentityMatrix4x4()

	m.Set(0, 0, math.Cos(rad))
	m.Set(1, 1, math.Cos(rad))
	m.Set(0, 1, -math.Sin(rad))
	m.Set(1, 0, math.Sin(rad))

	return m
}

func NewShearing(xy, xz, yx, yz, zx, zy float64) *matrix.Matrix {
	m := matrix.NewIdentityMatrix4x4()

	m.Set(0, 1, xy)
	m.Set(0, 2, xz)
	m.Set(1, 0, yx)
	m.Set(1, 2, yz)
	m.Set(2, 0, zx)
	m.Set(2, 1, zy)

	return m
}
