package main

import (
	"fmt"
	"log"
	"math"
)

const (
	MAP_WIDTH  = 23
	MAP_HEIGHT = 21
)

type actionType int

const (
	MOVE   actionType = 0
	SLOWER actionType = 1
	WAIT   actionType = 2
	FIRE   actionType = 3
	MINE   actionType = 4
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
func (p Point) equalsTo(p2 Point) bool {
	return p.x == p2.x && p.y == p2.y
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

type Entity struct {
	id         int
	entityType string
	pos        Point
}

func (e *Entity) updateEntity() {
}
func (e Entity) getPosition() Point {
	return e.pos
}
func (e Entity) distanceTo(e2 Entity) float64 {
	return e.pos.distanceTo(e2.getPosition())
}

type Ship struct {
	Entity
	orientation        int
	speed              int
	rum                int
	owner              int
	hasFiredCannonBall bool

	//Action
	actionType int
	target     Entity

	//sim? And what if I go 2 instead of 1 etc...
	newOrientation int
}

func (s Ship) stern() Point {
	return s.pos.neighbour((s.orientation + 3) % 6)
}
func (s Ship) bow() Point {
	return s.pos.neighbour(s.orientation)
}
func (s Ship) newStern() Point {
	return s.pos.neighbour((s.newOrientation + 3) % 6)
}
func (s Ship) newBow() Point {
	return s.pos.neighbour(s.newOrientation)
}
func (s Ship) printAction() {
	switch s.actionType {
	case 0:
		fmt.Println("MOVE", s.target.pos.x, s.target.pos.y)
	case 1:
		fmt.Println("SLOWER")
	case 2:
		fmt.Println("WAIT")
	case 3:
		fmt.Println("FIRE", s.target.pos.x, s.target.pos.y)
	case 4:
		fmt.Println("MINE")
	}
}

type Barrel struct {
	Entity
	rumAmount int
}
type Mine struct {
	Entity
}
type cannonBall struct {
	Entity
	fromShip          int
	turnsBeforeImpact int
}

type Player struct {
	id        int
	shipCount int
	ships     []Ship
}
type State struct {
	entityCount int
	players     [2]Player
	ships       []Ship
	barrels     []Barrel
	mines       []Mine
	cannonBalls []cannonBall
}

//WIP
type Turn struct {
	actionType int
	move       string
}

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
				s.players[1].ships = append(s.players[1].ships, Ship{Entity: Entity{entityId, entityType, Point{x, y}}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})
				s.ships = append(s.ships, Ship{Entity: Entity{entityId, entityType, Point{x, y}}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})
			} else if arg4 == 0 {
				s.players[0].ships = append(s.players[0].ships, Ship{Entity: Entity{entityId, entityType, Point{x, y}}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})
				s.ships = append(s.ships, Ship{Entity: Entity{entityId, entityType, Point{x, y}}, orientation: arg1, speed: arg2, rum: arg3, owner: arg4})

			}
		case "BARREL":
			s.barrels = append(s.barrels, Barrel{Entity: Entity{entityId, entityType, Point{x, y}}, rumAmount: arg1})
		case "MINE":
			s.mines = append(s.mines, Mine{Entity: Entity{entityId, entityType, Point{x, y}}})
		case "CANNONBALL":
			s.cannonBalls = append(s.cannonBalls, cannonBall{Entity: Entity{entityId, entityType, Point{x, y}}, fromShip: arg1, turnsBeforeImpact: arg2})
		}

	}

}

//should be in think?
//should pass the ship as arg!!
func (s *State) getNearestTarget() Entity {
	//width+1 as maxDist
	var maxDist = 24.0
	var shipPos = s.players[1].ships[0].bow()
	var target Entity

	//if nearestBarrel > nearestMine, fire mine instead?
	//if enemy ship is in range or i collide with must fire!!
	for _, barrel := range s.barrels {
		if d := shipPos.distanceTo(barrel.pos); d < maxDist {
			maxDist = d
			//ugly really and i lost info!!
			target = barrel.Entity
			s.players[1].ships[0].actionType = 0
		}
	}
	for _, mine := range s.mines {
		if d := shipPos.distanceTo(mine.pos); d < maxDist {
			maxDist = d
			target = mine.Entity
			s.players[1].ships[0].actionType = 3
		}
	}
	//it stuck my ship!!
	//if, really, we are closer to enemy ship just fire at it?
	if s.players[1].ships[0].pos.distanceTo(s.players[0].ships[0].pos) < maxDist {
		target = s.players[0].ships[0].Entity
		s.players[1].ships[0].actionType = 3
	}
	return target
}

//shouldnt that yield a turn? Then parse and display??
func (s *State) think() {
	//will become a queue string when multiple ships
	//var action string
	for i := 0; i < s.players[1].shipCount; i++ {
		//if my rum total < rum adv go to target, else wait?
		if s.players[1].ships[i].rum < s.players[0].ships[i].rum {
			s.players[1].ships[i].target = s.getNearestTarget()
			log.Println(s.players[1].ships[i].target)
			s.players[1].ships[i].printAction()
		} else {
			fmt.Println("WAIT")
		}
	}
	//clear state!!
	s.barrels = []Barrel{}
	s.mines = []Mine{}
	s.cannonBalls = []cannonBall{}
	s.players[0].ships = []Ship{}
	s.players[1].ships = []Ship{}
}
func main() {
	agent := State{}
	for {
		agent.readEntities()
		agent.think()
	}
}
