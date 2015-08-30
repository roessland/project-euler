package main

import "fmt"

// Find the unique positive integer whose square has the form
// 1_2_3_4_5_6_7_8_9_0, where each “_” is a single digit.

// Max integer: 1929394959697989990
// Sqrt(max) = 1389026623

// Min integer: 1020304050607080900
// Sqrt(min)  = 1010101010

// Range: 1389025941 ~ 10^9, which is feasible to loop trough

// Valid checks if a number is of the form 1_2_3_4_5_6_7_8_9_0, where _ can be
// any digits (not necessarily the same digit).
func Valid(n int64) bool {
	i := int64(0)
	for n > 0 {
		digit := n % 10
		if digit != i {
			return false
		}
		n = n / 10
		digit = n % 10
		n = n / 10
		i = ((i - 1) + 10) % 10
	}
	return true
}

func main() {
	for n := int64(1010101010); ; n += 2 {
		if Valid(n * n) {
			fmt.Printf("%v^2 = %v\n", n, n*n)
			break
		}
	}
}
