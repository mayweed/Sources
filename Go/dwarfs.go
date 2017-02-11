package main

import (
	"fmt"
	//	"log"
)

type graph struct {
	nodes []int
	edges map[int][]int
}

//init
func newGraph() graph {
	return graph{
		nodes: nil,
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

//a simple dfs
func (g graph) dfs(startNode int) {
	var visited = make(map[int]bool)
	var stack []int

	stack = g.edges[startNode]
	visited[startNode] = true

	for _, n := range stack {
		if !visited[n] {
			g.dfs(n)
		}
	}
}

//a dfs which gives path
var path []int

func (g graph) dfsPath(startNode, endNode int) []int {
	var visited = make(map[int]bool)
	var stack []int

	stack = g.edges[startNode]
	visited[startNode] = true

	//path = []int{startNode}

	for _, n := range stack {
		if n == endNode {
			path = append(path, n)
			return path
		}

		if _, ok := visited[n]; !ok {
			path = append(path, n)
			g.dfsPath(n, endNode)
		}

	}
	return path
}

// a simple bfs
func (g graph) bfs(startNode int) {
	var visited = make(map[int]bool)
	visited[startNode] = true

	var queue []int
	queue = append(queue, startNode)

	for 0 < len(queue) {
		//pop the first element
		v := queue[0]
		queue = queue[1:]

		for _, w := range g.edges[v] {
			if !visited[w] {
				visited[w] = true
				queue = append(queue, w)
			}
		}
	}
}

//a bfs which gives path
func (g graph) bfsPath(start, end int) []int {
	var queue [][]int
	node := []int{start}
	queue = append(queue, node)

	for 0 < len(queue) {
		//pop the first element
		path := queue[0]
		queue = queue[1:]

		//last node
		lastNode := path[len(path)-1]
		if lastNode == end {
			return path
		}

		for _, w := range g.edges[lastNode] {
			var new_path = path
			new_path = append(new_path, w)
			queue = append(queue, new_path)
			//log.Println(queue)
		}
	}
	//empty to return sth
	return []int{}

}

func main() {
	/*
	   // n: the number of relationships of influence
	   var n int
	   fmt.Scan(&n)

	   g:=newGraph()

	   for i := 0; i < n; i++ {
	       // x: a relationship of influence between two people (x influences y)
	       var x, y int
	       fmt.Scan(&x, &y)
	       g.edges[x]=append(g.edges[x],y)
	   }
	*/
	//test case
	g := newGraph()
	//g.dfs(1)
	//log.Println(g.nodes,g.edges,g.dfs(1))
	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// The number of people involved in the longest succession of influences
	//var max = 0
	//var height int
	//for k, _ := range g.edges {
	//height := g.dfs(k)
	//height := g.bfs(10, 5)

	//if height > max {
	//	max = height
	//}
	//fmt.Println(height)
	//for k, _ := range g.edges {
	//fmt.Println(g.bfsPath(9, 8))
	fmt.Println(g.dfsPath(5, 8))
	//}
	//}
	//fmt.Printf("Node 10 has depth %d\n",  height)
	//}
	//fmt.Println(max)
}
