package utils

import "math"

const EPSILON = 0.00001

func Compare(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}
