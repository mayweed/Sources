package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cell struct {
	x    int
	y    int
	what string
}

type Grid struct {
	h          int
	w          int
	c          [][]Cell
	startCells []Cell
}

func (g *Grid) printAnswer() {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			fmt.Print(g.c[x][y].what)
		}
		fmt.Println()
	}
	fmt.Println()
}
func (g *Grid) getStartCells() {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			if g.c[x][y].what != "." && g.c[x][y].what != "#" {
				g.startCells = append(g.startCells, g.c[x][y])
			}
		}
	}
}
func ff(c []Cell) {
	//for each startCells you get the neighbour
	//if that neighbour is == "." || != "#" you copy cell.what
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	g := Grid{}

	var W int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &W)
	g.w = W

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)
	g.h = H

	var row string
	for i := 0; i < H; i++ {
		scanner.Scan()
		row += scanner.Text()
	}

	g.c = make([][]Cell, g.w)
	for x := 0; x < g.w; x++ {
		g.c[x] = make([]Cell, g.h)
		for y := 0; y < g.h; y++ {
			//not sure i need x and y in cell...(g.c[x][y] and you got it)
			g.c[x][y] = Cell{x, y, string(row[y*g.w+x])}
		}
	}
	g.getStartCells()

	g.printAnswer()

}
