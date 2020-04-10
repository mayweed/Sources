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

//POINT
type Point struct {
	x, y int
}

//distance in a grid idea voronoi
func (p Point) manhattanDist(p2 Point) float64 {
	var dx = p2.x - p.x
	var dy = p2.y - p.y
	return math.Abs(float64(dx)) + math.Abs(float64(dy))
}

//CELL
type Cell struct {
	pos       Point //might use anon point but...
	what      string
	hasMe     bool
	hasPlayer bool
	hasBomb   bool
}

func (c Cell) isEmpty() bool {
	if c.what == "." {
		return true
	} else {
		return false
	}
}
func (c Cell) hasCrate() bool {
	if c.what == "0" {
		return true
	} else {
		return false
	}
}

//GRID
type Grid struct {
	w      int
	h      int
	rows   string
	c      [][]Cell
	crates []Cell
}

//It's NOT always 3 if you meet a crate before it's that range!!
//eg: if the nearest crate is in range 2, range for the bomb is two!!
func (g Grid) cratesAround(c Cell) (int, []Cell) {
	var bombRange = 3
	var numCrates int
	var crateCells []Cell
	//for any given free cell let's see how many crates are in range
	// DID I REALLY NEED Round here??
	for _, crate := range g.crates {
		if c.pos.x == crate.pos.x && math.Round(math.Abs(float64(c.pos.y-crate.pos.y))) <= bombRange {
			crateCells = append(crateCells, crate)
			numCrates += 1
			bombRange = math.Round(math.Abs(float64(c.pos.y - crate.pos.y))) //update brange
		}
		//||
		if c.pos.y == crate.pos.y && math.Round(math.Abs(float64(c.pos.x-crate.pos.x))) <= bombRange {
			crateCells = append(crateCells, crate)

			numCrates += 1
			bombRange = math.Round(math.Abs(float64(c.pos.x - crate.pos.x))) //update brange

		}
	}
	return numCrates, crateCells
}
func (g Grid) getCellFromXY(x, y int) Cell {
	return g.c[x][y]
}
func (g Grid) getCellFromPoint(p Point) Cell {
	return g.c[p.x][p.y]
}

type Player struct {
	position       Point
	id             int
	numOfBombsLeft int
	rangeOfBombs   int
}

type Bomb struct {
	position  Point
	ownerId   int
	countdown int
	expRange  int
}

type State struct {
	me      Player
	board   Grid
	players []Player
	bombs   []Bomb
}

//TURN and Action
//a turn is an action + a destination
type Turn struct {
	c             Cell
	possibleMoves []Cell
	//evalScore float64
}

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
		if c.pos.x == crate.pos.x && math.Round(math.Abs(float64(c.pos.y-crate.pos.y))) <= BOMB_RANGE ||
			c.pos.y == crate.pos.y && math.Round(math.Abs(float64(c.pos.x-crate.pos.x))) <= BOMB_RANGE {
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
				s.board.c[x][y] = Cell{Point{x, y}, string(s.board.rows[y*width+x]), false, false, false}
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
				s.me.position = Point{x, y}
				s.me.id = owner
				s.me.numOfBombsLeft = param1
				s.me.rangeOfBombs = param2
				s.board.c[x][y].hasMe = true
			} else {
				switch entityType {
				case 0:
					s.board.c[x][y].hasPlayer = true
					s.players = append(s.players, Player{position: Point{x, y}, id: owner, numOfBombsLeft: param1, rangeOfBombs: param2})
				case 1:
					s.board.c[x][y].hasBomb = true
					s.bombs = append(s.bombs, Bomb{position: Point{x, y}, ownerId: owner, countdown: param1, expRange: param2})
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
