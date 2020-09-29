package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/roessland/gopkg/mathutil"
)

func S(nSqrt int) bool {
	if nSqrt%9 > 1 {
		return false
	}
	n := nSqrt * nSqrt
	nDigits := mathutil.ToDigitsInt(n, 10)
	nPartitions := mathutil.SlicePartitions(nDigits)
	for _, nPartition := range nPartitions {
		if len(nPartition) == 1 {
			continue // must be split into 2 or more partitions
		}
		partitionSum := 0
		for _, part := range nPartition {
			partitionSum += mathutil.FromDigitsInt(part, 10)
			if partitionSum > nSqrt {
				break
			}
		}
		if partitionSum == nSqrt {
			return true
		}
	}
	return false
}

func T(N int) int {
	nSqrtChan := make(chan int)
	sumChan := make(chan int)

	var wg sync.WaitGroup
	nThreads := runtime.NumCPU()
	wg.Add(nThreads)

	for threadId := 0; threadId < nThreads; threadId++ {
		go func() {
			for nSqrt := range nSqrtChan {
				if S(nSqrt) {
					sumChan <- nSqrt * nSqrt
				}
			}
			wg.Done()
		}()
	}

	go func() {
		i := 0
		for nSqrt := 0; nSqrt*nSqrt <= N; nSqrt++ {
			if i%1000 == 0 {
				fmt.Println("Progress: ", 100*float64(nSqrt*nSqrt)/(float64(N)), "%")
			}
			nSqrtChan <- nSqrt

			i++
		}
		close(nSqrtChan)
		wg.Wait()
		close(sumChan)
	}()

	sum := 0
	for delta := range sumChan {
		sum += delta
	}

	return sum
}

func main() {

	start := time.Now()

	fmt.Println("T(10000) =", T(10000)) // 41333
	fmt.Println("T(1000000000000) =", T(1000000000000))

	end := time.Now()

	fmt.Println("Used", end.Sub(start), "secs")
}
