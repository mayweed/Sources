package main

import (
	"fmt"
	"log"
)

type graph struct {
	nodes []int
	edges map[int][]int
}

//init
func newGraph() graph {
	return graph{
		nodes: nil,
		//edges: make(map[int][]int),
        /*  depth max 6
		edges: map[int][]int{
			5: []int{3, 6},
			6: []int{1, 2},
			7: []int{4},
			9: []int{4},
			4: []int{5},
			2: []int{8},
		},
        */
        //depth max 4
        edges: map[int][]int{
			10: []int{1, 3, 11},
			1: []int{2, 3},
			3: []int{4},
			2: []int{4,5},
		},

		//TODO find a way to quickly parse that
		//oki string: first split on space then split num on ""
		//or num[0] => x num[1] =>y
		// 56  53  61  74  62  94  45  28
	}
}

// cf orderedSet: https://github.com/stevenle/topsort/blob/master/topsort.go
//a int that is the height of the graph
func (g graph) dfs(startNode int) int {
	var visited = make(map[int]bool)
	var stack []int
	var height int

	stack = g.edges[startNode]
	visited[startNode] = true
	//log.Println(visited, stack)

	for _, n := range stack {
		if !visited[n] {
			g.dfs(n)
			//height += 1
		}
	}
	return height
}

//very primitive should write that with 2 queues one 
//for the current and one for the next
// first comm http://stackoverflow.com/questions/10258305/how-to-implement-a-breadth-first-search-to-a-certain-depth/16923440#16923440
func (g graph) bfs(startNode int) int{
	var visited = make(map[int]bool)
	visited[startNode] = true

	var queue []int
	queue = append(queue, startNode)

    var depth int
    var pendingDepthIncrease bool

	for 0 < len(queue) {
        var timeToDepthIncrease=len(queue)
        log.Println("INIT",timeToDepthIncrease)
        //pop the first element
		v := queue[0]
        queue=queue[1:]

        timeToDepthIncrease-=1
        log.Println("END",timeToDepthIncrease)


		for _,w := range g.edges[v] {
			if !visited[w] {
				visited[w] = true
				queue = append(queue, w)
                if pendingDepthIncrease{
                    timeToDepthIncrease=len(queue)-1
                    pendingDepthIncrease=false
                    }
				//log.Println(queue,visited)
            }
        }
        if timeToDepthIncrease==0{
                depth+=1
                pendingDepthIncrease=true
        }

		//}

	}
    return depth
}

func main() {
	/*
	   // n: the number of relationships of influence
	   var n int
	   fmt.Scan(&n)

	   g:=newGraph()

	   for i := 0; i < n; i++ {
	       // x: a relationship of influence between two people (x influences y)
	       var x, y int
	       fmt.Scan(&x, &y)
	       g.edges[x]=append(g.edges[x],y)
	   }
	*/
	//test case
	g := newGraph()
	//g.dfs(1)
	//log.Println(g.nodes,g.edges,g.dfs(1))
	// fmt.Fprintln(os.Stderr, "Debug messages...")

	// The number of people involved in the longest succession of influences
	var max = 0
	//var height int
	for k, _ := range g.edges {
		//height := g.dfs(k)
		height:=g.bfs(k)

		if height > max {
			max = height
		}

		fmt.Printf("Node %d has depth %d\n", k, height)
	}
	fmt.Println(max)
}
