package main

import "fmt"
import "time"
import "github.com/roessland/gopkg/iterutil"
import "github.com/roessland/gopkg/sliceutil"

// 25 primes are less than 100
var primes []int64 = []int64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
}

// Products of four distinct primes less than 100
var pqrs []int64

func main() {
	pqrs := make([]int64, 12650) // Products of 4 distinct primes
	pos := 0
	for i := range iterutil.Subsets(25, 4) {
		pqrs[pos] = primes[i[0]] * primes[i[1]] * primes[i[2]] * primes[i[3]]
		pos++
	}
	sliceutil.SortInt64(pqrs)

	t0 := float64(time.Now().UnixNano()) / 1000000.0

	// Loop trough all numbers
	var maxNum int64 = 10000000
	var intervalSize int64 = 10000 // Must divide 10^6
	var numIntervals int64 = maxNum / intervalSize
	isPqrst := make([]bool, intervalSize)

	var intervalStart int64 = 0
	var intervalEnd int64 = intervalSize
	var numPqrst int64 = 0
	for i := int64(0); i < numIntervals; i++ {
		// Clear sieve for this range
		for j := int64(0); j < intervalSize; j++ {
			isPqrst[j] = false
		}

		// Fill sieve with multiples of pqrs
		for j := 0; j < len(pqrs); j++ {
			t0 := intervalStart / pqrs[j]
			t1 := intervalEnd/pqrs[j] + 1

			if t0 == 0 {
				t0 = 1
			}
			for t := t0; t < t1; t++ {
				pqrst := pqrs[j] * t
				if intervalStart <= pqrst && pqrst < intervalEnd {
					isPqrst[pqrst-intervalStart] = true
				}
			}
		}

		// Sum up the results for this interval
		for j := int64(0); j < intervalSize; j++ {
			if isPqrst[j] {
				numPqrst++
			}
		}

		intervalStart += intervalSize
		intervalEnd += intervalSize
	}

	t1 := float64(time.Now().UnixNano()) / 1000000.0
	fmt.Printf("time = %vms\n", t1-t0)
	fmt.Printf("N = %v\n", numPqrst)
}
