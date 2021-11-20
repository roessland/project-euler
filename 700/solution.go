package main

import (
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"sort"
	"time"
)

// Answer: 1517xxxxxxxxx556
// Time spent: 5.2299467s

var k int64 = 1504170715041707
var p int64 = 4503599627370517

func partialBruteforce(threshold int64) (sum, lastCoin int64) {
	seq := int64(0)
	lastCoin = p
	for {
		seq += k
		seq %= p
		if seq < lastCoin {
			sum += seq
			lastCoin = seq
			if lastCoin < threshold {
				return sum, lastCoin
			}
		}
	}
}

func hybridMethod(threshold int64) {
	var k int64 = 1504170715041707
	var p int64 = 4503599627370517
	var kinv = mathutil.ModularInverse(k, p) // 3451657199285664

	// Reminder: euler numbers are the LHS, if condition holds
	// Bruteforce really slows down after a while.
	// At that point, do it inversely by computing n for every sequence number below the threshold.
	// Use bruteforce to find the first euler number below a threshold, and the corresponding n.
	// Then for x==1..threshold, solve the equation  x=1504170715041707n mod 4503599627370517 for n.
	// Store solutions in a list of tuples (n, x).
	// Sort the list.
	// Loop through list, adding only eulerCoins and ignoring those that are not.
	sum, lastEulerCoin := partialBruteforce(threshold)
	// Then for x==lastEulerCoin..1, solve the equation  x=1504170715041707n mod 4503599627370517 for n.
	type tuple struct {
		n int64
		x int64
	}
	var solutions []tuple
	for x := lastEulerCoin -1; x > 0; x-- {
		// We have
		// 		x=kn mod p
		//
		// We know that k has an inverse mod p, since p is prime.
		//
		// Multiply both sides with the modular inverse.
		// 		inv(k) x = inv(k) k n (mod p)
		//
		// Simplify:
		// 		n = inv(k) x (mod p)
		n := mathutil.MulMod(kinv, x, p)
		// Small optimization:
		solutions = append(solutions, tuple{
			n: n,
			x: x,
		})
	}

	// We now have all the (n, x) tuples for x < threshold, but they are in random order.
	// The spec says we must loop through the sequence by increasing n, so do that.
	sort.Slice(solutions, func(i, j int)bool {
		return solutions[i].n < solutions[j].n
	})

	for _, sol := range solutions {
		if sol.x < lastEulerCoin {
			// sol.x is an eulerCoin.
			sum += sol.x
			lastEulerCoin = sol.x
		}
	}
	fmt.Println("Answer:", sum)
}

func main() {
	t0 := time.Now()
	hybridMethod(20000000)
	fmt.Println("Time spent:", time.Since(t0))
}