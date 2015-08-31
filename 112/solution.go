// Runs in 30 ms on i7-4770k CPU
package main

import "fmt"

func IsBouncy(n int) bool {
	var wentUp, wentDown bool
	q := n / 10
	d := n - 10*q
	n = q
	prevDigit := d
	for n > 0 {
		q := n / 10
		d := n - 10*q
		n = q

		if d > prevDigit {
			wentUp = true
			if wentDown {
				return true
			}
		}
		if d < prevDigit {
			wentDown = true
			if wentUp {
				return true
			}
		}
		prevDigit = d
	}
	return false
}

func main() {
	numBouncy := 0
	numTotal := 0

	for n := 1; ; n++ {
		numTotal++
		if IsBouncy(n) {
			numBouncy++
		}
		if float64(numBouncy)/float64(numTotal) >= float64(0.99) {
			fmt.Printf("%v\n", numTotal)
			break
		}
	}
}
