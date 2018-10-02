package main

import (
	"fmt"
	"math"
)

const (
	WIDTH            = 16001
	HEIGHT           = 7501
	MAX_MOVE_THRUST  = 150
	MAX_THROW_THRUST = 500
)

type Point struct {
	x, y int
}

func dist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2)) + float64((y1-y2)*(y1-y2)))
}

type Entity struct {
	id    int
	etype string
	vx    int
	vy    int
	s     int
	pos   Point
}

func distEntity(wiz, snaf Entity) float64 {
	return dist(wiz.pos.x, wiz.pos.y, snaf.pos.x, snaf.pos.y)
}

type Wizard struct {
	Entity
	action string
	target Point
	thrust int
}

func (w Wizard) isCarryingSnaffle() bool {
	return w.s == 1
}
func (w *Wizard) move(destination Point) {
	w.action = "MOVE"
	w.target.x = destination.x
	w.target.y = destination.y
	//to begin with
	w.thrust = MAX_MOVE_THRUST
}
func (w *Wizard) throw(destination Point) {
	w.action = "THROW"
	w.target.x = destination.x
	w.target.y = destination.y
	//to begin with
	w.thrust = MAX_THROW_THRUST
}
func (w Wizard) printAction() {
	switch w.action {
	case "MOVE":
		fmt.Println(w.action, w.target.x, w.target.y, w.thrust)
	case "THROW":
		fmt.Println(w.action, w.target.x, w.target.y, w.thrust)
	}
}

type State struct {
	entityCount int
	//me
	myTeamId int
	myGoal   Point
	myWiz    []Wizard
	myScore  int
	myMagic  int
	//opp
	oppGoal  Point
	oppWiz   []Wizard
	oppScore int
	oppMagic int
	//other entities
	snaffles []Entity
	bludgers []Entity
}

func (s *State) setGoals() {
	var myTeamId int
	fmt.Scan(&myTeamId)
	s.myTeamId = myTeamId

	switch myTeamId {
	case 0:
		s.myGoal = Point{0, 3750}
		s.oppGoal = Point{16000, 3750}
	case 1:
		s.myGoal = Point{16000, 3750}
		s.oppGoal = Point{0, 3750}
	}

}
func (s *State) readEntities() {
	var myScore, myMagic int
	fmt.Scan(&myScore, &myMagic)
	s.myScore = myScore
	s.myMagic = myMagic

	var opponentScore, opponentMagic int
	fmt.Scan(&opponentScore, &opponentMagic)
	s.oppScore = opponentScore
	s.oppMagic = opponentMagic

	// entities: number of entities still in game
	var entities int
	fmt.Scan(&entities)
	s.entityCount = entities

	for i := 0; i < entities; i++ {
		var entityId int
		var entityType string
		var x, y, vx, vy, state int
		fmt.Scan(&entityId, &entityType, &x, &y, &vx, &vy, &state)
		if entityType == "WIZARD" {
			s.myWiz = append(s.myWiz, Wizard{Entity: Entity{id: entityId, etype: entityType, vx: vx, vy: vy, s: state, pos: Point{x: x, y: y}}})
		} else if entityType == "OPPONENT_WIZARD" {
			s.oppWiz = append(s.oppWiz, Wizard{Entity: Entity{id: entityId, etype: entityType, vx: vx, vy: vy, s: state, pos: Point{x: x, y: y}}})
		} else if entityType == "SNAFFLE" {
			s.snaffles = append(s.snaffles, Entity{id: entityId, etype: entityType, vx: vx, vy: vy, s: state, pos: Point{x: x, y: y}})
		} else if entityType == "BLUDGER" {
			s.bludgers = append(s.bludgers, Entity{id: entityId, etype: entityType, vx: vx, vy: vy, s: state, pos: Point{x: x, y: y}})
		}
	}

}
func (s *State) clear() {
	s.myWiz = []Wizard{}
	s.oppWiz = []Wizard{}
	s.bludgers = []Entity{}
	s.snaffles = []Entity{}

}
func (s State) pickNearestSnaffle(wiz Wizard) Entity {
	var best = WIDTH + 1.0
	var nearestSnaffle Entity
	for _, snaffle := range s.snaffles {
		distance := distEntity(wiz.Entity, snaffle)
		if distance < best {
			best = distance
			nearestSnaffle = snaffle
		}
	}
	return nearestSnaffle
}
func (s State) think() {
	//pick the nearest, go for it...the dumbest of strat...
	var bestSnaffle Entity
	for _, wiz := range s.myWiz {
		if wiz.isCarryingSnaffle() {
			wiz.throw(s.oppGoal)
		} else {
			bestSnaffle = s.pickNearestSnaffle(wiz)
			wiz.move(bestSnaffle.pos)
		}
		wiz.printAction()
	}
}

func main() {
	s := State{}
	s.setGoals()
	for {
		s.readEntities()
		s.think()
		s.clear()
	}
}
