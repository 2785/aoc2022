package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestD5ParseStacks(t *testing.T) {
	inp := `[D]        
[N] [C]    
[Z] [M] [P]
 1   2   3 `

	stacks := d5ParseStacks(inp)

	chars := make([]string, len(stacks))

	for i := range stacks {
		stack := stacks[i]
		for stack.Len() > 0 {
			chars[i] = chars[i] + stacks[i].PopBack()
		}
	}

	require.Equal(t, "DNZ", chars[0])
	require.Equal(t, "CM", chars[1])
	require.Equal(t, "P", chars[2])
}
