package main

import "fmt"
import "math"
import "github.com/roessland/gopkg/mathutil"
import "github.com/roessland/gopkg/primegen"

var modPrime int64 = 1000000007 // Numbers are modded to this prime
var minPrime int64 = 10000000   // Sum G(p) for minPrime < p < maxPrime
var maxPrime int64 = 10010000
var factorialMod []int64
var inverseFactorialMod []int64

func BuildFactorialTables() {
	factorialMod = make([]int64, maxPrime+1)
	inverseFactorialMod = make([]int64, maxPrime+1)

	var prod int64 = 1
	var invProd int64 = 1
	for i := int64(1); i <= maxPrime; i++ {
		prod = (prod * i) % modPrime
		invProd = (invProd * mathutil.ModularInverse(i, modPrime)) % modPrime
		factorialMod[i] = prod
		inverseFactorialMod[i] = invProd
	}
}

func ChooseModPrime(n, k int64) int64 {
	return (((factorialMod[n] * inverseFactorialMod[k]) % modPrime) * inverseFactorialMod[n-k]) % modPrime
}

func G(a int64) int64 {
	m := int64(math.Ceil(math.Sqrt(float64(a))))
	S := int64(1)
	c := int64(0)
	b := int64(math.Ceil((float64(c) + 1.0) * math.Sqrt(float64(a))))
	c = 1

	for c := int64(1); c < m; c++ {
		//S = (S + mathutil.ChooseMod(c+a-b, c, modPrime)) % modPrime
		S = (S + ChooseModPrime(c+a-b, c)) % modPrime

		// Find the next special point by going down until sign changes
		b = int64(math.Ceil((float64(c) + 1.0) * math.Sqrt(float64(a))))
	}

	return S
}

func main() {
	BuildFactorialTables()
	fmt.Printf("Tables generated")

	isPrime := primegen.Map(maxPrime)
	fmt.Printf("Primes generated")

	var S int64 = 0
	var numPrimes int64 = 0
	for i := minPrime + 1; i < maxPrime; i++ {
		if isPrime[i] {
			numPrimes++
			S = (S + G(i)) % modPrime
			fmt.Printf(".")
		}
	}
	fmt.Printf("The final result is %v\n", S)
	fmt.Printf("Number of primes: %v\n", numPrimes)
}
