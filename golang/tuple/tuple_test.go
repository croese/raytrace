package tuple_test

import (
	"math"
	"testing"

	"github.com/croese/raytrace/rtmath"
	"github.com/croese/raytrace/tuple"
)

func checkFloatValue(t *testing.T, a, b float64) {
	t.Helper()
	if !rtmath.FloatEqualEpsilon(a, b) {
		t.Errorf("expected %f to equal %f", a, b)
	}
}

func TestCreation(t *testing.T) {
	t.Run("CreatePointTuple4", func(t *testing.T) {
		a := tuple.New(4.3, -4.2, 3.1, 1.0)

		checkFloatValue(t, a.X(), 4.3)
		checkFloatValue(t, a.Y(), -4.2)
		checkFloatValue(t, a.Z(), 3.1)
		checkFloatValue(t, a.W(), 1.0)

		if !a.IsPoint() {
			t.Errorf("expected %+v to be a point", a)
		}

		if a.IsVector() {
			t.Errorf("expected %+v to not be a point", a)
		}
	})

	t.Run("CreateVectorTuple4", func(t *testing.T) {
		a := tuple.New(4.3, -4.2, 3.1, 0.0)

		checkFloatValue(t, a.X(), 4.3)
		checkFloatValue(t, a.Y(), -4.2)
		checkFloatValue(t, a.Z(), 3.1)
		checkFloatValue(t, a.W(), 0.0)

		if a.IsPoint() {
			t.Errorf("expected %+v to be a vector", a)
		}

		if !a.IsVector() {
			t.Errorf("expected %+v to be a vector", a)
		}
	})

	t.Run("PointFactory", func(t *testing.T) {
		p := tuple.Point(4, -4, 3)

		checkFloatValue(t, p.X(), 4)
		checkFloatValue(t, p.Y(), -4)
		checkFloatValue(t, p.Z(), 3)
		checkFloatValue(t, p.W(), 1)
	})

	t.Run("VectorFactory", func(t *testing.T) {
		p := tuple.Vector(4, -4, 3)

		checkFloatValue(t, p.X(), 4)
		checkFloatValue(t, p.Y(), -4)
		checkFloatValue(t, p.Z(), 3)
		checkFloatValue(t, p.W(), 0)
	})
}

func checkTupleEquality(t *testing.T, a, b tuple.Tuple4) {
	t.Helper()
	if !a.Equals(b) {
		t.Errorf("expected %+v and %+v to be equal", a, b)
	}
}

func TestEquality(t *testing.T) {
	a, b := tuple.New(1, 2, 3, 4), tuple.New(0+1, 2.0, 4-1, 2+2)

	checkTupleEquality(t, a, b)
}

