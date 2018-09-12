package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type Point struct {
	x, y int
}

func setPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}
func (p Point) Move(x, y int) Point {
	x = p.x + x
	y = p.y + y
	return Point{x, y}
}

type State struct {
	busters []Entity
	ghosts  []Entity
	//grab them by team id?
	player []Player
}
type Player struct {
	teamId         int
	nbBusters      int
	homeBase       Point
	capturedGhosts []Entity
}
type Entity struct {
	entityId   int
	entityType int
	state      int
	value      int
	pos        Point
}

func readEntity() {
	// entities: the number of busters and ghosts visible to you
	var entities int
	fmt.Scan(&entities)
	var entityId, x, y, entityType, state, value int
	//must be in a state
	var ghosts_list []Entity
	var my_busters_list []Entity
	var foe_busters_list []Entity

	for i := 0; i < entities; i++ {
		// entityId: buster id or ghost id
		// y: position of this buster / ghost
		// entityType: the team id if it is a buster, -1 if it is a ghost.
		// state: For busters: 0=idle, 1=carrying a ghost.
		// value: For busters: Ghost id being carried. For ghosts: number of busters attempting to trap this ghost.

		fmt.Scan(&entityId, &x, &y, &entityType, &state, &value)
		if entityType == myTeamId {
			b := Entity{entityId, entityType, state, value, Point{x, y}}
			my_busters_list = append(my_busters_list, b)
		}
		if entityType == -1 {
			ghosts_list = append(ghosts_list, Entity{entityId, entityType, state, value, Point{x, y}})
		}
		if entityType != myTeamId && entityType != -1 {
			b := Entity{entityId, entityType, state, value, Point{x, y}}
			foe_busters_list = append(foe_busters_list, b)
		}

	}
}

func main() {
	//state.readParameters()
	// bustersPerPlayer: the amount of busters you control
	var bustersPerPlayer int
	fmt.Scan(&bustersPerPlayer)

	// ghostCount: the amount of ghosts on the map
	var ghostCount int
	fmt.Scan(&ghostCount)

	// myTeamId: if this is 0, your base is on the top left of the map, if it is one, on the bottom right
	var myTeamId int
	fmt.Scan(&myTeamId)

	//that switch should be in player init?
	var home Point
	switch myTeamId {
	case 0:
		home = Point{0, 0}
	case 1:
		home = Point{16000, 9000}
	}

	for {
		//state.readEntity()

		//state.think()
		//make them move to see the game
		for i := 0; i < len(my_busters_list); i++ {
			if u := len(ghosts_list); u != 0 {
				//should iterate through ghosts_lists(pop?)
				//then the buster should be blocked in "transport_state" bool?
				fmt.Printf("BUST %d\n", ghosts_list[0].entityId)
				u -= 1
				continue
			} else {
				fmt.Printf("MOVE %d %d\n", rand.Intn(16000), rand.Intn(9000))
			}
		}
		//state.act()
		//fmt.Printf("MOVE 8000 4500\n") // MOVE x y | BUST id | RELEASE
	}
}
