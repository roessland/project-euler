package main

import "fmt"
import "github.com/roessland/gopkg/mathutil"

func main() {
	fmt.Printf("%v\n", 7*(1-float64(mathutil.Choose(60, 20))/float64(mathutil.Choose(70, 20))))
}