func TestOperations(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		a1, a2 := tuple.New(3, -2, 5, 1), tuple.New(-2, 3, 1, 0)

		expected := tuple.New(1, 1, 6, 1)
		actual := a1.Plus(a2)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("SubtractionPoints", func(t *testing.T) {
		p1, p2 := tuple.Point(3, 2, 1), tuple.Point(5, 6, 7)

		expected := tuple.Vector(-2, -4, -6)
		actual := p1.Minus(p2)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("SubtractionVectorFromPoint", func(t *testing.T) {
		p, v := tuple.Point(3, 2, 1), tuple.Vector(5, 6, 7)

		expected := tuple.Point(-2, -4, -6)
		actual := p.Minus(v)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("SubtractionVectors", func(t *testing.T) {
		v1, v2 := tuple.Vector(3, 2, 1), tuple.Vector(5, 6, 7)

		expected := tuple.Vector(-2, -4, -6)
		actual := v1.Minus(v2)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("SubtractionFromZero", func(t *testing.T) {
		zero, v := tuple.Vector(0, 0, 0), tuple.Vector(1, -2, 3)

		expected := tuple.Vector(-1, 2, -3)
		actual := zero.Minus(v)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("Negation", func(t *testing.T) {
		a := tuple.New(1, -2, 3, -4)

		expected := tuple.New(-1, 2, -3, 4)
		actual := a.Negate()

		checkTupleEquality(t, expected, actual)
	})

	t.Run("ScalarMultiply", func(t *testing.T) {
		a := tuple.New(1, -2, 3, -4)

		expected := tuple.New(3.5, -7, 10.5, -14)
		actual := a.ScalarMult(3.5)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("ScalarMultiplyFraction", func(t *testing.T) {
		a := tuple.New(1, -2, 3, -4)

		expected := tuple.New(0.5, -1, 1.5, -2)
		actual := a.ScalarMult(0.5)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("ScalarDivision", func(t *testing.T) {
		a := tuple.New(1, -2, 3, -4)

		expected := tuple.New(0.5, -1, 1.5, -2)
		actual := a.Div(2)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("Dot", func(t *testing.T) {
		a := tuple.Vector(1, 2, 3)
		b := tuple.Vector(2, 3, 4)

		actual := a.Dot(b)

		checkFloatValue(t, 20, actual)
	})

	t.Run("Cross", func(t *testing.T) {
		a, b := tuple.Vector(1, 2, 3), tuple.Vector(2, 3, 4)

		checkTupleEquality(t, tuple.Vector(-1, 2, -1), a.Cross(b))
		checkTupleEquality(t, tuple.Vector(1, -2, 1), b.Cross(a))
	})
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		v        tuple.Tuple4
		expected float64
	}{
		{tuple.Vector(1, 0, 0), 1},
		{tuple.Vector(0, 1, 0), 1},
		{tuple.Vector(0, 0, 1), 1},
		{tuple.Vector(1, 2, 3), math.Sqrt(14)},
		{tuple.Vector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, tt := range tests {
		actual := tt.v.Magnitude()

		checkFloatValue(t, tt.expected, actual)
	}
}

func TestNorm(t *testing.T) {
	t.Run("Norm", func(t *testing.T) {
		tests := []struct {
			v        tuple.Tuple4
			expected tuple.Tuple4
		}{
			{tuple.Vector(4, 0, 0), tuple.Vector(1, 0, 0)},
			{tuple.Vector(1, 2, 3), tuple.Vector(0.26726, 0.53452, 0.80178)},
		}

		for _, tt := range tests {
			actual := tt.v.Norm()

			checkTupleEquality(t, tt.expected, actual)
		}
	})

	t.Run("NormMagnitude", func(t *testing.T) {
		v := tuple.Vector(1, 2, 3)
		norm := v.Norm()

		checkFloatValue(t, 1, norm.Magnitude())
	})
}

func TestColors(t *testing.T) {
	t.Run("ColorTuple", func(t *testing.T) {
		c := tuple.Color(-0.5, 0.4, 1.7)

		checkFloatValue(t, -0.5, c.Red())
		checkFloatValue(t, 0.4, c.Green())
		checkFloatValue(t, 1.7, c.Blue())
	})

	t.Run("Addition", func(t *testing.T) {
		c1, c2 := tuple.Color(0.9, 0.6, 0.75), tuple.Color(0.7, 0.1, 0.25)

		expected := tuple.Color(1.6, 0.7, 1.0)
		actual := c1.Plus(c2)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("Subtraction", func(t *testing.T) {
		c1, c2 := tuple.Color(0.9, 0.6, 0.75), tuple.Color(0.7, 0.1, 0.25)

		expected := tuple.Color(0.2, 0.5, 0.5)
		actual := c1.Minus(c2)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("ScalarMult", func(t *testing.T) {
		c := tuple.Color(0.2, 0.3, 0.4)

		expected := tuple.Color(0.4, 0.6, 0.8)
		actual := c.ScalarMult(2)

		checkTupleEquality(t, expected, actual)
	})

	t.Run("ElementwiseMult", func(t *testing.T) {
		c1, c2 := tuple.Color(1, 0.2, 0.4), tuple.Color(0.9, 1, 0.1)

		expected := tuple.Color(0.9, 0.2, 0.04)
		actual := c1.ElementwiseMult(c2)

		checkTupleEquality(t, expected, actual)
	})
}
