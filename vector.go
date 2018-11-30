package main

import (
	"math"
)

// Vector2f represents a 2d vector
type Vector2f struct {
	x float64
	y float64
}

// NewVector2f creates a vector given x and y coordinates
func NewVector2f(x float64, y float64) *Vector2f {
	return &Vector2f{x, y}
}

// VectorEquals compares two vectors for equality
func VectorEquals(left *Vector2f, right *Vector2f) bool {
	return left.x == right.x && left.y == right.y
}

// VectorAdd computes the addition of two vectors
func VectorAdd(left *Vector2f, right *Vector2f) *Vector2f {
	return NewVector2f(left.x+right.x, left.y+right.y)
}

// VectorSubtract computes the subtraction of two vectors
func VectorSubtract(left *Vector2f, right *Vector2f) *Vector2f {
	return NewVector2f(left.x-right.x, left.y-right.y)
}

// Magnitude computes the magnitude of the vector
func (v *Vector2f) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Angle computes the angle of the vector
func (v *Vector2f) Angle() float64 {
	return math.Atan2(v.y, v.x)
}

// VectorProduct computes the multiplication of two vectors
func VectorProduct(left *Vector2f, right *Vector2f) *Vector2f {
	return NewVector2f(left.x*right.x, left.y*right.y)
}

// ScalarProduct computes the scalar product of the vector
func ScalarProduct(a *Vector2f, b *Vector2f) float64 {
	return a.x*b.x + b.y*a.y
}

// ScalarMultiplication computes the mutiplication of a vector and scalar
func ScalarMultiplication(v *Vector2f, c float64) *Vector2f {
	return NewVector2f(v.x*c, v.y*c)
}

// ScalarQuotient computes the division of a vector and scalar
func ScalarQuotient(v *Vector2f, c float64) *Vector2f {
	return NewVector2f(v.x/c, v.y/c)
}
