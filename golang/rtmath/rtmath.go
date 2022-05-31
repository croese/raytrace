package rtmath

import "math"

const EPSILON = 0.00001

func FloatEqualEpsilon(a, b float64) bool {
	return FloatEqual(a, b, EPSILON)
}

func FloatEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}
