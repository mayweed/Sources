package main

import (
	"fmt"
	"log"
)

const (
	WIDTH  = 30
	HEIGHT = 20
)

type Cell struct {
	x, y  int
	owner int
}
type Board struct {
	width  int //=30
	height int //=20
	//list of points
	cells [][]Cell
}
type Player struct {
	id       int
	startPos Cell
	lastPos  Cell
}
type gameState struct {
	//where int is the id?
	players map[int]Player
	board   Board
}

//should be inited with w h // gameState??
func initBoard(width, height int) Board { //dont need that no?,start ...point)board{
	//a simple grid made of cells
	var i, j int
	//no player got -1 id
	var m = -1
	var grid = make([][]Cell, height)
	for i = 0; i < height; i++ {
		grid[i] = make([]Cell, width)
		for j = range grid[i] {
			grid[i][j] = Cell{i, j, m}
		}
	}

	return Board{
		width:  width,
		height: height,
		cells:  grid,
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

	g := gameState{players: make(map[int]Player)}
	board := initBoard(WIDTH, HEIGHT)
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
			//i== player, first is me that is 0
			g.players[i] = Player{i, Cell{X0, Y0, i}, Cell{X1, Y1, i}}
			board.cells[Y1][X1].owner = i
			log.Println(X0, Y0, X1, Y1)
		}

		//Does it work? Seems so
		//for dy:=0;dy<HEIGHT;dy++{
		//    for dx:=0;dx<WIDTH;dx++{
		//        if board.cells[dy][dx].owner==0{
		//            log.Println(board.cells[dy][dx])
		//        }
		//    }
		//}
		log.Println(g.players[0])
		fmt.Println("LEFT") // A single line with UP, DOWN, LEFT or RIGHT
	}
}
