package main

import (
	"fmt"
	"log"
	"strings"
)

type Point struct {
	x, y int
}

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	pos  Point
	what string
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
			s.grid[i][j].pos = Point{j, i}
			s.grid[i][j].what = inputs[j]
		}
	}

}

//for a given pos is there any nodes and how many?
//should handle better indestructible nodes
//this func must be used after to eval boardstate
func (s State) nodesInRange(x, y int) (bool, int) {
	var count int
	var sn bool
	var pn bool
	//first X
	if x-1 > 0 {
		if s.grid[y][x-1].what == "@" {
			count += 1
			sn = true
		} else if s.grid[y][x-1].what == "#" {
			pn = true
		}
	}
	if x-2 > 0 {
		if s.grid[y][x-2].what == "@" && !pn {
			count += 1
			//must reinit it if there is any more in other dirs
			pn = false
		}
	}
	if x+1 < s.width {
		if s.grid[y][x+1].what == "@" {
			count += 1
			sn = true
		} else if s.grid[y][x+1].what == "#" {
			pn = true
		}
	}
	if x+2 < s.width {
		if s.grid[y][x+2].what == "@" && !pn {
			count += 1
			//must reinit it if there is any more in other dirs
			pn = false
		}
	}
	//then Y
	if y-1 > 0 {
		if s.grid[y-1][x].what == "@" {
			count += 1
			sn = true
		} else if s.grid[y-1][x].what == "#" {
			pn = true
		}
	}
	if y-2 > 0 {
		if s.grid[y-2][x].what == "@" && !pn {
			count += 1
			//must reinit it if there is any more in other dirs
			pn = false
		}
	}
	if y+1 < s.height {
		if s.grid[y+1][x].what == "@" {
			count += 1
			sn = true
		} else if s.grid[y+1][x].what == "#" {
			pn = true
		}
	}
	if y+2 < s.height {
		if s.grid[y+2][x].what == "@" && !pn {
			count += 1
			//must reinit it if there is any more in other dirs
			pn = false
		}
	}
	return sn, count
}

//should build a tree of possible result if i place a bomb here
func simulateTurn(s State) []int {
	//g := s.grid
	//generate random moves
	//eval the moves or better the board
	//would it be too costly to eval nodes in range of each cell??
	//this is a bloody test ;)
	var count []int
	for y := 0; y < s.height; y++ {
		for _, cell := range s.grid[y] {
			b, c := s.nodesInRange(cell.pos.x, cell.pos.y)
			if b {
				//all the count by cell, shouldn't I use map[Cell]int here??
				count = append(count, c)
			}
		}
	}
	return count
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

		log.Println(simulateTurn(s))
		log.Println(s.nodesInRange(2, 1))
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("2 1") // Write action to stdout
		s.turn += 1
	}
}
