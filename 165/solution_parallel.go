package main

import "fmt"
import "math/big"

import "crypto/sha1"

var srand int = 629527 // 290797. Skipped to next point for correct start.

var sha1nil = [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func Rand() int64 {
	t := srand % 500
	srand = srand * srand % 50515093
	return int64(t)
}

type Vec struct {
	X, Y int64
}

func (v1 Vec) Sub(v2 Vec) Vec {
	return Vec{v1.X - v2.X, v1.Y - v2.Y}
}

// Returns hash if they intersect, otherwise nil
func Intersect(P1, Q1, P2, Q2 Vec) ([20]byte, bool) {
	// Vectors with same length and direction as line segments
	v1 := Q1.Sub(P1)
	v2 := Q2.Sub(P2)

	// Parallel vectors never intersect
	det := v2.X*v1.Y - v1.X*v2.Y
	if det == 0 {
		return sha1nil, false
	}

	// Intersection is calculated using linear algebra
	var k1, k2 big.Rat
	Pdx, Pdy := P2.X-P1.X, P2.Y-P1.Y
	invDet := big.NewRat(1, det)
	k1.Mul(invDet, big.NewRat(-v2.Y*Pdx+v2.X*Pdy, 1))
	k2.Mul(invDet, big.NewRat(-v1.Y*Pdx+v1.X*Pdy, 1))

	// Check if intersection is inside both segments
	zero := big.NewRat(0, 1)
	one := big.NewRat(1, 1)
	k1valid := k1.Cmp(zero) == 1 && k1.Cmp(one) == -1
	k2valid := k2.Cmp(zero) == 1 && k2.Cmp(one) == -1

	// Return hash of intersection coordinate if it was
	if k1valid && k2valid {
		var Ix, Iy big.Rat
		Ix.Mul(big.NewRat(v1.X, 1), &k1).Add(&Ix, big.NewRat(P1.X, 1))
		Iy.Mul(big.NewRat(v1.Y, 1), &k1).Add(&Iy, big.NewRat(P1.Y, 1))
		return sha1.Sum([]byte(fmt.Sprintf("%v,%v", Ix.String(), Iy.String()))), true
	} else {
		return sha1nil, false
	}
}

func Sum(P, Q []Vec, i0, i1 int, pointChan chan<- [20]byte, done chan<- bool) {
	// N^2 loop to check for intersections
	for i := i0; i < i1; i++ {
		for j := 0; j < i; j++ {
			point, intersect := Intersect(P[i], Q[i], P[j], Q[j])
			if intersect {
				pointChan <- point
			}
		}
	}
	done <- true
}

func main() {
	// sha1(intersection point)
	intersectionPoints := make(map[[20]byte]bool)

	N := 5000

	// Generate all segments
	P, Q := make([]Vec, N), make([]Vec, N)
	for i := 0; i < N; i++ {
		P[i], Q[i] = Vec{Rand(), Rand()}, Vec{Rand(), Rand()}
	}

	// Retrieve points and put them in the map
	pointChan := make(chan [20]byte, 100)

	// For finding out when goroutines are finished
	doneChan := make(chan bool, 1)

	// Start four goroutines with equal amount of work.
	// Indices chosen so each goroutine has an equal amount of work.
	go Sum(P, Q, 0000, 2500, pointChan, doneChan)
	go Sum(P, Q, 2500, 3535, pointChan, doneChan)
	go Sum(P, Q, 3535, 4330, pointChan, doneChan)
	go Sum(P, Q, 4330, 5000, pointChan, doneChan)

	numRemaining := 4
	for numRemaining > 0 {
		select {
		case point := <-pointChan:
			intersectionPoints[point] = true
		case <-doneChan:
			numRemaining--
		}
	}

	fmt.Printf("Number of distinct intersection points: %v\n", len(intersectionPoints))
}
