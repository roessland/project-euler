package main

import "fmt"
import "gitub.com/roessland/gopkg/mathutil"

func main() {
	// Problem reformulation: Find the number of coefficients that ARE NOT
	// multiples of 10^12. Then the number of coefficients that are multiples
	// of 10^12 is given by the formula N(mult) (n+1)(n+2)/2 - N(notmult).

	// Since the factorial is involved, numbers including >= 12 ten factors
	// will be congruent to zero mod 10^12. This leaves the problem:
	//
	//		Given i,j,k, n=i+j+k, determine if the multionomial coefficient
	//		n!/(i!j!k!) is divisible by 10^12

	n := 4

	// Walk trough all partitions n=ijk
	// The total when multiplied by its symmetry value should be the triangular number
	// (n+1)(n+2)/2
	sum := int64(0)
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			k := n - i - j
			if k < 0 || k > j {
				continue
			}
			multiplicity :=
				fmt.Printf("i=%v, j=%v, k=%v\n", i, j, k)
		}
	}
}
