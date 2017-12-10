package main

import "testing"
import "math"

//import "math"
import "github.com/stretchr/testify/assert"

func TestLineLineIntersectDiagonal(t *testing.T) {
	start := Vec{0, 0}
	dir := Vec{0.4, -0.4}
	line := Line{A: 1, B: -2}
	pos := LineLineIntersect(start, dir, line)
	assert.Equal(t, 1.0, pos.X)
	assert.Equal(t, -1.0, pos.Y)
}

func TestLineLineIntersectHorizontal(t *testing.T) {
	start := Vec{0, 0}
	dir := Vec{0.4, 0.0}
	line := Line{A: 1, B: -2.3}
	pos := LineLineIntersect(start, dir, line)
	assert.Equal(t, 2.3, pos.X)
	assert.Equal(t, 0.0, pos.Y)

	// 13.4036593619 is wrong
	// 13.3787906416 is wrong
}

func TestRefract(t *testing.T) {
	l := Vec{math.Sqrt(2) / 2, -math.Sqrt(2) / 2}
	ref := Refract(l, 0, 10, 9)
	assert.True(t, math.Abs(0.636396-ref.X) < 1e-6)
	assert.True(t, math.Abs(-0.771362-ref.Y) < 1e-6)
}
