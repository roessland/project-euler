package main

import "fmt"
import "time"

// 25 primes are less than 100
var p []int64 = []int64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
}

func main() {
	t0 := float64(time.Now().UnixNano()) / 1000000.0

	// Loop trough all numbers
	var maxNum int64 = 10000000000
	var intervalSize int64 = 10000000 // Must divide 10^6
	var numIntervals int64 = maxNum / intervalSize
	numFactors := make([]int64, intervalSize)

	var intervalStart int64 = 0
	var intervalEnd int64 = intervalSize
	var numPqrst int64 = 0
	for i := int64(0); i < numIntervals; i++ {
		// Clear sieve for this range
		for j := int64(0); j < intervalSize; j++ {
			numFactors[j] = 0
		}

		// Increment sieve with multiples of small primes
		for j := 0; j < len(p); j++ {
			t0 := intervalStart / p[j]
			t1 := intervalEnd/p[j] + 1

			if t0 == 0 {
				t0 = 1
			}
			for t := t0; t < t1; t++ {
				pt := p[j] * t
				if intervalStart <= pt && pt < intervalEnd {
					numFactors[pt-intervalStart]++
				}
			}
		}

		// Sum up the results for this interval
		for j := int64(0); j < intervalSize; j++ {
			if numFactors[j] >= 4 {
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
