// Runs instantly
package main

import "fmt"
import "math"

func PickFurthest(x0, y0, xa, ya, xb, yb float64) (float64, float64) {
	distA := (xa-x0)*(xa-x0) + (ya-y0)*(ya-y0)
	distB := (xb-x0)*(xb-x0) + (yb-y0)*(yb-y0)
	if distA >= distB {
		return xa, ya
	} else {
		return xb, yb
	}
}

// Intersect find the intersection point between y=ax+b and 4x^2 + y^2=100 that
// is furthest from (x0, y0).
func Intersect(x0, y0, a, b float64) (float64, float64) {
	A := 4 + a*a
	B := 2 * a * b
	C := b*b - 100
	x1 := (-B + math.Sqrt(B*B-4*A*C)) / (2 * A)
	x2 := (-B - math.Sqrt(B*B-4*A*C)) / (2 * A)
	y1 := a*x1 + b
	y2 := a*x2 + b
	return PickFurthest(x0, y0, x1, y1, x2, y2)
}

// Reflect finds the coefficients a and b for the reflection of the line
// y=a0*x + b0 in a point on the ellipse (x, y).
func Reflect(x, y, a0, b0 float64) (float64, float64) {
	// Inbound vector
	dx := float64(1)
	dy := a0

	// Normal vector
	norm := math.Sqrt(1 + 16*x*x/(y*y))
	nx := 4 * (x / y) / norm
	ny := 1 / norm

	// Reflection vector
	dot := dx*nx + dy*ny
	rx := dx - 2*dot*nx
	ry := dy - 2*dot*ny

	// Convert to line coefficients
	a := ry / rx
	b := y - a*x

	return a, b
}

func main() {
	// Initial line and starting point
	a := float64(-9.6-10.1) / float64(1.4)
	b := float64(10.1)
	x := float64(0.0)
	y := float64(10.1)

	numReflections := 0
	for i := 0; ; i++ {
		x, y = Intersect(x, y, a, b)
		a, b = Reflect(x, y, a, b)
		if -0.01 <= x && x <= 0.01 && y > 0 {
			fmt.Printf("%v (%0.3f, %0.3f)\n", numReflections, x, y)
			break
		}
		numReflections++
	}
}
