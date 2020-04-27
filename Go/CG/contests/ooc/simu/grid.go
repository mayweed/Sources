package main

import "bytes"

type Tile struct {
	Point
	what string
}

type Grid struct {
	carte        [HEIGHT][WIDTH]Tile
	wtiles       []Tile
	visitedTiles []Tile
	neighbours   map[Tile][]Tile
}

func (g *Grid) NewGrid(c string) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			g.carte[x][y] = Tile{Point{x, y}, string(c[(y*WIDTH)+x]), []Tile{}}
		}
	}
}

func (g *Grid) getWalkableTiles() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if g.carte[x][y].what == "." {
				g.wtiles = append(g.wtiles, g.carte[x][y])
			}
		}
	}
}

//shouldnt be t passes as a pointer?
func (g *Grid) getTileNeighbours(t *Tile) {
	//var t.neighbours = list.NewList()
	if t.x-1 >= 0 && g.carte[t.x-1][t.y].what == "." { //&& !g[t.x-1][t.y].visited { just neigh can check that later no?
		t.neighbours = append(t.neighbours, g.carte[t.x-1][t.y])
	}
	if t.x+1 < WIDTH && g.carte[t.x+1][t.y].what == "." {
		t.neighbours = append(t.neighbours, g.carte[t.x+1][t.y])
	}
	if t.y-1 >= 0 && g.carte[t.x][t.y-1].what == "." {
		t.neighbours = append(t.neighbours, g.carte[t.x][t.y-1])
	}
	if t.y+1 < HEIGHT && g.carte[t.x][t.y+1].what == "." {
		t.neighbours = append(t.neighbours, g.carte[t.x][t.y+1])
	}
}

func (g *Grid) updateGrid(posPlayer Point) {

}
func (g *Grid) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			buf.WriteString(g.carte[x][y].what)
		}
		buf.WriteString("\n")
	}
	//fmt.Println()
	return buf.String()
}

/*
func (g *Grid) printGrid() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			fmt.Printf("%v", g.carte[x][y].what)
		}
		fmt.Println()
	}
}
*/
