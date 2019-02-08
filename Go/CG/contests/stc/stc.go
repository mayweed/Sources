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

type Point struct {
	X, Y int
}
type Block struct {
	Point
	ColorFirstBlock  int
	ColorSecondBlock int
}

type Grid [HEIGHT][WIDTH]Block

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
		s.queueBlock = append(s.queueBlock, Block{Point: Point{}, ColorFirstBlock: colorA, ColorSecondBlock: colorB})
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
			s.myGrid[y][x] = Block{Point{x, y}, -1, -1}
			switch c {
			case ".":
				s.myGrid[y][x].ColorFirstBlock = -1
				s.myGrid[y][x].ColorSecondBlock = -1
			default:
				s.myGrid[y][x].ColorFirstBlock, _ = strconv.Atoi(c)
				s.myGrid[y][x].ColorSecondBlock, _ = strconv.Atoi(c)
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
			s.oppGrid[y][x] = Block{Point{x, y}, -1, -1}
			switch c {
			case ".":
				s.oppGrid[y][x].ColorFirstBlock = -1
				s.oppGrid[y][x].ColorSecondBlock = -1
			default:
				s.oppGrid[y][x].ColorFirstBlock, _ = strconv.Atoi(c)
				s.oppGrid[y][x].ColorSecondBlock, _ = strconv.Atoi(c)
			}
		}
	}

}

func (s *State) think() {
	nextBlock := s.queueBlock[0]
}

func main() {
	s := State{}
	for {
		s.initQueueBlock()
		s.initPlayer1()
		s.initPlayer2()
		fmt.Printf("0\n") // "x": the column in which to drop your blocks
		s.Turn += 1
		//must clear state at one moment...
		log.Println(s.myGrid)
	}
}
