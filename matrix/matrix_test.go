package matrix

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goray/tuple"
	"testing"
)

func TestConstructing4x4Matrix(t *testing.T) {
	matrix := NewMatrix(4, 4, 1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5)

	assert.Equal(t, float64(1), matrix.At(0, 0))
	assert.Equal(t, float64(4), matrix.At(0, 3))
	assert.Equal(t, 5.5, matrix.At(1, 0))
	assert.Equal(t, 7.5, matrix.At(1, 2))
	assert.Equal(t, float64(11), matrix.At(2, 2))
	assert.Equal(t, 13.5, matrix.At(3, 0))
	assert.Equal(t, 15.5, matrix.At(3, 2))
	assert.Equal(t, 16.5, matrix.At(3, 3))
}

func TestConstructing2x2Matrix(t *testing.T) {
	matrix := NewMatrix(2, 2, -3, 5, 1, -2)

	assert.Equal(t, float64(-3), matrix.At(0, 0))
	assert.Equal(t, float64(5), matrix.At(0, 1))
	assert.Equal(t, float64(1), matrix.At(1, 0))
	assert.Equal(t, float64(-2), matrix.At(1, 1))
}

func TestConstructing3x3Matrix(t *testing.T) {
	matrix := NewMatrix(3, 3, -3, 5, 0, 1, -2, -7, 0, 1, 1)

	assert.Equal(t, float64(-3), matrix.At(0, 0))
	assert.Equal(t, float64(-2), matrix.At(1, 1))
	assert.Equal(t, float64(1), matrix.At(2, 2))
}

