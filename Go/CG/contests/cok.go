package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	x, y int
	what string
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

func (c Cell) findWhat() {}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)

	var height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)

	gState := State{}
	gState.board = make([][]Cell, height)

	for y := 0; y < height; y++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), "")
		gState.board[y] = make([]Cell, width)
		for x := range gState.board[y] {
			gState.board[y][x] = Cell{x: x, y: y, what: inputs[x]}
			if inputs[x] == "." {
				gState.emptyCells = append(gState.emptyCells, Cell{x: x, y: y, what: inputs[x]})
			}
		}
	}
	// sanityLossLonely: how much sanity you lose every turn when alone, always 3 until wood 1
	// sanityLossGroup: how much sanity you lose every turn when near another player, always 1 until wood 1
	// wandererSpawnTime: how many turns the wanderer take to spawn, always 3 until wood 1
	// wandererLifeTime: how many turns the wanderer is on map after spawning, always 40 until wood 1
	var sanityLossLonely, sanityLossGroup, wandererSpawnTime, wandererLifeTime int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &sanityLossLonely, &sanityLossGroup, &wandererSpawnTime, &wandererLifeTime)

	for {
		// entityCount: the first given entity corresponds to your explorer
		var entityCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &entityCount)

		for i := 0; i < entityCount; i++ {
			var entityType string
			var id, x, y, param0, param1, param2 int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &entityType, &id, &x, &y, &param0, &param1, &param2)
			switch entityType {
			case "EXPLORER":
				gState.explorers = append(gState.explorers, Entity{id: id, pos: Cell{x: x, y: y, what: entityType}, param1: param1, param2: param2})
			case "WANDERER":
				gState.wanderers = append(gState.wanderers, Entity{id: id, pos: Cell{x: x, y: y, what: entityType}, param1: param1, param2: param2})
			}
		}
		log.Println(gState.explorers)
		fmt.Println("WAIT") // MOVE <x> <y> | WAIT
	}
}
