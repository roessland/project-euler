// 60 ms on a nice CPU -- 2015
// The most interesting part here is the keys used in the hashmap
package main

import "fmt"

var m [][]int = [][]int{
	[]int{7, 53, 183, 439, 863, 497, 383, 563, 79, 973, 287, 63, 343, 169, 583},
	[]int{627, 343, 773, 959, 943, 767, 473, 103, 699, 303, 957, 703, 583, 639, 913},
	[]int{447, 283, 463, 29, 23, 487, 463, 993, 119, 883, 327, 493, 423, 159, 743},
	[]int{217, 623, 3, 399, 853, 407, 103, 983, 89, 463, 290, 516, 212, 462, 350},
	[]int{960, 376, 682, 962, 300, 780, 486, 502, 912, 800, 250, 346, 172, 812, 350},
	[]int{870, 456, 192, 162, 593, 473, 915, 45, 989, 873, 823, 965, 425, 329, 803},
	[]int{973, 965, 905, 919, 133, 673, 665, 235, 509, 613, 673, 815, 165, 992, 326},
	[]int{322, 148, 972, 962, 286, 255, 941, 541, 265, 323, 925, 281, 601, 95, 973},
	[]int{445, 721, 11, 525, 473, 65, 511, 164, 138, 672, 18, 428, 154, 448, 848},
	[]int{414, 456, 310, 312, 798, 104, 566, 520, 302, 248, 694, 976, 430, 392, 198},
	[]int{184, 829, 373, 181, 631, 101, 969, 613, 840, 740, 778, 458, 284, 760, 390},
	[]int{821, 461, 843, 513, 17, 901, 711, 993, 293, 157, 274, 94, 192, 156, 574},
	[]int{34, 124, 4, 878, 450, 476, 712, 914, 838, 669, 875, 299, 823, 329, 699},
	[]int{815, 559, 813, 459, 522, 788, 168, 586, 966, 232, 308, 833, 251, 631, 107},
	[]int{813, 883, 451, 509, 615, 77, 281, 613, 459, 205, 380, 274, 302, 35, 805},
}

func Max(n []int) int {
	max := 0
	for _, num := range n {
		if num > max {
			max = num
		}
	}
	return max
}

func Print(M [][]int, i, j []int) {
	fmt.Printf("--------\n")
	for t, ik := range i {
		fmt.Printf(" [ ")
		for _, jk := range j {
			fmt.Printf("%3d ", M[ik][jk])
		}
		fmt.Printf(" ] ")
		if t == 0 {
			fmt.Printf("i: %v, j: %v", i, j)
		}
		fmt.Printf("\n")
	}
}

var cache map[Val]int

type Val struct {
	X uint
	Y uint
}

func GetVal(i, j []int) Val {
	var x, y uint
	for k := 0; k < len(i); k++ {
		x += 1 << uint(i[k])
		y += 1 << uint(j[k])
	}
	return Val{x, y}
}

func Sum(M [][]int, i, j []int) int {

	n := len(i)
	if len(i) != len(j) {
		panic("unequal i and j arrays")
	}

	if n == 0 {
		panic("wtf dood")
	}

	val := GetVal(i, j)
	cachedSum, ok := cache[val]
	if ok {
		return cachedSum
	}

	if n == 1 {
		return M[i[0]][j[0]]
	}
	sums := make([]int, n)
	for k := 0; k < n; k++ {
		var iNext []int
		iNext = append(append(iNext, i[:k]...), i[k+1:]...)
		sums[k] = M[i[k]][j[0]] + Sum(M, iNext, j[1:])
	}
	max := Max(sums)
	cache[val] = max
	return max
}

func main() {
	cache = make(map[Val]int)
	i := make([]int, 15)
	for k := 0; k < 15; k++ {
		i[k] = k
	}
	fmt.Printf("%v\n", Sum(m, i, i))
}
