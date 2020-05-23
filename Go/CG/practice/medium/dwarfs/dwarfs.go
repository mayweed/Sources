package main

import (
	"fmt"
	"os"
)

//GRAPH
//type Graph map[int][]int ??
type Graph struct {
	nodes []int
	edges map[int][]int
}

//count directly here?
func (g Graph) maxDepth(node int) int {
	var depth int
	//This one should stop when g.edges[node] is empty!!
	//There is no more nodes to visit
	if len(g.edges[node]) > 0 {
		for _, n := range g.edges[node] {
			depth = g.maxDepth(n) + 1
		}
	} else {
		return depth
	}
	return depth
}

//https://www.geeksforgeeks.org/print-paths-given-source-destination-using-bfs/
/*
Algorithm :

create a queue which will store path(s) of type vector
initialise the queue with first path starting from src

Now run a loop till queue is not empty
   get the frontmost path from queue
   check if the lastnode of this path is destination
       if true then print the path
   run a loop for all the vertices connected to the
   current vertex i.e. lastnode extracted from path
      if the vertex is not visited in current path
         a) create a new path from earlier path and
             append this vertex
         b) insert this new path to queue
*/
func (g Graph) bfs(node int) [][]int {
	//this one stops where there is no node left to go to...
	var visited = make(map[int]bool)
	var queue [][]int
	var path []int
	path = append(path, node)
	queue = append(queue, path)

	for len(queue) > 0 {
		startPath := queue[0]
		queue = queue[1:]
		for _, neigh := range g.edges[startPath[len(startPath)-1]] {
			if !visited[neigh] {
				visited[neigh] = true
				//var newpath = append(newpath, neigh)
				queue = append(queue, startPath)
			}
		}
	}
	//should yield all the possible paths in a graph from a given node
	//return //what??
}

//MAIN
func main() {
	// n: the number of relationships of influence
	var n int
	fmt.Scan(&n)
	g := Graph{edges: make(map[int][]int)}

	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
		g.edges[x] = append(g.edges[x], y)
	}

	//I should ( must?) put that in maxDepth no?
	var max = 0
	for n, _ := range g.edges {
		if g.maxDepth(n) >= max {
			max = g.maxDepth(n)
		}
	}

	//LOGS
	fmt.Fprintln(os.Stderr, g.edges, g.bfs(1))
	// The number of people involved in the longest succession of influences
	fmt.Println(max)
}
