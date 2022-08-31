package main

import (
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"github.com/roessland/gopkg/primegen"
	"os"
	"sort"
)

const N = 12
const Mod = 500500507

type Val struct {
	mul  int64
	idx  int
	cost int
}

var counter int

func find(ps []int, as []int, sum int, prod int64) {
	counter++
	if counter%10000 == 0 {
		fmt.Printf(".")
	}
	if sum < 0 {
		return
	}
	if sum == 0 {
		fmt.Println("Boom!", prod, as)
		// os.Exit(0)
		return
	}

	vals := []Val{}

	// Increase existing
	var idx = 0
	for ; idx < len(as) && as[idx] != 0; idx++ {
		a := as[idx]
		aPrime := 2*(a+1) - 1
		if idx == 0 || as[idx-1] >= aPrime {
			vals = append(vals, Val{
				mul:  mathutil.Pow(int64(ps[idx]), int64(aPrime-a)),
				idx:  idx,
				cost: a + 1,
			})
		}
	}

	// Add new
	vals = append(vals, Val{
		mul:  int64(ps[idx]),
		idx:  idx,
		cost: 1,
	})

	// Sort
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].mul > vals[j].mul
	})

	// BFS
	for _, val := range vals {
		as[val.idx] += val.cost
		find(ps, as, sum-val.cost, (prod * val.mul))
		//find(ps, as, sum-val.cost, (prod*val.mul)%Mod)
		as[val.idx] -= val.cost
	}
}

func main() {
	primes := primegen.SliceFromMapInt(primegen.MapInt(20 * N))
	if len(primes) < N+1 {
		fmt.Println("need more primes")
		os.Exit(1)
	}

	init := make([]int, 2*N)
	init[0] = 1
	find(primes, init, N, 2)

}
