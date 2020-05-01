package main

import (
	"bytes"
	"container/list"
	"fmt"
	"os"
)

const (
	WIDTH  = 6
	HEIGHT = 12
)

type Cell struct {
	x    int
	y    int
	what string
}

type Block struct {
	ColorFirstBlock  int
	ColorSecondBlock int
}

type State struct {
	Turn       int
	myGrid     [][]Cell
	oppGrid    [][]Cell
	myScore    int
	oppScore   int
	queueBlock list.List
}

func (s *State) initQueueBlock() {
	for i := 0; i < 8; i++ {
		// colorA: color of the first block
		// colorB: color of the attached block
		var colorA, colorB int
		fmt.Scan(&colorA, &colorB)
		s.queueBlock.PushBack(Block{ColorFirstBlock: colorA, ColorSecondBlock: colorB})
	}

}
func (s *State) initPlayer1() {
	var score1 int
	fmt.Scan(&score1)
	s.myScore = score1

	var row string
	for y := 0; y < 12; y++ {
		// row: One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
		var line string
		fmt.Scan(&line)
		row = row + line
	}

	s.myGrid = make([][]Cell, WIDTH)
	for x := 0; x < WIDTH; x++ {
		s.myGrid[x] = make([]Cell, HEIGHT)
		for y := 0; y < HEIGHT; y++ {
			s.myGrid[x][y] = Cell{x, y, string(row[y*WIDTH+x])}
		}
	}
}

func (s *State) initPlayer2() {
	var score2 int
	fmt.Scan(&score2)
	s.oppScore = score2

	var row string
	for y := 0; y < 12; y++ {
		// row: One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
		var line string
		fmt.Scan(&line)
		row = row + line
	}

	s.oppGrid = make([][]Cell, WIDTH)
	for x := 0; x < WIDTH; x++ {
		s.oppGrid[x] = make([]Cell, HEIGHT)
		for y := 0; y < HEIGHT; y++ {
			s.oppGrid[x][y] = Cell{x, y, string(row[y*WIDTH+x])}
		}
	}

}

func printGrid(g [][]Cell) string {
	var buf bytes.Buffer
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			buf.WriteString(g[x][y].what)
		}
		buf.WriteString("\n")
	}
	//fmt.Println()
	return buf.String()
}

/*
//better write a sim func, to output what the grid looks like in different cases
func (s *State) simu() {
	//a copy of the grid
	cpGrid := s.myGrid
	//take the blocks
	//only one at first to see how it goes
	block := s.queueBlocks[0]
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
		}
	}
}
func (s *State) think() {
	nextBlock := s.queueBlock[0]
	var result int
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			//if s.myGrid[y][x]==
		}
	}

	log.Println(result)
}
*/
func main() {
	s := State{}
	for {
		s.initQueueBlock()
		s.initPlayer1()
		s.initPlayer2()
		//s.think()
		fmt.Fprintln(os.Stderr, printGrid(s.oppGrid))
		fmt.Printf("0\n") // "x": the column in which to drop your blocks
		s.Turn += 1
	}
}
