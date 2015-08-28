package main

import "math/big"
import "strconv"
import "fmt"

func main() {
	e := []int{2}
	for i := 1; i < 100; i++ {
		e = append(e, 1, 2*i, 1)
	}

	N := 100
	f := big.NewRat(int64(e[N-1]), 1)
	for i := N - 1; i > 0; i-- {
		f.Inv(f).Add(f, big.NewRat(int64(e[i-1]), 1))
	}

	fmt.Printf("%v\n", f)

	digitSum := 0
	digitStr := f.Num().String()
	for _, digit := range digitStr {
		i, _ := strconv.Atoi(string(digit))
		digitSum += i
	}

	fmt.Printf("Sum of digits is %v\n", digitSum)
}
