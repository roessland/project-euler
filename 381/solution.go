package main

import "fmt"
import "github.com/roessland/gopkg/primegen"
import "github.com/roessland/gopkg/mathutil"

func S(p int64) int64 {
	// p1 = (p-1)! = p-1                                             (mod p)
	// p2 = (p-2)! = 1                                               (mod p)
	// p3 = (p-3)! = inv(p-2)                                        (mod p)
	// p4 = (p-4)! = inv(p-3) * inv(p-2)            =  inv(p-3) * p3 (mod p)
	// p5 = (p-5)! = inv(p-4) * inv(p-3) * inv(p-2) =  inv(p-4) * p4 (mod p)
	var p3 int64 = mathutil.ModularInverse(p-2, p)
	var p4 int64 = (mathutil.ModularInverse(p-3, p) * p3) % p
	var p5 int64 = (mathutil.ModularInverse(p-4, p) * p4) % p
	return (p3 + p4 + p5) % p
}

func main() {
	limit := int64(100000000)
	isPrime := primegen.Map(limit)

	sumS := int64(0)
	for p := int64(5); p < limit; p++ {
		if isPrime[p] {
			sumS += S(p)
		}
	}
	fmt.Printf("SumS up to %v is %v\n", limit, sumS)
}
