package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	position Point
	what     string
}

type State struct {
	width  int
	height int
	grid   [][]Cell
	turn   int
}

func (s *State) initGrid() {
	// width: width of the firewall grid
	// height: height of the firewall grid
	var width, height int
	fmt.Scan(&width, &height)
	s.width = width
	s.height = height

	s.grid = make([][]Cell, height)
	var mapRow string
	for i := 0; i < height; i++ {
		fmt.Scan(&mapRow)
		inputs := strings.Split(mapRow, "")
		s.grid[i] = make([]Cell, width)
		for j := 0; j < width; j++ {
			s.grid[i][j].position = Point{j, i}
			s.grid[i][j].what = inputs[j]
		}
	}

}

//should build a tree of possible result if i place a bomb here
func simulateTurn(s State) {
	//g := s.grid

}

func wait() string {
	return "WAIT"
}
func bomb(x, y int) string {
	s := fmt.Sprintf("%d %d", x, y)
	return s
}

func main() {
	s := State{}
	s.initGrid()
	for {
		// rounds: number of rounds left before the end of the game
		// bombs: number of bombs left
		var rounds, bombs int
		fmt.Scan(&rounds, &bombs)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("3 0") // Write action to stdout
		s.turn += 1
	}
}
