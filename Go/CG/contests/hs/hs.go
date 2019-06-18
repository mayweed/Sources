package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	WIDTH  = 13
	HEIGHT = 11
	//entityType
	EMPTY_CELL = -1
	PLAYER     = 0
	BOMB       = 1
	CRATE      = 2
)

type Grid [WIDTH][HEIGHT]Cell

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	x, y int
	what Entity
}

func move(c Cell) string {
	s := fmt.Sprintf("MOVE %d %d\n", c.x, c.y)
	return s
}
func bomb(c Cell) string {
	s := fmt.Sprintf("BOMB %d %d\n", c.x, c.y)
	return s
}

type Entity struct {
	entityType int
	owner      int //id of the player or id of the one who sets the bomb
	param1     int //for bomb, bomb ticker
	param2     int
}

type Player struct {
	//my position should i use anon struct position here?
	Entity
	//my num of bombs left
	myBombs int
	//can only place one bomb here
	bombsPlaced Cell
}

type State struct {
	board   Grid
	crates  []Cell
	bombs   []Entity
	players []Player
}

func (s *State) readGrid() {
	var width, height, myId int
	fmt.Scan(&width, &height, &myId)

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

	s.bombs = []Entity{}

	for i := 0; i < entities; i++ {
		var entityType, owner, x, y, param1, param2 int
		fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)
		switch entityType {
		case 0:
			s.board[x][y].what.entityType = entityType
			s.board[x][y].what.owner = owner
			s.board[x][y].what.param1 = param1
		case 1:
			s.board[x][y].what.entityType = entityType
			s.board[x][y].what.owner = owner
			s.board[x][y].what.param1 = param1
			s.bombs = append(s.bombs, Entity{entityType, owner, param1, param2})
		}

	}
	//log.Println(s.bombs)
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
				result.WriteString("P")
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

func think() {
	//so first move to a cell

	//then pose a bomb

	//then get out of bomb range

	//wait til param1 of bombPlaced == 0

	//and loop til no bombs left on map?
}
func main() {
	s := State{}

	for {
		s.readGrid()
		s.readEntities()

		//MOVE test
		//should list possible moves+simulate where to leave bombs to get
		//more boxes destroy
		fmt.Println("MOVE 11 13") // Write action to stdout

	}
}
