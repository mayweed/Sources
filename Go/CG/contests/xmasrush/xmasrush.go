package main

import (
	"fmt"
	"log"
)

/**
 * Help the Christmas elves fetch presents in a magical labyrinth!
 **/

type Grid [7][7]Tile

type Tile struct {
	direction    string
	itemName     string
	itemPlayerId int
}
type Point struct {
	x, y int
}

type Player struct {
	totalOfquests int
	revealedQuest int
	tile          string
	position      Point
}

type Turn struct {
}

type State struct {
	players  []Player
	grid     Grid
	numItems int
	turn     Turn
}

func main() {

	for {
		//clean state to begin with
		s := State{}

		var turnType int
		fmt.Scan(&turnType)

		//cf GameBoard => sendMapToPlayer()
		for y := 0; y < 7; y++ {
			for x := 0; x < 7; x++ {
				var col string
				fmt.Scan(&col)
				log.Println(col, y, x)
				s.grid[y][x].direction = col
			}
		}
		log.Println(s.grid[5][4])

		for i := 0; i < 2; i++ {
			// numPlayerCards: the total number of quests for a player (hidden and revealed)
			var numPlayerCards, playerX, playerY int
			var playerTile string
			fmt.Scan(&numPlayerCards, &playerX, &playerY, &playerTile)
			log.Println("PlayerTile:", playerTile)
		}

		// numItems: the total number of items available on board and on player tiles
		var numItems int
		fmt.Scan(&numItems)
		s.numItems = numItems

		for i := 0; i < numItems; i++ {
			var itemName string
			var itemX, itemY, itemPlayerId int
			fmt.Scan(&itemName, &itemX, &itemY, &itemPlayerId)
			log.Println("name", itemName, "x", itemX, "y", itemY)
			s.grid[itemY][itemX].itemName = itemName
			s.grid[itemY][itemX].itemPlayerId = itemPlayerId
		}

		// numQuests: the total number of revealed quests for both players
		var numQuests int
		fmt.Scan(&numQuests)

		for i := 0; i < numQuests; i++ {
			var questItemName string
			var questPlayerId int
			fmt.Scan(&questItemName, &questPlayerId)
		}

		//ternary op would be great here, to test only
		if turnType == 0 {
			fmt.Println("PUSH 3 RIGHT") // PUSH <id> <direction> | MOVE <direction> | PASS
		} else {
			fmt.Println("MOVE RIGHT")
		}
	}
}
