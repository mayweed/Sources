package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HEIGHT = 15
	WIDTH  = 30
)

type Cell struct {
	x, y    int
	ore     int
	hasHole bool //1 has, 0 hasnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// height: size of the map
	var width, height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)

	var b [HEIGHT][WIDTH]Cell

	for {
		// myScore: Amount of ore delivered
		var myScore, opponentScore int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &myScore, &opponentScore)

		for i := 0; i < height; i++ {
			scanner.Scan()
			inputs := strings.Split(scanner.Text(), " ")
			for j := 0; j < width; j++ {
				// ore: amount of ore or "?" if unknown
				// hole: 1 if cell has a hole
				ore := inputs[2*j]
				if ore == "?" {
					b[i][j].ore = -1
				} else {
					b[i][j].ore, _ = strconv.Atoi(ore)
				}
				hole, _ := strconv.ParseInt(inputs[2*j+1], 10, 32)
				if hole == 1 {
					b[i][j].hasHole = true
				} else {
					b[i][j].hasHole = false
				}
			}

		}
		// entityCount: number of entities visible to you
		// radarCooldown: turns left until a new radar can be requested
		// trapCooldown: turns left until a new trap can be requested
		var entityCount, radarCooldown, trapCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &entityCount, &radarCooldown, &trapCooldown)
		for i := 0; i < entityCount; i++ {
			// id: unique id of the entity
			// entityType: 0 for your robot, 1 for other robot, 2 for radar, 3 for trap
			// y: position of the entity
			// item: if this entity is a robot, the item it is carrying (-1 for NONE, 2 for RADAR, 3 for TRAP, 4 for ORE)
			var id, entityType, x, y, item int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &id, &entityType, &x, &y, &item)
		}
		for i := 0; i < 5; i++ {
			fmt.Println("WAIT") // WAIT|MOVE x y|DIG x y|REQUEST item
		}
	}
}
