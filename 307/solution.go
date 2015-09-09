// Runs in 3 minutes and 30 seconds on a Xeon X7542 @ 2.67 GHz
// Needs more than 16 GB of memory! I ran it on a computer with 256 GB memory.
package main

import "fmt"
import "math/rand"

var k int32 = 20000
var n int32 = 1000000

type Args struct {
	ballsLeft, zeroBins, oneBins int32
}

var cache map[Args]float64

// Estimate(20000, 1000000) = 0.7267
func Estimate(k, n int) {
	numTrials := 10000
	numTrialsWithAtLeast3 := 0

	for trial := 0; trial < numTrials; trial++ {
		bins := make([]int, n)
		for i := 0; i < k; i++ {
			binIndex := rand.Int31n(int32(n))
			bins[binIndex]++
			if bins[binIndex] == 3 {
				numTrialsWithAtLeast3++
				break
			}
		}
	}
	P := float32(numTrialsWithAtLeast3) / float32(numTrials)
	fmt.Printf("Estimate for %v balls and %v bins = %v\n", k, n, P)
}

// The probability that adding ballsLeft balls does not end up with a
// configuration where there are three or more balls in a single bin, given
// that zeroBins are bins with zero balls in, oneBins are bins with one ball
// in.
func P(ballsLeft, zeroBins, oneBins int32) float64 {
	args := Args{ballsLeft, zeroBins, oneBins}
	p, ok := cache[args]
	if ok {
		return p
	}
	if ballsLeft == 0 {
		return float64(1)
	}
	pZero := float64(zeroBins) / float64(n)
	pOne := float64(oneBins) / float64(n)
	prob := P(ballsLeft-1, zeroBins-1, oneBins+1)*pZero + P(ballsLeft-1, zeroBins, oneBins-1)*pOne
	cache[args] = prob
	return prob
}

func main() {
	cache = make(map[Args]float64)
	fmt.Printf("Running with k = %v\n", k)
	fmt.Printf("%v\n", 1-P(k, n, 0))
}
