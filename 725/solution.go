package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/roessland/gopkg/mathutil"
)

const modulus = 10000000000000000

func S(n int) *big.Int {
	sum := big.NewInt(0)
	for d := 1; d <= 9; d++ {
		sum.Add(sum, f(n, d))
	}
	return sum
}

func multinomialMod(ms mathutil.MultisetIntInt) *big.Int {
	n := 0
	kNonZero := 0
	y := big.NewInt(1)
	for part, multiplicity := range ms {
		n += multiplicity
		if part == 0 {
			continue
		}
		kNonZero += multiplicity
		multFact := mathutil.Factorial(int64(multiplicity))
		y.Mul(y, big.NewInt(multFact))
	}
	x := big.NewInt(1)
	x = x.MulRange(int64(n-kNonZero+1), int64(n))
	x.Quo(x, y)
	return x
}

func f(n int, d int) *big.Int {
	partitions := mathutil.PartitionsInt(d)
	totalReps := big.NewInt(0)
	for _, partition := range partitions {
		if len(partition) >= n {
			continue
		}
		partition = append(partition, d)
		partitionMs := mathutil.PartitionIntToMultisetIntInt(partition)
		partitionMs[0] += n - len(partition)
		reps := multinomialMod(partitionMs)
		totalReps.Add(totalReps, reps)
		partitionSum := big.NewInt(1)
		partitionSum.Mul(partitionSum, big.NewInt(int64(2*d)))
		partitionSum.Mul(partitionSum, reps)
	}
	digitSum := int64(2 * d)
	oneoneone := big.NewInt(0)
	oneoneone.SetString(strings.Repeat("1", n), 10)
	res := big.NewInt(1)
	res.Mul(res, totalReps)
	res.Mul(res, big.NewInt(digitSum))
	res.Mul(res, oneoneone)
	res.Div(res, big.NewInt(int64(n)))
	return res
}

func main() {
	var s, smod *big.Int

	e16 := big.NewInt(0)
	e16.SetString("10000000000000000", 10)

	fmt.Println("S(3) is", S(3))

	fmt.Println("S(7) is", S(7))

	s = S(2020)
	smod = big.NewInt(0)
	smod.Set(s)
	smod.Mod(smod, e16)

	fmt.Println("S(2020) is", s, "and S(2020) mod 1e16 is", smod)
}
