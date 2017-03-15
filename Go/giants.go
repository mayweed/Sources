package main

import (
	"fmt"
	"log"
)

type graph struct {
	nodes []int
	edges map[int][]int
}

//init
func newGraph() graph {
	return graph{
		nodes: nil,
		edges: make(map[int][]int),
	}
}

//collect all the nodes
func (g *graph) collectNode() {
	for k, _ := range g.edges {
		g.nodes = append(g.nodes, k)
	}
	for _, neighbours := range g.edges {
		for _, w := range neighbours {
			if !checkList(w, g.nodes) {
				g.nodes = append(g.nodes, w)
			}
		}
	}
}
func checkList(c int, d []int) bool {
	for _, v := range d {
		if v == c {
			return true
		}
	}
	return false
}

//simple bfs
func (g graph) bfs(startNode, endNode int) []int {
	var visited = make(map[int]bool)
	visited[startNode] = true

	var node = []int{startNode}
	var queue = [][]int{node}

	for 0 < len(queue) {
		//pop the first element
		path := queue[0]
		queue = queue[1:]

		lastNode := path[len(path)-1]
		if lastNode == endNode {
			return path
		}

		for _, w := range g.edges[lastNode] {
			var newPath = path
			if !visited[w] {
				visited[w] = true
				newPath = append(newPath, w)
				queue = append(queue, newPath)
			}
		}
	}
	return []int{}
}

//MAIN
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
	(&g).collectNode()

	var queue = g.nodes
	var max = 0
	//var longestPath []int //for sake of it

	for 0 < len(queue) {
		startNode := queue[0]
		queue = queue[1:]
		for i := 0; i < len(queue); i++ {
			path := g.bfs(startNode, queue[i])
			//log.Println(path)
			if len(path) > max {
				max = len(path)
				//longestPath=path
			}
		}
	}

	log.Println(g.nodes, g.edges)
	// The number of people involved in the longest succession of influences
	//Should apply bfs from one node to all the others and take the longest one
	fmt.Println(max)
}
