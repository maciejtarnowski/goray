package tuple

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestPoint(t *testing.T) {
	tuple := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}

	require.InDelta(t, 4.3, tuple.X, 0.00001)
	require.InDelta(t, -4.2, tuple.Y, 0.00001)
	require.InDelta(t, 3.1, tuple.Z, 0.00001)
	require.True(t, tuple.IsPoint())
	require.False(t, tuple.IsVector())
}

func TestVector(t *testing.T) {
	tuple := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0.0}

	require.InDelta(t, 4.3, tuple.X, 0.00001)
	require.InDelta(t, -4.2, tuple.Y, 0.00001)
	require.InDelta(t, 3.1, tuple.Z, 0.00001)
	require.False(t, tuple.IsPoint())
	require.True(t, tuple.IsVector())
}

func TestPointFactory(t *testing.T) {
	point := NewPoint(4, -4, 3)

	expected := Tuple{X: 4, Y: -4, Z: 3, W: 1}

	require.True(t, point.Equals(&expected))
}

func TestVectorFactory(t *testing.T) {
	vector := NewVector(4, -4, 3)

	expected := Tuple{X: 4, Y: -4, Z: 3, W: 0}

	require.True(t, vector.Equals(&expected))
}

func TestAddingTuples(t *testing.T) {
	point := NewPoint(3, -2, 5)
	vector := NewVector(-2, 3, 1)

	expected := NewPoint(1, 1, 6)

	require.True(t, point.Add(vector).Equals(expected))
}

func TestSubtractingPoints(t *testing.T) {
	point1 := NewPoint(3, 2, 1)
	point2 := NewPoint(5, 6, 7)

	expected := NewVector(-2, -4, -6)

	require.True(t, point1.Sub(point2).Equals(expected))
}

func TestSubtractingVectorFromPoint(t *testing.T) {
	point := NewPoint(3, 2, 1)
	vector := NewVector(5, 6, 7)

	expected := NewPoint(-2, -4, -6)

	require.True(t, point.Sub(vector).Equals(expected))
}

func TestSubtractingVectors(t *testing.T) {
	vector1 := NewVector(3, 2, 1)
	vector2 := NewVector(5, 6, 7)

	expected := NewVector(-2, -4, -6)

	require.True(t, vector1.Sub(vector2).Equals(expected))
}

func TestSubtractingVectorFromZeroVector(t *testing.T) {
	zeroVector := NewVector(0, 0, 0)
	vector := NewVector(1, -2, 3)

	expected := NewVector(-1, 2, -3)

	require.True(t, zeroVector.Sub(vector).Equals(expected))
}

func TestNegatingTuple(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)

	expected := NewTuple(-1, 2, -3, 4)

	require.True(t, tuple.Negate().Equals(expected))
}

func TestMultiplyingTupleByScalar(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)

	expected := NewTuple(3.5, -7, 10.5, -14)

	require.True(t, tuple.Multiply(3.5).Equals(expected))
}

func TestMultiplyingTupleByScalarFraction(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)

	expected := NewTuple(0.5, -1, 1.5, -2)

	require.True(t, tuple.Multiply(0.5).Equals(expected))
}

func TestDividingTupleByScalar(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)

	expected := NewTuple(0.5, -1, 1.5, -2)

	require.True(t, tuple.Divide(2).Equals(expected))
}

func TestVectorMagnitude(t *testing.T) {
	tests := []struct {
		vector            *Tuple
		expectedMagnitude float64
	}{
		{
			NewVector(1, 0, 0),
			1,
		},
		{
			NewVector(0, 1, 0),
			1,
		},
		{
			NewVector(0, 0, 1),
			1,
		},
		{
			NewVector(1, 2, 3),
			math.Sqrt(14),
		},
		{
			NewVector(-1, -2, -3),
			math.Sqrt(14),
		},
	}

	for _, tt := range tests {
		assert.InDelta(t, tt.expectedMagnitude, tt.vector.Magnitude(), 0.00001)
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		vector           *Tuple
		normalizedVector *Tuple
	}{
		{
			NewVector(4, 0, 0),
			NewVector(1, 0, 0),
		},
		{
			NewVector(1, 2, 3),
			NewVector(0.26726, 0.53452, 0.80178),
		},
	}

	for _, tt := range tests {
		assert.True(t, tt.vector.Normalize().Equals(tt.normalizedVector))
		assert.InDelta(t, 1, tt.vector.Normalize().Magnitude(), 0.00001)
	}
}

func TestDotProduct(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	require.InDelta(t, 20, vector1.Dot(vector2), 0.00001)
}

func TestCrossProduct(t *testing.T) {
	vector1 := NewVector(1, 2, 3)
	vector2 := NewVector(2, 3, 4)

	expected1 := NewVector(-1, 2, -1)
	expected2 := expected1.Negate()

	require.True(t, vector1.Cross(vector2).Equals(expected1))
	require.True(t, vector2.Cross(vector1).Equals(expected2))
}
