package main

import (
	"fmt"
	"log"
)

//UTILS
func In(c int, d []int) bool {
	for _, v := range d {
		if v == c {
			return true
		}
	}
	return false
}

//GRAPH
type graph struct {
	nodes []int
	edges map[int][]int
}

//count directly here?
func (g graph) maxDepth(node int) int {
	//var acc=0
	//This one should stop when g.edges[node] is empty!!
	//There is no more nodes to visit
	if len(g.edges[node]) > 0 {
		for _, n := range g.edges[node] {
			return g.maxDepth(n) + 1
		}
	} else {
		return 1
	}
	//return 0
}

//MAIN
func main() {
	// n: the number of relationships of influence
	var n int
	fmt.Scan(&n)
	g := graph{edges: make(map[int][]int)}

	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
		//dont think it's necessary!!take the nodes from map key!!
		if !In(x, g.nodes) {
			g.nodes = append(g.nodes, x)
		}
		if !In(y, g.nodes) {
			g.nodes = append(g.nodes, y)
		}
		g.edges[x] = append(g.edges[x], y)
	}

	//I should ( must?) put that in maxDepth no?
	var max = 0
	for n, _ := range g.edges {
		log.Println(n, g.maxDepth(n))
		if g.maxDepth(n) >= max {
			max = g.maxDepth(n)
		}
	}

	//LOGS
	log.Println(g.nodes, g.edges)

	// The number of people involved in the longest succession of influences
	fmt.Println(max)
}
