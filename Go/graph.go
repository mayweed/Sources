// cf https://www.hackerrank.com/challenges/bfsshortreach
// for input examples

package main
import "fmt"
//import "bufio" >> scanner to read file
//import "bytes"
//import "strings"
//import "os" >> needs Open() for file
import "io/ioutil"
/*
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
    g.nodes=append(g.nodes,node)
}

func (g *Graph) addEdge(e Edge){
    g.edges[e.last_node]=e.next_node
    //if it's no directed graph:
    g.edges[e.next_node]=e.last_node
}
*/
func main(){
    //Scanner better to split stdin
    //s:=bufio.NewScanner(os.Stdin)
    //for s.Scan() {
    //    fmt.Println(scanner.Text)
    //Should do the same for a file
    //var g Graph
    //g=new(&Graph)
}
