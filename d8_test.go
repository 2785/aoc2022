package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const d8TestInput = `
30373
25512
65332
33549
35390
`

func TestD8ParseTrees(t *testing.T) {
	parsed := d8SortNodes(d8TestInput)

	require.Equal(t, 21, d8CountVisible(parsed))

	require.Equal(t, 6, parsed[2][1].ScenicScore)
}
