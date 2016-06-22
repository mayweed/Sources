// cf https://www.hackerrank.com/challenges/bfsshortreach
// for input examples

package main
import "fmt"
//import "bufio" >> scanner to read file
//import "bytes"
//import "strings"
//import "os" >> needs Open() for file
//import "io/ioutil"

type Graph struct{
    //num nodes
    nodes int
    //num vertices
    vertices int
    //a node: a list of connected nodes
    edges map[int][]int
}

func (g *Graph) addEdge(node,node2 int){
    g.edges[node]=append(g.edges[node],node2)
    //if it's no directed graph:
    g.edges[node2]=append(g.edges[node2],node)
}

func NewGraph(n,v int) *Graph{
    return &Graph{
        nodes:n,
        vertices : v,
        edges : make(map[int][]int,n),
    }
}

func main(){
    //Scanner better to split stdin
    //s:=bufio.NewScanner(os.Stdin)
    //for s.Scan() {
    //    fmt.Println(scanner.Text)
    //Should do the same for a file
    //var g Graph
    g:=NewGraph(4,5)
    g.addEdge(0,1)
    fmt.Println(g.edges)
}
