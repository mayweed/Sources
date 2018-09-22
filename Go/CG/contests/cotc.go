package main

import (
	"fmt"
	//	"log"
	"math"
)

type actionType string

const (
	MAP_WIDTH  = 23
	MAP_HEIGHT = 21
	//shouldn't be there action..
	move   actionType = "MOVE"
	wait   actionType = "WAIT"
	slower actionType = "SLOWER"
)

var DIRECTIONS_EVEN = [6][2]int{{1, 0}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}}
var DIRECTIONS_ODD = [6][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, 0}, {0, 1}, {1, 1}}

type Point struct {
	x, y int
}

func (p Point) neighbour(orientation int) Point {
	var newY, newX int
	var neighbour Point
	if p.y%2 == 1 {
		newY = p.y + DIRECTIONS_ODD[orientation][1]
		newX = p.x + DIRECTIONS_ODD[orientation][0]
	} else {
		newY = p.y + DIRECTIONS_EVEN[orientation][1]
		newX = p.x + DIRECTIONS_EVEN[orientation][0]
	}

	neighbour.x = newX
	neighbour.y = newY

	return neighbour
}

func (p Point) isInsideMap() bool {
	return p.x >= 0 && p.x < MAP_WIDTH && p.y >= 0 && p.y < MAP_HEIGHT
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
	id        int
	shipCount int
	ships     []Ship
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
	// myShipCount: the number of remaining ships
	var myShipCount int
	fmt.Scan(&myShipCount)
	s.players[1].shipCount = myShipCount

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

func (s *State) think() {
	test := s.getNearestBarrel()
	for i := 0; i < s.players[1].shipCount; i++ {
		fmt.Println("MOVE", test.x, test.y)
	}
	//clear state!!
	s.barrels = []Barrel{}
}
func main() {
	agent := State{}
	for {
		agent.readEntities()
		agent.think()
	}
}
