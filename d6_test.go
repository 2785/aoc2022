package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindStart(t *testing.T) {
	testStr := "bvwbjplbgvbhsrlpgdmjqwftvncz"
	found := d6FindStart(testStr, 4)

	require.Equal(t, 1, found)

	require.Equal(t, 2, d6FindStart("nppdvjthqldpwncqszvftbrmjlhg", 4))
	require.Equal(t, 6, d6FindStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))

	require.Equal(t, 5, d6FindStart("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
}
