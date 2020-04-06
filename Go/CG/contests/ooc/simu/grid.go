package main

import (
	"errors"
	"fmt"
	"log"
)

type Tile struct {
	x    int
	y    int
	what string
}
type Grid struct {
	height   int
	width    int
	contents [][]Tile
}

func (g *Grid) NewGrid(h int, w int, c string) {
	g.height = h
	g.width = w

	g.contents = make([][]Tile, g.height)
	for y := range g.contents {
		g.contents[y] = make([]Tile, g.width)
		for x := range g.contents[y] {
			g.contents[y][x] = Tile{x, y, string(c[(y*g.width)+x])}
		}
	}
}

func (g *Grid) Get(x, y int) (Tile, error) {
	if x < 0 || x > g.width || y < 0 || y > g.height {
		return Tile{}, errors.New("out of bound")
	} else {
		return g.contents[y][x], nil
	}

}
func (g Grid) getNeigh(x, y int) (Tile, error) {
	t, err := g.Get(x, y+1)
	if err != nil {
		return Tile{}, errors.New("out of bound")
	} else {
		return t, nil
	}
}

func (g *Grid) printGrid() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			fmt.Printf("%v", g.contents[x][y].what)
			//log.Println(x, y, g.contents[x][y].what)
		}
		fmt.Println()
	}
}

func main() {
	var board = "xx..xxx............xxx.............xxx.....xx.....xxx.....xx.....xxx.......................................................xx.............xx...............xxx............xxx......x.....xxx......x......xx......................"
	var g Grid
	g.NewGrid(15, 15, board)
	//g.NewGrid(15, 15, "xx..xxx............xxx.............xxx.....xx.....xxx.....xx.....xxx.......................................................xx.............xx...............xxx............xxx......x.....xxx......x......xx......................")
	g.printGrid()
	log.Println(g.getNeigh(0, 14))
	fmt.Printf("%+v\n", g.contents[5][5])
}
