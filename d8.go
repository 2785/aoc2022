package main

import (
	"strconv"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d8p1)
	rootCmd.AddCommand(d8p2)
}

var d8p1 = &cobra.Command{
	Use: "d8p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(8)

		parsed := d8SortNodes(inp)

		s.Infof("total: %d", d8CountVisible(parsed))
	},
}

var d8p2 = &cobra.Command{
	Use: "d8p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(8)

		parsed := d8SortNodes(inp)

		nodes := lo.Flatten(parsed)

		max := lo.MaxBy(nodes, func(a, b *d8Node) bool {
			return a.ScenicScore > b.ScenicScore
		})

		s.Infof("max: %d", max.ScenicScore)
	},
}

type d8Node struct {
	Height int

	Up, Down, Left, Right int

	ScenicScore int
}

func d8SortNodes(inp string) [][]*d8Node {
	lines := splitAndSanitize(inp, "\n")

	height := len(lines)
	width := 0

	nodes := lo.Map(lines, func(l string, _ int) []*d8Node {
		rows := splitAndSanitize(l, "")

		w := len(rows)
		if width == 0 {
			width = w
		} else if w != width {
			panic("width mismatch")
		}

		return lo.Map(rows, func(r string, _ int) *d8Node {
			h, err := strconv.Atoi(r)
			c(err)

			return &d8Node{Height: h, Up: -1, Down: -1, Left: -1, Right: -1}
		})
	})

	for i := 0; i < height; i++ {
		// L -> R
		for j := 0; j < width-1; j++ {
			if j == 0 {
				continue
			}

			curr := nodes[i][j]
			left := nodes[i][j-1]

			effectiveLeft := lo.Ternary(left.Left > left.Height, left.Left, left.Height)

			curr.Left = effectiveLeft
		}

		// R -> L
		for j := width - 1; j > 0; j-- {
			if j == width-1 {
				continue
			}

			curr := nodes[i][j]
			right := nodes[i][j+1]

			effectiveRight := lo.Ternary(right.Right > right.Height, right.Right, right.Height)

			curr.Right = effectiveRight
		}
	}

	for j := 0; j < width; j++ {
		// T -> B
		for i := 0; i < height-1; i++ {
			if i == 0 {
				continue
			}

			curr := nodes[i][j]
			up := nodes[i-1][j]

			effectiveUp := lo.Ternary(up.Up > up.Height, up.Up, up.Height)

			curr.Up = effectiveUp
		}

		// B -> T
		for i := height - 1; i > 0; i-- {
			if i == height-1 {
				continue
			}

			curr := nodes[i][j]
			down := nodes[i+1][j]

			effectiveDown := lo.Ternary(down.Down > down.Height, down.Down, down.Height)

			curr.Down = effectiveDown
		}
	}

	// sort out the scenic scores. why is aoc annoying

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			curr := nodes[i][j]

			var left, right, up, down int

			// left
			for jp := j - 1; jp >= 0; jp-- {
				left++
				if nodes[i][jp].Height >= curr.Height {
					break
				}
			}

			// right
			for jp := j + 1; jp < width; jp++ {
				right++
				if nodes[i][jp].Height >= curr.Height {
					break
				}
			}

			// up
			for ip := i - 1; ip >= 0; ip-- {
				up++
				if nodes[ip][j].Height >= curr.Height {
					break
				}
			}

			// down
			for ip := i + 1; ip < height; ip++ {
				down++
				if nodes[ip][j].Height >= curr.Height {
					break
				}
			}

			curr.ScenicScore = left * right * up * down
		}
	}

	return nodes
}

func d8CountVisible(grid [][]*d8Node) int {
	nodes := lo.Flatten(grid)

	return lo.CountBy(nodes, func(n *d8Node) bool {
		return n.Height > lo.Min([]int{n.Left, n.Right, n.Down, n.Up})
	})
}
