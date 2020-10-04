package transformation

import (
	"github.com/stretchr/testify/assert"
	"goray/matrix"
	"goray/tuple"
	"testing"
)

func TestTransformationMatrixForDefaultOrientation(t *testing.T) {
	from := tuple.NewPoint(0, 0, 0)
	to := tuple.NewPoint(0, 0, -1)
	up := tuple.NewVector(0, 1, 0)

	transform := ViewTransform(from, to, up)

	assert.True(t, matrix.NewIdentityMatrix4x4().Equals(transform))
}

func TestViewTransformationMatrixLookingInPositiveZDirection(t *testing.T) {
	from := tuple.NewPoint(0, 0, 0)
	to := tuple.NewPoint(0, 0, 1)
	up := tuple.NewVector(0, 1, 0)

	transform := ViewTransform(from, to, up)

	assert.True(t, NewScaling(-1, 1, -1).Equals(transform))
}

func TestViewTransformationMovesTheWorld(t *testing.T) {
	from := tuple.NewPoint(0, 0, 8)
	to := tuple.NewPoint(0, 0, 0)
	up := tuple.NewVector(0, 1, 0)

	transform := ViewTransform(from, to, up)

	assert.True(t, NewTranslation(0, 0, -8).Equals(transform))
}

func TestArbitraryViewTransformation(t *testing.T) {
	from := tuple.NewPoint(1, 3, 2)
	to := tuple.NewPoint(4, -2, 8)
	up := tuple.NewVector(1, 1, 0)

	transform := ViewTransform(from, to, up)

	expected := matrix.NewMatrix(4, 4, -0.50709, 0.50709, 0.67612, -2.36643, 0.76772, 0.60609, 0.12122, -2.82843, -0.35857, 0.59761, -0.71714, 0.0, 0.0, 0.0, 0.0, 1.0)

	assert.True(t, expected.Equals(transform))
}
