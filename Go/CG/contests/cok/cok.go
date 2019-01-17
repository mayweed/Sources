package main

import (
	"fmt"
	"log"
	"strings"
)

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	x, y int
	what string
}

func (c Cell) findWhat() string {
	return c.what
}

type Entity struct {
	id     int
	pos    Cell
	param1 int
	param2 int
}

//quick
func turnWait() string {
	return "WAIT"
}
func turnMove(c Cell) string {
	s := fmt.Sprintf("MOVE %d %d", c.x, c.y)
	return s
}

type State struct {
	mapHeight      int
	mapWidth       int
	board          [][]Cell
	visited        map[Cell]bool
	me             Entity
	wanderers      []Entity
	explorers      []Entity
	wandererSpawns []Cell
	turn           int
}

//should test linear+create a file map reusable!!
func (s *State) initBoard() {
	var width int
	fmt.Scan(&width)
	s.mapWidth = width

	var height int
	fmt.Scan(&height)
	s.mapHeight = height

	s.board = make([][]Cell, height)

	for y := 0; y < height; y++ {
		s.board[y] = make([]Cell, width)
		var row string
		fmt.Scan(&row)
		for x := range s.board[y] {
			item := strings.Split(row, "")
			s.board[y][x] = Cell{x: x, y: y, what: item[x]}
			if item[x] == "w" {
				s.wandererSpawns = append(s.wandererSpawns, Cell{x: x, y: y, what: item[x]})
			}
		}
	}

}

//quick and dirty to check
func (s State) printBoard() {
	var row string
	for y, _ := range s.board {
		for _, it := range s.board[y] {
			//fmt.Sprintf(row, it.what)
			row += it.what
		}
		row += "\n"
	}
	log.Println(row)
}

//for a given cell got the neighbours
func (s State) getNeighbours(c Cell) []Cell {
	var neighbours []Cell
	//for _, cell := range s.board {
	if s.board[c.y+1][c.x].what == "." && c.y+1 < s.mapHeight {
		neighbours = append(neighbours, s.board[c.y+1][c.x])
	}
	if s.board[c.y-1][c.x].what == "." && c.y-1 > 0 {
		neighbours = append(neighbours, s.board[c.y-1][c.x])
	}
	if s.board[c.y][c.x+1].what == "." && c.x+1 < s.mapWidth {
		neighbours = append(neighbours, s.board[c.y][c.x+1])
	}
	if s.board[c.y][c.x-1].what == "." && c.x-1 > 0 {
		neighbours = append(neighbours, s.board[c.y][c.x-1])
	}
	//}
	return neighbours
}

func (s *State) think() string {
	//simply go to next free?
	var freePos = s.getNeighbours(s.me.pos)
	var action string
	//should score which one is best??? And REALLY, which one is best?? :)))
	if len(freePos) == 0 {
		action = turnWait()
		s.visited[s.me.pos] = true
	} else {
		if s.turn == 0 {
			action = turnMove(freePos[1])
			s.visited[freePos[1]] = true
		} else {
			//should factorize that
			//and examine all cells!!
			for _, cell := range freePos {
				if !s.visited[cell] {
					action = turnMove(cell)
					s.visited[cell] = true
					break
				} else {
					//am cornered!!
					action = turnWait()
				}
			}
		}
	}
	return action
}

func main() {
	s := State{}
	//is this really a good idea in the end??
	s.visited = make(map[Cell]bool)

	s.initBoard()
	s.printBoard()

	// sanityLossLonely: how much sanity you lose every turn when alone, always 3 until wood 1
	// sanityLossGroup: how much sanity you lose every turn when near another player, always 1 until wood 1
	// wandererSpawnTime: how many turns the wanderer take to spawn, always 3 until wood 1
	// wandererLifeTime: how many turns the wanderer is on map after spawning, always 40 until wood 1
	var sanityLossLonely, sanityLossGroup, wandererSpawnTime, wandererLifeTime int
	fmt.Scan(&sanityLossLonely, &sanityLossGroup, &wandererSpawnTime, &wandererLifeTime)

	for {
		// entityCount: the first given entity corresponds to your explorer
		var entityCount int
		fmt.Scan(&entityCount)

		for i := 0; i < entityCount; i++ {
			var entityType string
			var id, x, y, param0, param1, param2 int
			fmt.Scan(&entityType, &id, &x, &y, &param0, &param1, &param2)
			switch entityType {
			case "EXPLORER":
				if i == 0 {
					s.me = Entity{id: id, pos: Cell{x: x, y: y, what: entityType}, param1: param1, param2: param2}
				}
				s.explorers = append(s.explorers, Entity{id: id, pos: Cell{x: x, y: y, what: entityType}, param1: param1, param2: param2})
			case "WANDERER":
				s.wanderers = append(s.wanderers, Entity{id: id, pos: Cell{x: x, y: y, what: entityType}, param1: param1, param2: param2})
			}
		}
		log.Println(s.visited)
		log.Println(s.getNeighbours(s.me.pos))
		output := s.think()
		fmt.Println(output)

		s.explorers = []Entity{}
		s.wanderers = []Entity{}
		s.turn += 1
	}
}
