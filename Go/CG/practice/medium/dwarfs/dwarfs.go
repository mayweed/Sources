//https://www.geeksforgeeks.org/longest-path-in-a-directed-acyclic-graph-dynamic-programming/
//https://gist.github.com/joaquinvanschoren/5006291
//https://tylercipriani.com/blog/2017/09/13/topographical-sorting-in-golang/
package main

import (
	"fmt"
)

//GRAPH
type Graph struct {
	nodes []int
	edges map[int][]int
	sort  []int
}

func (g *Graph) dfs(src int, dp map[int]int) {
	visited := make(map[int]bool)
	visited[src] = true
	for _, n := range g.edges[src] {
		if !visited[n] {
			g.dfs(n, dp)
			dp[n] = dp[src] + 1
		}
	}
}

//MAIN
func main() {
	// n: the number of relationships of influence
	var n int
	fmt.Scan(&n)
	g := Graph{edges: make(map[int][]int)}

	var inDegree = make(map[int]int)
	var seen = make(map[int]bool)
	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
		//log.Println("relation ", i, "x ", x, "y ", y)
		g.edges[x] = append(g.edges[x], y)
		inDegree[y] += 1
		if !seen[x] {
			g.nodes = append(g.nodes, x)
			seen[x] = true
		} else if !seen[y] {
			g.nodes = append(g.nodes, y)
			seen[y] = true
		}
	}

	//find a root (or roots)
	var roots []int
	for _, n := range g.nodes {
		if inDegree[n] == 0 {
			roots = append(roots, n)
		}
	}
	var dp = make(map[int]int)
	//	for _, n := range g.nodes {
	g.dfs(roots[0], dp)
	//}
	//fmt.Fprintln(os.Stderr, n, g.edges, g.nodes, dp, g.sort)
	// The number of people involved in the longest succession of influences
}
