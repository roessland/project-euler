// Runs in 6 seconds on an i7-4770k
package main

import "fmt"

func main() {
	sum := uint64(0)
	for a := uint64(3); a <= 1000; a++ {
		a2 := a * a
		aPrev := uint64(1)
		aNext := uint64(1)

		rMax := uint64(0)
		for n := uint64(0); n < a2; n++ {
			r := (aPrev + aNext) % a2
			if r > rMax {
				rMax = r
			}

			aPrev = (aPrev * (a - 1)) % a2
			aNext = (aNext * (a + 1)) % a2
		}
		sum += rMax
	}
	fmt.Printf("sum = %v\n", sum)
}
