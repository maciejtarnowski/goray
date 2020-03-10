package transformation

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goray/tuple"
	"goray/utils"
	"math"
	"testing"
)

func TestMultiplyingByTranslationMatrix(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	point := tuple.NewPoint(-3, 4, 5)

	expected := tuple.NewPoint(2, 1, 7)

	assert.True(t, transform.MultiplyTuple(point).Equals(expected))
}

func TestMultiplyingByInverseOfTranslationMatrix(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	inverse := transform.Invert()
	point := tuple.NewPoint(-3, 4, 5)

	expected := tuple.NewPoint(-8, 7, 3)

	assert.True(t, inverse.MultiplyTuple(point).Equals(expected))
}

func TestMultiplyingVectorDoesNotChangeIt(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	vector := tuple.NewVector(-3, 4, 5)

	assert.True(t, transform.MultiplyTuple(vector).Equals(vector))
}

func TestScalingPoint(t *testing.T) {
	scaling := NewScaling(2, 3, 4)
	point := tuple.NewPoint(-4, 6, 8)

	expected := tuple.NewPoint(-8, 18, 32)

	assert.True(t, scaling.MultiplyTuple(point).Equals(expected))
}

func TestScalingVector(t *testing.T) {
	scaling := NewScaling(2, 3, 4)
	vector := tuple.NewVector(-4, 6, 8)

	expected := tuple.NewVector(-8, 18, 32)

	assert.True(t, scaling.MultiplyTuple(vector).Equals(expected))
}

func TestMultiplyByInverseOfScalingMatrix(t *testing.T) {
	scaling := NewScaling(2, 3, 4)
	inverse := scaling.Invert()
	vector := tuple.NewVector(-4, 6, 8)

	expected := tuple.NewVector(-2, 2, 2)

	assert.True(t, inverse.MultiplyTuple(vector).Equals(expected))
}

func TestReflectionAsScalingByNegativeValue(t *testing.T) {
	scaling := NewScaling(-1, 1, 1)
	point := tuple.NewPoint(2, 3, 4)

	expected := tuple.NewPoint(-2, 3, 4)

	assert.True(t, scaling.MultiplyTuple(point).Equals(expected))
}

func TestRotatingPointAroundXAxis(t *testing.T) {
	point := tuple.NewPoint(0, 1, 0)
	halfQuarter := NewRotationX(utils.Degrees2Radians(45))
	fullQuarter := NewRotationX(utils.Degrees2Radians(90))

	expected1 := tuple.NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	expected2 := tuple.NewPoint(0, 0, 1)

	assert.True(t, halfQuarter.MultiplyTuple(point).Equals(expected1))
	assert.True(t, fullQuarter.MultiplyTuple(point).Equals(expected2))
}

func TestRotatingPointAroundXAxisWithInverse(t *testing.T) {
	point := tuple.NewPoint(0, 1, 0)
	halfQuarter := NewRotationX(utils.Degrees2Radians(45))
	inverse := halfQuarter.Invert()

	expected := tuple.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	assert.True(t, inverse.MultiplyTuple(point).Equals(expected))
}

func TestRotatingPointAroundYAxis(t *testing.T) {
	point := tuple.NewPoint(0, 0, 1)
	halfQuarter := NewRotationY(utils.Degrees2Radians(45))
	fullQuarter := NewRotationY(utils.Degrees2Radians(90))

	expected1 := tuple.NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)
	expected2 := tuple.NewPoint(1, 0, 0)

	assert.True(t, halfQuarter.MultiplyTuple(point).Equals(expected1))
	assert.True(t, fullQuarter.MultiplyTuple(point).Equals(expected2))
}

func TestRotatingPointAroundZAxis(t *testing.T) {
	point := tuple.NewPoint(0, 1, 0)
	halfQuarter := NewRotationZ(utils.Degrees2Radians(45))
	fullQuarter := NewRotationZ(utils.Degrees2Radians(90))

	expected1 := tuple.NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	expected2 := tuple.NewPoint(-1, 0, 0)

	assert.True(t, halfQuarter.MultiplyTuple(point).Equals(expected1))
	assert.True(t, fullQuarter.MultiplyTuple(point).Equals(expected2))
}

func TestShearingMovesXInProportionToY(t *testing.T) {
	shearing := NewShearing(1, 0, 0, 0, 0, 0)
	point := tuple.NewPoint(2, 3, 4)

	expected := tuple.NewPoint(5, 3, 4)

	assert.True(t, shearing.MultiplyTuple(point).Equals(expected))
}

func TestShearingMovesXInProportionToZ(t *testing.T) {
	shearing := NewShearing(0, 1, 0, 0, 0, 0)
	point := tuple.NewPoint(2, 3, 4)

	expected := tuple.NewPoint(6, 3, 4)

	assert.True(t, shearing.MultiplyTuple(point).Equals(expected))
}

func TestShearingMovesYInProportionToX(t *testing.T) {
	shearing := NewShearing(0, 0, 1, 0, 0, 0)
	point := tuple.NewPoint(2, 3, 4)

	expected := tuple.NewPoint(2, 5, 4)

	assert.True(t, shearing.MultiplyTuple(point).Equals(expected))
}

func TestShearingMovesYInProportionToZ(t *testing.T) {
	shearing := NewShearing(0, 0, 0, 1, 0, 0)
	point := tuple.NewPoint(2, 3, 4)

	expected := tuple.NewPoint(2, 7, 4)

	assert.True(t, shearing.MultiplyTuple(point).Equals(expected))
}

func TestShearingMovesZInProportionToX(t *testing.T) {
	shearing := NewShearing(0, 0, 0, 0, 1, 0)
	point := tuple.NewPoint(2, 3, 4)

	expected := tuple.NewPoint(2, 3, 6)

	assert.True(t, shearing.MultiplyTuple(point).Equals(expected))
}

func TestShearingMovesZInProportionToY(t *testing.T) {
	shearing := NewShearing(0, 0, 0, 0, 0, 1)
	point := tuple.NewPoint(2, 3, 4)

	expected := tuple.NewPoint(2, 3, 7)

	assert.True(t, shearing.MultiplyTuple(point).Equals(expected))
}

func TestApplyingTransformationsInSequence(t *testing.T) {
	p := tuple.NewPoint(1, 0, 1)

	a := NewRotationX(utils.Degrees2Radians(90))
	b := NewScaling(5, 5, 5)
	c := NewTranslation(10, 5, 7)

	p2 := a.MultiplyTuple(p)
	require.True(t, p2.Equals(tuple.NewPoint(1, -1, 0)))

	p3 := b.MultiplyTuple(p2)
	require.True(t, p3.Equals(tuple.NewPoint(5, -5, 0)))

	p4 := c.MultiplyTuple(p3)
	require.True(t, p4.Equals(tuple.NewPoint(15, 0, 7)))
}

func TestApplyingTransformationsInChain(t *testing.T) {
	p := tuple.NewPoint(1, 0, 1)

	a := NewRotationX(utils.Degrees2Radians(90))
	b := NewScaling(5, 5, 5)
	c := NewTranslation(10, 5, 7)

	expected := tuple.NewPoint(15, 0, 7)

	assert.True(t, c.MultiplyMatrix(b).MultiplyMatrix(a).MultiplyTuple(p).Equals(expected))
}
