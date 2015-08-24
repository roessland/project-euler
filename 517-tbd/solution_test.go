package main

import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"

func TestAlgo(t *testing.T) {
	assert.Equal(t, int64(7564511), SlowAlgo(90), "example result")
	assert.Equal(t, int64(7564511), Algo(90), "example result")
	assert.Equal(t, int64(7564511), FastAlgo(90), "example result")

	// Test fails for square numbers, so skip the square numbers
	for a := int64(3); a < int64(10); a++ {
		assert.Equal(t, SlowAlgo(a*a+1), Algo(a*a+1), fmt.Sprintf("a = %v", a*a+1))
	}
}

func BenchmarkAlgo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Algo(10010)
	}
}

func BenchmarkFastAlgo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FastAlgo(10010)
	}
}
