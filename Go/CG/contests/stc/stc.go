package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	WIDTH  = 6
	HEIGHT = 12
)

type Block struct {
	ColorFirstBlock  int
	ColorSecondBlock int
}

type Grid [HEIGHT][WIDTH]int

type State struct {
	Turn       int
	myGrid     Grid
	oppGrid    Grid
	myScore    int
	oppScore   int
	queueBlock []Block
}

func (s *State) initQueueBlock() {
	for i := 0; i < 8; i++ {
		// colorA: color of the first block
		// colorB: color of the attached block
		var colorA, colorB int
		fmt.Scan(&colorA, &colorB)
		s.queueBlock = append(s.queueBlock, Block{ColorFirstBlock: colorA, ColorSecondBlock: colorB})
	}

}
func (s *State) initPlayer1() {
	var score1 int
	fmt.Scan(&score1)
	s.myScore = score1

	for y := 0; y < 12; y++ {
		// row: One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
		var row string
		fmt.Scan(&row)

		r := strings.Split(row, "")
		for x, c := range r {
			switch c {
			case ".":
				s.myGrid[y][x] = -1
			default:
				s.myGrid[y][x], _ = strconv.Atoi(c)

			}
		}
	}
}

func (s *State) initPlayer2() {
	var score2 int
	fmt.Scan(&score2)
	s.oppScore = score2

	for y := 0; y < 12; y++ {
		// row: One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
		var row string
		fmt.Scan(&row)

		r := strings.Split(row, "")
		for x, c := range r {
			switch c {
			case ".":
				s.oppGrid[y][x] = -1
			default:
				s.oppGrid[y][x], _ = strconv.Atoi(c)
			}
		}
	}

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
*/
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

func main() {
	s := State{}
	for {
		s.initQueueBlock()
		s.initPlayer1()
		s.initPlayer2()
		s.think()
		fmt.Printf("0\n") // "x": the column in which to drop your blocks
		s.Turn += 1
		//must clear state at one moment...
		//write a proper func to print grid
		//log.Println(s.myGrid)
	}
}
