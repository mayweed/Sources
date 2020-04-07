package main

import (
	"errors"
	"fmt"
)

type Tile struct {
	x    int
	y    int
	what string
}

type Grid [HEIGHT][WIDTH]Tile

func (g *Grid) NewGrid(c string) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			g[x][y] = Tile{x, y, string(c[(y*WIDTH)+x])}
		}
	}
}

func (g *Grid) Get(x, y int) (Tile, error) {
	if x < 0 || x > WIDTH || y < 0 || y > HEIGHT {
		return Tile{}, errors.New("out of bound")
	} else {
		return g[x][y], nil
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
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			fmt.Printf("%v", g[x][y].what)
		}
		fmt.Println()
	}
}

func main() {
	var board = "xx..xxx............xxx.............xxx.....xx.....xxx.....xx.....xxx.......................................................xx.............xx...............xxx............xxx......x.....xxx......x......xx......................"
	var g Grid
	g.NewGrid(board)
	g.printGrid()
}
