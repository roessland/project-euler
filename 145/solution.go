package main

import "github.com/roessland/gopkg/mathutil"

/*
Some positive integers n have the property that the sum [ n + reverse(n) ]
consists entirely of odd (decimal) digits. For instance, 36 + 63 = 99 and 409 +
904 = 1313. We will call such numbers reversible; so 36, 63, 409, and 904 are
reversible. Leading zeroes are not allowed in either n or reverse(n).

There are 120 reversible numbers below one-thousand.

How many reversible numbers are there below one-billion (109)?
*/

import "fmt"

func EntirelyOddSlow(n int64) bool {
	digits := mathutil.ToDigits(n, 10)
	for _, d := range digits {
		if d%2 == 0 {
			return false
		}
	}
	return true
}

// EntirelyOdd checks if all digits in number are odd
func EntirelyOdd(n int64) bool {
	for n > 0 {
		if n%2 == 0 {
			return false
		}
		n = n / 10
	}
	return true
}

func Reversible(n int64) bool {
	if n%10 == 0 {
		return false
	}
	r := mathutil.Reverse(n, 10)
	return EntirelyOdd(n + r)
}

func ReversibleFast(n int64, N int64) bool {
	if n%10 == 0 {
		return false
	}
	nr := n
	for n > 0 {
		nr += (n % 10) * N
		N /= 10
		n /= 10
	}
	return EntirelyOdd(nr)
}

func BruteForce(N int64) int64 {
	var numReversible int64 = 0
	for n := int64(1); n < N; n++ {
		if Reversible(n) {
			numReversible++
		}
	}
	return numReversible
}

func Smart(Npow int64) int64 {
	var numReversible int64 = 0
	for npow := int64(1); npow <= Npow; npow++ {
		N := mathutil.Pow(10, npow)
		for n := N / 10; n < N; n++ {
			if ReversibleFast(n, N/10) {
				numReversible++
			}
		}
	}
	return numReversible
}

func main() {
	fmt.Println("10^2", BruteForce(100))
	fmt.Println("10^3", BruteForce(1000))
	fmt.Println("10^4", Smart(4))
	fmt.Println("10^5", Smart(5))
	fmt.Println("10^6", Smart(6))
	fmt.Println("10^7", Smart(7))
	fmt.Println("10^8", Smart(8))
	fmt.Println("10^9", Smart(9))
}
