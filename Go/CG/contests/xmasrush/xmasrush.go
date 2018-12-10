package main

import (
	"fmt"
	"log"
)

/* not good
const (
	UP    = Point{0, -1}
	RIGHT = Point{+1, 0}
	DOWN  = Point{0, +1}
	LEFT  = Point{-1, 0}
)
*/
type Point struct {
	x, y int
}

type Grid [7][7]Tile

type Tile struct {
	direction    string
	position     Point
	hasItem      bool
	itemName     string
	itemPlayerId int
}

type Player struct {
	totalOfQuests int
	quests        []string
	tile          Tile
	position      Point
}

type Turn struct {
	turnType int
	//action Action
}

type State struct {
	players  [2]Player
	grid     Grid
	numItems int
	turn     Turn
}

func (s State) getNeighbours(t Tile) []Tile {
	var neighbours []Tile
	//check borders!!
	if t.direction[0] == 1 {
		neighbours = append(neighbours, s.grid[t.position.y-1][t.position.x])
	}
	if t.direction[1] == 1 {
		neighbours = append(neighbours, s.grid[t.position.y][t.position.x+1])
	}
	if t.direction[2] == 1 {
		neighbours = append(neighbours, s.grid[t.position.y+1][t.position.x])
	}
	if t.direction[3] == 1 {
		neighbours = append(neighbours, s.grid[t.position.y][t.position.x-1])
	}
	log.Println(neighbours)
	return neighbours
}

func (s *State) read() {
	var turnType int
	fmt.Scan(&turnType)
	s.turn.turnType = turnType

	//cf GameBoard => sendMapToPlayer()
	for y := 0; y < 7; y++ {
		for x := 0; x < 7; x++ {
			var row string
			fmt.Scan(&row)
			s.grid[y][x].direction = row
			s.grid[y][x].position.x = x
			s.grid[y][x].position.y = y
		}
	}

	for i := 0; i < 2; i++ {
		// numPlayerCards: the total number of quests for a player (hidden and revealed)
		var numPlayerCards, playerX, playerY int
		var playerTile string
		fmt.Scan(&numPlayerCards, &playerX, &playerY, &playerTile)
		s.players[i].totalOfQuests = numPlayerCards
		s.players[i].position.x = playerX
		s.players[i].position.y = playerY
		s.players[i].tile.direction = playerTile
	}

	// numItems: the total number of items available on board and on player tiles
	var numItems int
	fmt.Scan(&numItems)
	s.numItems = numItems

	for i := 0; i < numItems; i++ {
		var itemName string
		var itemX, itemY, itemPlayerId int
		fmt.Scan(&itemName, &itemX, &itemY, &itemPlayerId)
		switch itemX {
		case -1:
			s.players[0].tile.hasItem = true
		case -2:
			s.players[1].tile.hasItem = true
		default:
			s.grid[itemY][itemX].hasItem = true
			s.grid[itemY][itemX].itemName = itemName
			s.grid[itemY][itemX].itemPlayerId = itemPlayerId
		}
	}

	// numQuests: the total number of revealed quests for both players
	var numQuests int
	fmt.Scan(&numQuests)

	for i := 0; i < numQuests; i++ {
		var questItemName string
		var questPlayerId int
		fmt.Scan(&questItemName, &questPlayerId)
		switch questPlayerId {
		case 0:
			s.players[0].quests = append(s.players[0].quests, questItemName)
		case 1:
			s.players[1].quests = append(s.players[1].quests, questItemName)
		}
	}
}

func main() {

	for {
		//clean state to begin with
		s := State{}
		s.read()
		log.Println(s.players[0].quests)
		for _, t := range s.getNeighbours(s.grid[3][3]) {
			log.Println(t.direction)
		}
		log.Println(s.grid[3][3].direction)

		//ternary op would be great here, to test only
		if s.turn.turnType == 0 {
			fmt.Println("PUSH 3 RIGHT") // PUSH <id> <direction> | MOVE <direction> | PASS
		} else {
			fmt.Println("MOVE RIGHT")
		}
	}
}
