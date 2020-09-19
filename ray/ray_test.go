package ray

import (
	"github.com/stretchr/testify/assert"
	"goray/transformation"
	"goray/tuple"
	"testing"
)

func TestCreatingAndQueryingRay(t *testing.T) {
	origin := tuple.NewPoint(1, 2, 3)
	dir := tuple.NewVector(4, 5, 6)

	r := NewRay(origin, dir)

	assert.Equal(t, origin, r.Origin)
	assert.Equal(t, dir, r.Direction)
}

func TestComputingPointFromDistance(t *testing.T) {
	r := NewRay(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0))

	assert.True(t, r.Position(0).Equals(tuple.NewPoint(2, 3, 4)))
	assert.True(t, r.Position(1).Equals(tuple.NewPoint(3, 3, 4)))
	assert.True(t, r.Position(-1).Equals(tuple.NewPoint(1, 3, 4)))
	assert.True(t, r.Position(2.5).Equals(tuple.NewPoint(4.5, 3, 4)))
}

func TestTranslatingRay(t *testing.T) {
	r := NewRay(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	m := transformation.NewTranslation(3, 4, 5)

	r2 := r.Transform(m)

	assert.Equal(t, tuple.NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, tuple.NewVector(0, 1, 0), r2.Direction)
}

func TestScalingRay(t *testing.T) {
	r := NewRay(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	m := transformation.NewScaling(2, 3, 4)

	r2 := r.Transform(m)

	assert.Equal(t, tuple.NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, tuple.NewVector(0, 3, 0), r2.Direction)
}
