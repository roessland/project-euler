package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	sum := 0
	for d := 2; d <= 12000; d++ {
		for n := d/3 + 1; n <= d/2; n++ {
			if 2*n == d {
				continue
			}
			if GCD(n, d) == 1 {
				sum++
			}
		}
	}
	fmt.Printf("\n%v\n", sum)
}
