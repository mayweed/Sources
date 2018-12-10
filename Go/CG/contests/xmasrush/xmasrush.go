package main

import (
	"fmt"
	"log"
)

type Point struct {
	x, y int
}

type Grid [7][7]Tile

type Tile struct {
	direction    string
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

		//ternary op would be great here, to test only
		if s.turn.turnType == 0 {
			fmt.Println("PUSH 3 RIGHT") // PUSH <id> <direction> | MOVE <direction> | PASS
		} else {
			fmt.Println("MOVE RIGHT")
		}
	}
}
