package main

import (
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d9p1)
	rootCmd.AddCommand(d9p2)
}

var d9p1 = &cobra.Command{
	Use: "d9p1",
	Run: func(cmd *cobra.Command, args []string) {
		h, t := D9Cord{0, 0}, D9Cord{0, 0}

		inp := mustLoadInput(9)

		lines := splitAndSanitize(inp, "\n")

		tailVisited := map[D9Cord]bool{
			t: true,
		}

		for _, line := range lines {
			parts := splitAndSanitize(line, " ")
			if len(parts) != 2 {
				panic("invalid input")
			}

			dir := parts[0]
			count, err := strconv.Atoi(parts[1])
			c(err)

			var dalta D9Cord
			switch dir {
			case "U":
				dalta = D9Up
			case "D":
				dalta = D9Down
			case "L":
				dalta = D9Left
			case "R":
				dalta = D9Right
			default:
				panic("invalid input")
			}

			for i := 0; i < count; i++ {
				h, t = D9MoveHead(h, t, dalta)
				tailVisited[t] = true
			}
		}

		s.Infof("tail visited: %d", len(tailVisited))
	},
}

var d9p2 = &cobra.Command{
	Use: "d9p2",
	Run: func(cmd *cobra.Command, args []string) {
		h := D9Cord{0, 0}
		rope := make([]D9Cord, 9)
		for i := range rope {
			rope[i] = D9Cord{0, 0}
		}

		inp := mustLoadInput(9)

		lines := splitAndSanitize(inp, "\n")

		tailVisited := map[D9Cord]bool{
			rope[8]: true,
		}

		for _, line := range lines {
			parts := splitAndSanitize(line, " ")
			if len(parts) != 2 {
				panic("invalid input")
			}

			dir := parts[0]
			count, err := strconv.Atoi(parts[1])
			c(err)

			var dalta D9Cord
			switch dir {
			case "U":
				dalta = D9Up
			case "D":
				dalta = D9Down
			case "L":
				dalta = D9Left
			case "R":
				dalta = D9Right
			default:
				panic("invalid input")
			}

			for i := 0; i < count; i++ {

				h = D9Add(h, dalta)

				for j := 0; j < len(rope); j++ {
					curr := rope[j]
					currHead := h
					if j > 0 {
						currHead = rope[j-1]
					}

					rope[j] = D9Follow(currHead, curr)
				}

				tail := rope[len(rope)-1]
				tailVisited[tail] = true
			}
		}

		s.Infof("tail visited: %d", len(tailVisited))
	},
}

type D9Cord struct {
	X, Y int
}

func D9Add(a, b D9Cord) D9Cord {
	return D9Cord{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

var (
	D9Up    = D9Cord{0, 1}
	D9Down  = D9Cord{0, -1}
	D9Left  = D9Cord{-1, 0}
	D9Right = D9Cord{1, 0}
)

func D9MoveHead(h, t, delta D9Cord) (rh, rt D9Cord) {
	rh = D9Cord{
		X: h.X + delta.X,
		Y: h.Y + delta.Y,
	}

	rt = D9Follow(rh, t)

	return
}

func D9Follow(target, curr D9Cord) D9Cord {
	// if h and t are touching, do nothing
	if touching(target, curr) {
		return curr
	}

	// if same x, then move y closer by 1
	if target.X == curr.X && dist(target.Y, curr.Y) >= 2 {
		if target.Y > curr.Y {
			return D9Cord{
				X: curr.X,
				Y: curr.Y + 1,
			}
		} else {
			return D9Cord{
				X: curr.X,
				Y: curr.Y - 1,
			}
		}
	}

	// if same y, then move x closer by 1
	if target.Y == curr.Y && dist(target.X, curr.X) >= 2 {
		if target.X > curr.X {
			return D9Cord{
				X: curr.X + 1,
				Y: curr.Y,
			}
		} else {
			return D9Cord{
				X: curr.X - 1,
				Y: curr.Y,
			}
		}
	}

	// otherwise move in the direction of the head
	rt := curr

	if target.X > curr.X {
		rt.X = curr.X + 1
	}

	if target.X < curr.X {
		rt.X = curr.X - 1
	}

	if target.Y > curr.Y {
		rt.Y = curr.Y + 1
	}

	if target.Y < curr.Y {
		rt.Y = curr.Y - 1
	}

	return rt
}

func dist(a, b int) int {
	d := a - b
	if d < 0 {
		d = -d
	}

	return d
}

func touching(a, b D9Cord) bool {
	return dist(a.X, b.X) <= 1 && dist(a.Y, b.Y) <= 1
}
