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

func main() {
	var g Grid
	var queueBlock []Block
	for {
		for i := 0; i < 8; i++ {
			// colorA: color of the first block
			// colorB: color of the attached block
			var colorA, colorB int
			fmt.Scan(&colorA, &colorB)
			queueBlock = append(queueBlock, Block{Point: Point{}, ColorFirstBlock: colorA, ColorSecondBlock: colorB})
		}
		var score1 int
		fmt.Scan(&score1)

		for y := 0; y < 12; y++ {
			// row: One line of the map ('.' = empty, '0' = skull block, '1' to '5' = colored block)
			var row string
			fmt.Scan(&row)

			r := strings.Split(row, "")
			for x, c := range r {
				g[y][x] = Block{Point{x, y}, -1, -1}
				switch c {
				case ".":
					g[y][x].ColorFirstBlock = -1
					g[y][x].ColorSecondBlock = -1
				default:
					g[y][x].ColorFirstBlock, _ = strconv.Atoi(c)
					g[y][x].ColorSecondBlock, _ = strconv.Atoi(c)
				}
			}
		}
		var score2 int
		fmt.Scan(&score2)

		for i := 0; i < 12; i++ {
			var row string
			fmt.Scan(&row)
		}
		log.Println(g)

		fmt.Printf("0\n") // "x": the column in which to drop your blocks
	}
}
