package dimension

import "fmt"

// Vec is a container of 2 ints, X and Y (a 2D integer vector).
type Vec struct {
	X, Y int
}

// V returns a new Vec.
func V(x, y int) Vec {
	return Vec{X: x, Y: y}
}

// Add returns the sum of the vectors.
func (u Vec) Add(v Vec) Vec {
	return V(u.X+v.X, u.Y+v.Y)
}

// Dot returns the dot product of the vector.
func (u Vec) Dot() int {
	return u.X * u.Y
}

// String returns the string representation of the vector.
func (u Vec) String() string {
	return fmt.Sprintf("Vec(%v, %v)", u.X, u.Y)
}

// Inside returns true if 0 < u < v
func (u Vec) Inside(v Vec) bool {
	if (u.X <= v.X) && (u.X >= 0) && (u.Y <= v.Y) && (u.Y >= 0) {
		return true
	}
	return false
}
