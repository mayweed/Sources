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

/*
//a int that is the height of the graph
func (g graph) dfs(startNode int) int{
    var visited=make[int]bool
    var stack []int
    stack=g.edges[startNode]
    //!!!LIFO!!!
    //for _,node := range stack{
    for i:=len(stack);i>=0;i--{
        if visited[i]{continue}
        else{
            visited[i]=true
            dfs(i)
        }
    }
}
*/

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
