package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	//entityType
	EMPTY_CELL = -1
	PLAYER     = 0
	BOMB       = 1
	CRATE      = 2
	//wood league
	BOMB_RANGE = 3
)

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
	c Cell
	//evalScore float64
}
type State struct {
	myId    int  //who i am?
	me      Cell //where i am
	board   [][]Cell
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
}

func (s State) printBoard() string {
	var result bytes.Buffer
	for y := 0; y < 11; y += 1 {
		for x := 0; x < 13; x += 1 {
			switch s.board[y][x].what.entityType {
			case EMPTY_CELL:
				result.WriteString(".")
			case PLAYER:
				//should differentiate between me and opp?
				if s.board[y][x].what.owner == s.myId {
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

/*
func (s *State) think() string {
	//first select a batch of random possible move
	//evaluate them: is this enough far from any given bomb? is this close to
	// a foe (bombs will become lethal)? AND is there any crates around?

	var cells []Cell
	//so first move to a random cell
	//select 10 cells check cratesAround
	for i := 0; i < 10; i++ {
		//x := rand.Intn(12)
		//y := rand.Intn(10)
		//if s.board[y][x].isEmpty() {
		//		cells = append(cells, s.board[y][x])
		//	}
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
	if s.me.what.param1 > 0 {
		res = bomb(cell)
	} else {
		res = move(cell)
	}

	//wait til param1 of bombPlaced == 0
	return res
}
*/

func main() {
	rand.Seed(time.Now().Unix())
	var s State
	for {

		//read Grid
		var width, height, myId int
		fmt.Scan(&width, &height, &myId)
		s.myId = myId

		s.board = make([][]Cell, height)
		for y := 0; y < height; y++ {
			var row string
			fmt.Scan(&row)
			s.board[y] = make([]Cell, width)
			for x := range s.board[y] {
				var c int
				if row[x] == '.' {
					c = EMPTY_CELL
					//and so what??
				} else {
					c = CRATE
					s.crates = append(s.crates, Cell{x, y, Entity{entityType: CRATE}})
				}
				s.board[y][x] = Cell{x: x, y: y, what: Entity{entityType: c}}
			}
		}

		s.printBoard()

		//read Entities
		var entities int
		fmt.Scan(&entities)

		s.bombs = []Cell{}
		s.players = []Cell{}

		for i := 0; i < entities; i++ {
			var entityType, owner, x, y, param1, param2 int
			fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)

			log.Println("x", x, "y", y)

			if owner == s.myId {
				s.me = Cell{x: x, y: y, what: Entity{entityType, owner, param1, param2}}
			} else {
				switch entityType {
				case 0:
					s.board[y][x].what.entityType = entityType
					s.board[y][x].what.owner = owner
					s.board[y][x].what.param1 = param1
					s.board[y][x].what.param2 = param2
					s.players = append(s.players, Cell{x: x, y: y, what: Entity{entityType, owner, param1, param2}})
				case 1:
					s.board[y][x].what.entityType = entityType
					s.board[y][x].what.owner = owner
					s.board[y][x].what.param1 = param1
					s.board[y][x].what.param2 = param2
					s.bombs = append(s.bombs, Cell{x: x, y: y, what: Entity{entityType, owner, param1, param2}})
				}
			}

		}

		//res := s.think()
		//fmt.Println(res)
		//MOVE test
		//should list possible moves+simulate where to leave bombs to get
		//more boxes destroy
		fmt.Println("MOVE 10 10") // Write action to stdout

		//LOGS
		//num := s.cratesAround(s.board[x][y])
		//log.Println("param1:", s.me.what.param1) //, "x:", x, "y:", y, "num", num)
	}
}
