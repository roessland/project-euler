package main

import "fmt"

func main() {
	num := int64(0)
	for n := int64(1); n <= int64(1<<30); n++ {
		if (n)^(2*n)^(3*n) == 0 {
			num++
		}
	}
	fmt.Printf("%v\n", num)
}
