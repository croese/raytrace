package core

import "testing"

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
