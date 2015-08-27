package main

import "fmt"
import "github.com/roessland/gopkg/primegen"
import "github.com/roessland/gopkg/mathutil"
import "github.com/roessland/gopkg/sliceutil"

var isPrime []bool
var S map[int][]int

func Twins(a, b int) bool {
	a_ := int64(a)
	b_ := int64(b)
	if isPrime[mathutil.FromDigits(append(mathutil.ToDigits(a_, 10),
		mathutil.ToDigits(b_, 10)...), 10)] &&
		isPrime[mathutil.FromDigits(append(mathutil.ToDigits(b_, 10),
			mathutil.ToDigits(a_, 10)...), 10)] {
		return true
	}
	return false
}

func main() {
	// Generate primes
	N := int64(10000)
	isPrime = primegen.Map(99999999)
	p := primegen.SliceFromMap(primegen.Map(N))

	// Store the twin primes of every small prime
	S = make(map[int][]int)
	for i := 0; i < len(p); i++ {
		pi := int(p[i])
		Spi := []int{}
		for j := i + 1; j < len(p); j++ {
			pj := int(p[j])
			if Twins(pi, pj) {
				S[pi] = append(S[pi], pj)
			}
		}
		if len(Spi) > 0 {
			S[pi] = Spi
		}
	}

	// Try reading this out loud
	for pi, Si := range S {
		for _, pii := range Si {
			Sii := sliceutil.IntersectInt(S[pii], Si)
			for _, piii := range Sii {
				Siii := sliceutil.IntersectInt(S[piii], Sii)
				for _, piiii := range Siii {
					Siiii := sliceutil.IntersectInt(S[piiii], Siii)
					for _, piiiii := range Siiii {
						fmt.Printf("%v, %v, %v, %v, %v = %v\n", pi, pii, piii, piiii, piiiii, pi+pii+piii+piiii+piiiii)
					}
				}
			}
		}
	}

}
