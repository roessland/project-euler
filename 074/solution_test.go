package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestNext(t *testing.T) {
	assert.Equal(t, int64(363601), Next(169), "")
	assert.Equal(t, int64(1454), Next(363601), "")
	assert.Equal(t, int64(169), Next(1454), "")
}

func TestChainLength(t *testing.T) {
	assert.Equal(t, 5, ChainLength(69), "")
	assert.Equal(t, 3, ChainLength(169), "")
	assert.Equal(t, 4, ChainLength(78), "")
	assert.Equal(t, 2, ChainLength(540), "")
}
