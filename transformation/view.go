package transformation

import (
	"goray/matrix"
	"goray/tuple"
)

func ViewTransform(from *tuple.Tuple, to *tuple.Tuple, up *tuple.Tuple) *matrix.Matrix {
	forward := to.Sub(from).Normalize()
	upn := up.Normalize()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	orientation := matrix.NewMatrix(4, 4, left.X, left.Y, left.Z, 0, trueUp.X, trueUp.Y, trueUp.Z, 0, -forward.X, -forward.Y, -forward.Z, 0, 0, 0, 0, 1)

	return orientation.MultiplyMatrix(NewTranslation(-from.X, -from.Y, -from.Z))
}
