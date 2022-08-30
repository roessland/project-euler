package main

import (
	"fmt"
)

const mod = 1000000

func dp() int {

	type row struct {
		pals, palSum, twoPals, twoPalSum int
	}

	r6 := row{0, 0, 0, 0}
	r5 := row{0, 0, 0, 0}
	r4 := row{0, 0, 0, 0}
	r3 := row{1, 1, 0, 0}
	r2 := row{1, 1, 0, 0}
	r1 := row{2, 3, 1, 1}

	for i := 2; ; {
		rn := row{}
		rn.pals = 1 + r2.palSum
		rn.palSum = (rn.pals + r2.palSum) % mod
		rn.twoPals = (r2.twoPals + r4.pals + r6.twoPalSum) % mod
		rn.twoPalSum = (rn.twoPals + r2.twoPalSum) % mod

		i, r6, r5, r4, r3, r2, r1 = i+1, r5, r4, r3, r2, r1, rn

		if rn.twoPals%mod == 0 && i > 42 {
			return rn.twoPals
		}
	}
}

func main() {
	fmt.Println(dp())
}
