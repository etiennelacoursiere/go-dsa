package vector2

import (
	"math"
)

type Vector2 struct {
	x float64
	y float64
}

func (v *Vector2) Add(other Vector2) *Vector2 {
	v.x = v.x + other.x
	v.y = v.y + other.y

	return v
}

func (v *Vector2) Sub(other Vector2) *Vector2 {
	v.x = v.x - other.x
	v.y = v.y - other.y

	return v
}

func (v *Vector2) Div(other Vector2) *Vector2 {
	// TODO: What to do in the case of other having zero values?
	v.x = v.x / other.x
	v.y = v.y / other.y

	return v
}

func (v *Vector2) Length() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func (v *Vector2) Magnitude() float64 {
	return v.Length()
}

func (v *Vector2) Normalize() *Vector2 {
	length := v.Length()
	if length != 0 {
		v.x = v.x / v.Length()
		v.y = v.y / v.Length()

		return v
	}

	return v.Scale(1.0 / math.Sqrt(length))
}

func (v *Vector2) Normalized() Vector2 {
	if v.Length() != 0 {
		return Vector2{
			x: v.x / v.Length(),
			y: v.y / v.Length(),
		}
	}

	return Vector2{x: 0, y: 0}
}

func (v *Vector2) Scale(s float64) *Vector2 {
	v.x *= s
	v.y *= s
	return v
}

func (v *Vector2) Scaled(s float64) Vector2 {
	return Vector2{
		x: v.x * s,
		y: v.y * s,
	}
}

func Distance(a Vector2, b Vector2) float64 {
	c := a.Sub(b)
	return c.Length()
}
