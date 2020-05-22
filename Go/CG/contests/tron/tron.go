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
type Cell struct {
	Point
	owner int
}

func free(c Cell) bool {
	return c.owner == -1
}

func getDir(from, to Cell) string {
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

type Grid [][]Cell

type Turn struct {
	id int
	//action Action
	board   [][]Cell
	nbTurns int
}

func initGrid(width, height int) Grid {
	//no player got -1 id
	var m = -1
	var grid = make([][]Cell, width)
	for x := 0; x < width; x++ {
		grid[x] = make([]Cell, height)
		for y := 0; y < height; y++ {
			grid[x][y] = Cell{Point{x, y}, m}
		}
	}
	return grid
}

//idea : for each adjacent of a given cell do a floodfill and go
//for the max one
func (g Grid) getAdjacent(c Cell) []Cell {
	var adjacents []Cell
	if c.x+1 < WIDTH && g[c.x+1][c.y].owner == -1 {
		adjacents = append(adjacents, g[c.x+1][c.y])
	}
	if c.y+1 < HEIGHT && g[c.x][c.y+1].owner == -1 {
		adjacents = append(adjacents, g[c.x][c.y+1])
	}
	if c.x-1 >= 0 && g[c.x-1][c.y].owner == -1 {
		adjacents = append(adjacents, g[c.x-1][c.y])
	}
	if c.y-1 >= 0 && g[c.x][c.y-1].owner == -1 {
		adjacents = append(adjacents, g[c.x][c.y-1])
	}
	return adjacents
}
func (g Grid) fill(from Cell) int {
	var fillableCell int

	var visited = make(map[Cell]bool)

	var queue []Cell
	queue = append(queue, from)

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		for _, adj := range g.getAdjacent(start) {
			if !visited[adj] {
				queue = append(queue, adj)
				fillableCell += 1
				visited[adj] = true
			}
		}
	}
	return fillableCell
}
func main() {
	//is this the right struct for this?
	var actions = make(map[string][]int)
	//left: x-=1 :)
	actions["LEFT"] = []int{-1, 0}
	actions["RIGHT"] = []int{1, 0}
	actions["UP"] = []int{0, -1}
	actions["DOWN"] = []int{0, 1}

	t := Turn{}
	board := initGrid(WIDTH, HEIGHT)

	var startPos Cell
	for {
		var myPos Cell

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
			if i == P {
				startPos = board[X0][Y0]
				myPos = board[X1][Y1]
			}
			board[X1][Y1].owner = i
		}
		adj := board.getAdjacent(myPos)
		fmt.Fprintln(os.Stderr, startPos, adj[0], board.fill(adj[0]))
		t.nbTurns += 1
		//testing
		if free(adj[0]) {
			fmt.Println(getDir(myPos, adj[0]))
		}
		//fmt.Println("LEFT") // A single line with UP, DOWN, LEFT or RIGHT
	}
}
