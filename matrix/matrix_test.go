package matrix

import (
	"github.com/stretchr/testify/assert"
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
