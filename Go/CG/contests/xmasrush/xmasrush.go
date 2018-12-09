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
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Buffer(make([]byte, 1000000), 1000000)

	for {
		//clean state to begin with
		s := State{}

		var turnType int
		//scanner.Scan()
		//fmt.Sscan(scanner.Text(), &turnType)
		fmt.Scan(&turnType)

		//BOARD: first tile dir then item
		for i := 0; i < 7; i++ {
			//scanner.Scan()
			var row string
			fmt.Scan(&row)
			//inputs := strings.Split(scanner.Text(), " ")
			//inputs := strings.Split(row, " ")
			for j := 0; j < 7; j++ {
				//tile := inputs[j]
				//s.grid[i][j].direction = inputs[j]
				s.grid[i][j].direction = row
				log.Println(s.grid[i][j], i, j)
			}
		}
		// numItems: the total number of items available on board and on player tiles
		var numItems int
		//scanner.Scan()
		//fmt.Sscan(scanner.Text(), &numItems)
		fmt.Scan(&numItems)
		s.numItems = numItems

		for i := 0; i < numItems; i++ {
			var itemName string
			var itemX, itemY, itemPlayerId int
			//scanner.Scan()
			//fmt.Sscan(scanner.Text(), &itemName, &itemX, &itemY, &itemPlayerId)
			fmt.Scan(&itemName, &itemX, &itemY, &itemPlayerId)
			//s.grid[itemX][itemY].itemName = itemName
			//s.grid[itemX][itemY].itemPlayerId = itemPlayerId
		}

		//PLAYER: first pos tile num quest
		for i := 0; i < 2; i++ {
			// numPlayerCards: the total number of quests for a player (hidden and revealed)
			var numPlayerCards, playerX, playerY int
			var playerTile string
			//scanner.Scan()
			//fmt.Sscan(scanner.Text(), &numPlayerCards, &playerX, &playerY, &playerTile)
			fmt.Scan(&numPlayerCards, &playerX, &playerY, &playerTile)
		}
		// numQuests: the total number of revealed quests for both players
		var numQuests int
		//scanner.Scan()
		//fmt.Sscan(scanner.Text(), &numQuests)
		fmt.Scan(&numQuests)

		for i := 0; i < numQuests; i++ {
			var questItemName string
			var questPlayerId int
			//scanner.Scan()
			//fmt.Sscan(scanner.Text(), &questItemName, &questPlayerId)
			fmt.Scan(&questItemName, &questPlayerId)
		}
		fmt.Println("PUSH 3 RIGHT") // PUSH <id> <direction> | MOVE <direction> | PASS
	}
}
