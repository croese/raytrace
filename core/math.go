package core

import "math"

const epsilon = 0.00001

func compareFloat64s(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
