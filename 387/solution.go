package main

import "fmt"
import "github.com/roessland/gopkg/mathutil"
import "github.com/roessland/gopkg/sliceutil"
import "github.com/roessland/gopkg/primegen"

var isPrime []bool
var primes []int64

func IsPrime(n int64) bool {
	if n < int64(len(isPrime)) {
		return isPrime[n]
	}

	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := int64(0)
	for primes[i]*primes[i] <= n {
		if n%primes[i] == 0 {
			return false
		}
		i += 1
	}
	return true
}

func DigitSum(n int64) int64 {
	return sliceutil.SumInt64(mathutil.ToDigits(n, 10))
}

func IsHarshad(n int64) bool {
	if n == int64(0) {
		return true
	}
	return n%DigitSum(n) == 0
}

func IsStrongHarshad(n int64) bool {
	if n == int64(0) {
		return false
	}
	return IsPrime(n / DigitSum(n))
}

func ExpandRight(in []int64) []int64 {
	out := []int64{}
	for _, num := range in {
		for i := int64(0); i < int64(10); i++ {
			expandedNum := 10*num + i
			if IsHarshad(expandedNum) {
				out = append(out, expandedNum)
			}
		}
	}
	return out
}

func ExpandPrimeSum(left int64) int64 {
	sum := int64(0)
	if IsPrime(left*10 + 1) {
		sum += left*10 + 1
	}
	if IsPrime(left*10 + 3) {
		sum += left*10 + 3
	}
	if IsPrime(left*10 + 7) {
		sum += left*10 + 7
	}
	if IsPrime(left*10 + 9) {
		sum += left*10 + 9
	}
	return sum
}

func main() {
	isPrime = primegen.Map(10000000) // 10^7 = sqrt(10^14)
	primes = primegen.SliceFromMap(isPrime)

	N := 14

	// Find all Harshad numbers with N or fewer digits
	rightTruncatable := make([][]int64, N+1)
	rightTruncatable[1] = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; i < N-1; i++ {
		rightTruncatable[i+1] = ExpandRight(rightTruncatable[i])
	}

	// Keep only the strong ones
	strongTruncatable := []int64{}
	for _, num := range rightTruncatable[N-1] {
		if IsStrongHarshad(num) {
			strongTruncatable = append(strongTruncatable, num)
		}
	}

	// Add 1, 3, 7, 9 to the end of all these numbers and check if they are prime.
	// If so, add them to the sum.
	sum := int64(0)
	for _, num := range strongTruncatable {
		sum += ExpandPrimeSum(num)
	}

	fmt.Printf("%v\n", sum)
}
