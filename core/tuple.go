package core

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
