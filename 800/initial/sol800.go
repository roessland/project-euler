package main

import (
	"fmt"
	"github.com/roessland/gopkg/primegen"
	"math"
	"time"
)

func solveEquation(N float64, p float64) float64 {
	// p^q q^p = N^N
	// log(p^q q^p) = log(N^N)
	// q log(p) + p log(q) = N log N
	f := func(q float64) float64 {
		return q*math.Log(p) + p*math.Log(q) - N*math.Log(N)
	}
	dfdq := func(q float64) float64 {
		return math.Log(p) + p/q
	}
	next := func(q float64) float64 {
		return q - f(q)/dfdq(q)
	}
	// estimate x0
	// q + p = N log N => q = N log N - p
	q := N*math.Log(N) - p
	for math.Abs(f(q)) > 0.0001 {
		q = next(q)
	}
	return q
}

func main() {
	t0 := time.Now()
	N := 800800.0
	q2Max := solveEquation(N, 2.0)
	primes := primegen.SliceFromMap(primegen.Map(int64(q2Max + 2)))
	count := 0
	for i := 0; float64(primes[i]) < q2Max; i++ {
		qMax := solveEquation(N, float64(primes[i]))
		for j := i + 1; float64(primes[j]) < qMax; j++ {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(time.Since(t0))
}
