package main

import "testing"
import "github.com/stretchr/testify/assert"

var m_ [][]int = [][]int{
	//      0    1    2    3    4
	[]int{7, 53, 183, 439, 863},    // 0
	[]int{497, 383, 563, 79, 973},  // 1
	[]int{287, 63, 343, 169, 583},  // 2
	[]int{627, 343, 773, 959, 943}, // 3
	[]int{767, 473, 103, 699, 303}, // 4
}

func TestSum(t *testing.T) {
	cache = make(map[Val]int)
	// 1x1
	assert.Equal(t, 7, Sum(m_, []int{0}, []int{0}))
	assert.Equal(t, 773, Sum(m_, []int{3}, []int{2}))

	// 2x2
	assert.Equal(t, 903, Sum(m_, []int{0, 1, 2}, []int{0, 1, 2}))

	// 5x5
	assert.Equal(t, 3315, Sum(m_, []int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4}))

}
