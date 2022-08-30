package main

import (
	"fmt"
)

const mod = 1000000

/*

6 -> 3

divide evenly
1 1 1, 1 2, 2 1, 3

divide odd, middle 1 impossible since 6 is even
divide odd, middle 2, (6 - 2 = 4), so 4/2 = 2 on each side
1 1 (1 1 2 1 1), 2 (2 2 2)

divide odd, middle 3 impossible since 6 even
divide odd, middle 4, (6-4 = 2), so 2/2 = 1 on each side
1 (1 4 1)

divide odd, middle 6
, (6)


So for 6 we need:
- entire number = 1
- compositions of 6/2=3 = 4
- compositions of 4/2=2 = 2
- compositions of 2/2=1 = 1
- compositions of 0/2=0 = 0 (by definition)

And for 7 we need:
- compositions of 7/2=3 = 4
- compositions of 5/2=2 = 2
- compositions of 3/2=1 = 1
- +1 for entire number
7, 1 1 1 1 1 1 1
     2 1 1 1 2
     1 2 1 2 1
       3 1 3
     1 1 3 1 1
       2 3 2
       1 5 1
         7

So C(n+1) = C(n) if n is even.

c(n) = 2^(n-1) (all combinations with size above 0)
How to filter for those containing a 2?
Say n=6
1 1 1 1 1 1
Two in middle: Each side is 2.
2 2 2
1 1 2 1 1.
So c((6-2)/2)=2 for those ones.

Two on sides: (6/2-1) possibilities
2 1 1 2: c(6/2-2)=1
1 2 2 1: c(6/2-2)=1
Total: (6/2-1)*c(6/2-2)

n=8
1 1 1 1 1 1 1 1
 x x x
2 1 1 = c(0) + c(2)
1 2 1 = c(1) + c(1)
1 1 2 = c(2) + c(0)

For n=8:
Two in middle:

For n=20:
Two in middle:
c((20-2)/2)=256
On sides:
(20/2-1)*c(20/2-2 )=9*128=1152

c(0) = 0
c(1) = 1
c(2) = 2
c(3) = 4
c(4) = 8
c(5) = 16
c(6) = 32
c(7) = 64
c(8) = 128


Each side is 10, put 2 somewhere. 8 left. Partitions of 8 are:
0, 8, c(0) + c(8) = 0 + 128 = 128
1, 7, c(1) + c(7) = 1 + 64 = 65
2, 6, c(2) + c(6) = 2 + 32 = 34
3, 5, c(3) + c(5) = 4 + 16 = 20
4, 4, c(4) + c(4) = 8 + 8 = 16
5, 3, c(5) + c(3) = 16 + 4 = 20
6, 2, c(6) + c(2) = 32 + 2 = 34
7, 1, c(7) + c(1) = 64 + 1 = 65
8, 0, c(8) + c(0) = 128 + 0 = 128
*/

/*

3
2 1
1 2
1 1 1

3
pick 1
	2 left -> 2
pick 2
	1 left -> 1
pick 3
	0 left -> 0
*/

var compsCache = map[int]int{}

func Comps(n int) int {
	var comps func(int) int
	comps = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}

		cached, ok := compsCache[n]
		if ok {
			return cached
		}

		ways := 1
		for i := 1; i < n; i++ {
			ways += comps(n - i)
		}

		if ways == 0 {
			panic("overflow")
		}
		compsCache[n] = ways
		return compsCache[n]
	}
	return comps(n)
}

var palCompsCache = map[int]int{}
var palCompSumCache = map[int]int{}

func PalComps(n int) int {
	var palComps func(int) int
	var palCompSum func(int) int

	palCompSum = func(n int) int {
		//waysA := 0
		//for i := 2; i <= n; i += 2 {
		//	waysA = (waysA + palComps(n-i)) % mod
		//}
		//return waysA

		//+ palComps(n-2)
		//+ palComps(n-4)
		//+ palComps(n-6)
		// ...
		// //+ palComps(0)

		if n < 0 {
			return 0
		}
		if n == 0 {
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

	palComps = func(n int) int {
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 1
		}
		if n == 1 {
			return 1
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
	return palComps(n)
}

type twoPalValue struct {
	hasTwo bool
	comps  int
}

var twoPalCompsCache = map[int]int{}
var twoPalCompSumCache = map[int]int{}

func TwoPalComps(n int) int {
	var twoPalCompSum func(n int) int
	var twoPalComps func(int) int

	twoPalCompSum = func(n int) int {
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

	twoPalComps = func(n int) int {
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 0
		}
		if n == 1 {
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
		ways = (ways + PalComps(n-4)) % mod      // Subtract 2! on each side
		ways = (ways + twoPalCompSum(n-6)) % mod // Subtract 3 or more on each side

		twoPalCompsCache[n] = ways
		return twoPalCompsCache[n]
	}
	return twoPalComps(n)
}

func main() {
	for i := 42; ; i++ {
		if TwoPalComps(i) == 0 {
			fmt.Printf("t(n) mod 1000000 = 0, n = %d", i)
			break
		}
	}
}
