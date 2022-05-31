package rtmath_test

import (
	"testing"

	"github.com/croese/raytrace/rtmath"
)

func TestFloatEqualEpsilon(t *testing.T) {
	t.Run("Equality", func(t *testing.T) {
		a, b := 4.3, 4+0.3

		if !rtmath.FloatEqualEpsilon(a, b) {
			t.Errorf("expected %f to equal %f", a, b)
		}
	})

	t.Run("Inequality", func(t *testing.T) {
		a, b := 4.12345, 4+0.12347

		if rtmath.FloatEqualEpsilon(a, b) {
			t.Errorf("expected %f to not equal %f", a, b)
		}
	})
}
