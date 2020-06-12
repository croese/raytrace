package raytrace

import "math"

type Tuple4 struct {
	x, y, z, w float64
}

func (t Tuple4) X() float64 {
	return t.x
}

func (t Tuple4) Y() float64 {
	return t.y
}

func (t Tuple4) Z() float64 {
	return t.z
}

func (t Tuple4) W() float64 {
	return t.w
}

func NewTuple4(x, y, z, w float64) Tuple4 {
	return Tuple4{x, y, z, w}
}

func (t Tuple4) IsPoint() bool {
	return t.w == 1.0
}

func (t Tuple4) IsVector() bool {
	return t.w == 0.0
}

func NewPoint(x, y, z float64) Tuple4 {
	return NewTuple4(x, y, z, 1.0)
}

func NewVector(x, y, z float64) Tuple4 {
	return NewTuple4(x, y, z, 0.0)
}

func (t Tuple4) Equals(other Tuple4) bool {
	return compareFloat64s(t.x, other.x) &&
		compareFloat64s(t.y, other.y) &&
		compareFloat64s(t.z, other.z) &&
		compareFloat64s(t.w, other.w)
}

func (t Tuple4) Plus(other Tuple4) Tuple4 {
	return NewTuple4(t.x+other.x, t.y+other.y, t.z+other.z, t.w+other.w)
}

func (t Tuple4) Minus(other Tuple4) Tuple4 {
	return NewTuple4(t.x-other.x, t.y-other.y, t.z-other.z, t.w-other.w)
}

func (t Tuple4) Negate() Tuple4 {
	return NewTuple4(-t.x, -t.y, -t.z, -t.w)
}

func (t Tuple4) Scale(by float64) Tuple4 {
	return NewTuple4(t.x*by, t.y*by, t.z*by, t.w*by)
}

func (t Tuple4) Divide(by float64) Tuple4 {
	return NewTuple4(t.x/by, t.y/by, t.z/by, t.w/by)
}

func (t Tuple4) Magnitude() float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z + t.w*t.w)
}

func (t Tuple4) Normalize() Tuple4 {
	magnitude := t.Magnitude()
	return NewTuple4(t.x/magnitude, t.y/magnitude,
		t.z/magnitude, t.w/magnitude)
}

func (t Tuple4) Dot(other Tuple4) float64 {
	return t.x*other.x +
		t.y*other.y +
		t.z*other.z +
		t.w*other.w
}

func (t Tuple4) Cross(other Tuple4) Tuple4 {
	return NewVector(t.y*other.z-t.z*other.y,
		t.z*other.x-t.x*other.z,
		t.x*other.y-t.y*other.x)
}

func NewColor(red, green, blue float64) Tuple4 {
	return NewVector(red, green, blue)
}

func (t Tuple4) Red() float64 {
	return t.x
}

func (t Tuple4) Green() float64 {
	return t.y
}

func (t Tuple4) Blue() float64 {
	return t.z
}

func (t Tuple4) Times(other Tuple4) Tuple4 {
	return NewColor(t.Red()*other.Red(),
		t.Green()*other.Green(), t.Blue()*other.Blue())
}
