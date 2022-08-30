package main

import (
	"fmt"
)

const mod = 1000000

var palCompsCache = map[int]int{}
var palCompSumCache = map[int]int{}

var twoPalCompsCache = map[int]int{}
var twoPalCompSumCache = map[int]int{}

func palCompSum(n int) int {
	if n < 0 {
		return 0
	}
	cached, ok := palCompSumCache[n]
	if ok {
		return cached
	}

	ways := (palComps(n-2) + palCompSum(n-2)) % mod
	palCompSumCache[n] = ways
	return ways
}

func palComps(n int) int {
	if n < 0 {
		return 0
	}
	cached, ok := palCompsCache[n]
	if ok {
		return cached
	}
	ways := 1
	ways += palCompSum(n)

	palCompsCache[n] = ways
	return palCompsCache[n]
}

func twoPalCompSum(n int) int {
	if n < 0 {
		return 0
	}
	cached, ok := twoPalCompSumCache[n]
	if ok {
		return cached
	}

	ret := (twoPalComps(n) + twoPalCompSum(n-2)) % mod
	twoPalCompSumCache[n] = ret
	return ret
}

func twoPalComps(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	cached, ok := twoPalCompsCache[n]
	if ok {
		return cached
	}

	ways := 0
	ways = (ways + twoPalComps(n-2)) % mod   // Subtract 1 on each side
	ways = (ways + palComps(n-4)) % mod      // Subtract 2! on each side
	ways = (ways + twoPalCompSum(n-6)) % mod // Subtract 3 or more on each side

	twoPalCompsCache[n] = ways
	return twoPalCompsCache[n]
}

func main() {
	for i := 42; ; i++ {
		if twoPalComps(i) == 0 {
			fmt.Printf("t(n) mod 1000000 = 0, n = %d", i)
			break
		}
	}

}
