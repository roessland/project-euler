package main

import "fmt"
import "github.com/roessland/gopkg/iterutil"
import "github.com/roessland/gopkg/sliceutil"

// 0.5729 <-- somewhere around there

func main() {

	// Create result frequency tables for Colin and Peter
	freqsC := make([]int, 37)
	freqsP := make([]int, 37)
	for i := range iterutil.CartesianPower(6, 6) {
		sum := 6 + sliceutil.SumInt64(i)
		freqsC[sum]++
	}
	for i := range iterutil.CartesianPower(4, 9) {
		sum := 9 + sliceutil.SumInt64(i)
		freqsP[sum]++
	}

	// Create probability distributions for Colin and Peter
	probsC := make([]float64, 37)
	probsP := make([]float64, 37)
	cOutcomes := float64(6 * 6 * 6 * 6 * 6 * 6)
	pOutcomes := float64(4 * 4 * 4 * 4 * 4 * 4 * 4 * 4 * 4)
	for i, _ := range freqsC {
		probsC[i] = float64(freqsC[i]) / cOutcomes
		probsP[i] = float64(freqsP[i]) / pOutcomes
	}

	// Calculate P(P > C) = Sum[tot=0..36] P(C < P|P=tot) P(P=tot)
	sum := float64(0)
	for totP := 0; totP <= 36; totP++ {
		for totC := 0; totC < totP; totC++ {
			sum += probsC[totC] * probsP[totP]
		}
	}

	fmt.Printf("%v\n", sum)
}
