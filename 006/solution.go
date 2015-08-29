package main

import "fmt"

func main() {
	S2 := 0
	S := 0

	for i := 1; i <= 100; i++ {
		S2 += i * i
		S += i
	}
	fmt.Printf("%v\n", S2-S*S)
}
