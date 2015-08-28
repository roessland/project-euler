package main

import "fmt"

func Multiplicity(i, j, k int) int64 {
	ret := int64(6)
	if i == j || i == k || j == k {
		ret = 3
	}
	if i == j && j == k {
		ret = 1
	}
	return ret
}

func TwosInFactorial(n int) int {
	numTwos := 0
	for k := n / 2; k != 0; k /= 2 {
		numTwos += k
	}
	return numTwos
}

func FivesInFactorial(n int) int {
	numFives := 0
	for k := n / 5; k != 0; k /= 5 {
		numFives += k
	}
	return numFives
}

func X(i, j, k int) bool {
	n := i + j + k
	twos := TwosInFactorial(n) - TwosInFactorial(i) - TwosInFactorial(j) - TwosInFactorial(k)
	fives := FivesInFactorial(n) - FivesInFactorial(i) - FivesInFactorial(j) - FivesInFactorial(k)

	if twos > fives {
		return fives >= 12
	} else {
		return twos >= 12
	}
}

func main() {
	n := 200000

	sum := int64(0)
	for i := 0; i <= n; i++ {
		if i%10000 == 0 {
			fmt.Printf("%v\n", (100*i)/n)
		}
		for j := (n - i) / 2; j <= i; j++ {
			k := n - i - j
			if k < 0 || k > j {
				continue
			}
			if X(i, j, k) {
				sum += Multiplicity(i, j, k)
			}
		}
	}
	fmt.Printf("%v\n", sum)
}
