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

type State struct {
	board      [][]Cell
	wanderers  []Entity
	explorers  []Entity
	emptyCells []Cell
}

//should test linear+create a file map reusable!!
func (s *State) initBoard() {
	var width int
	fmt.Scan(&width)

	var height int
	fmt.Scan(&height)

	s.board = make([][]Cell, height)

	for y := 0; y < height; y++ {
		s.board[y] = make([]Cell, width)
		var row string
		fmt.Scan(&row)
		for x := range s.board[y] {
			item := strings.Split(row, "")
			s.board[y][x] = Cell{x: x, y: y, what: item[x]}
			if item[x] == "." {
				s.emptyCells = append(s.emptyCells, Cell{x: x, y: y, what: item[x]})
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

func main() {
	gState := State{}
	gState.initBoard()
	gState.printBoard()

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
				gState.explorers = append(gState.explorers, Entity{id: id, pos: Cell{x: x, y: y, what: entityType}, param1: param1, param2: param2})
			case "WANDERER":
				gState.wanderers = append(gState.wanderers, Entity{id: id, pos: Cell{x: x, y: y, what: entityType}, param1: param1, param2: param2})
			}
		}
		log.Println(gState.explorers)
		fmt.Println("WAIT") // MOVE <x> <y> | WAIT

		gState.explorers = []Entity{}
		gState.wanderers = []Entity{}
	}
}
