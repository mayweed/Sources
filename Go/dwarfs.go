package main

import (
	"fmt"
	"log"
)

type graph struct {
	nodes []int
	edges map[int][]int
}

func checkList(c int, d []int) bool {
	for _, v := range d {
		if v == c {
			return true
		}
	}
	return false
}

//count directly here?
func (g graph) dfs(node int) int {
	var acc = 1
	if len(g.edges[node]) > 0 {
		visited := make(map[int]bool)
		visited[node] = true
		for n := range g.edges[node] {
			acc += 1
			if !visited[n] {
				g.dfs(n)
			}
		}
	}
	return acc + 1
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
		if !checkList(x, g.nodes) {
			g.nodes = append(g.nodes, x)
		}
		if !checkList(y, g.nodes) {
			g.nodes = append(g.nodes, y)
		}
		g.edges[x] = append(g.edges[x], y)
	}

	var max = 0
	for n, _ := range g.edges {
		log.Println(g.dfs(n))
		if g.dfs(n) > max {
			max = g.dfs(n)
		}
	}

	//LOGS
	//log.Println(g.nodes,g.edges)

	// The number of people involved in the longest succession of influences
	fmt.Println(max)
}
