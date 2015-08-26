package main

import "fmt"
import "github.com/roessland/gopkg/mathutil"

func Next(n int64) int64 {
	factorialSum := int64(0)
	for _, digit := range mathutil.ToDigits(n, 10) {
		factorialSum += mathutil.Factorial(digit)
	}
	return factorialSum
}

func ChainLength(n int64) int64 {
	terms := make(map[int64]bool)
	terms[n] = true
	prevTerm := n
	for {
		term := Next(prevTerm)

		// if term has already been added, return
		_, ok := terms[term]
		if ok {
			return int64(len(terms))
		}

		// else, add term and keep going
		terms[term] = true
		prevTerm = term
	}
}

func ChainsWithLength(maxNumber, chainLength int64) int64 {
	chainsWithThatLength := int64(0)
	for n := int64(0); n < maxNumber; n++ {
		if ChainLength(n) == chainLength {
			chainsWithThatLength++
		}
	}
	return chainsWithThatLength
}

func main() {
	fmt.Printf("Chains with length 60: %v\n", ChainsWithLength(1000000, 60))
}
