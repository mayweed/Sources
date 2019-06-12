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
	x, y  int
	Color int
}
type Block struct {
	ColorFirstBlock  int
	ColorSecondBlock int
}

type Grid [HEIGHT][WIDTH]Point

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
				s.myGrid[y][x] = Point{y, x, -1}
			default:
				d, _ := strconv.Atoi(c)
				s.myGrid[y][x] = Point{y, x, d}
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
				s.oppGrid[y][x] = Point{y, x, -1}
			default:
				d, _ := strconv.Atoi(c)
				s.oppGrid[y][x] = Point{y, x, d}
			}
		}
	}

}

func (s *State) think() {
	nextBlock := s.queueBlock[0]
	var result int
	//better
	cpGrid := s.myGrid
	//then you make the block falling *time and evaluate...
	for _, col := range s.myGrid {
		//BEURK
		for _, p := range col {
			if p.Color == nextBlock.ColorFirstBlock {
				if p.y+1 == -1 && p.y+1 < WIDTH {
					result = p.y + 1
				}
			}
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
