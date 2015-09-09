// Runs in 333 ms on an i7-4770k
package main

import "fmt"

import "github.com/roessland/gopkg/mathutil"

func main() {
	squareNums := make(map[int64]bool)

	// Starting numbers
	for n0 := int64(1); n0 < 10000; n0++ {
		// Sum of squares n0^2 + ...
		S := n0 * n0
		for n := n0 + 1; S < 100000000; n++ {
			S += n * n
			squareNums[S] = true
		}
	}

	// Sum the numbers who are also palindrome
	sum := int64(0)
	for N := range squareNums {
		if mathutil.IsPalindrome(N) {
			sum += N
		}
	}
	fmt.Printf("%v\n", sum)
}