func TestComparingIdenticalMatrices(t *testing.T) {
	matrix1 := NewMatrix(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	matrix2 := NewMatrix(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)

	assert.True(t, matrix1.Equals(matrix2))
	assert.True(t, matrix2.Equals(matrix1))
}

func TestComparingDifferentMatrices(t *testing.T) {
	matrix1 := NewMatrix(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	matrix2 := NewMatrix(4, 4, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1)

	assert.False(t, matrix1.Equals(matrix2))
	assert.False(t, matrix2.Equals(matrix1))
}

func TestMultiplyingTwoMatrices(t *testing.T) {
	matrix1 := NewMatrix(4, 4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2)
	matrix2 := NewMatrix(4, 4, -2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8)

	expected := NewMatrix(4, 4, 20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42)

	assert.True(t, matrix1.MultiplyMatrix(matrix2).Equals(expected))
}

func TestMultiplyingMatrixByTuple(t *testing.T) {
	matrix := NewMatrix(4, 4, 1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1)
	tuple1 := tuple.NewTuple(1, 2, 3, 1)

	expected := tuple.NewTuple(18, 24, 33, 1)

	assert.True(t, matrix.MultiplyTuple(tuple1).Equals(expected))
}

func TestMultiplyingMatrixByIdentityMatrix(t *testing.T) {
	matrix := NewMatrix(4, 4, 0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32)
	identity := NewIdentityMatrix4x4()

	expected := NewMatrix(4, 4, 0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32)

	assert.True(t, matrix.MultiplyMatrix(identity).Equals(expected))
}

func TestMultiplyingIdentityMatrixByTuple(t *testing.T) {
	tuple1 := tuple.NewTuple(1, 2, 3, 4)
	identity := NewIdentityMatrix4x4()

	expected := tuple.NewTuple(1, 2, 3, 4)

	assert.True(t, identity.MultiplyTuple(tuple1).Equals(expected))
}

func TestTransposingMatrix(t *testing.T) {
	matrix := NewMatrix(4, 4, 0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8)

	expected := NewMatrix(4, 4, 0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8)

	assert.True(t, matrix.Transpose().Equals(expected))
}

func TestTransposingIdentityMatrix(t *testing.T) {
	matrix := NewIdentityMatrix4x4()

	expected := NewIdentityMatrix4x4()

	assert.True(t, matrix.Transpose().Equals(expected))
}

func TestDeterminantOf2x2Matrix(t *testing.T) {
	matrix := NewMatrix(2, 2, 1, 5, -3, 2)

	assert.InDelta(t, 17, matrix.Determinant(), 0.00001)
}

func TestSubmatrixOf3x3Matrix(t *testing.T) {
	matrix := NewMatrix(3, 3, 1, 5, 0, -3, 2, 7, 0, 6, -3)

	expected := NewMatrix(2, 2, -3, 2, 0, 6)

	assert.True(t, matrix.Submatrix(0, 2).Equals(expected))
}

func TestSubmatrixOf4x4Matrix(t *testing.T) {
	matrix := NewMatrix(4, 4, -6, 1, 1, 6, -8, 5, 8, 6, -1, 0, 8, 2, -7, 1, -1, 1)

	expected := NewMatrix(3, 3, -6, 1, 6, -8, 8, 6, -7, -1, 1)

	assert.True(t, matrix.Submatrix(2, 1).Equals(expected))
}

func TestMinorOf3x3Matrix(t *testing.T) {
	matrix1 := NewMatrix(3, 3, 3, 5, 0, 2, -1, -7, 6, -1, 5)
	matrix2 := matrix1.Submatrix(1, 0)

	require.InDelta(t, 25, matrix2.Determinant(), 0.00001)
	assert.InDelta(t, 25, matrix1.Minor(1, 0), 0.00001)
}

func TestCofactorOf3x3Matrix(t *testing.T) {
	matrix := NewMatrix(3, 3, 3, 5, 0, 2, -1, -7, 6, -1, 5)

	assert.InDelta(t, -12, matrix.Minor(0, 0), 0.00001)
	assert.InDelta(t, -12, matrix.Cofactor(0, 0), 0.00001)
	assert.InDelta(t, 25, matrix.Minor(1, 0), 0.00001)
	assert.InDelta(t, -25, matrix.Cofactor(1, 0), 0.00001)
}

func TestDeterminantOf3x3Matrix(t *testing.T) {
	matrix := NewMatrix(3, 3, 1, 2, 6, -5, 8, -4, 2, 6, 4)

	assert.InDelta(t, 56, matrix.Cofactor(0, 0), 0.00001)
	assert.InDelta(t, 12, matrix.Cofactor(0, 1), 0.00001)
	assert.InDelta(t, -46, matrix.Cofactor(0, 2), 0.00001)
	assert.InDelta(t, -196, matrix.Determinant(), 0.00001)
}

func TestDeterminantOf4x4Matrix(t *testing.T) {
	matrix := NewMatrix(4, 4, -2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9)

	assert.InDelta(t, 690, matrix.Cofactor(0, 0), 0.00001)
	assert.InDelta(t, 447, matrix.Cofactor(0, 1), 0.00001)
	assert.InDelta(t, 210, matrix.Cofactor(0, 2), 0.00001)
	assert.InDelta(t, 51, matrix.Cofactor(0, 3), 0.00001)
	assert.InDelta(t, -4071, matrix.Determinant(), 0.00001)
}

func TestInvertibleMatrixIsInvertible(t *testing.T) {
	matrix := NewMatrix(4, 4, 6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6)

	assert.InDelta(t, -2120, matrix.Determinant(), 0.00001)
	assert.True(t, matrix.IsInvertible())
}

func TestNonInvertibleMatrixIsNotInvertible(t *testing.T) {
	matrix := NewMatrix(4, 4, -4, 2, -2, 3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0)

	assert.InDelta(t, 0, matrix.Determinant(), 0.00001)
	assert.False(t, matrix.IsInvertible())
}

func TestInvertingMatrix(t *testing.T) {
	matrix := NewMatrix(4, 4, -5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4)

	expected := NewMatrix(4, 4, 0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068, -0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639)

	inverted := matrix.Invert()

	assert.InDelta(t, 532, matrix.Determinant(), 0.00001)
	assert.InDelta(t, -160, matrix.Cofactor(2, 3), 0.00001)
	assert.InDelta(t, float64(-160)/float64(532), inverted.At(3, 2), 0.00001)
	assert.InDelta(t, 105, matrix.Cofactor(3, 2), 0.00001)
	assert.InDelta(t, float64(105)/float64(532), inverted.At(2, 3), 0.00001)
	assert.True(t, inverted.Equals(expected))
}
