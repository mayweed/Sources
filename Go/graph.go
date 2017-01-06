// cf https://www.hackerrank.com/challenges/bfsshortreach
//First post very insightful here: http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	//num nodes
	vertices int
	//a node: a list of connected nodes
	edges map[int][]int
}

func (g Graph) addEdge(node, node2 int) {
	g.edges[node] = append(g.edges[node], node2)
	//if it's no directed graph:
	g.edges[node2] = append(g.edges[node2], node)
}

//want to make a variadic func here: multiple nodes
//should be passed has args
func (g Graph) String() string {
	var s string
	for k, v := range g.edges {
		for _, val := range v {
			s += fmt.Sprintf("%d -> %d\n", k, val)
		}
	}
	return s
}

//same here should be able to pass one or + nodes
func (g Graph) degree() {
	for k, v := range g.edges {
		fmt.Printf("Node n°%d degree:%d\n", k, len(v))
	}
}

/*
//should add an endNode arg later on
func (g Graph) dfs(startNode int){
	var queue []int
	visited=make(map[int]bool)
	queue=g.edges[startNode]
	for m:=range (queue){
		if !visited[m]{
			dfs(m)
		}
		visited[m]=true
	}
}
*/

func main() {
	fi, err := os.Open("mediumG.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	//first two lines: v==num of vertices, e==num of edges
	scanner := bufio.NewScanner(fi)
	scanner.Scan()
	V, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	E, _ := strconv.Atoi(scanner.Text())

	//should put that in a func but test..
	g := new(Graph)
	g.vertices = V
	g.edges = make(map[int][]int)

	for i := 0; i < E; i++ {
		scanner.Scan()
		edges := strings.Split(scanner.Text(), " ")
		node1, _ := strconv.Atoi(edges[0])
		node2, _ := strconv.Atoi(edges[1])
		g.addEdge(node1, node2)
	}

	//fmt.Println(len(g.edges))
	fmt.Println(len(g.edges))

	fmt.Println(g.String())

}
