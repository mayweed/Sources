package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//GRAPH
type Node struct {
	id        string
	name      string
	latitude  float64
	longitude float64
}
type Edge struct {
	to     Node
	weight float64
}
type Graph struct {
	//string==node id in each cases?
	nodes map[string]Node
	//a map of node to slice of edge struct
	route map[string][]Edge
}

//a constructor to handle map init correctly
func newGraph() Graph {
	return Graph{
		nodes: make(map[string]Node),
		route: make(map[string][]Edge),
	}
}

//GRAPH HANDLERS
func (g Graph) addNode(n Node) {
	//g.nodes=make(map[string]Node)
	g.nodes[n.id] = n
}
func (g Graph) addEdge(from, to string) {
	//g.route=make(map[string]Edge)
	e := Edge{g.nodes[to], distNodes(g.nodes[from], g.nodes[to])}
	g.route[from] = append(g.route[from], e)
}

//func (g Graph) deleteEdge

//check func TO REWRITE FOR GRAPHS
//should handle error cases and bette than if!?!?
func (g Graph) toString() {
	if g.nodes != nil {
		for key, n := range g.nodes {
			fmt.Printf("Node %s: %s, %f, %f\n", key, n.name, n.latitude, n.longitude)
		}
	} else {
		fmt.Println("Empty map?")
	}
}

// DISTANCE
//To convert degrees in radians: degrees*PI/180
func degreesToRad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

//calculate dist per formula
func distance(latA, longA, latB, longB float64) float64 {
	x := (longB - longA) * math.Cos(latA+latB/2)
	y := latB - latA
	return math.Sqrt(x*x+y*y) * 6371
}
func distNodes(from, to Node) float64 {
	return distance(from.latitude, from.longitude, to.latitude, to.longitude)
}

//MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var startPoint string
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &startPoint)

	var endPoint string
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &endPoint)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	var myGraph = newGraph()
	for i := 0; i < N; i++ {
		scanner.Scan()
		//clean things up
		st := strings.TrimPrefix(scanner.Text(), "StopArea:")
		stopName := strings.Split(st, ",")

		//multiple value context...+convert
		lat, _ := strconv.ParseFloat(stopName[3], 64)
		lat1 := degreesToRad(lat)
		long, _ := strconv.ParseFloat(stopName[4], 64)
		long1 := degreesToRad(long)

		myGraph.addNode(Node{stopName[0], stopName[1], lat1, long1})
	}

	var M int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &M)

	for i := 0; i < M; i++ {
		scanner.Scan()
		route := strings.Split(scanner.Text(), " ")
		var edge []string
		for _, r := range route {
			edge = append(edge, strings.TrimPrefix(r, "StopArea:"))
		}

		myGraph.addEdge(edge[0], edge[1])
		//log.Println(edge)
	}
	//myGraph.toString()
	//to test if every route is indeed in it!!
	//countMul must be equal to M
	var countMul int
	for _, v := range myGraph.route {
		countMul += len(v)
	}
	fmt.Println(countMul)
	log.Println("My Num of stops:", len(myGraph.nodes), "My num of routes:", len(myGraph.route))
	log.Println("Number of stops:", N, "number of routes:", M)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	//fmt.Println("IMPOSSIBLE")// Write answer to stdout
}
