package main

import "fmt"
import "math"
import "github.com/roessland/gopkg/mathutil"

var cache []float64

func E(k int64) float64 {
	if k == 0 {
		return 0
	}

	e := cache[k]
	if e != 0 {
		return e
	}

	M := float64(1)
	for i := k - 1; i >= 0; i-- {
		M += float64(mathutil.Choose(k, i)) * (1 + E(i))
	}
	ret := M / (math.Pow(2, float64(k)) - 1)
	cache[k] = ret
	return ret
}

func main() {
	cache = make([]float64, 33)
	fmt.Printf("%v\n", E(32))
}
