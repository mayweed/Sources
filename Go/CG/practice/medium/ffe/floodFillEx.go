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

	log.Println(g, g.c[3][4])
	fmt.Println("answer") // Write answer to stdout
}
