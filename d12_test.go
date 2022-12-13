package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCharToInt(t *testing.T) {
	require.Equal(t, 0, d12CharInt('a'))
	require.Equal(t, 1, d12CharInt('b'))

	require.Panics(t, func() { d12CharInt('A') })
	require.Panics(t, func() { d12CharInt('0') })
}

const d12Sample = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestD12ParseNodes(t *testing.T) {
	nodes, src, dst := d12ParseInput(d12Sample)

	require.Equal(t, d12Coord{x: 0, y: 0}, src)

	require.Equal(t, d12Coord{x: 2, y: 5}, dst)

	require.Equal(t, 1, nodes[d12Coord{x: 0, y: 2}].h)

	require.Len(t, nodes[d12Coord{x: 0, y: 5}].neighbours, 2)

	require.Equal(t, 31, d12FuckingDijkstra(nodes, src, dst))
}

func TestD12ShortestDijkstra(t *testing.T) {
	nodes, src, dst := d12ParseInput(d12Sample)

	require.Equal(t, 29, d12ShortestDijkstra(nodes, src, dst))
}
