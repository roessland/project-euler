package main

import (
	"github.com/roessland/gopkg/mathutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkEntirelyOddSlow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for num := int64(1); num < 10000; num++ {
			EntirelyOddSlow(num)
		}
	}
}

func BenchmarkEntirelyOdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for num := int64(1); num < 10000; num++ {
			EntirelyOdd(num)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for num := int64(1); num < 10000; num++ {
			mathutil.Reverse(num, 10)
		}
	}
}

func TestEntirelyOdd(t *testing.T) {
	for n := int64(1); n < 100000; n++ {
		assert.Equal(t, EntirelyOddSlow(n), EntirelyOdd(n))
	}
}

func BenchmarkBruteForce6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BruteForce(1000000)
	}
}

func BenchmarkSmart6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Smart(6)
	}
}

func TestCorrect(t *testing.T) {
	assert.Equal(t, int64(120), BruteForce(1000))
	assert.Equal(t, BruteForce(10), Smart(1))
	assert.Equal(t, BruteForce(100), Smart(2))
	assert.Equal(t, BruteForce(1000), Smart(3))
	assert.Equal(t, BruteForce(10000), Smart(4))
	assert.Equal(t, BruteForce(100000), Smart(5))
	assert.Equal(t, BruteForce(1000000), Smart(6))
}
