package main

import "fmt"

func main() {
	for c := 0; c < 1000; c++ {
		for b := 0; b < c; b++ {
			a := 1000 - b - c
			if a < 0 {
				continue
			}
			if a*a+b*b == c*c {
				fmt.Printf("%v\n", a*b*c)
			}
		}
	}
}
