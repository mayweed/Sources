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

type Point struct {
	x, y int
}

func (p Point) angle(targetPosition Point) float64 {
	var dy = float64(targetPosition.y-p.y) * math.Sqrt(3) / 2
	var dx = float64(targetPosition.x-p.x) + float64(float64((p.y-targetPosition.y)&1)*0.5)
	var angle = -math.Atan2(dy, dx) * 3 / math.Pi
	if angle < 0 {
		angle += 6
	} else if angle >= 6 {
		angle -= 6
	}
	return angle
}

func (p Point) neighbour(orientation int) Point {
	var DIRECTIONS_EVEN = [6][2]int{{1, 0}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}}
	var DIRECTIONS_ODD = [6][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, 0}, {0, 1}, {1, 1}}
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
func (p Point) toCubeCoordinate() cubeCoord {
	var c cubeCoord
	c.x = p.x - (p.y-(p.y&1))>>1
	c.z = p.y
	c.y = -(c.x + c.z)
	return c
}
func (p Point) isInsideMap() bool {
	return p.x >= 0 && p.x < MAP_WIDTH && p.y >= 0 && p.y < MAP_HEIGHT
}
func (p Point) distanceTo(dst Point) float64 {
	return p.toCubeCoordinate().distanceTo(dst.toCubeCoordinate())
}

type cubeCoord struct {
	x, y, z int
}

func (c cubeCoord) toOffsetCoord() Point {
	x := c.x + (c.z-(c.z&1))>>1
	y := c.z
	return Point{x, y}
}
func (c cubeCoord) neighbour(orientation int) cubeCoord {
	var directions = [6][3]int{{1, -1, 0}, {+1, 0, -1}, {0, +1, -1}, {-1, +1, 0}, {-1, 0, +1}, {0, -1, +1}}
	nx := c.x + directions[orientation][0]
	ny := c.y + directions[orientation][1]
	nz := c.z + directions[orientation][2]
	return cubeCoord{nx, ny, nz}
}
func (c cubeCoord) distanceTo(dst cubeCoord) float64 {
	return (math.Abs(float64(c.x-dst.x)) + math.Abs(float64(c.y-dst.y)) + math.Abs(float64(c.z-dst.z))) / 2.0
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
		if d := s.players[1].ships[0].pos.distanceTo(barrel.pos); d < maxDist {
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
