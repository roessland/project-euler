package main

import "fmt"

func main() {
	N := 2000000
	primes := []int{2}
	isComposite := make([]bool, N+1)
	isComposite[0] = true
	isComposite[1] = true
	isComposite[2] = false

	// Take last prime, set multiples in sieve to true
	// Next prime is the next index in array which is still false.
	// Repeat
	for it := 0; it < N; it++ {
		p := primes[len(primes)-1]
		for i := 1; i*p <= N; i++ {
			isComposite[i*p] = true
		}

		// Find next prime
		for i := p + 1; i < N; i++ {
			if isComposite[i] == false {
				primes = append(primes, i)
				break
			}
		}

		if primes[len(primes)-1] > N {
			primes = primes[0 : len(primes)-1]
			break
		}

	}

	sum := 0
	for _, p := range primes {
		sum += p
	}
	fmt.Printf("%v\n", sum)

}
