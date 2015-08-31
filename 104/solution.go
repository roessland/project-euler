// i7-4770
// real    0m15.898s
// user    0m16.601s
// sys     0m0.122s

package main

import "fmt"
import "math/big"

func IsPandigital(num string) bool {
	found := make(map[rune]bool)
	for _, d := range num {
		if d == '0' {
			return false
		}
		found[d] = true
	}
	return len(found) == 9
}

func main() {
	lastDigits := big.NewInt(1337)
	mod := big.NewInt(1000000000)
	Fprev := big.NewInt(1)
	Fcurr := big.NewInt(1)
	Fnext := big.NewInt(35)

	n := 2
	for {
		Fnext.Add(Fprev, Fcurr)
		Fprev.Add(big.NewInt(0), Fcurr)
		Fcurr.Add(big.NewInt(0), Fnext)
		n++

		lastDigits.Mod(Fcurr, mod)
		if IsPandigital(lastDigits.String()) {
			fmt.Printf("F(%v) is pandigital in the final digits.\n", n)
			if IsPandigital(Fcurr.String()[0:9]) {
				fmt.Printf("F(%v) is also pandigital in the first digits!\n", n)
				break
			}
		}
	}

}
