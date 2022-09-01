package main

import (
	"container/heap"
	"fmt"
	"github.com/roessland/gopkg/primegen"
	"os"
	"sort"
)

const N = 500500
const Mod = 500500507

type PrimeFactor struct {
	Prime int64
}

type PrimeFactors []PrimeFactor

func (pfs *PrimeFactors) Len() int {
	return len(*pfs)
}

func (pfs *PrimeFactors) Less(i, j int) bool {
	return (*pfs)[i].Prime < (*pfs)[j].Prime
}

func (pfs *PrimeFactors) Swap(i, j int) {
	(*pfs)[i], (*pfs)[j] = (*pfs)[j], (*pfs)[i]
}

func (pfs *PrimeFactors) Push(x any) {
	*pfs = append(*pfs, x.(PrimeFactor))
}

func (pfs *PrimeFactors) Pop() any {
	last := (*pfs)[len(*pfs)-1]
	*pfs = (*pfs)[:len(*pfs)-1]
	return last
}

func find(ps []int64, N int64) {
	var num int64 = 1
	var divisors int64 = 0
	pfs := make(PrimeFactors, len(ps))
	for i, p := range ps {
		pfs[i] = PrimeFactor{Prime: p}
	}

	heap.Init(&pfs)

	for divisors < N {
		pf := heap.Pop(&pfs).(PrimeFactor)
		num = (num * pf.Prime) % Mod
		pf.Prime *= pf.Prime
		heap.Push(&pfs, pf)
		divisors += 1
	}

	sort.Stable(&pfs)
	fmt.Println(num)
}

func main() {
	primes := primegen.SliceFromMap(primegen.Map(20 * N))
	if len(primes) < N+1 {
		fmt.Println("need more primes")
		os.Exit(1)
	}

	find(primes, N)
}
