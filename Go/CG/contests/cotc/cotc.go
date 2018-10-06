package main

import (
	"container/list"
	"fmt"
	"log"
	"math"
)

const (
	MAP_WIDTH             = 23
	MAP_HEIGHT            = 21
	INITIAL_SHIP_HEALTH   = 100
	MAX_SHIP_HEALTH       = 100
	MAX_SHIP_SPEED        = 2
	MINE_VISIBILITY_RANGE = 5
	MINE_DAMAGE           = 25
	NEAR_MINE_DAMAGE      = 10
	FIRE_DISTANCE_MAX     = 10
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
	actionType string
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

func (s *Ship) fire(pos Point) {
	s.actionType = "FIRE"
	s.target.pos.x = pos.x
	s.target.pos.y = pos.y
}
func (s *Ship) mine() {
	s.actionType = "MINE"
}
func (s *Ship) wait() {
	s.actionType = "WAIT"
}
func (s *Ship) slower() {
	s.actionType = "SLOWER"
}
func (s *Ship) move(pos Point) {
	s.actionType = "MOVE"
	s.target.pos.x = pos.x
	s.target.pos.y = pos.y
}
func (s *Ship) faster() {
	s.actionType = "FASTER"
}
func (s *Ship) port() {
	s.actionType = "PORT"
}
func (s *Ship) starboard() {
	s.actionType = "STARBOARD"
}
func (s Ship) printAction() {
	switch s.actionType {
	case "MOVE":
		fmt.Println("MOVE", s.target.pos.x, s.target.pos.y)
	case "SLOWER":
		fmt.Println("SLOWER")
	case "WAIT":
		fmt.Println("WAIT")
	case "FIRE":
		fmt.Println("FIRE", s.target.pos.x, s.target.pos.y)
	case "MINE":
		fmt.Println("MINE")
	}
}

type Barrel struct {
	Entity
	health     int
	isTargeted bool
}
type Mine struct {
	Entity
	isTargeted bool
}
type cannonball struct {
	Entity
	fromShip       int
	remainingTurns int
	explosions     Point
}

type State struct {
	entityCount int
	myShipCount int
	enemyShips  []Ship
	allyShips   []Ship
	ships       []Ship
	barrels     []Barrel
	mines       []Mine
	cannonballs *list.List
}

func (s State) isaMine(dest Point) bool {
	for _, mine := range s.mines {
		if mine.Entity.pos == dest {
			return true
		}
	}
	return false
}
func (s *State) moveShips() {
	for i := 1; i <= MAX_SHIP_SPEED; i++ {
		for _, ship := range s.ships {
			if ship.isDead || i > ship.speed {
				continue
			}
			ship.newPosition = ship.Entity.pos
			ship.newBowCoordinate = ship.bow()
			ship.newSternCoordinate = ship.stern()
			var newCoordinate = ship.Entity.pos.neighbour(ship.orientation)
			if newCoordinate.isInsideMap() {
				ship.newPosition = newCoordinate
				ship.newBowCoordinate = newCoordinate.neighbour(ship.orientation)
				ship.newSternCoordinate = newCoordinate.neighbour((ship.orientation + 3) % 6)
			} else {
				//"If a ship attempts to leave the map, it is stopped and its speed is set to 0."
				ship.speed = 0
			}
		}
		//check for collisions
		var collisions []Ship
		var collisionDetected = true
		for collisionDetected {
			collisionDetected = false
			for _, ship := range s.ships {
				if ship.isDead {
					continue
				}
				if ship.newBowsIntersect(s.ships) {
					collisions = append(collisions, ship)
				}
			}
			for _, ship := range collisions {
				//Revert last move coz you collide cant move fw
				ship.newPosition = ship.Entity.pos
				ship.newBowCoordinate = ship.bow()
				ship.newSternCoordinate = ship.stern()
				ship.speed = 0
				collisionDetected = true
			}
			//just clear it before exit from loop
			collisions = []Ship{}
		}
		for _, ship := range s.ships {
			if ship.isDead {
				continue
			}
			ship.Entity.pos = ship.newPosition
		}
		//checkCollisions()
	}
}

func (s *State) checkCollisions() {
	var funcRumBarrels []Barrel
	//var funcMines []Mine
	for _, ship := range s.ships {
		if ship.isDead {
			continue
		}
		var bow = ship.bow()
		var stern = ship.stern()
		var center = ship.Entity.pos
		for _, barrel := range s.barrels {
			if barrel.Entity.pos == bow || barrel.Entity.pos == stern || barrel.Entity.pos == center {
				ship.heal(barrel.health)
				//Any other way to remove that which will be less costly and easy to work with that container/list?
				for _, b := range s.barrels {
					if b != barrel {
						funcRumBarrels = append(funcRumBarrels, b)
					}
				}
				s.barrels = funcRumBarrels
			}
		}
		for _, mine := range s.mines {
			if mine.Entity.pos == bow || mine.Entity.pos == stern || mine.Entity.pos == center {
				ship.damage(MINE_DAMAGE)
				for _, otherShip := range s.ships {
					if otherShip.isDead || otherShip == ship {
						continue
					}
					if mine.Entity.pos.distanceTo(otherShip.Entity.pos) == 1 ||
						mine.Entity.pos.distanceTo(otherShip.stern()) == 1 ||
						mine.Entity.pos.distanceTo(otherShip.bow()) == 1 {
						otherShip.damage(NEAR_MINE_DAMAGE)
					}
				}
				//here stuff to remove mine from s.mines...
			}

		}

	}
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
PARTIAL        this->moveCannonballs();
DONE!    this->decrementRum();
        this->applyActions();
PARTIAL        this->moveShips();
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

//test container/list!!
func (s *State) movecannonballs() {
	for ball := s.cannonballs.Front(); ball != nil; ball = ball.Next() {
		if ball.Value.(cannonball).remainingTurns == 0 {
			s.cannonballs.Remove(ball)
			continue
		} else if ball.Value.(cannonball).remainingTurns > 0 {
			//cannot assign??
			//ball.Value.(cannonball).remainingTurns--
			log.Println(ball.Value)
		}
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
			s.barrels = append(s.barrels, Barrel{Entity: Entity{entityId, entityType, Point{x, y}}, health: arg1})
		case "MINE":
			s.mines = append(s.mines, Mine{Entity: Entity{entityId, entityType, Point{x, y}}})
		case "CANNONBALL":
			//s.cannonballs = append(s.cannonballs, cannonBall{Entity: Entity{entityId, entityType, Point{x, y}}, fromShip: arg1, remainingTurns: arg2})
			s.cannonballs = list.New()
			//pass pointer to modify struct
			s.cannonballs.PushBack(&cannonball{Entity: Entity{entityId, entityType, Point{x, y}}, fromShip: arg1, remainingTurns: arg2})
		}
	}
}

func (s *State) clear() {
	s.enemyShips = []Ship{}
	s.allyShips = []Ship{}
	s.ships = []Ship{}
	s.barrels = []Barrel{}
	s.mines = []Mine{}
	//list.New() yields a pointer on list?
	//s.cannonballs = []cannonball{}

}

func computeScore(ships []Ship) int {
	var score int
	for _, ship := range ships {
		score += ship.health
	}
	return score
}

//to fire where the ship will be
func (s Ship) nextPosShip(inTurns int) Point {
	var nextPos = s.pos
	for t := 0; t < inTurns; t++ {
		nextPos = nextPos.neighbour(s.orientation)
	}
	return nextPos
}

//read https://github.com/gyuho/learn/tree/master/doc/go_function_method_pointer_nil_map_slice
func (s State) isCannonballComing(sp Ship) bool {
	for _, enemyShip := range s.enemyShips {
		//for _, cannonball := range s.cannonballs {
		//"panic: runtime error: invalid memory address or nil pointer dereference" :''(
		//this is a pointer on a list of pointers...
		for ball := *s.cannonballs.Front(); &ball != nil; ball.Next() {
			if ball.Value.(cannonball).fromShip == enemyShip.id && ball.Value.(cannonball).Entity.pos == sp.pos {
				return true
			}
		}
	}
	return false
}

func (s *State) think() {
	var maxDist = MAP_WIDTH + 1.0 //24.0
	var target Entity

	for _, myShip := range s.allyShips {
		var shipPos = myShip.bow()
		//iif x y de cannon enn == my pos move your ass

		if s.isCannonballComing(myShip) {
			//fuckin move!!
			myShip.move(Point{myShip.pos.x + 3, myShip.pos.y + 3})
		} else if computeScore(s.allyShips) < computeScore(s.enemyShips) {
			for _, barrel := range s.barrels {
				//check if barrel is not already targeted (by another boat later on)
				if d := shipPos.distanceTo(barrel.pos); d < maxDist {
					maxDist = d
					target = barrel.Entity
					barrel.isTargeted = true
					myShip.move(target.pos)
				}
			}
		} else {
			//if, really, we are closer to enemy ship just fire at it?
			for _, enemyShip := range s.enemyShips {
				//idem : si dans 5 cases (tours?) c'est une mine soit je dodge
				//soit je fire!!
				var targetShip = enemyShip //.Entity
				var numTurns = int(1 + targetShip.Entity.pos.distanceTo(myShip.pos)/3)
				if myShip.pos.distanceTo(enemyShip.pos) < FIRE_DISTANCE_MAX {
					myShip.fire(targetShip.nextPosShip(numTurns))
				} else if s.isaMine(myShip.nextPosShip(2)) {
					myShip.fire(myShip.nextPosShip(numTurns))
				} else {
					//go in range!!
					myShip.move(enemyShip.pos)
				}
			}
		}
		//TEST really not conclusive...
		//if s.cannonballs.Len() > 0 {
		//s.movecannonballs()
		//}

		//in the end if ship.Action is empty, just wait?
		//gosh ugly!!
		if myShip.actionType == "" {
			fmt.Println("WAIT")
		} else {
			myShip.printAction()
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
