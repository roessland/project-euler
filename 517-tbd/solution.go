package main

import "fmt"
import "time"
import "math"
import "github.com/roessland/gopkg/mathutil"

// Numeric value = b + c * sqrt(a)
func g(a, b, c int64) int64 {
	var ret int64
	if b*b < (c+1)*(c+1)*a {
		ret = 1
	} else {
		ret = g(a, b-1, c) + g(a, b, c+1)
	}
	return ret
}

func SlowAlgo(a int64) int64 {
	return g(a, a, 0)
}

// b - c sqrt(a) < a
func Under(a, b, c int64) bool {
	return b*b < (c+1)*(c+1)*a
}

func OverEqual(a, b, c int64) bool {
	return !Under(a, b, c)
}

func Number(a, b, c int64) float64 {
	return float64(b) - float64(c)*math.Sqrt(float64(a))
}

func Algo(a int64) int64 {
	// Fast algorithmc
	t0 := time.Now().UnixNano()
	m := int64(math.Ceil(math.Sqrt(float64(a))))
	prev := make([]int64, m)
	curr := make([]int64, m)

	// Special positions: (3,0), (4,1), (7,2), (10,1)
	for b := int64(1); b <= a; b++ {
		for c := m - 2; c >= 0; c-- {
			north := prev[c]
			if north == 0 && Under(a, b, c) != Under(a, b-1, c) {
				north = 1
			}
			east := curr[c+1]
			if east == 0 && Under(a, b, c) != Under(a, b, c+1) {
				east = 1
			}
			curr[c] = (north + east) % 1000000007
			prev = curr
		}
	}

	t1 := time.Now().UnixNano()
	fmt.Printf("Norm algorithm, %0.3fms = %v = ", float64(t1-t0)/1000000.0, curr[0])
	fmt.Printf("461204842\n") // Answer for 1001000

	return curr[0]
}

func FastAlgo(a int64) int64 {
	t0 := time.Now().UnixNano()
	m := int64(math.Ceil(math.Sqrt(float64(a))))

	// Step 1: Detect all special points.
	// A point is special if:
	//
	// 1. The sign changes either west and/or south
	// 2. There are no special points above it
	S := int64(1)
	c := int64(0)
	b := int64(0)

	// Find the initial special point
	b = int64(math.Ceil((float64(c) + 1.0) * math.Sqrt(float64(a))))
	c = 1

	for c := int64(1); c < m; c++ {
		fmt.Printf("(%v,%v)\n", c+a-b, c)
		S = (S + mathutil.ChooseMod(c+a-b, c, 1000000007)) % 1000000007

		// Find the next special point by going down until sign changes
		b = int64(math.Ceil((float64(c) + 1.0) * math.Sqrt(float64(a))))
	}

	t1 := time.Now().UnixNano()
	fmt.Printf("Fast algorithm, %0.3fms = %v", float64(t1-t0)/1000000.0, S)
	return S
}

func main() {
	// Around 10000 primes, so this algorithm should run in around 6 ms for each iteration!!!
	//Algo(1001)
	Algo(90)
	FastAlgo(1001000)
}
