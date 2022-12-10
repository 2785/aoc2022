package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d10p1)
	rootCmd.AddCommand(d10p2)
}

var d10p1 = &cobra.Command{
	Use: "d10p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(10)

		xVals := d10RunProgram(inp)

		if len(xVals) < 221 {
			panic("invalid input")
		}

		sum := 0
		for _, cyc := range []int{20, 60, 100, 140, 180, 220} {
			sum += cyc * xVals[cyc-1]
		}

		s.Infof("sum: %d", sum)
	},
}

var d10p2 = &cobra.Command{
	Use: "d10p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(10)

		xVals := d10RunProgram(inp)

		fmt.Println(d10DrawThing(xVals))
	},
}

func d10RunProgram(inp string) []int {
	instructions := splitAndSanitize(inp, "\n")

	xVals := []int{1}

	for _, line := range instructions {
		if line == "noop" {
			xVals = append(xVals, xVals[len(xVals)-1])
		}

		if strings.HasPrefix(line, "addx") {
			parts := splitAndSanitize(line, " ")
			if len(parts) != 2 {
				panic("invalid input")
			}
			val, err := strconv.Atoi(parts[1])
			c(err)
			// cycle 1
			xVals = append(xVals, xVals[len(xVals)-1])
			// cycle 2
			xVals = append(xVals, xVals[len(xVals)-1]+val)
		}
	}

	return xVals
}

func d10DrawThing(xVals []int) string {
	// why this works I've no idea. Reading comprehension bad, words are hard
	xVals = xVals[0 : len(xVals)-1]

	if len(xVals) != 40*6 {
		panic("invalid input")
	}

	lines := []string{
		"", "", "", "", "", "",
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			x := xVals[i*40+j]
			if x-1 <= j && j <= x+1 {
				lines[i] += "#"
			} else {
				lines[i] += "."
			}
		}
	}

	return strings.Join(lines, "\n")
}
