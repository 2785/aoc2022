package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d11p1)
	rootCmd.AddCommand(d11p2)
}

var d11p1 = &cobra.Command{
	Use: "d11p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(11)

		monkeysWithN := d11ParseMonkeys(inp)

		monkeys := lo.Map(monkeysWithN, func(m lo.Tuple2[day11Monkey, int], _ int) day11Monkey {
			return m.A
		})

		for i := 0; i < 20; i++ {
			d11Round(monkeys)
		}

		sort.Slice(monkeys, func(i, j int) bool {
			return monkeys[i].inspectCount > monkeys[j].inspectCount
		})

		res := monkeys[0].inspectCount * monkeys[1].inspectCount

		s.Infof("res: %d", res)
	},
}

var d11p2 = &cobra.Command{
	Use: "d11p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(11)

		monkeysWithN := d11ParseMonkeys(inp)

		monkeys := lo.Map(monkeysWithN, func(m lo.Tuple2[day11Monkey, int], _ int) day11Monkey {
			return m.A
		})

		divisors := lo.Map(monkeysWithN, func(m lo.Tuple2[day11Monkey, int], _ int) int {
			return m.B
		})

		base := lcm(divisors...)

		for i := 0; i < 10000; i++ {
			d11RoundWithBase(monkeys, base)
		}

		sort.Slice(monkeys, func(i, j int) bool {
			return monkeys[i].inspectCount > monkeys[j].inspectCount
		})

		res := monkeys[0].inspectCount * monkeys[1].inspectCount

		s.Infof("res: %d", res)
	},
}

type day11Monkey struct {
	items []int
	// what to modify the item by
	op func(int) int
	// number of money to throw the thing to (?)
	rules func(int) int

	inspectCount int
}

var opRe = regexp.MustCompile(`^new = old (?P<op>\S) (?P<pa>\S+)$`)
var condRe = regexp.MustCompile(`^(?P<op>\w+) by (?P<pa>\d+)$`)
var throwRe = regexp.MustCompile(`^throw to monkey (?P<monkey>\d+)$`)

func day11ParseMonkey(inp string) (day11Monkey, int) {
	lines := splitAndSanitize(inp, "\n")

	lines = lines[1:]

	if len(lines) != 5 {
		panic("invalid input")
	}

	itemsS := splitAndSanitize(strings.TrimPrefix(lines[0], "Starting items:"), ",")

	items := lo.Map(itemsS, func(itemS string, _ int) int {
		n, err := strconv.Atoi(itemS)
		c(err)
		return n
	})

	var operator func(int) int

	opStr := strings.TrimPrefix(lines[1], "Operation: ")

	matches, err := NamedCaptureGroup(opRe, opStr)
	c(err)

	op := matches["op"]
	pa := matches["pa"]

	switch op {
	case "+":
		if pa == "old" {
			operator = opMul(2)
		} else if n, err := strconv.Atoi(pa); err == nil {
			operator = opAdd(n)
		} else {
			panic("invalid input")
		}
	case "*":
		if pa == "old" {
			operator = opSq()
		} else if n, err := strconv.Atoi(pa); err == nil {
			operator = opMul(n)
		} else {
			panic("invalid input")
		}
	default:
		panic("invalid input")
	}

	condStr := strings.TrimPrefix(lines[2], "Test: ")
	tStr := strings.TrimPrefix(lines[3], "If true: ")
	fStr := strings.TrimPrefix(lines[4], "If false: ")

	tMatches, err := NamedCaptureGroup(throwRe, tStr)
	c(err)

	fMatches, err := NamedCaptureGroup(throwRe, fStr)
	c(err)

	t := tMatches["monkey"]
	f := fMatches["monkey"]

	if t == "" || f == "" {
		panic("invalid input")
	}

	tN, err := strconv.Atoi(t)
	c(err)

	fN, err := strconv.Atoi(f)
	c(err)

	var rules func(int) int

	cond, err := NamedCaptureGroup(condRe, condStr)
	c(err)

	var n int

	switch cond["op"] {
	case "divisible":
		n, err = strconv.Atoi(cond["pa"])
		c(err)

		rules = func(x int) int {

			if x%n == 0 {
				return tN
			} else {
				return fN
			}

		}
	default:
		panic("invalid input")
	}

	return day11Monkey{
		items: items,
		op:    operator,
		rules: rules,
	}, n
}

func d11Round(monkeys []day11Monkey) {
	for i := range monkeys {
		for _, item := range monkeys[i].items {
			monkeys[i].inspectCount++

			newWorryLevel := monkeys[i].op(item)

			nll := newWorryLevel / 3

			newMonkey := monkeys[i].rules(nll)

			monkeys[newMonkey].items = append(monkeys[newMonkey].items, nll)
		}

		monkeys[i].items = []int{}
	}
}

func d11RoundWithBase(monkeys []day11Monkey, base int) {
	for i := range monkeys {
		for _, item := range monkeys[i].items {
			monkeys[i].inspectCount++

			newWorryLevel := monkeys[i].op(item)

			nll := newWorryLevel % base

			newMonkey := monkeys[i].rules(nll)

			monkeys[newMonkey].items = append(monkeys[newMonkey].items, nll)
		}

		monkeys[i].items = []int{}
	}
}

func opAdd(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func opMul(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func opSq() func(int) int {
	return func(y int) int {
		return y * y
	}
}

func d11ParseMonkeys(inp string) []lo.Tuple2[day11Monkey, int] {
	parts := splitAndSanitize(inp, "\n\n")

	return lo.Map(parts, func(part string, _ int) lo.Tuple2[day11Monkey, int] {
		m, n := day11ParseMonkey(part)
		return lo.T2(m, n)
	})
}
