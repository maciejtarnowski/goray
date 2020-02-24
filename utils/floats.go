package utils

import "math"

const EPSILON = 0.00001

func Compare(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func Float64ToUint(value float64, maxInt int) int {
	product := int(math.Round(value * float64(maxInt)))

	if product > maxInt {
		return maxInt
	}
	if product < 0 {
		return 0
	}

	return product
}
