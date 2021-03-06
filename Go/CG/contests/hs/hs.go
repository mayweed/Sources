package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	MAX_BOMB_RANGE = 3
)

type (
	point struct {
		x, y int
	}
	Bomb struct {
		position  point
		ownerId   int
		countdown int
		expRange  int
	}
	//a turn is an action + a destination
	Turn struct {
		c             Cell
		possibleMoves []Cell
		//evalScore float64
	}
	Cell struct {
		pos       point //might use anon point but...
		what      string
		hasMe     bool
		hasPlayer bool
		hasBomb   bool
	}
	Grid struct {
		w      int
		h      int
		rows   string
		c      [][]Cell
		crates []Cell
	}
	Player struct {
		position       point
		id             int
		numOfBombsLeft int
		rangeOfBombs   int
	}
	State struct {
		me      Player
		board   Grid
		players []Player
		bombs   []Bomb
	}
)

//distance in a grid idea voronoi
func (p point) manhattanDist(p2 point) float64 {
	var dx = p2.x - p.x
	var dy = p2.y - p.y
	return math.Abs(float64(dx)) + math.Abs(float64(dy))
}

func (c Cell) isEmpty() bool {
	if c.what == "." {
		return true
	}
	return false
}

func (c Cell) hasCrate() bool {
	if c.what == "0" {
		return true
	}
	return false
}

func (g *Grid) printAnswer() {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			fmt.Print(g.c[x][y].what)
		}
		fmt.Println()
	}
	fmt.Println()
}

//It's NOT always 3 if you meet a crate before it's that range!!
//eg: if the nearest crate is in range 2, range for the bomb is two!!
func (g Grid) cratesAround(c Cell) (int, []Cell) {
	var bombRange = 3.0
	var numCrates int
	var crateCells []Cell
	for _, crate := range g.crates {
		//if it's equal to the current br
		if c.pos.x == crate.pos.x && math.Round(math.Abs(float64(c.pos.y-crate.pos.y))) >= bombRange || c.pos.y == crate.pos.y && math.Round(math.Abs(float64(c.pos.x-crate.pos.x))) >= bombRange {
			crateCells = append(crateCells, crate)
			numCrates++
		}
	}
	return numCrates, crateCells
}
func (g Grid) getCellFromXY(x, y int) Cell {
	return g.c[x][y]
}
func (g Grid) getCellFromPoint(p point) Cell {
	return g.c[p.x][p.y]
}

//choose also moves by dist
//can calc num bomb on path?
func getPossibleMoves(s State) Cell {
	var cells []Cell
	//get ten possible actions
	for i := 0; i < 10; i++ {
		x := rand.Intn(12)
		y := rand.Intn(10)
		if s.board.c[x][y].isEmpty() {
			cells = append(cells, s.board.c[x][y])
		}
	}
	var max int
	var dest Cell
	for _, c := range cells {
		num, _ := s.board.cratesAround(c)
		if num > max {
			max = num
			dest = c
		}
	}
	return dest

}

//should make it generic, so that it works for every player
//idea: simulate bomb explosion
func (s *State) simBombTurn(c Cell) {
	g := s.board //copy grid
	for _, crate := range g.crates {
		if c.pos.x == crate.pos.x && math.Round(math.Abs(float64(c.pos.y-crate.pos.y))) <= MAX_BOMB_RANGE ||
			c.pos.y == crate.pos.y && math.Round(math.Abs(float64(c.pos.x-crate.pos.x))) <= MAX_BOMB_RANGE {
			g.c[crate.pos.x][crate.pos.y].what = "." //wipe out crate
		}
	}
}

//should be in action type
func move(c Cell) string {
	s := fmt.Sprintf("MOVE %d %d", c.pos.x, c.pos.y)
	return s
}
func bomb(c Cell) string {
	s := fmt.Sprintf("BOMB %d %d", c.pos.x, c.pos.y)
	return s
}

/*
func (s *State) think() string {
	//first select a batch of random possible move
	//evaluate them: is this enough far from any given bomb? is this close to
	// a foe (bombs will become lethal)? AND is there any crates around?

	//so first move to a random cell
	//select 10 cells check cratesAround
	//LIST POSSIBLE ACTIONS!!

	//very light eval, should take into account the range of others players bomb (and
	//mine too, watch out not be killed by my own bombs!!)
	//var res string
	//attempt to understand
	//OOPS you go first and bomb then!!
	//so calculate a path with most crates and bomb at a given time?
	//if s.me.numOfBombsLeft > 0 {
	//j		res = bomb(cell)
	//	} else {
	//		res = move(cell)
	////	}

	//wait til param1 of bombPlaced == 0
	//return res
}
*/

func main() {
	//read Grid
	var width, height, myId int
	fmt.Scan(&width, &height, &myId)

	rand.Seed(time.Now().Unix())
	var s State
	//init Grid
	//var g Grid
	s.board.h = height
	s.board.w = width

	for {
		for y := 0; y < height; y++ {
			var row string
			fmt.Scan(&row)
			s.board.rows += row
		}
		s.board.c = make([][]Cell, s.board.w)
		for x := 0; x < s.board.w; x++ {
			s.board.c[x] = make([]Cell, s.board.h)
			for y := 0; y < s.board.h; y++ {
				s.board.c[x][y] = Cell{point{x, y}, string(s.board.rows[y*width+x]), false, false, false}
			}
		}

		//isolate crates
		for x := 0; x < s.board.w; x++ {
			for y := 0; y < s.board.h; y++ {

				if s.board.c[x][y].what == "0" {
					s.board.crates = append(s.board.crates, s.board.c[x][y])
				}
			}
		}

		//read Entities
		var entities int
		fmt.Scan(&entities)

		for i := 0; i < entities; i++ {
			var entityType, owner, x, y, param1, param2 int
			fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)

			if owner == myId {
				s.me.position = point{x, y}
				s.me.id = owner
				s.me.numOfBombsLeft = param1
				s.me.rangeOfBombs = param2
				s.board.c[x][y].hasMe = true
			} else {
				switch entityType {
				case 0:
					s.board.c[x][y].hasPlayer = true
					s.players = append(s.players, Player{position: point{x, y}, id: owner, numOfBombsLeft: param1, rangeOfBombs: param2})
				case 1:
					s.board.c[x][y].hasBomb = true
					s.bombs = append(s.bombs, Bomb{position: point{x, y}, ownerId: owner, countdown: param1, expRange: param2})
				}
			}

		}
		c := getPossibleMoves(s)
		log.Println("DEST: ", c)

		_, cs := s.board.cratesAround(c)
		log.Println("CRATES IN RANGE: ", cs)
		log.Println(len(cs))
		//res := s.think()
		//fmt.Println(res)

		//MOVE test
		//should list possible moves+simulate where to leave bombs to get
		//more boxes destroy
		fmt.Println("MOVE 10 10") // Write action to stdout
		s.board.crates = []Cell{} //empties it!!
		s.board.rows = ""         //to update map when crates get bombed
	}
}
