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

//a turn is an action + a destination
type Turn struct {
	actionType int
	c          Cell
	//evalScore float64
}
type State struct {
	myId    int //who i am?
	board   Grid
	crates  []Cell
	bombs   []Entity
	players []Entity
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

	s.bombs = []Entity{}
	s.players = []Entity{}

	for i := 0; i < entities; i++ {
		var entityType, owner, x, y, param1, param2 int
		fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)
		switch entityType {
		case 0:
			s.board[x][y].what.entityType = entityType
			s.board[x][y].what.owner = owner
			s.board[x][y].what.param1 = param1
			s.board[x][y].what.param2 = param2
			s.players = append(s.players, Entity{entityType, owner, param1, param2})
		case 1:
			s.board[x][y].what.entityType = entityType
			s.board[x][y].what.owner = owner
			s.board[x][y].what.param1 = param1
			s.board[x][y].what.param2 = param2
			s.bombs = append(s.bombs, Entity{entityType, owner, param1, param2})
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
