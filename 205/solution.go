package main

import "fmt"
import "github.com/roessland/gopkg/iterutil"

// 0.5729 <-- somewhere around there

func main() {
	// Dice values
	d4 := []int{1, 2, 3, 4}       // Peter
	d6 := []int{1, 2, 3, 4, 5, 6} // Colin

	// Create result frequency tables for Colin and Peter
	freqsC := make([]int, 37)
	for i := range iterutil.CartesianPower(6, 6) {
		sum := d6[i[0]] + d6[i[1]] + d6[i[2]] + d6[i[3]] + d6[i[4]] + d6[i[5]]
		freqsC[sum]++
	}
	freqsP := make([]int, 37)
	for i := range iterutil.CartesianPower(4, 9) {
		sum := d4[i[0]] + d4[i[1]] + d4[i[2]] + d4[i[3]] + d4[i[4]] + d4[i[5]] + d4[i[6]] + d4[i[7]] + d4[i[8]]
		freqsP[sum]++
	}

	// Create probability distributions for Colin and Peter
	probsC := make([]float64, 37)
	cOutcomes := float64(6 * 6 * 6 * 6 * 6 * 6)
	cumSumC := float64(0)
	for i, _ := range freqsC {
		probsC[i] = float64(freqsC[i]) / cOutcomes
		cumSumC += probsC[i]
	}
	fmt.Printf("Colin's distribution sums to %v\n", cumSumC)

	// Create probability distribution for Peter
	probsP := make([]float64, 37)
	pOutcomes := float64(4 * 4 * 4 * 4 * 4 * 4 * 4 * 4 * 4)
	cumSumP := float64(0)
	for i, _ := range freqsP {
		probsP[i] = float64(freqsP[i]) / pOutcomes
		cumSumP += probsP[i]
	}
	fmt.Printf("Peter's distribution sums to %v\n", cumSumP)

	// Calculate P(P > C) = Sum[tot=0..36] P(C < P|P=tot) P(P=tot)
	sum := float64(0)
	for totP := 0; totP <= 36; totP++ {
		for totC := 0; totC < totP; totC++ {
			sum += probsC[totC] * probsP[totP]
		}
	}

	fmt.Printf("%v\n", sum)

}
