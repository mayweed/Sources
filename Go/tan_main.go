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

//check func
func (n Node) toString() {
	fmt.Printf("Node %s: %s, %f, %f\n", n.id, n.name, n.latitude, n.longitude)
}

type Graph struct {
	nodes []Node
	//a map of map to got dist via route[ADBLA][ADBU] for ex
	route map[Node]map[Node]float64
}

///To write+read soniakeys lib on github for implementation idea
//dont forget weight
//func (g Graph) addEdge(from,to Node)

/// DISTANCE
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

	var nodes []Node
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

		nodes = append(nodes, Node{stopName[0], stopName[1], lat1, long1})
	}

	//here should instantiate my graph...
	var M int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &M)

	var edge []string
	for i := 0; i < M; i++ {
		scanner.Scan()
		route := strings.Split(scanner.Text(), " ")
		for _, r := range route {
			edge = append(edge, strings.TrimPrefix(r, "StopArea:"))
		}
		log.Println(edge)
	}

	var edgeWeight float64
	edgeWeight = distance(nodes[edge[0]].latitude, nodes[edge[0]].longitude, nodes[edge[1]].latitude, nodes[edge[1]].longitude)
	log.Println(edgeWeight)
	//fmt.Println("IMPOSSIBLE")// Write answer to stdout
}
