package main
import "fmt"

type Edge struct{
    last_node int
    next_node int
}

type Graph struct{
    nodes []int
    //a node: a list of connected nodes
    edges map[int][]int
}

func (g *Graph) addNode(node int){
    //add the node
    g.nodes=append(nodes,node)
}

func (g *Graph) addEdge(e Edge){
    g.edges[e.last_node]=e.next_node
    //if it's no directed graph:
    g.edges[e.next_node]=e.last_node
}

func main(){
    var g Graph
    g=new(Graph)
}
