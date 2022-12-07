package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWindow(t *testing.T) {
	thing := []int{1, 2, 3, 4, 5}

	window, err := NewSlidingWindow(thing, 3)
	require.NoError(t, err)

	ind, w1 := window.Next()
	require.Equal(t, 0, ind)
	require.Equal(t, []int{1, 2, 3}, w1)

	ind, w2 := window.Next()
	require.Equal(t, 1, ind)
	require.Equal(t, []int{2, 3, 4}, w2)

	ind, w3 := window.Next()
	require.Equal(t, 2, ind)
	require.Equal(t, []int{3, 4, 5}, w3)

	_, w4 := window.Next()
	require.Nil(t, w4)
}
