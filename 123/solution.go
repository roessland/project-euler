// Runs in 150 ms on an i7-4770k
package main

import "fmt"
import "github.com/roessland/gopkg/primegen"

import mu "github.com/cznic/mathutil"

func main() {
	isPrime := primegen.Map(1000000)
	p := append([]int64{0}, primegen.SliceFromMap(isPrime)...)

	for n := uint64(1); ; n++ {
		p2 := uint64(p[n] * p[n])
		r := (mu.ModPowUint64(uint64(p[n]-1), n, p2) + mu.ModPowUint64(uint64(p[n]+1), n, p2)) % p2
		if r > 10000000000 {
			fmt.Printf("%v\n", n)
			break
		}
	}
}
