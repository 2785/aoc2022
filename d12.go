package main

import (
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(d12p1)
	rootCmd.AddCommand(d12p2)
}

var d12p1 = &cobra.Command{
	Use: "d12p1",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(12)

		nodes, src, dst := d12ParseInput(inp)

		dist := d12FuckingDijkstra(nodes, src, dst)

		s.Infof("dist: %d", dist)
	},
}

var d12p2 = &cobra.Command{
	Use: "d12p2",
	Run: func(cmd *cobra.Command, args []string) {
		inp := mustLoadInput(12)

		nodes, src, dst := d12ParseInput(inp)

		dist := d12ShortestDijkstra(nodes, src, dst)

		s.Infof("dist: %d", dist)
	},
}

type d12Coord struct {
	x, y int
}

type d12Node struct {
	cord d12Coord
	h    int
	dist int

	visited bool

	neighbours []*d12Node
}

func d12ParseInput(inp string) (nodes map[d12Coord]*d12Node, src, dst d12Coord) {
	lines := splitAndSanitize(inp, "\n")

	if len(lines) == 0 {
		panic("invalid input")
	}

	width := len(lines[0])

	for _, line := range lines {
		if len(line) != width {
			panic("invalid input")
		}
	}

	nodes = make(map[d12Coord]*d12Node)

	for i, line := range lines {
		for j, char := range []byte(line) {
			cord := d12Coord{x: i, y: j}

			node := d12Node{
				cord:       cord,
				neighbours: make([]*d12Node, 0),
			}

			switch char {
			case 'S':
				node.dist = 0
				src = cord
				node.h = 0
			case 'E':
				node.dist = -1
				dst = cord
				node.h = 25
			default:
				node.dist = -1
				node.h = d12CharInt(char)
			}

			nodes[cord] = &node
		}
	}

	for _, node := range nodes {
		for _, cord := range []d12Coord{
			{x: node.cord.x - 1, y: node.cord.y},
			{x: node.cord.x + 1, y: node.cord.y},
			{x: node.cord.x, y: node.cord.y - 1},
			{x: node.cord.x, y: node.cord.y + 1},
		} {
			if n, ok := nodes[cord]; ok && n.h <= node.h+1 {
				node.neighbours = append(node.neighbours, n)
			}
		}
	}

	return nodes, src, dst
}

func d12CharInt(char byte) int {
	i := int(char) - 97
	if i < 0 || i > 25 {
		panic("invalid input")
	}

	return i
}

func d12FuckingDijkstra(nodes map[d12Coord]*d12Node, src, dst d12Coord) int {
	if src == dst {
		return 0
	}

	curr := nodes[src]

	for {
		if nodes[dst].visited {
			return nodes[dst].dist
		}

		for _, n := range curr.neighbours {
			if n.visited {
				continue
			}

			if n.dist == -1 || n.dist > curr.dist+1 {
				n.dist = curr.dist + 1
			}
		}

		curr.visited = true

		next := lo.MinBy(lo.Filter(lo.Values(nodes), func(n *d12Node, _ int) bool { return !n.visited }), func(a, b *d12Node) bool {
			if a.dist == b.dist {
				return true
			}

			if a.dist == -1 {
				return false
			}

			if b.dist == -1 {
				return true
			}

			return a.dist < b.dist
		})

		curr = next
	}
}

func d12ShortestDijkstra(nodes map[d12Coord]*d12Node, src, dst d12Coord) int {
	nodes[src].dist = -1

	sets := lo.FilterMap(lo.Entries(nodes), func(n lo.Entry[d12Coord, *d12Node], _ int) (lo.Entry[d12Coord, map[d12Coord]*d12Node], bool) {
		if n.Value.h == 0 {
			// if all the neighbours are ground level, why even bother checking. this apparently
			// reduces the runtime from ~ 15 minutes to ~ 10 seconds. I imagine this is intentional
			// by AOC
			nHights := lo.Map(n.Value.neighbours, func(n *d12Node, _ int) int { return n.h })

			if lo.Max(nHights) == 0 {
				return lo.Entry[d12Coord, map[d12Coord]*d12Node]{Key: n.Key, Value: nodes}, false
			}

			cloned := d12CloneNodes(nodes)

			newSrc := cloned[n.Key]

			newSrc.dist = 0

			cloned[n.Key] = newSrc

			return lo.Entry[d12Coord, map[d12Coord]*d12Node]{Key: n.Key, Value: cloned}, true
		}

		return lo.Entry[d12Coord, map[d12Coord]*d12Node]{}, false
	})

	// progress bar because I hate waiting
	pbar := progressbar.Default(int64(len(sets)))

	defer func() { _ = pbar.Finish() }()

	dists := parallel.Map(sets, func(set lo.Entry[d12Coord, map[d12Coord]*d12Node], _ int) int {
		defer pbar.Add(1)

		return d12FuckingDijkstra(set.Value, set.Key, dst)
	})

	return lo.Min(dists)
}

func d12CloneNodes(nodes map[d12Coord]*d12Node) map[d12Coord]*d12Node {
	newNodes := make(map[d12Coord]*d12Node)

	for _, node := range nodes {
		newNodes[node.cord] = &d12Node{
			cord:    node.cord,
			h:       node.h,
			dist:    node.dist,
			visited: node.visited,
		}
	}

	for _, node := range nodes {
		for _, n := range node.neighbours {
			newNodes[node.cord].neighbours = append(newNodes[node.cord].neighbours, newNodes[n.cord])
		}
	}

	return newNodes
}
