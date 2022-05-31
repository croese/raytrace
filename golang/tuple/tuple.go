package tuple

import (
	"math"

	"github.com/croese/raytrace/rtmath"
)

type Tuple4 struct {
	x, y, z, w float64
}

func New(x, y, z, w float64) Tuple4 {
	return Tuple4{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

func Point(x, y, z float64) Tuple4 {
	return New(x, y, z, 1.0)
}

func Vector(x, y, z float64) Tuple4 {
	return New(x, y, z, 0.0)
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

func (t Tuple4) IsPoint() bool {
	return rtmath.FloatEqualEpsilon(t.w, 1.0)
}

func (t Tuple4) IsVector() bool {
	return rtmath.FloatEqualEpsilon(t.w, 0.0)
}

func (t Tuple4) Equals(b Tuple4) bool {
	return rtmath.FloatEqualEpsilon(t.x, b.x) &&
		rtmath.FloatEqualEpsilon(t.y, b.y) &&
		rtmath.FloatEqualEpsilon(t.z, b.z) &&
		rtmath.FloatEqualEpsilon(t.w, b.w)
}

func (t Tuple4) Plus(b Tuple4) Tuple4 {
	return Tuple4{
		x: t.x + b.x,
		y: t.y + b.y,
		z: t.z + b.z,
		w: t.w + b.w,
	}
}

func (t Tuple4) Minus(b Tuple4) Tuple4 {
	return Tuple4{
		x: t.x - b.x,
		y: t.y - b.y,
		z: t.z - b.z,
		w: t.w - b.w,
	}
}

func (t Tuple4) Negate() Tuple4 {
	return Tuple4{
		x: -t.x,
		y: -t.y,
		z: -t.z,
		w: -t.w,
	}
}

func (t Tuple4) Times(scalar float64) Tuple4 {
	return Tuple4{
		x: t.x * scalar,
		y: t.y * scalar,
		z: t.z * scalar,
		w: t.w * scalar,
	}
}

func (t Tuple4) Div(scalar float64) Tuple4 {
	return Tuple4{
		x: t.x / scalar,
		y: t.y / scalar,
		z: t.z / scalar,
		w: t.w / scalar,
	}
}

func (t Tuple4) Magnitude() float64 {
	return math.Sqrt(t.x*t.x +
		t.y*t.y +
		t.z*t.z +
		t.w*t.w)
}

func (t Tuple4) Norm() Tuple4 {
	m := t.Magnitude()

	return New(t.x/m,
		t.y/m,
		t.z/m,
		t.w/m)
}

func (t Tuple4) Dot(b Tuple4) float64 {
	return t.x*b.x +
		t.y*b.y +
		t.z*b.z +
		t.w*b.w
}

func (t Tuple4) Cross(b Tuple4) Tuple4 {
	return Vector(t.y*b.z-t.z*b.y,
		t.z*b.x-t.x*b.z,
		t.x*b.y-t.y*b.x)
}
