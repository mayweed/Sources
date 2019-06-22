package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	WIDTH  = 13
	HEIGHT = 11
	//entityType
	EMPTY_CELL = -1
	PLAYER     = 0
	BOMB       = 1
	CRATE      = 2
	//wood league
	BOMB_RANGE = 3
	//actionType
	m actionType = 3
	b actionType = 4
)

type Grid [WIDTH][HEIGHT]Cell
type actionType int

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	x, y int
	what Entity
}

func (c *Cell) isEmpty() bool {
	return c.what.entityType == -1
}
func move(c Cell) string {
	s := fmt.Sprintf("MOVE %d %d", c.x, c.y)
	return s
}
func bomb(c Cell) string {
	s := fmt.Sprintf("BOMB %d %d", c.x, c.y)
	return s
}

type Entity struct {
	entityType int
	owner      int //id of the player or id of the one who sets the bomb
	param1     int //for bomb, bomb ticker
	param2     int
}

//a turn is an action + a destination
type Turn struct {
	actionType int
	c          Cell
	//evalScore float64
}
type State struct {
	myId    int  //who i am?
	me      Cell //where i am
	board   Grid
	crates  []Cell
	bombs   []Cell
	players []Cell
}

func (s *State) cratesAround(c Cell) int {
	var numCrates int
	//for any given free cell let's see how many crates are in range
	for _, crate := range s.crates {
		if c.x == crate.x && math.Round(math.Abs(float64(c.y-crate.y))) <= BOMB_RANGE ||
			c.y == crate.y && math.Round(math.Abs(float64(c.x-crate.x))) <= BOMB_RANGE {
			numCrates += 1
		}
	}
	return numCrates
}

//should make it generic, so that it works for every player
func (s *State) applyTurn(t Turn) {
	//copy of the state of the current board
	cpBoard := s.board

	switch t.actionType {
	case 3:
	case 4:
		//place a bomb
		cpBoard[t.c.x][t.c.y] = Cell{x: t.c.x, y: t.c.y, what: Entity{entityType: BOMB}}
	}
}
func (s *State) readGrid() {
	var width, height, myId int
	fmt.Scan(&width, &height, &myId)

	s.myId = myId

	for y := 0; y < height; y++ {
		var row string
		fmt.Scan(&row)
		inputs := strings.Split(row, "")
		for x := 0; x < width; x++ {
			var c int
			if inputs[x] == "." {
				c = EMPTY_CELL
			} else {
				c = CRATE
				s.crates = append(s.crates, Cell{x, y, Entity{entityType: CRATE}})
			}
			s.board[x][y] = Cell{x: x, y: y, what: Entity{entityType: c}}
		}
	}
}
func (s *State) readEntities() {
	var entities int
	fmt.Scan(&entities)

	s.bombs = []Cell{}
	s.players = []Cell{}

	for i := 0; i < entities; i++ {
		var entityType, owner, x, y, param1, param2 int
		fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)
		if owner == s.myId {
			s.me = Cell{x: x, y: y, what: Entity{entityType, owner, param1, param2}}
		} else {
			switch entityType {
			case 0:
				s.board[x][y].what.entityType = entityType
				s.board[x][y].what.owner = owner
				s.board[x][y].what.param1 = param1
				s.board[x][y].what.param2 = param2
				s.players = append(s.players, Cell{x: x, y: y, what: Entity{entityType, owner, param1, param2}})
			case 1:
				s.board[x][y].what.entityType = entityType
				s.board[x][y].what.owner = owner
				s.board[x][y].what.param1 = param1
				s.board[x][y].what.param2 = param2
				s.bombs = append(s.bombs, Cell{x: x, y: y, what: Entity{entityType, owner, param1, param2}})
			}
		}

	}
}
func (s State) printBoard() string {
	var result bytes.Buffer
	for y := 0; y < HEIGHT; y += 1 {
		for x := 0; x < WIDTH; x += 1 {
			switch s.board[x][y].what.entityType {
			case EMPTY_CELL:
				result.WriteString(".")
			case PLAYER:
				//should differentiate between me and opp?
				if s.board[x][y].what.owner == s.myId {
					result.WriteString("Me")
				} else {
					result.WriteString("P")
				}
			case BOMB:
				result.WriteString("B")
			case CRATE:
				result.WriteString("C")
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func (s *State) think() string {
	//first select a batch of random possible move
	//evaluate them: is this enough far from any given bomb? is this close to
	// a foe (bombs will become lethal)? AND is there any crates around?

	var cells []Cell

	//so first move to a random cell
	//select 10 cells check cratesAround
	for i := 0; i < 10; i++ {
		x := rand.Intn(12)
		y := rand.Intn(10)
		if s.board[x][y].isEmpty() {
			cells = append(cells, s.board[x][y])
		}
	}

	//very light eval, should take into account the range of others players bomb (and
	//mine too, watch out not be killed by my own bombs!!)
	var max int
	var cell Cell
	for _, c := range cells {
		num := s.cratesAround(c)
		if num > max {
			max = num
			cell = c
		}
	}
	var res string
	//attempt to understand
	if s.me.what.param1 == 1 {
		res = bomb(cell)
	} else {
		res = move(cell)
	}

	//wait til param1 of bombPlaced == 0
	return res
}
func main() {
	rand.Seed(time.Now().Unix())
	s := State{}
	for {
		s.readGrid()
		s.readEntities()

		res := s.think()
		fmt.Println(res)
		//MOVE test
		//should list possible moves+simulate where to leave bombs to get
		//more boxes destroy
		//fmt.Println("MOVE 10 10") // Write action to stdout

		//LOGS
		//num := s.cratesAround(s.board[x][y])
		log.Println("param1:", s.me.what.param1) //, "x:", x, "y:", y, "num", num)

	}
}
