package main

import "fmt"
import "math"

func Theta(x, y float64) float64 {
	dotNum := x*x + y*y - 4*x - 3*y
	dotDenom := math.Sqrt((x*x + 9 - 6*y + y*y) * (16 - 8*x + x*x + y*y))
	ret := math.Acos(dotNum / dotDenom)
	if math.IsNaN(ret) {
		return float64(0.0)
	}
	return ret
}

func Theta2(y float64) float64 {
	dotNum := y*y - 3*y
	dotDenom := math.Sqrt((9 - 6*y + y*y) * (16 + y*y))
	ret := math.Acos(dotNum / dotDenom)
	if math.IsNaN(ret) {
		return float64(0.0)
	}
	return ret
}

func One(y float64) float64 {
	return 1.0
}

func main() {
	fmt.Println("Approximation using 2D integration")
	dx := 1e-2
	dy := 1e-2

	var tot, c float64
	for x := float64(0.0); x < float64(4); x += dx {
		for y := float64(0.0); y < 3-0.75*x; y += dy {
			delta := Theta(x, y)*dy*dx - c
			t := tot + delta
			c = (t - tot) - delta
			tot = t
		}
	}
	fmt.Println(tot / (12 * math.Pi))

	fmt.Println("Smarter integration using symmetry/substitution")
	tot = float64(0.0)
	c = float64(0.0)
	dy = 1e-9
	dl := dy * 4.0 / 5.0
	for y := float64(0.0); y < 3; y += dy {
		delta := Theta2(y)*5.0/3.0*y*dl - c
		t := tot + delta
		c = (t - tot) - delta
		tot = t
	}
	fmt.Println(tot / (12 * math.Pi))
}
