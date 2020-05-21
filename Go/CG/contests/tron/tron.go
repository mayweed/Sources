package main

import (
	"fmt"
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
type Grid [][]Cell

type Turn struct {
	id int
	//action Action
	board   [][]Cell
	nbTurns int
}

func initGrid(width, height int) Grid {
	//a simple grid made of cells
	var i, j int
	//no player got -1 id
	var m = -1
	var grid = make([][]Cell, height)
	for i = 0; i < height; i++ {
		grid[i] = make([]Cell, width)
		for j = range grid[i] {
			grid[i][j] = Cell{Point{i, j}, m}
		}
	}
	return grid
}
func (g Grid) possibleMoves(c Cell) /*???*/ {
	if c.x+1 < WIDTH && g[c.x+1][c.y].owner == -1 {
	}
	if c.y+1 < HEIGHT && g[c.x][c.y+1].owner == -1 {
	}
	if c.x-1 >= 0 && g[c.x-1][c.y].owner == -1 {
	}
	if c.y-1 >= 0 && g[c.x][c.y-1].owner == -1 {
	}
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

	for {
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
			board[Y1][X1].owner = i
		}
		t.nbTurns += 1
		//fmt.Println("LEFT") // A single line with UP, DOWN, LEFT or RIGHT
	}
}
