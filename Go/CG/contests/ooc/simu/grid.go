package main

import (
	"errors"
	"fmt"
)

type Tile struct {
	//x    int
	//y    int
	Point
	what string
}

type Grid [HEIGHT][WIDTH]Tile

func (g *Grid) NewGrid(c string) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			g[x][y] = Tile{Point{x, y}, string(c[(y*WIDTH)+x])}
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
func (g *Grid) getWalkableTiles() []Tile {
	var wtiles []Tile
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if g[x][y].what == "." {
				wtiles = append(wtiles, g[x][y])
			}
		}
	}
	return wtiles
}

func (g *Grid) updateGrid(posPlayer Point) {

}
func (g *Grid) printGrid() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			fmt.Printf("%v", g[x][y].what)
		}
		fmt.Println()
	}
}
