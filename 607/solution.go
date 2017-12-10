package main

import "fmt"
import "math"

// Speed of walking in normal terrain
const c float64 = 10.0

type Line struct {
	A, B float64
}

type Vec struct {
	X, Y float64
}

func (v Vec) Mul(a float64) Vec {
	return Vec{v.X * a, v.Y * a}
}

func (v1 Vec) Add(v2 Vec) Vec {
	return Vec{v1.X + v2.X, v1.Y + v2.Y}
}

func (v Vec) Norm2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec) Normalized() Vec {
	length := v.Norm2()
	return Vec{v.X / length, v.Y / length}
}

func (v1 Vec) Distance(v2 Vec) float64 {
	dx := v1.X - v2.X
	dy := v1.Y - v2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

var A Vec = Vec{0, 0}

var B Vec = Vec{100, 0}

var Y []Line = []Line{
	{A: 1, B: 0},                     // Y0
	{A: 1, B: -50 + 25*math.Sqrt(2)}, // Y1
	{A: 1, B: -50 + 15*math.Sqrt(2)}, // Y2
	{A: 1, B: -50 + 5*math.Sqrt(2)},  // Y3
	{A: 1, B: -50 - 5*math.Sqrt(2)},  // Y4
	{A: 1, B: -50 - 15*math.Sqrt(2)}, // Y5
	{A: 1, B: -50 - 25*math.Sqrt(2)}, // Y6
	{A: 1, B: -100},                  // Y7
}

var Velocity []float64 = []float64{10, 9, 8, 7, 6, 5, 10}

func LineLineIntersect(start Vec, dir Vec, line Line) Vec {
	x0, y0, dx, dy, a, b := start.X, start.Y, dir.X, dir.Y, line.A, line.B
	r := (b - y0 + a*x0) / (dy - a*dx)
	return Vec{x0 + r*dx, y0 + r*dy}
}

func Refract(dir Vec, a float64, v1, v2 float64) Vec {
	r := v2 / v1
	l := dir.Normalized()
	n := Vec{-a, 1}.Normalized()
	cos := -n.X*l.X - n.Y*l.Y
	tmp := r*cos - math.Sqrt(1-r*r*(1-cos*cos))
	return l.Mul(r).Add(n.Mul(tmp))
}

func RayTrace(azimuth float64, verbose bool) (Vec, float64) {
	azimuthRads := azimuth / 180 * math.Pi
	time := 0.0
	pos := Vec{0, 0}
	dir := Vec{math.Sin(azimuthRads), math.Cos(azimuthRads)}
	for i := 1; i < len(Y); i++ {
		nextPos := LineLineIntersect(pos, dir, Y[i])
		if verbose {
			fmt.Println("Going to", nextPos, "with velocity", Velocity[i-1])
			fmt.Println("(", nextPos.X, ", ", nextPos.Y, ")")
		}
		time += pos.Distance(nextPos) / Velocity[i-1]
		pos = nextPos
		if i < len(Y)-1 {
			dir = Refract(dir, Y[i].A, Velocity[i-1], Velocity[i])
		}
	}
	return pos, time
}

func Error(azimuth float64) float64 {
	pos, _ := RayTrace(azimuth, false)
	return pos.Y - B.Y
}

func ArgMin(F func(float64) float64, a, b float64) float64 {
	tol := 1e-14
	for {
		mid := (a + b) / 2
		if (b-a)/2 < tol {
			return mid
		}
		if math.Signbit(F(mid)) == math.Signbit(F(a)) {
			a = mid
		} else {
			b = mid
		}
	}
}

func main() {
	bestAzimuth := ArgMin(Error, 50.0, 90.0)
	pos, time := RayTrace(bestAzimuth, false)
	fmt.Println("Choosing azimuth of", bestAzimuth, "gets you to", pos, "in", time, "days")
	fmt.Printf("Answer rounded to 10 decimals: %.10f\n", time)
}
