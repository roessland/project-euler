package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test800(t *testing.T) {
	require.Equal(t, 10790, solve(800))
}
