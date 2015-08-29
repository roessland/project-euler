package main

import "fmt"
import "math"

func main() {
	p := int64(600851475143)
	for i := int64(2); i < int64(math.Sqrt(float64(p))); i++ {
		if p%i == 0 {
			// Check if it is prime
			isPrime := true
			for j := int64(2); j < int64(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isPrime = false
				}
			}
			if isPrime {
				fmt.Printf("%v\n", i)
			}
		}
	}
}
