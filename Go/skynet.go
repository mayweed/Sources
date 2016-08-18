package main

import "fmt"
import "os"

type Graph struct{
    //if true, node is an EI
    exit_gw map[int]bool

    //adjency list
    connections map[int][]int
}
func (g Graph) addLink(n1,n2 int){
    g.connections[n1]=append(g.connections[n1],n2)
    g.connections[n2]=append(g.connections[n2],n1)
}
func (g Graph) addEI(node int){
    g.exit_gw[node]=true
}

func main() {
    // N: the total number of nodes in the level, including the gateways
    // L: the number of links
    // E: the number of exit gateways
    var N, L, E int
    fmt.Scan(&N, &L, &E)

    //init things up w/o a custom func newGraph?
    graph:=new(Graph)
    graph.exit_gw=make(map[int]bool)
    graph.connections=make(map[int][]int)

    for i := 0; i < L; i++ {
        // N1: N1 and N2 defines a link between these nodes
        var N1, N2 int
        fmt.Scan(&N1, &N2)
        graph.addLink(N1,N2)
    }
    for i := 0; i < E; i++ {
        // EI: the index of a gateway node
        var EI int
        fmt.Scan(&EI)
        graph.addEI(EI)
    }

    for {
        // SI: The index of the node on which the Skynet agent is positioned this turn
        var SI int
        fmt.Scan(&SI)
        for _,v := range(graph.connections[SI]){
            if graph.exit_gw[v]{
                fmt.Println(SI,v)
            }else{
                fmt.Println(SI,graph.connections[SI][0])
            }
        }

    fmt.Fprintln(os.Stderr, "Debug messages...",SI,graph.connections[SI],graph.exit_gw)
    }
}
