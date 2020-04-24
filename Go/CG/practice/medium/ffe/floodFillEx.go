package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Cell struct {
	x    int
	y    int
	what string
}

type Grid struct {
	h int
	w int
	c [][]Cell
}

func (g *Grid) printAnswer() {
	for y := 0; y < g.w; y++ {
		for x := 0; x < g.h; x++ {
			fmt.Print(g.c[y][x].what)
		}
		fmt.Println()
	}
	fmt.Println()
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

	g.c = make([][]Cell, g.h)
	for y := 0; y < g.h; y++ {
		g.c[y] = make([]Cell, g.w)
		for x := 0; x < g.w; x++ {
			//not sure i need x and y in cell...(g.c[x][y] and you got it)
			g.c[y][x] = Cell{x, y, string(row[y*g.w+x])}
		}
	}

	log.Println(g.c[5][4])
	g.printAnswer()
	//fmt.Println("answer") // Write answer to stdout
}
