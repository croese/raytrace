package raytrace

import (
	"math"
	"testing"
)

func verifyFloat(expected, actual float64, t *testing.T) {
	if !compareFloat64s(expected, actual) {
		t.Fatalf("expected=%g. got=%g", expected, actual)
	}
}

func TestPointTuple(t *testing.T) {
	a := NewTuple4(4.3, -4.2, 3.1, 1.0)
	verifyFloat(4.3, a.X(), t)
	verifyFloat(-4.2, a.Y(), t)
	verifyFloat(3.1, a.Z(), t)
	verifyFloat(1.0, a.W(), t)

	if !a.IsPoint() {
		t.Fatalf("expected a to be a point. a=%v", a)
	}

	if a.IsVector() {
		t.Fatalf("expected a to not be a vector. a=%v", a)
	}
}

func TestVectorTuple(t *testing.T) {
	a := NewTuple4(4.3, -4.2, 3.1, 0.0)
	verifyFloat(4.3, a.X(), t)
	verifyFloat(-4.2, a.Y(), t)
	verifyFloat(3.1, a.Z(), t)
	verifyFloat(0.0, a.W(), t)

	if a.IsPoint() {
		t.Fatalf("expected a to not be a point. a=%v", a)
	}

	if !a.IsVector() {
		t.Fatalf("expected a to be a vector. a=%v", a)
	}
}

func TestNewPointCreatesPointTuple(t *testing.T) {
	p := NewPoint(4, -4, 3)

	expected := NewTuple4(4, -4, 3, 1)
	if !p.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, p)
	}
}

func TestNewVectorCreatesVectorTuple(t *testing.T) {
	p := NewVector(4, -4, 3)

	expected := NewTuple4(4, -4, 3, 0)
	if !p.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, p)
	}
}

func TestAddingTwoTuples(t *testing.T) {
	a1 := NewTuple4(3, -2, 5, 1)
	a2 := NewTuple4(-2, 3, 1, 0)
	actual := a1.Plus(a2)
	expected := NewTuple4(1, 1, 6, 1)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}

	actual = a2.Plus(a1)
	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestSubtractingTwoPoints(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)
	actual := p1.Minus(p2)
	expected := NewVector(-2, -4, -6)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestSubtractingVectorFromPoint(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v := NewVector(5, 6, 7)
	actual := p.Minus(v)
	expected := NewPoint(-2, -4, -6)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestSubtractingTwoVectors(t *testing.T) {
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)
	actual := v1.Minus(v2)
	expected := NewVector(-2, -4, -6)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestSubtractingVectorFromZeroVector(t *testing.T) {
	zero := NewVector(0, 0, 0)
	v := NewVector(1, -2, 3)
	actual := zero.Minus(v)
	expected := NewVector(-1, 2, -3)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestNegatingTuple(t *testing.T) {
	a := NewTuple4(1, -2, 3, -4)
	actual := a.Negate()
	expected := NewTuple4(-1, 2, -3, 4)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestScalarMultiplication(t *testing.T) {
	a := NewTuple4(1, -2, 3, -4)
	actual := a.Scale(3.5)
	expected := NewTuple4(3.5, -7, 10.5, -14)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}

	actual = a.Scale(0.5)
	expected = NewTuple4(0.5, -1, 1.5, -2)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestScalarDivision(t *testing.T) {
	a := NewTuple4(1, -2, 3, -4)
	actual := a.Divide(2)
	expected := NewTuple4(0.5, -1, 1.5, -2)

	if expected != actual {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestVectorMagnitude(t *testing.T) {
	tests := []struct {
		x, y, z  float64
		expected float64
	}{
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
		{1, 2, 3, math.Sqrt(14)},
		{-1, -2, -3, math.Sqrt(14)},
	}

	for i, tt := range tests {
		v := NewVector(tt.x, tt.y, tt.z)
		actual := v.Magnitude()
		if tt.expected != actual {
			t.Fatalf("magnitude tests[%d]: expected=%v. got=%v", i,
				tt.expected, actual)
		}
	}
}

func TestNormalizingVector(t *testing.T) {
	v := NewVector(4, 0, 0)
	actual := v.Normalize()
	expected := NewVector(1, 0, 0)

	if !actual.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}

	v = NewVector(1, 2, 3)
	actual = v.Normalize()
	expected = NewVector(0.26726, 0.53452, 0.80178)

	if !actual.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestMagnitudeOfNormalizedVector(t *testing.T) {
	v := NewVector(1, 2, 3)
	norm := v.Normalize()

	if norm.Magnitude() != 1 {
		t.Fatalf("expected magnitude to be 1. got=%v", norm.Magnitude())
	}
}

func TestDotProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	actual := a.Dot(b)
	expected := 20.0

	verifyFloat(expected, actual, t)
}

func TestCrossProduct(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	actual := a.Cross(b)
	expected := NewVector(-1, 2, -1)

	if !actual.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}

	actual = b.Cross(a)
	expected = NewVector(1, -2, 1)

	if !actual.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestColorTuple(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)
	verifyFloat(-0.5, c.Red(), t)
	verifyFloat(0.4, c.Green(), t)
	verifyFloat(1.7, c.Blue(), t)
}

func TestAddingColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	actual := c1.Plus(c2)
	expected := NewColor(1.6, 0.7, 1.0)

	if !actual.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestSubtractingColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	actual := c1.Minus(c2)
	expected := NewColor(0.2, 0.5, 0.5)

	if !actual.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}

func TestMultiplyingColors(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	actual := c1.Times(c2)
	expected := NewColor(0.9, 0.2, 0.04)

	if !actual.Equals(expected) {
		t.Fatalf("expected=%v. got=%v", expected, actual)
	}
}
