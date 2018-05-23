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
//func newGraph() graph {
//	return graph{
//		nodes: nil,
//		edges: make(map[int][]int),
//	}
//}
//init
func newGraph() graph {
	return graph{
		nodes: nil,
		/*
			//edges: make(map[int][]int),
			//  depth max 6
			edges: map[int][]int{
				5: []int{3, 6},
				6: []int{1, 2},
				7: []int{4},
				9: []int{4},
				4: []int{5},
				2: []int{8},
			},
		*/
		//(1,2),(2,3),(2,4),(3,4),(4,5)
		//should be five...
		edges: map[int][]int{
			1: []int{2},
			2: []int{3, 4},
			3: []int{4},
			4: []int{5},
		},

		/*
			//depth max 4
			edges: map[int][]int{
				10: []int{1, 3, 11},
				1:  []int{2, 3},
				3:  []int{4},
				2:  []int{4, 5},
			},

		*/
		//TODO find a way to quickly parse that
		//oki string: first split on space then split num on ""
		//or num[0] => x num[1] =>y
		// 56  53  61  74  62  94  45  28
	}
}

//oki and what's the base case??
//https://myarch.com/treeiter/traditways/ ex 1
func (g graph) traverseGraph(node, depth int) int {
	for _, v := range g.edges[node] {
		log.Println(v)
		g.traverseGraph(v, depth+1)
	}
	return depth
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
	/*
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
	*/
	g := newGraph()

	(&g).collectNode()

	var queue = g.nodes
	var max = 0
	//var longestPath []int //for sake of it
	for 0 < len(queue) {
		startNode := queue[0]
		queue = queue[1:]
		for i := 0; i < len(queue); i++ {
			path := g.bfs(startNode, queue[i])
			log.Println(path)
			if len(path) > max {
				max = len(path)
				//longestPath=path
			}
		}
	}
	depth := g.traverseGraph(1, 0)
	log.Println(g.nodes, g.edges, depth)
	// The number of people involved in the longest succession of influences
	//Should apply bfs from one node to all the others and take the longest one
	fmt.Println(max)
}
