package matrix

import (
	"goray/tuple"
	"goray/utils"
)

type Matrix struct {
	Rows     int
	Cols     int
	Elements []float64
}

func (m *Matrix) At(row, col int) float64 {
	return m.Elements[row*m.Cols+col]
}

func (m *Matrix) Equals(other *Matrix) bool {
	if m.Rows != other.Rows || m.Cols != other.Cols || len(m.Elements) != len(other.Elements) {
		return false
	}

	for i, val := range m.Elements {
		if !utils.Compare(val, other.Elements[i]) {
			return false
		}
	}

	return true
}

func (m *Matrix) MultiplyMatrix(other *Matrix) *Matrix {
	productElements := make([]float64, 16)
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			productElements[row*m.Cols+col] = m.At(row, 0)*other.At(0, col) +
				m.At(row, 1)*other.At(1, col) +
				m.At(row, 2)*other.At(2, col) +
				m.At(row, 3)*other.At(3, col)
		}
	}

	return NewMatrix(4, 4, productElements...)
}

func (m *Matrix) MultiplyTuple(other *tuple.Tuple) *tuple.Tuple {
	productElements := make([]float64, 16)
	for row := 0; row < 4; row++ {
		productElements[row] = m.At(row, 0)*other.X +
			m.At(row, 1)*other.Y +
			m.At(row, 2)*other.Z +
			m.At(row, 3)*other.W
	}

	return tuple.NewTuple(productElements[0], productElements[1], productElements[2], productElements[3])
}

func (m *Matrix) Transpose() *Matrix {
	productElements := make([]float64, 16)
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			productElements[col*m.Cols+row] = m.At(row, col)
		}
	}

	return NewMatrix(m.Rows, m.Cols, productElements...)
}

func NewMatrix(rows, cols int, elements ...float64) *Matrix {
	return &Matrix{Rows: rows, Cols: cols, Elements: elements}
}

func NewIdentityMatrix4x4() *Matrix {
	return NewMatrix(4, 4, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)
}
