// cf https://www.hackerrank.com/challenges/bfsshortreach
//First post very insightful here: http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file
//well...
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type graph struct {
	//num nodes
	nodes []int
	//a node: a list of connected nodes
	edges map[int][]int
}

func (g graph) addEdge(node, node2 int) {
	g.edges[node] = append(g.edges[node], node2)
	//if it's no directed graph:
	g.edges[node2] = append(g.edges[node2], node)
}

//want to make a variadic func here: multiple nodes
//should be passed has args
func (g graph) String() string {
	var s string
	for k, v := range g.edges {
		for _, val := range v {
			s += fmt.Sprintf("%d -> %d\n", k, val)
		}
	}
	return s
}

//same here should be able to pass one or + nodes
func (g graph) degree() {
	for k, v := range g.edges {
		fmt.Printf("Node n°%d degree:%d\n", k, len(v))
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
				parent[w] = v
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
		}
	}
	//empty to return sth
	return []int{}

}

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
	g := new(graph)
	g.vertices = V
	g.edges = make(map[int][]int)

	for i := 0; i < E; i++ {
		scanner.Scan()
		edges := strings.Split(scanner.Text(), " ")
		node1, _ := strconv.Atoi(edges[0])
		node2, _ := strconv.Atoi(edges[1])
		g.addEdge(node1, node2)
	}

	fmt.Println(len(g.edges))

	fmt.Println(g.String())

}
