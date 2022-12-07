package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gammazero/deque"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d5p1)
	rootCmd.AddCommand(d5p2)
}

var d5p1 = &cobra.Command{
	Use: "d5p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(5)
		parts := strings.Split(inp, "\n\n")

		if len(parts) != 2 {
			panic("invalid input, expected 2 parts")
		}

		stacks := d5ParseStacks(parts[0])

		instructions := splitAndSanitize(parts[1], "\n")

		for _, i := range instructions {
			count, from, to := d5ParseInstruction(i)
			from--
			to--

			if count > stacks[from].Len() {
				panic("invalid instruction, moving more than available")
			}

			for i := 0; i < count; i++ {
				stacks[to].PushBack(stacks[from].PopBack())
			}
		}

		res := ""

		for i := range stacks {
			stack := stacks[i]
			if stack.Len() == 0 {
				panic("invalid input, stack is empty")
			}

			res = res + stack.PopBack()
		}

		s.Infof("result: %s", res)
	},
}

var d5p2 = &cobra.Command{
	Use: "d5p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(5)
		parts := strings.Split(inp, "\n\n")

		if len(parts) != 2 {
			panic("invalid input, expected 2 parts")
		}

		stacks := d5ParseStacks(parts[0])

		instructions := splitAndSanitize(parts[1], "\n")

		for _, i := range instructions {
			count, from, to := d5ParseInstruction(i)
			from--
			to--

			if count > stacks[from].Len() {
				panic("invalid instruction, moving more than available")
			}

			tmpQ := deque.New[string]()

			for i := 0; i < count; i++ {
				tmpQ.PushBack(stacks[from].PopBack())
			}

			for tmpQ.Len() > 0 {
				stacks[to].PushBack(tmpQ.PopBack())
			}
		}

		res := ""

		for i := range stacks {
			stack := stacks[i]
			if stack.Len() == 0 {
				panic("invalid input, stack is empty")
			}

			res = res + stack.PopBack()
		}

		s.Infof("result: %s", res)
	},
}

func d5ParseStacks(inp string) []*deque.Deque[string] {
	lines := strings.Split(inp, "\n")
	firstLength := len(lines[0])
	for _, l := range lines {
		if firstLength != len(l) {
			panic("invalid input")
		}
	}

	if (firstLength+1)%4 != 0 {
		panic("invalid input, not a multiple of 4")
	}

	stackSize := (firstLength + 1) / 4

	stacks := make([]*deque.Deque[string], stackSize)
	for i := range stacks {
		stacks[i] = deque.New[string]()
	}

	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for j := 0; j < stackSize; j++ {
			char := line[j*4+1 : j*4+2]
			if char != " " {
				stacks[j].PushBack(char)
			}
		}
	}

	return stacks
}

var d5InstructionRe = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

func d5ParseInstruction(s string) (count, from, to int) {
	matches := d5InstructionRe.FindStringSubmatch(s)
	if len(matches) != 4 {
		panic("invalid instruction")
	}

	count, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}

	from, err = strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	to, err = strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}

	return
}
