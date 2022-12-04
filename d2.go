package main

import "github.com/spf13/cobra"

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
	Win      = "win"
	Lose     = "lose"
	Draw     = "draw"
)

func init() {
	rootCmd.AddCommand(d2p1)
	rootCmd.AddCommand(d2p2)
}

var d2p1 = &cobra.Command{
	Use: "d2p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(2)
		lines := splitAndSanitize(inp, "\n")

		total := 0

		for _, line := range lines {
			parts := splitAndSanitize(line, " ")
			if len(parts) != 2 {
				panic("invalid line")
			}

			var you, opp string

			switch parts[0] {
			case "A":
				opp = Rock
			case "B":
				opp = Paper
			case "C":
				opp = Scissors
			default:
				panic("invalid opp")
			}

			switch parts[1] {
			case "X":
				you = Rock
			case "Y":
				you = Paper
			case "Z":
				you = Scissors
			default:
				panic("invalid you")
			}

			total += roundScore(you, opp)
		}

		s.Infof("total: %d", total)
	},
}

var d2p2 = &cobra.Command{
	Use: "d2p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(2)
		lines := splitAndSanitize(inp, "\n")

		total := 0

		for _, line := range lines {
			parts := splitAndSanitize(line, " ")
			if len(parts) != 2 {
				panic("invalid line")
			}

			var opp, goal string

			switch parts[0] {
			case "A":
				opp = Rock
			case "B":
				opp = Paper
			case "C":
				opp = Scissors
			default:
				panic("invalid opp")
			}

			switch parts[1] {
			case "X":
				goal = Lose
			case "Y":
				goal = Draw
			case "Z":
				goal = Win
			default:
				panic("invalid goal")
			}

			you := playThisTo(opp, goal)
			total += roundScore(you, opp)
		}

		s.Infof("total: %d", total)
	},
}

func roundScore(you, opp string) int {
	switch you {
	case Rock:
		switch opp {
		case Rock:
			return 1 + 3
		case Paper:
			return 1
		case Scissors:
			return 1 + 6
		default:
			panic("invalid opp")
		}
	case Paper:
		switch opp {
		case Rock:
			return 2 + 6
		case Paper:
			return 2 + 3
		case Scissors:
			return 2
		default:
			panic("invalid opp")
		}
	case Scissors:
		switch opp {
		case Rock:
			return 3
		case Paper:
			return 3 + 6
		case Scissors:
			return 3 + 3
		default:
			panic("invalid opp")
		}
	default:
		panic("invalid you")
	}
}

func playThisTo(opp string, goal string) string {
	switch goal {
	case Win:
		switch opp {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		default:
			panic("invalid opp")
		}
	case Lose:
		switch opp {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		default:
			panic("invalid opp")
		}
	case Draw:
		switch opp {
		case Rock:
			return Rock
		case Paper:
			return Paper
		case Scissors:
			return Scissors
		default:
			panic("invalid opp")
		}
	default:
		panic("invalid goal")
	}
}
