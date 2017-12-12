package main

import "fmt"
import "math/big"
import "github.com/roessland/gopkg/primegen"

func main() {
	isPrime := primegen.Map(500)
	curr := make([]*big.Rat, 501)
	next := make([]*big.Rat, 501)
	for i := 1; i <= 500; i++ {
		if isPrime[i] {
			curr[i] = big.NewRat(1, 3)
		} else {
			curr[i] = big.NewRat(2, 3)
		}
		next[i] = big.NewRat(0, 1)
	}
	oneThird := big.NewRat(1, 3)
	twoThirds := big.NewRat(2, 3)
	oneHalf := big.NewRat(1, 2)
	for _, np := range "PNPPNPPPNNPPPP" {
		if np == 'N' {
			next[1].Mul(twoThirds, curr[2])
			next[500].Mul(twoThirds, curr[499])
		} else {
			next[1].Mul(oneThird, curr[2])
			next[500].Mul(oneThird, curr[499])
		}
		for i := 2; i <= 499; i++ {
			var P *big.Rat
			if (isPrime[i] && np == 'P') || (!isPrime[i] && np == 'N') {
				P = twoThirds
			} else {
				P = oneThird
			}
			next[i].Add(curr[i-1], curr[i+1])
			next[i].Mul(next[i], oneHalf)
			next[i].Mul(next[i], P)
		}
		curr, next = next, curr
	}
	sum := big.NewRat(0, 1)
	for i := 1; i <= 500; i++ {
		sum.Add(sum, curr[i])
	}
	fiveHundred := big.NewRat(500, 1)
	sum.Quo(sum, fiveHundred)
	fmt.Println(sum)
}
