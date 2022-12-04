package main

import (
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d4p1)
	rootCmd.AddCommand(d4p2)
}

type d2Range struct {
	start, end int
}

var d4p1 = &cobra.Command{
	Use: "d4p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(4)
		lines := splitAndSanitize(inp, "\n")

		total := 0

		for _, l := range lines {
			items := splitAndSanitize(l, ",")

			if len(items) != 2 {
				panic("invalid input")
			}

			r1 := parseRange(items[0])
			r2 := parseRange(items[1])

			if d2Contains(r1, r2) {
				total++
			}
		}

		s.Infof("total: %d", total)
	},
}

var d4p2 = &cobra.Command{
	Use: "d4p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(4)
		lines := splitAndSanitize(inp, "\n")

		total := 0

		for _, l := range lines {
			items := splitAndSanitize(l, ",")

			if len(items) != 2 {
				panic("invalid input")
			}

			r1 := parseRange(items[0])
			r2 := parseRange(items[1])

			if d2Overlaps(r1, r2) {
				total++
			}
		}

		s.Infof("total: %d", total)
	},
}

func d2Contains(a, b d2Range) bool {
	if a.start <= b.start && a.end >= b.end {
		return true
	}

	if b.start <= a.start && b.end >= a.end {
		return true
	}

	return false
}

func d2Overlaps(a, b d2Range) bool {
	if a.start <= b.start && a.end >= b.start {
		return true
	}

	if b.start <= a.start && b.end >= a.start {
		return true
	}

	return false
}

func parseRange(s string) d2Range {
	parts := splitAndSanitize(s, "-")
	if len(parts) != 2 {
		panic("invalid input")
	}

	l, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	r, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return d2Range{
		start: l,
		end:   r,
	}
}
