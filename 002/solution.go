package main

import "fmt"

func main() {
	prev := int64(1)
	curr := int64(1)
	sum := int64(0)
	for {
		prev, curr = curr, prev+curr

		if curr > int64(4000000) {
			break
		}
		if curr%2 == 0 {
			sum += curr
		}
	}
	fmt.Printf("%v\n", sum)

}
