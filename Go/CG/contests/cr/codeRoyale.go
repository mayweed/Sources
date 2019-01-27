package main

import (
	"fmt"
	"log"
)

//constants
const (
	MAP_HEIGHT   = 1000
	MAP_WIDTH    = 1920
	QUEEN_RADIUS = 30
	QUEEN_SPEED  = 60
)

//MAP-Grid
type Point struct {
	x, y int
}

//SITES
type Site interface {
	Id() int
	Pos() Point
	Type() string
	Radius() int
}
type (
	Cell struct {
		id     int
		pos    Point
		what   string
		radius int
	}

	Goldmine struct {
		siteId        int
		pos           Point
		radius        int
		structureType int
		owner         int
		incomeRate    int
		param2        int
	}
	Tower struct {
		siteId        int
		pos           Point
		radius        int
		structureType int
		owner         int
		healthPoint   int
		attackRadius  int
	}
	Barrack struct {
		siteId        int
		pos           Point
		radius        int
		structureType int
		owner         int
		turnsInactive int
		creepType     int
	}
)

//UNITS
type Unit struct {
	pos      Point
	owner    int
	unitType int
	health   int
}

//PLAYER
type Player struct {
	//a player sites
	goldmine []Goldmine
	towers   []Tower
	barracks []Barrack
	//units by type
	queen   Unit
	knights []Unit
	archers []Unit
	giants  []Unit
}

//GAME STATE
type State struct {
	players [2]Player
	grid    [MAP_HEIGHT][MAP_WIDTH]Cell
}

func (s *State) initBoard() {
	for y := 0; y < MAP_HEIGHT; y++ {
		for x := 0; x < MAP_WIDTH; x++ {
			s.grid[y][x] = Cell{0, Point{x, y}, ".", 0}
		}
	}
}

func (s *State) initSite(numSites int) {
	for i := 0; i < numSites; i++ {
		var x, y int
		var siteId, radius int
		fmt.Scan(&siteId, &x, &y, &radius)
		s.grid[y][x].id = siteId
		s.grid[y][x].pos = Point{x, y}
		s.grid[y][x].what = "B"
		s.grid[y][x].radius = radius
	}
}

//quick and dirty to check
func (s State) printBoard() {
	var row string
	for y, _ := range s.grid {
		for _, it := range s.grid[y] {
			//row += fmt.Sprintf(row, it.what)
			row += it.what
		}
		row += "\n"
	}
	log.Println(row)
}

//MAIN
func main() {

	var s = State{}
	//s.initBoard()
	var numSites int
	fmt.Scan(&numSites)

	s.initSite(numSites)
	//me := s.players[0]
	//opp := s.players[1]

	s.printBoard()

	for {
		// touchedSite: -1 if none
		var gold, touchedSite int
		fmt.Scan(&gold, &touchedSite)

		for i := 0; i < numSites; i++ {
			// ignore1: used in future leagues
			// ignore2: used in future leagues
			// structureType: -1 = No structure, 0=Goldmine, 1= Tower, 2 = Barracks
			// owner: -1 = No structure, 0 = Friendly, 1 = Enemy
			var siteId, ignore1, ignore2, structureType, owner, param1, param2 int
			fmt.Scan(&siteId, &ignore1, &ignore2, &structureType, &owner, &param1, &param2)
		}
		var numUnits int
		fmt.Scan(&numUnits)

		for i := 0; i < numUnits; i++ {
			// unitType: -1 = QUEEN, 0 = KNIGHT, 1 = ARCHER, 2=GIANT
			var x, y int
			var owner, unitType, health int
			fmt.Scan(&x, &y, &owner, &unitType, &health)
		}
		fmt.Println("WAIT")
		fmt.Println("TRAIN")
	}
}
