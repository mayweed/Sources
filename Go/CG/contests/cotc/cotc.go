package main

import (
	"fmt"
	"log"
	"math"
)

const (
	MAP_WIDTH           = 23
	MAP_HEIGHT          = 21
	INITIAL_SHIP_HEALTH = 100
	MAX_SHIP_HEALTH     = 100
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
	health             int
	owner              int
	hasFiredCannonBall bool

	//Action
	actionType int
	target     Entity

	//sim? And what if I go 2 instead of 1 etc...
	newOrientation     int
	newPosition        Point
	newBowCoordinate   Point
	newSternCoordinate Point
	//why? should be able to revert back or what?
	initialHealth int

	isDead bool
}

func (s Ship) isAlly() bool {
	return s.owner == 1
}
func (s Ship) heal(amount int) {
	s.health += amount
	if s.health > MAX_SHIP_HEALTH {
		s.health = MAX_SHIP_HEALTH
	}
}
func (s Ship) damage(amount int) {
	s.health -= amount
	if s.health <= 0 {
		s.health = 0
		s.isDead = true
	}
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

// I MUST TEST THOSE!! new* values are supposed correctly init...
//first check is problematic: will be initialised to null?? Should init in readEntity
//to another value? We suppose newBowCoord is not nil!!
func (s Ship) newBowIntersect(other Ship) bool {
	return s.newBowCoordinate == other.newBowCoordinate ||
		s.newBowCoordinate == other.newPosition ||
		s.newBowCoordinate == other.newSternCoordinate
}
func (s Ship) newBowsIntersect(ships []Ship) bool {
	for _, otherShip := range ships {
		if otherShip.isDead {
			continue
		}
		if s != otherShip && s.newBowIntersect(otherShip) {
			return true
		}
	}
	return false
}

func (s Ship) newPositionIntersect(other Ship) bool {
	var sternCollision = s.newSternCoordinate == other.newBowCoordinate ||
		s.newSternCoordinate == other.newPosition ||
		s.newSternCoordinate == other.newSternCoordinate
	var centerCollision = s.newPosition == other.newBowCoordinate ||
		s.newPosition == other.newPosition ||
		s.newPosition == other.newSternCoordinate
	return s.newBowIntersect(other) || sternCollision || centerCollision
}

func (s Ship) newPositionsIntersect(ships []Ship) bool {
	for _, otherShip := range ships {
		if s != otherShip && s.newPositionIntersect(otherShip) {
			return true
		}
	}
	return false
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

type State struct {
	entityCount int
	myShipCount int
	enemyShips  []Ship
	allyShips   []Ship
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

/*
It will apply on a newState==current state
   void simulateTurn() {
DONE!        this->updateInitialRum();
        this->moveCannonballs();
DONE!    this->decrementRum();
        this->applyActions();
        this->moveShips();
        this->rotateShips();
        this->explodeShips();
        this->explodeMines();
        this->explodeBarrels();
        this->createDroppedRum();
        ++turn;
    }
*/
func (s *State) decrementRum() {
	for _, ship := range s.ships {
		ship.damage(1)
	}
}
func (s *State) updateInitialRum() {
	for _, ship := range s.ships {
		ship.initialHealth = ship.health
	}
}

func (s *State) readEntities() {
	// myShipCount: the number of remaining ships
	var myShipCount int
	fmt.Scan(&myShipCount)
	s.myShipCount = myShipCount

	var entityCount int
	fmt.Scan(&entityCount)
	s.entityCount = entityCount

	for i := 0; i < entityCount; i++ {
		var entityId int
		var entityType string
		var x, y, arg1, arg2, arg3, arg4 int
		fmt.Scan(&entityId, &entityType, &x, &y, &arg1, &arg2, &arg3, &arg4)
		switch entityType {
		case "SHIP":
			if arg4 == 1 {
				s.allyShips = append(s.allyShips, Ship{Entity: Entity{entityId, entityType, Point{x, y}}, orientation: arg1, speed: arg2, health: arg3, owner: arg4})
				s.ships = append(s.ships, s.allyShips...)
			} else if arg4 == 0 {
				s.enemyShips = append(s.enemyShips, Ship{Entity: Entity{entityId, entityType, Point{x, y}}, orientation: arg1, speed: arg2, health: arg3, owner: arg4})
				s.ships = append(s.ships, s.enemyShips...)
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

func (s *State) clear() {
	s.barrels = []Barrel{}
	s.mines = []Mine{}
	s.cannonBalls = []cannonBall{}
	s.enemyShips = []Ship{}
	s.allyShips = []Ship{}
}

//should be in think?
//should pass the ship as arg!!
func (s *State) getNearestTarget() Entity {
	//width+1 as maxDist
	var maxDist = 24.0
	var shipPos = s.allyShips[0].bow()
	//log.Println(s.allyShips[0].bow())
	var target Entity

	//if nearestBarrel > nearestMine, fire mine instead?
	//if enemy ship is in range or i collide with must fire!!
	for _, barrel := range s.barrels {
		if d := shipPos.distanceTo(barrel.pos); d < maxDist {
			maxDist = d
			//ugly really and i lost info!!
			target = barrel.Entity
			s.allyShips[0].actionType = 0
		}
	}
	for _, mine := range s.mines {
		if d := shipPos.distanceTo(mine.pos); d < maxDist {
			maxDist = d
			target = mine.Entity
			s.allyShips[0].actionType = 3
		}
	}
	//it stuck my ship!!
	//if, really, we are closer to enemy ship just fire at it?
	if s.allyShips[0].pos.distanceTo(s.enemyShips[0].pos) < maxDist {
		target = s.enemyShips[0].Entity
		log.Println(target)
		s.allyShips[0].actionType = 3
	}
	return target
}

//shouldnt that yield a turn? Then parse and display??
func (s *State) think() {
	//will become a queue string when multiple ships
	//var action string
	for i := 0; i < s.myShipCount; i++ {
		//if my rum total < rum adv go to target, else wait?
		if s.allyShips[i].health < s.enemyShips[i].health {
			s.allyShips[i].target = s.getNearestTarget()
			s.allyShips[i].printAction()
		} else {
			fmt.Println("WAIT")
		}
	}
	//clear state!!
	s.clear()
}
func main() {
	agent := State{}
	for {
		agent.readEntities()
		agent.think()
	}
}
