package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * Help the Christmas elves fetch presents in a magical labyrinth!
 **/

//use matrix
type Grid [7][7]Cell

type Cell struct {
}

type Player struct {
}

type Turn struct {
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	for {
		var turnType int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &turnType)

		for i := 0; i < 7; i++ {
			scanner.Scan()
			inputs := strings.Split(scanner.Text(), " ")
			for j := 0; j < 7; j++ {
				tile, _ := inputs[j]
			}
		}
		for i := 0; i < 2; i++ {
			// numPlayerCards: the total number of quests for a player (hidden and revealed)
			var numPlayerCards, playerX, playerY int
			var playerTile string
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &numPlayerCards, &playerX, &playerY, &playerTile)
		}
		// numItems: the total number of items available on board and on player tiles
		var numItems int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &numItems)

		for i := 0; i < numItems; i++ {
			var itemName string
			var itemX, itemY, itemPlayerId int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &itemName, &itemX, &itemY, &itemPlayerId)
		}
		// numQuests: the total number of revealed quests for both players
		var numQuests int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &numQuests)

		for i := 0; i < numQuests; i++ {
			var questItemName string
			var questPlayerId int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &questItemName, &questPlayerId)
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("PUSH 3 RIGHT") // PUSH <id> <direction> | MOVE <direction> | PASS
	}
}
