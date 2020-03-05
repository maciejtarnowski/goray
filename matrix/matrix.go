package matrix

import (
	"goray/tuple"
	"goray/utils"
)

func calcIndex(cols, x, y int) int {
	return x*cols + y
}

type Matrix struct {
	Rows     int
	Cols     int
	Elements []float64
}

func (m *Matrix) At(row, col int) float64 {
	return m.Elements[calcIndex(m.Cols, row, col)]
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
			productElements[calcIndex(m.Cols, row, col)] = m.At(row, 0)*other.At(0, col) +
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
			productElements[calcIndex(m.Cols, col, row)] = m.At(row, col)
		}
	}

	return NewMatrix(m.Rows, m.Cols, productElements...)
}

func (m *Matrix) Determinant() float64 {
	det := float64(0)

	if m.Cols == 2 && m.Rows == 2 {
		det = m.At(0, 0)*m.At(1, 1) - m.At(0, 1)*m.At(1, 0)
	}
	if m.Cols > 2 || m.Rows > 2 {
		for col := 0; col < m.Cols; col++ {
			det += m.At(0, col) * m.Cofactor(0, col)
		}
	}
	return det
}

func (m *Matrix) Submatrix(excludedRow, excludedCol int) *Matrix {
	productElements := make([]float64, (m.Cols-1)*(m.Rows-1))
	for row := 0; row < m.Rows; row++ {
		if row == excludedRow {
			continue
		}

		for col := 0; col < m.Cols; col++ {
			if col == excludedCol {
				continue
			}

			targetRow := row
			if row > excludedRow {
				targetRow = row - 1
			}
			targetCol := col
			if col > excludedCol {
				targetCol = col - 1
			}

			productElements[calcIndex(m.Cols-1, targetRow, targetCol)] = m.At(row, col)
		}
	}

	return NewMatrix(m.Rows-1, m.Cols-1, productElements...)
}

func (m *Matrix) Minor(row, column int) float64 {
	return m.Submatrix(row, column).Determinant()
}

func (m *Matrix) Cofactor(row, column int) float64 {
	minor := m.Minor(row, column)

	if (row+column)%2 == 0 {
		return minor
	}
	return -minor
}

func (m *Matrix) IsInvertible() bool {
	return !utils.Compare(m.Determinant(), 0)
}

func (m *Matrix) Invert() *Matrix {
	if !m.IsInvertible() {
		panic("trying to invert non-invertible matrix")
	}

	inverseElements := make([]float64, m.Cols*m.Rows)
	det := m.Determinant()

	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			c := m.Cofactor(row, col)

			inverseElements[calcIndex(m.Cols, col, row)] = c / det
		}
	}

	return NewMatrix(m.Rows, m.Cols, inverseElements...)
}

func NewMatrix(rows, cols int, elements ...float64) *Matrix {
	return &Matrix{Rows: rows, Cols: cols, Elements: elements}
}

func NewIdentityMatrix4x4() *Matrix {
	return NewMatrix(4, 4, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1)
}
