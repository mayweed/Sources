// cf https://www.hackerrank.com/challenges/bfsshortreach
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
	nodes int
	//num vertices
	vertices int
	//a node: a list of connected nodes
	edges map[int][]int
}

func (g Graph) addEdge(node, node2 int) {
	g.edges[node] = append(g.edges[node], node2)
	//if it's no directed graph:
	g.edges[node2] = append(g.edges[node2], node)
}

func main() {
	//Scanner better to split stdin
	//s:=bufio.NewScanner(os.Stdin)
	//for s.Scan() {
	//    fmt.Println(scanner.Text)

	//Should do the same for a file
	fi, err := os.Open("mediumG.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	//first two lines: v==num of vertices, e== num of edges
	scanner := bufio.NewScanner(fi)
	scanner.Scan()
	V, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	E, _ := strconv.Atoi(scanner.Text())

	//should put that in a func but test..
	g := new(Graph)
	g.nodes = V
	g.edges = make(map[int][]int)

	for i := 0; i < E; i++ {
		scanner.Scan()
		edges := strings.Split(scanner.Text(), " ")
		node1, _ := strconv.Atoi(edges[0])
		node2, _ := strconv.Atoi(edges[1])
		g.addEdge(node1, node2)
	}

	//fmt.Println(len(g.edges))
	fmt.Println(g.edges)

}
