package main

import "testing"

/*
//some graphs to test maxDepth()
func newGraph() graph {
	return graph{
		nodes: nil,
		//  depth max 6
		edges: map[int][]int{
			5: []int{3, 6},
			6: []int{1, 2},
			7: []int{4},
			9: []int{4},
			4: []int{5},
			2: []int{8},
		},
		//(1,2),(2,3),(2,4),(3,4),(4,5)
		//should be five...
		edges: map[int][]int{
			1: []int{2},
			2: []int{3, 4},
			3: []int{4},
			4: []int{5},
		},
		//depth max 4
		edges: map[int][]int{
			10: []int{1, 3, 11},
			1:  []int{2, 3},
			3:  []int{4},
			2:  []int{4, 5},
		},
	}
}
*/

func TestMaxDepth(t *testing.T) {
	//a graph to test
	g := graph{
		nodes: nil,
		edges: map[int][]int{
			5: []int{3, 6},
			6: []int{1, 2},
			7: []int{4},
			9: []int{4},
			4: []int{5},
			2: []int{8},
		},
	}

	total := g.maxDepth(5)
	if total != 6 {
		t.Errorf("Depth was incorrect, got : %d, want : %d.", total, 6)
	}
}
