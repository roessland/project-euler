package main

import "fmt"
import "github.com/roessland/gopkg/mathutil"

func main() {
	p := int64(1009)
	q := int64(3643)
	phi := (p - 1) * (q - 1)

	sum := int64(0)
	for e := int64(2); e < phi; e++ {
		if mathutil.GCD(e, phi) != 1 {
			continue
		}
		numUnconcealed := (mathutil.GCD(e-1, p-1) + 1) * (mathutil.GCD(e-1, q-1) + 1)
		if numUnconcealed == 9 {
			sum += e
		}
	}
	fmt.Printf("%v\n", sum)
}
