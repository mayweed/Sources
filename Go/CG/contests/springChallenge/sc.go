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
func (g Grid) getNeighbours(p Pellet) []Pellet {
	var neighbours []Pellet
	if p.x+1 < g.w && g.c[p.x+1][p.y].what != -1 {
		neighbours = append(neighbours, g.c[p.x+1][p.y])
	}
	if p.x-1 >= 0 && g.c[p.x-1][p.y].what != -1 {
		neighbours = append(neighbours, g.c[p.x-1][p.y])
	}
	if p.y+1 < g.h && g.c[p.x][p.y+1].what != -1 {
		neighbours = append(neighbours, g.c[p.x][p.y+1])
	}
	if p.y-1 >= 0 && g.c[p.x][p.y-1].what != -1 {
		neighbours = append(neighbours, g.c[p.x][p.y-1])
	}
	return neighbours

}
func (g Grid) bfs(startPoint, endPoint Pellet) []Pellet {
	var visited = make(map[Pellet]bool)
	var pred = make(map[Pellet]Pellet)
	var path []Pellet
	var queue []Pellet
	queue = append(queue, startPoint)

	for len(queue) != 0 {
		s := queue[0]
		queue = queue[1:]
		if s == endPoint {
			path = append(path, s)
			for pred[endPoint] != startPoint {
				path = append(path, pred[endPoint])
				endPoint = pred[endPoint]
			}
			path = append(path, startPoint)
			break //???
		}
		for _, n := range g.getNeighbours(s) {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
				pred[n] = s
			}

		}
	}
	return path
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

		var myPos Pellet
		for i := 0; i < visiblePacCount; i++ {
			// pacId: pac number (unique within a team)
			// mine: true if this pac is yours
			// x: position in the grid
			// y: position in the grid
			// typeId: unused in wood leagues
			// speedTurnsLeft: unused in wood leagues
			// abilityCooldown: unused in wood leagues
			var pacId int
			var mine int
			var x, y int
			var typeId string
			var speedTurnsLeft, abilityCooldown int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &pacId, &mine, &x, &y, &typeId, &speedTurnsLeft, &abilityCooldown)
			if mine == 1 {
				myPos = g.c[x][y]
			}
		}
		// visiblePelletCount: all pellets in sight
		var visiblePelletCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePelletCount)

		var valuablePellets []Pellet
		for i := 0; i < visiblePelletCount; i++ {
			// value: amount of points this pellet is worth
			var x, y, value int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &x, &y, &value)
			g.c[x][y].what = value
			if value == 10 {
				valuablePellets = append(valuablePellets, g.c[x][y])
			}
		}
		//must extract the most valuable and the shortest
		//Best are those with multiple big pac
		var paths [][]Pellet
		for _, vp := range valuablePellets {
			p := g.bfs(myPos, vp)
			paths = append(paths, p)
		}

		//p := g.bfs(g.c[19][2], g.c[25][7])
		//fmt.Fprintln(os.Stderr, g.String())
		fmt.Fprintln(os.Stderr, myPos, len(paths), paths)
		fmt.Println("MOVE 0 15 10") // MOVE <pacId> <x> <y>
	}
}
