package main

import "fmt"
import "math/big"

var cache map[Arg]*big.Int

type Arg struct {
	n, a, b int
}

func StartingDigits() ([]int, []int) {
	as := []int{}
	bs := []int{}
	for a := 1; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			if a+b < 10 {
				as = append(as, a)
				bs = append(bs, b)
			}
		}
	}
	return as, bs
}

func Count(n, a, b int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	// Check cache
	cacheCount, ok := cache[Arg{n, a, b}]
	if ok {
		return cacheCount
	}

	// Possible next digits
	count := big.NewInt(int64(0))
	for c := 0; c <= 9-a-b; c++ {
		count.Add(count, Count(n-1, b, c))
	}

	// Set cache and return
	cache[Arg{n, a, b}] = count
	return count
}

func main() {
	cache = make(map[Arg]*big.Int)

	sum := big.NewInt(0)
	as, bs := StartingDigits()
	for i := 0; i < len(as); i++ {
		sum.Add(sum, Count(18, as[i], bs[i]))
	}
	fmt.Printf("%v\n", sum.String())
}
