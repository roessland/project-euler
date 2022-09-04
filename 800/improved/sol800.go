package main

import (
	"fmt"
	"github.com/roessland/gopkg/optimize"
	"github.com/roessland/gopkg/primegen"
	"math"
	"sort"
)

// Runs in 200 ms.

func solveEquation(N float64, p float64) float64 {
	// p^q q^p = N^N
	// log(p^q q^p) = log(N^N)
	// q log(p) + p log(q) = N log N
	f := func(q float64) float64 {
		return q*math.Log(p) + p*math.Log(q) - N*math.Log(N)
	}
	fprime := func(q float64) float64 {
		return math.Log(p) + p/q
	}
	// estimate x0
	// q + p = N log N => q = N log N - p
	q0 := N*math.Log(N) - p

	return optimize.FindRootNewtons1D(f, fprime, q0, 0.000001, nil)
}

func solve(N float64) int {
	q2Max := solveEquation(N, 2.0)
	primes := primegen.SliceFromMap(primegen.Map(int64(q2Max + N)))

	count := 0
	for i := 0; float64(primes[i]) < q2Max; i++ {
		qMax := solveEquation(N, float64(primes[i]))
		j := sort.Search(len(primes), func(j int) bool {
			return float64(primes[j]) > qMax
		})
		if j > i {
			count += j - i - 1
		}
	}
	return count
}

func main() {
	fmt.Println(solve(800800.0))
}
