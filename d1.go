package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d1p1)
	rootCmd.AddCommand(d1p2)
}

var d1p1 = &cobra.Command{
	Use: "d1p1",
	Run: func(cmd *cobra.Command, args []string) {
		sums := d1GetSums()

		max := lo.Max(sums)

		s.Infof("max: %d", max)
	},
}

var d1p2 = &cobra.Command{
	Use: "d1p2",
	Run: func(cmd *cobra.Command, args []string) {
		sums := d1GetSums()

		sort.Ints(sums)

		top3Sum := lo.Sum(sums[len(sums)-3:])

		s.Infof("top3Sum: %d", top3Sum)
	},
}

func d1GetSums() []int {
	content := mustLoadInput(1)

	elves := lo.FilterMap(splitAndSanitize(content, "\n\n"), func(elf string, _ int) ([]int, bool) {
		o := strings.TrimSpace(elf)
		if o == "" {
			return nil, false
		}

		lines := splitAndSanitize(o, "\n")

		return lo.Map(lines, func(line string, _ int) int {
			num, err := strconv.Atoi(line)
			c(err)
			return num
		}), true
	})

	return lo.Map(elves, func(elf []int, _ int) int {
		return lo.Sum(elf)
	})
}
