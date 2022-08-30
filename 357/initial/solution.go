package main

import (
	"fmt"
	"github.com/roessland/gopkg/primegen"
	"time"
)

func sieve(N int64) int64 {
	// Generate primes up to and including N+1
	isPrime := primegen.Map(N + 1)

	// Assume all numbers are prime generating
	isGen := make([]bool, N+1)
	for i := range isGen {
		isGen[i] = true
	}

	// Sieve
	for d := int64(1); d <= N+1; d++ {
		for i := int64(1); i*d <= N; i++ {
			if !isPrime[d+i] {
				isGen[i*d] = false
			}
		}
	}

	sum := int64(0)
	for n := int64(1); n <= N; n++ {
		if isGen[n] {
			sum += n
		}
	}
	return sum
}

func main() {
	N := int64(100_000_000)

	t1 := time.Now()
	a := sieve(N)
	fmt.Println(time.Since(t1))

	fmt.Println(a)

}

func sieve(N int64) int64 {
	// Generate primes up to and including N+1
	isPrime := primegen.Map(N + 1)

	// Assume all numbers are prime generating
	isGen := make([]bool, N+1)
	for i := range isGen {
		isGen[i] = true
	}

	// Sieve
	for d := int64(1); d <= N+1; d++ {
		for i := int64(1); i*d <= N; i++ {
			if !isPrime[d+i] {
				isGen[i*d] = false
			}
		}
	}

	sum := int64(0)
	for n := int64(1); n <= N; n++ {
		if isGen[n] {
			sum += n
		}
	}
	return sum
}
