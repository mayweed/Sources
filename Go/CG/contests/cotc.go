package main

import (
	"fmt"
	//	"log"
	"math"
)

type actionType string

const (
	move   actionType = "MOVE"
	wait   actionType = "WAIT"
	slower actionType = "SLOWER"
)

type Point struct {
	x, y int
}

func distance2(p1 Point, p2 Point) int {
	x := p2.x - p1.x
	x = x * x
	y := p2.y - p1.y
	y = y * y
	return x + y
}

func distance(p1 Point, p2 Point) float64 {
	return (math.Sqrt(float64(distance2(p1, p2))))
}

type Player struct {
	id    int
	ships []Ship
}
type Ship struct {
	pos         Point
	orientation int
	speed       int
	rum         int
	owner       int
}

type Barrel struct {
	pos       Point
	rumAmount int
}

type State struct {
	entityCount int
	players     [2]Player
	ships       []Ship
	barrels     []Barrel
}

//WIP
type Turn struct{}

func (s *State) readEntities() {
	var entityCount int
	fmt.Scan(&entityCount)

	for i := 0; i < entityCount; i++ {
		var entityId int
		var entityType string
		var x, y, arg1, arg2, arg3, arg4 int
		fmt.Scan(&entityId, &entityType, &x, &y, &arg1, &arg2, &arg3, &arg4)
		switch entityType {
		case "SHIP":
			if arg4 == 1 {
				s.players[1].ships = append(s.players[1].ships, Ship{pos: Point{x, y}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})
				s.ships = append(s.ships, Ship{pos: Point{x, y}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})
			} else if arg4 == 0 {
				s.players[0].ships = append(s.players[0].ships, Ship{pos: Point{x, y}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})
				s.ships = append(s.ships, Ship{pos: Point{x, y}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})

			}
		case "BARREL":
			s.barrels = append(s.barrels, Barrel{pos: Point{x, y}, rumAmount: arg1})
		}

	}

}

//should write test for that!!
func (s *State) getNearestBarrel() Point {
	//width+1 as maxDist
	var maxDist = 24.0
	var pos Point
	for _, barrel := range s.barrels {
		if d := distance(s.players[1].ships[0].pos, barrel.pos); d < maxDist {
			maxDist = d
			pos = barrel.pos
		}
	}
	return pos
}

func main() {
	agent := State{}
	for {
		// myShipCount: the number of remaining ships
		var myShipCount int
		fmt.Scan(&myShipCount)

		agent.readEntities()
		test := agent.getNearestBarrel()
		for i := 0; i < myShipCount; i++ {
			//fmt.Printf("MOVE 11 10\n") // Any valid action, such as "WAIT" or "MOVE x y"
			fmt.Println("MOVE %d %d", test.x, test.y)
		}

	}
}
