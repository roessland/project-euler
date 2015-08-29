package main

import "fmt"
import "math"

func main() {
	found := 0
	for i := 2; ; i++ {
		isPrime := true
		for d := 2; d <= int(math.Sqrt(float64(i))); d++ {
			if i%d == 0 {
				isPrime = false
			}
		}
		if isPrime {
			found++
			if found == 10001 {
				fmt.Printf("%v\n", i)
				break
			}
		}
	}
}
