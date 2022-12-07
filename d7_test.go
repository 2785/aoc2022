package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const samp = `
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func TestParseTree(t *testing.T) {
	root, dirs := d7ParseTree(samp)

	require.NotNil(t, root)
	require.NotNil(t, dirs)

	require.Equal(t, 584, dirs["a/e"].Size())
	require.Equal(t, 94853, dirs["a"].Size())
}
