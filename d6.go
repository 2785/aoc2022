package main

import (
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d6p1)
	rootCmd.AddCommand(d6p2)
}

var d6p1 = &cobra.Command{
	Use: "d6p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(6)

		start := d6FindStart(inp, 4)
		s.Infof("start: %d", start+4)
	},
}

var d6p2 = &cobra.Command{
	Use: "d6p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(6)

		start := d6FindStart(inp, 14)
		s.Infof("start: %d", start+14)
	},
}

func d6FindStart(s string, size int) int {
	parts := strings.Split(s, "")

	wind, err := NewSlidingWindow(parts, size)
	c(err)

	fi := -1

	for {
		ind, w := wind.Next()
		if w == nil {
			break
		}

		if len(lo.Uniq(w)) == size {
			// found it
			fi = ind
			break
		}
	}

	if fi == -1 {
		panic("could not find start")
	}

	return fi
}
