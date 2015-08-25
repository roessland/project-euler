package main

import "github.com/roessland/gopkg/iterutil"
import "github.com/roessland/gopkg/sliceutil"
import "fmt"

func main() {
	p := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89, 97}
	P := make([]int64, 12650)
	pqrs := make([][]int64, 12650)
	pos := 0
	for i := range iterutil.Subsets(25, 4) {
		P[pos] = p[i[0]] * p[i[1]] * p[i[2]] * p[i[3]]
		pqrs[pos] = make([]int64, 4)
		for j := 0; j < 4; j++ {
			pqrs[pos][j] = p[i[j]]
		}
		pos++
	}

	N := int64(1000)
	num := int64(0)
	for i := 0; i < len(P); i++ {
		fmt.Printf("%v%%\n", float32(i)/12650.0*100.0)
		num += (N - 1) / P[i]
		uniqueFactors := []int64{}
		//LCMoverflow := false
		for j := 0; j < i; j++ {
			fmt.Printf("j=%v\n", j)
			uniqueFactors := sliceutil.Union(uniqueFactors, pqrs[j])
			fmt.Printf("%v\n", uniqueFactors)
			LCM := sliceutil.ProductInt64(uniqueFactors)
			if num < 0 {
				panic("num wtf")
			}
			if LCM > 100000000000000000 || LCM < 0 {
				panic("LCM overflow")
			}
			num = num - ((N - 1) / LCM)
		}
	}
	fmt.Printf("habeeb it %v, %v\n", N, num)

}
