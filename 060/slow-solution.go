// Slow solution.
// Works for example case with 4 primes, but too slow for real solution.
package main

import "fmt"
import "github.com/roessland/gopkg/primegen"
import "github.com/roessland/gopkg/mathutil"

var isPrime []bool

func C(a, b int64) bool {
	if isPrime[mathutil.FromDigits(append(mathutil.ToDigits(a, 10), mathutil.ToDigits(b, 10)...), 10)] && isPrime[mathutil.FromDigits(append(mathutil.ToDigits(b, 10), mathutil.ToDigits(a, 10)...), 10)] {
		return true
	}
	return false
}

func AwesomePrimes(primes []int64) bool {
	a, b, c, d, e := primes[0], primes[1], primes[2], primes[3], primes[4]

	if C(a, b) && C(a, c) && C(a, d) && C(a, e) && C(b, c) && C(b, d) && C(b, e) && C(c, d) && C(c, e) && C(d, e) {
		return true
	} else {
		return false
	}
}

func AwesomePrimes4(primes []int64) bool {
	a, b, c, d := primes[0], primes[1], primes[2], primes[3]

	if C(a, b) && C(a, c) && C(a, d) && C(b, c) && C(b, d) && C(c, d) {
		return true
	} else {
		return false
	}
}

func main() {
	// For checking if large numbers are prime
	isPrime = primegen.Map(100000000)

	// For looping trough all combinations of
	smallPrimes := primegen.SliceFromMap(primegen.Map(999))

	// Loop trough possibilites by increasing sum
	for sum := int64(1); sum < 1000; sum++ {
		if sum%10 == 0 {
			fmt.Printf("%v\n", sum)
		}
		setsOfFive := mathutil.Partition(smallPrimes, sum, 5, 0)

		// Check if any of the possibilites satisfy the critera
		for _, setOfFive := range setsOfFive {
			if AwesomePrimes(setOfFive) {
				fmt.Printf("sum=%v, primes=%v\n", sum, setOfFive)
				return
			}
		}
	}
}
