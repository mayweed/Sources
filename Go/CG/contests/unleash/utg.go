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
	RADAR  = 2
	ORE    = 4
	TRAP   = 3
)

type Point struct {
	x, y int
}

type Cell struct {
	pos     Point
	ore     int
	hasHole bool //1 has, 0 hasnt
}

type Grid [HEIGHT][WIDTH]Cell

type Entity struct {
	id    int
	etype int
	pos   Point
	item  int
}

func (e Entity) hasRadar() bool {
	if e.item == RADAR {
		return true
	}
	return false
}

func (e Entity) hasOre() bool {
	if e.item == ORE {
		return true
	}
	return false
}

type Team struct {
	bots  []Entity
	score int
}
type GameState struct {
	board       Grid
	numEntities int
	myTeam      Team
	oppTeam     Team
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	s := GameState{}

	// height: size of the map
	var width, height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)

	for {
		// myScore: Amount of ore delivered
		var myScore, opponentScore int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &myScore, &opponentScore)
		s.myTeam.score = myScore
		s.oppTeam.score = opponentScore

		for i := 0; i < HEIGHT; i++ {
			scanner.Scan()
			inputs := strings.Split(scanner.Text(), " ")
			for j := 0; j < WIDTH; j++ {
				// ore: amount of ore or "?" if unknown
				// hole: 1 if cell has a hole
				s.board[i][j].pos = Point{i, j}
				ore := inputs[2*j]
				if ore == "?" {
					s.board[i][j].ore = -1
				} else {
					s.board[i][j].ore, _ = strconv.Atoi(ore)
				}
				hole, _ := strconv.ParseInt(inputs[2*j+1], 10, 32)
				if hole == 1 {
					s.board[i][j].hasHole = true
				} else {
					s.board[i][j].hasHole = false
				}
			}
		}

		// entityCount: number of entities visible to you
		// radarCooldown: turns left until a new radar can be requested
		// trapCooldown: turns left until a new trap can be requested
		var entityCount, radarCooldown, trapCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &entityCount, &radarCooldown, &trapCooldown)

		s.numEntities = entityCount
		for i := 0; i < entityCount; i++ {
			// id: unique id of the entity
			// entityType: 0 for your robot, 1 for other robot, 2 for radar, 3 for trap
			// y: position of the entity
			// item: if this entity is a robot, the item it is carrying (-1 for NONE, 2 for RADAR, 3 for TRAP, 4 for ORE)
			var id, entityType, x, y, item int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &id, &entityType, &x, &y, &item)
			if entityType == 0 {
				s.myTeam.bots = append(s.myTeam.bots, Entity{id: id, etype: entityType, pos: Point{x, y}, item: item})
			} else {
				s.oppTeam.bots = append(s.oppTeam.bots, Entity{id: id, etype: entityType, pos: Point{x, y}, item: item})
			}
		}

		for i := 0; i < 5; i++ {
			fmt.Println("WAIT") // WAIT|MOVE x y|DIG x y|REQUEST item
		}
		//clear state
		s.myTeam.bots = []Entity{}
	}
}
