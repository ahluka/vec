package vec

import "math"

type Vec2 struct {
	X, Y float64
}

// MagSq returns the squared magnitude of the Vec2
// Useful if you only want to compare two of them but don't care about their actual lengths.
func (v Vec2) MagSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Mag returns the magnitude (length) of the Vec2
func (v Vec2) Mag() float64 {
	return math.Sqrt(v.MagSq())
}

// Add one Vec2 to this one, returning the resulting Vec2
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

// Subtract one Vec2 from this one, returning the resulting Vec2
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}
