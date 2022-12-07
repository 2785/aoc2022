package main

import (
	"path"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d7p1)
	rootCmd.AddCommand(d7p2)
}

var d7p1 = &cobra.Command{
	Use: "d7p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(7)
		_, dirs := d7ParseTree(inp)

		tot := 0

		for _, v := range dirs {
			siz := v.Size()
			if siz <= 100000 {
				tot += siz
			}
		}

		s.Infof("total: %d", tot)
	},
}

var d7p2 = &cobra.Command{
	Use: "d7p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(7)
		_, dirs := d7ParseTree(inp)

		root, ok := dirs[""]
		if !ok {
			panic("could not find root")
		}

		rootSize := root.Size()

		spaceAvailable := 70000000 - rootSize
		needSpace := 30000000 - spaceAvailable

		if needSpace <= 0 {
			panic("no need to move")
		}

		bigEnoughDirs := lo.Filter(lo.Entries(dirs), func(e lo.Entry[string, *d7NodeDir], _ int) bool {
			return e.Value.Size() >= needSpace
		})

		smallest := lo.MinBy(bigEnoughDirs, func(a, b lo.Entry[string, *d7NodeDir]) bool {
			return a.Value.Size() < b.Value.Size()
		})

		s.Infof("smallest: %s (%d)", smallest.Key, smallest.Value.Size())
	},
}

type d7Node interface {
	Type() string
	Size() int
}

type d7NodeFile struct {
	name string
	size int
}

func (n *d7NodeFile) Type() string {
	return "file"
}

func (n *d7NodeFile) Size() int {
	return n.size
}

type d7NodeDir struct {
	name     string
	children []d7Node
	size     int
}

func (n *d7NodeDir) Type() string {
	return "dir"
}

func (n *d7NodeDir) Size() int {
	if n.size != 0 {
		return n.size
	}

	total := 0
	for _, c := range n.children {
		total += c.Size()
	}

	n.size = total

	return total
}

func d7ParseTree(s string) (root *d7NodeDir, dirs map[string]*d7NodeDir) {
	lines := splitAndSanitize(s, "\n")

	if len(lines) == 0 {
		panic("empty input")
	}

	if lines[0] != "$ cd /" {
		panic("invalid input, must start with $ cd /")
	}

	dirStack := make(dirStack, 0)

	root = &d7NodeDir{
		name:     "",
		children: []d7Node{},
	}

	dirMap := map[string]*d7NodeDir{
		"": root,
	}

	currDir := root

	for _, l := range lines[1:] {
		// if we're cd'ing, update pwd
		if strings.HasPrefix(l, "$ cd ") {
			dir := strings.TrimPrefix(l, "$ cd ")
			if dir == ".." {
				dirStack.Pop()
			} else {
				dirStack.Push(dir)
				// update curr dir
				newDir, ok := dirMap[dirStack.String()]
				if !ok {
					panic("invalid dir")
				}

				currDir = newDir
			}
			continue
		}

		// if we're ls-ing, meh, don't care
		if strings.HasPrefix(l, "$ ls") {
			continue
		}

		// otherwise we're listing stuff, if we're a dir, make a new dir entry, if not exist, & add
		// it to the current dir
		if strings.HasPrefix(l, "dir") {
			newDir := path.Join(dirStack.String(), strings.TrimPrefix(l, "dir "))

			if _, ok := dirMap[newDir]; ok {
				panic("duplicate dir")
			}

			newNode := &d7NodeDir{
				name:     newDir,
				children: []d7Node{},
			}

			dirMap[newDir] = newNode
			currDir.children = append(currDir.children, newNode)
			continue
		}

		// otherwise we're dealing with a file, extract the size
		parts := strings.Split(l, " ")
		if len(parts) != 2 {
			panic("invalid file")
		}

		size, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("invalid file")
		}

		name := path.Join(dirStack.String(), parts[1])

		newNode := &d7NodeFile{
			name: name,
			size: size,
		}

		currDir.children = append(currDir.children, newNode)
	}

	return root, dirMap
}

type dirStack []string

func (s *dirStack) Push(d string) {
	*s = append(*s, d)
}

func (s *dirStack) Pop() string {
	if len(*s) == 0 {
		panic("empty stack")
	}

	l := len(*s)
	d := (*s)[l-1]
	*s = (*s)[:l-1]
	return d
}

func (s *dirStack) String() string {
	return strings.Join(*s, "/")
}
