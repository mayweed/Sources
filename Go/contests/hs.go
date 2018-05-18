package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	WIDTH  = 13
	HEIGHT = 11
	//entityType
	PLAYER = 0
	BOMB   = 1
	CRATE  = 2
)

type Grid [WIDTH][HEIGHT]Cell

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	x, y int
	what Entity
}

type Entity struct {
	entityType int
	owner      int //id of the player or id of the one who sets the bomb
	param1     int
	param2     int
}

type Player struct {
	//my position should i use anon struct position here?
	Entity
	//my num of bombs left
	myBombs int
}

type World struct {
	board   Grid
	crates  []Cell
	bombs   []Cell
	players []Cell
}

func (g Grid) String() string {
	var result bytes.Buffer
	for y := 0; y < HEIGHT; y += 1 {
		for x := 0; x < WIDTH; x += 1 {
			if g[x][y].what.entityType == -1 {
				result.WriteString(".")
			} else {
				if g[x][y].what.entityType == 0 {
					result.WriteString(strconv.Itoa(g[x][y].what.owner))
				} else {
					result.WriteString(strconv.Itoa(g[x][y].what.entityType))
				}
			}
		}
		result.WriteString("\n")
	}
	return result.String()
}

func main() {
	//method readGrid from world?
	var width, height, myId int
	fmt.Scan(&width, &height, &myId)
	var board Grid
	for {
		for y := 0; y < HEIGHT; y++ {
			var row string
			fmt.Scan(&row)
			inputs := strings.Split(row, "")
			for x := 0; x < WIDTH; x++ {
				//THAT WORKS BUT THAT SUCKS!!
				var c int
				if inputs[x] == "." {
					c = -1
				} else {
					//c,_=strconv.Atoi(inputs[x])
					c = 2 //crate!!
				}
				board[x][y] = Cell{x: x, y: y, what: Entity{entityType: c}}
			}
		}

		var entities int
		fmt.Scan(&entities)

		//method updateEntities from world!!
		for i := 0; i < entities; i++ {
			var entityType, owner, x, y, param1, param2 int
			fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)
			//update
			//example:
			board[x][y].what.entityType = entityType
			board[x][y].what.owner = owner
			board[x][y].what.param1 = param1

			//if entityType == 0 {
			//	players = append(players, entity{entityType, owner, x, y, param1, param2})
			//} else if entityType == 1 {
			//	bombs = append(bombs, entity{entityType, owner, x, y, param1, param2})
			//}

		}

		//LOGS
		log.Println(board.String())

		//MOVE test
		//should list possible moves+simulate where to leave bombs to get
		//more boxes destroy
		fmt.Println("MOVE 11 13") // Write action to stdout
	}
}
