package main

import "fmt"

func f(n, maxPartition int64) int64 {
	// There is no way to partition a positive number using a sum of zero...
	if maxPartition == 0 {
		return 0
	}

	// There is exactly one way to partition 0
	if n == 0 {
		return 1
	}

	// There is no way to partition a negative number.
	if n < 0 {
		return 0
	}

	// The number of partitions is equal to the number of partitions where the
	// largest partition has size  maxPartition + the number of partitions
	// where the largest partition has size maxPartition-1 or less.
	return f(n-maxPartition, maxPartition) + f(n, maxPartition-1)
}

func main() {
	// The number itself is included in the partitioning, so subtract one
	fmt.Printf("ANSWER = %v\n", f(100, 100)-1)
}
