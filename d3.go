package main

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d3p1)
	rootCmd.AddCommand(d3p2)
}

var d3p1 = &cobra.Command{
	Use: "d3p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(3)
		lines := splitAndSanitize(inp, "\n")

		total := 0

		for _, l := range lines {
			items := splitAndSanitize(l, "")

			firstHalf := items[:len(items)/2]
			secondHalf := items[len(items)/2:]

			if len(firstHalf) != len(secondHalf) || len(firstHalf)+len(secondHalf) != len(items) {
				panic("invalid split")
			}

			common := lo.Intersect(firstHalf, secondHalf)
			common = lo.Uniq(common)
			if len(common) != 1 {
				s.Infof("l: %s, r: %s, common: %s", firstHalf, secondHalf, common)
				panic("invalid input, not only one common char")
			}

			total += getPrio(common[0])
		}

		s.Infof("total: %d", total)
	},
}

var d3p2 = &cobra.Command{
	Use: "d3p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(3)
		lines := splitAndSanitize(inp, "\n")

		if len(lines)%3 != 0 {
			panic("invalid input")
		}

		chunks := lo.Chunk(lines, 3)

		total := 0

		for _, chunk := range chunks {
			e1 := splitAndSanitize(chunk[0], "")
			e2 := splitAndSanitize(chunk[1], "")
			e3 := splitAndSanitize(chunk[2], "")

			e1 = lo.Uniq(e1)
			e2 = lo.Uniq(e2)
			e3 = lo.Uniq(e3)

			common := lo.Intersect(e1, e2)
			common = lo.Intersect(common, e3)

			if len(common) != 1 {
				panic("invalid input, not only one common char")
			}

			total += getPrio(common[0])
		}

		s.Infof("total: %d", total)
	},
}

func getPrio(c string) int {
	p, ok := prioLookup[c]
	if !ok {
		panic(fmt.Sprintf("invalid char %s", c))
	}

	return p
}

var prioLookup = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}
