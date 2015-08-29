package main

import "strconv"
import "fmt"

func IsPalindrome(n int) bool {
	str := strconv.Itoa(n)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	largest := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			ij := i * j
			if IsPalindrome(ij) && ij > largest {
				largest = ij
			}
		}
	}
	fmt.Printf("%v\n", largest)
}
