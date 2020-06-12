package core

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
