package main

import (
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"time"
)

func S(nSqrt int) bool {
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

	start := time.Now()

	nSqrtChan := make(chan int)
	sumChan := make(chan int)
	doneChan := make(chan struct{})

	go func() {
		i := 0
		for nSqrt := 0; nSqrt*nSqrt <= N; nSqrt++ {
			if i % 1000 == 0 {
				fmt.Println("Progress: ", 100*float64(nSqrt*nSqrt)/(float64(N)), "%")
			}
			nSqrtChan <- nSqrt

			i++
		}
		time.Sleep(2*time.Second)
		doneChan <- struct{}{}
	}()

	nThreads := 12
	for threadId := 0; threadId < nThreads; threadId++ {
		go func() {
			for nSqrt := range nSqrtChan {
				if S(nSqrt) {
					sumChan <- nSqrt*nSqrt
				}
			}
		}()
	}

	go func() {
		sum := 0
		for delta := range sumChan {
			sum += delta
			fmt.Println("Sum: ", sum)
		}
	}()

	<-doneChan

	end := time.Now()

	fmt.Println("Used", end.Sub(start), "secs with", nThreads, "threads")
	return 0
}

func main() {
	//T(10000) // 41333
	//T(10000000000) // 41333

	T(1000000000000) // 41333
}
