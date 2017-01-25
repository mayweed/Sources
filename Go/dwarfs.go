package main

import (
	"fmt"
	"log"
)

type edge struct {
	from int
	to   int
}
type graph struct {
	nodes int
	edges map[int][]int
}

//init
func newGraph() graph {
	return graph{
		nodes: 0,
		edges: make(map[int][]int),
	}
}

func (g *graph) pickUpNode() {
	for k, _ := range g.edges {
		g.nodes = append(g.nodes, k)
	}
	//does not work: should append only nodes that are not already
	//in the list
	for _, v := range g.edges {
		var notIn bool
		for _, node := range v {
			for n := range g.nodes {
				if node == n {
					notIn = true
				}
			}
			if !notIn {
				g.nodes = append(g.nodes, node)
			}
		}
	}
}

//a int that is the height of the graph
func (g graph) dfs(startNode int) int {
	var visited = make(map[int]bool)
	var stack []int
	var height = 1

	//i dont need that in a rec func non?
	stack = g.edges[startNode]
	visited[startNode] = true
	//log.Println(visited)

	//is this correct? no LIFO no?
	for _, n := range stack {
		if visited[n] {
			continue
		} else {
			g.dfs(n)
			height += 1
		}
	}
	return height
}

func main() {
	// n: the number of relationships of influence
	var n int
	fmt.Scan(&n)

	g := newGraph()

	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
		g.edges[x] = append(g.edges[x], y)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// The number of people involved in the longest succession of influences
	log.Println(g.nodes, g.edges)
}
