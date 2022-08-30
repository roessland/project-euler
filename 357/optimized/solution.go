package main

import (
	"fmt"
	"github.com/roessland/gopkg/primegen"
	"time"
)

func main() {
	t0 := time.Now()

	N := int64(100_000_000)

	// Generate primes up to and including N+1
	isPrime := primegen.Map(N + 1)

	// Assume all even numbers are prime generating.
	// Store only even numbers, since 1+n/1 must be prime.
	// All primes other than 2 are odd, so n=p-1 means n is even,
	// except for 1.
	// If isGen[i]=true then 2*i is prime generating.
	isGen := make([]bool, N/2+1)
	for i := range isGen {
		isGen[i] = true
	}

	// Sieve out those that aren't prime generating.
	// Loop through all divisors and multiply them up.
	for d := int64(1); d <= N+1; d++ {
		for i := int64(1); i*d <= N; i++ {
			// Ignore odd n since they can't be prime generating.
			if i*d%2 == 0 && !isPrime[d+i] {
				isGen[(i*d)/2] = false
			}
		}
	}

	sum := int64(1) // Add 1 since 2 is an even prime.
	for n := int64(2); n <= N; n += 2 {
		if isGen[n/2] {
			sum += n
		}
	}

	fmt.Println(time.Since(t0))
	fmt.Println(sum)
}
