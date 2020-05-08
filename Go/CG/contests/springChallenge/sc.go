package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}
type Pellet struct {
	Point
	what int
}
type Grid struct {
	h int
	w int
	c [][]Pellet
}

func (g *Grid) NewGrid(c string) {
	g.c = make([][]Pellet, g.w)
	for x := 0; x < g.w; x++ {
		g.c[x] = make([]Pellet, g.h)
		for y := 0; y < g.h; y++ {
			if string(c[(y*g.w)+x]) == "#" {
				g.c[x][y] = Pellet{Point{x, y}, -1}
			} else {
				g.c[x][y] = Pellet{Point{x, y}, 0}
			}

		}
	}
}
func (g *Grid) updateGrid(x, y, what int) {
}

//no pointer: Grid has a string meth, not *Grid !!
func (g *Grid) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			if g.c[x][y].what == -1 {
				buf.WriteString("#")
			} else if g.c[x][y].what != 0 {
				buf.WriteString(strconv.Itoa(g.c[x][y].what))

			} else {
				buf.WriteString(" ")
			}
		}
		buf.WriteString("\n")
	}
	//fmt.Println()
	return buf.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	g := Grid{}
	// width: size of the grid
	// height: top left corner is (x=0, y=0)
	var width, height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)
	g.h = height
	g.w = width

	var line string
	for i := 0; i < height; i++ {
		scanner.Scan()
		line = line + scanner.Text() // one line of the grid: space " " is floor, pound "#" is wall
	}

	g.NewGrid(line)

	for {
		var myScore, opponentScore int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &myScore, &opponentScore)
		// visiblePacCount: all your pacs and enemy pacs in sight
		var visiblePacCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePacCount)

		for i := 0; i < visiblePacCount; i++ {
			// pacId: pac number (unique within a team)
			// mine: true if this pac is yours
			// x: position in the grid
			// y: position in the grid
			// typeId: unused in wood leagues
			// speedTurnsLeft: unused in wood leagues
			// abilityCooldown: unused in wood leagues
			var pacId int
			//var mine bool
			var _mine int
			var x, y int
			var typeId string
			var speedTurnsLeft, abilityCooldown int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &pacId, &_mine, &x, &y, &typeId, &speedTurnsLeft, &abilityCooldown)
			//mine = _mine != 0
		}
		// visiblePelletCount: all pellets in sight
		var visiblePelletCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePelletCount)

		for i := 0; i < visiblePelletCount; i++ {
			// value: amount of points this pellet is worth
			var x, y, value int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &x, &y, &value)
			g.c[x][y].what = value
		}

		fmt.Fprintln(os.Stderr, g.String())
		fmt.Println("MOVE 0 15 10") // MOVE <pacId> <x> <y>
	}
}
