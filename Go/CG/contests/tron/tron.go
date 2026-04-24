package main

import (
	"fmt"
	"os"
)

const (
	WIDTH  = 30
	HEIGHT = 20
)

type Point struct {
	x, y int
}

type TronState struct {
    Width, Height int
    Walls  map[Point]bool

    myPos  Point
    oppPos Point

    MeDead  bool
    OppDead bool
}

func NewTronState(width, height int) TronState {
    return TronState{
        Width:   width,
        Height: height,
        Walls:  make(map[Point]bool),
    }
}

func (t TronState) isFree(c Point) bool {
	return !t.Walls[c]
}

func getDir(from, to Point) string {
	var dir string
	if to.x < from.x {
		dir = "LEFT"
	}
	if to.x > from.x {
		dir = "RIGHT"
	}
	if to.y < from.y {
		dir = "UP"
	}
	if to.y > from.y {
		dir = "DOWN"
	}
	return dir
}

//idea : for each adjacent of a given cell do a floodfill and go
//for the max one
func (t TronState) getAdjacent(c Point) []Point {
    var adj []Point

    dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    for _, d := range dirs {
        n := Point{c.x + d.x, c.y + d.y}

        if n.x < 0 || n.x >= WIDTH || n.y < 0 || n.y >= HEIGHT {
            continue
        }
        if t.Walls[n] {
            continue
        }
        adj = append(adj, n)
    }
    return adj
}
//replace by a voronoi?
//Your voronoiScore:
/*
Start BFS from:
You
All enemies -> prdre en cpte qu’il peut y en avoir plus d’un.
Expand simultaneously
Each cell gets an owner:
Whoever reaches it first
Count:
if mine → +1
if enemy → -1

func (g Grid) voronoi(myPos Point, oppPos Point) int {

	//do i need that? owner?
	var claimed = make(map[Point.Point]bool)
	claimed[myPos.Point] = true
	claimed[oppPos.Point] = true

	//one queue for me, one for opp? or all in the same queue?
	var queue []Point
	queue = append(queue, myPos, oppPos)

		for len(queue) > 0 {
				start := queue[0]
				queue = queue[1:]

				if !claimed[start]{
					start.owner = 0
				}

				for _, adj := range g.getAdjacent(start) {
					if !claimed[adj] {
						queue = append(queue, adj)
						adj.owner = 1 //one is for me
						claimed[adj] = true
					}
				}
			}

		}
	}
	return 0
}

*/
func (t TronState) fill(from Point) int {
	var fillablePoint int

	var visited = make(map[Point]bool)
	visited[from] = true

	var queue []Point
	queue = append(queue, from)

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for _, adj := range t.getAdjacent(start) {
			if !visited[adj] {
				queue = append(queue, adj)
				fillablePoint += 1
				visited[adj] = true
			}
		}
	}
	return fillablePoint
}
func main() {
	//is this the right struct for this?
	var actions = make(map[string][]int)
	//left: x-=1 :)
	actions["LEFT"] = []int{-1, 0}
	actions["RIGHT"] = []int{1, 0}
	actions["UP"] = []int{0, -1}
	actions["DOWN"] = []int{0, 1}

	state := NewTronState(WIDTH,HEIGHT)

	for {
		//state.MeDead = false
		//state.OppDead = false
		// N: total number of players (2 to 4).
		// P: your player number (0 to 3).
		var N, P int
		fmt.Scan(&N, &P)

		for i := 0; i < N; i++ {
			// X0: starting X coordinate of lightcycle (or -1)
			// Y0: starting Y coordinate of lightcycle (or -1)
			// X1: starting X coordinate of lightcycle (can be the same as X0 if you play before this player)
			// Y1: starting Y coordinate of lightcycle (can be the same as Y0 if you play before this player)
			var X0, Y0, X1, Y1 int
			fmt.Scan(&X0, &Y0, &X1, &Y1)

			if X0 == -1 && Y0 == -1 && X1 == -1 && Y1 == -1{
				if i == P{
					state.MeDead = true //bad luck happens
				}
				continue
			}

			// si je vis
			if i == P {
				state.myPos = Point{X1,Y1}
			}else{
				state.oppPos = Point{X1,Y1}
			}

			//si x0 le bot tente de revenir sur ses pas!!
			state.Walls[Point{X1,Y1}] = true //on ne distingue pas moi/adversaires
			
			fmt.Fprintln(
				os.Stderr,
				"player", i,
				"X0", X0, "Y0", Y0,
				"X1", X1, "Y1", Y1,
			)
			
		}
		adj := state.getAdjacent(state.myPos)
		
		bestScore := -1
		bestPoint := Point{}

		for _, cell := range adj {
			score := state.fill(cell)

			if score > bestScore {
				bestScore = score
				bestPoint = cell
			}
		}

		if bestScore != -1 {
			fmt.Println(getDir(state.myPos, bestPoint))
		} else {
			fmt.Println("UP") // fallback (avoid crash)
		}
	}
}
