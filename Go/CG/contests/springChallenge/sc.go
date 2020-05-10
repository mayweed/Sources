package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type (
	Point struct {
		x, y int
	}
	Grid struct {
		h               int
		w               int
		c               [][]Pellet
		valuablePellets []Pellet
		pellets         []Pellet
	}
	Pellet struct {
		Point
		what int
	}
	Pac struct {
		Point
		id int
	}
	//what about a player id == pacid to send comm?
	Player struct {
		score int
		pacs  []Pac
	}
	//state of the game per turn
	Turn struct {
		g        Grid
		me       Player
		p        []Player
		commands []string
	}
)

func move(id, x, y int) string {
	return fmt.Sprintf("MOVE %d %d %d", id, x, y)
}
func (t Turn) sendTurn() {
	if len(t.commands) == 1 {
		fmt.Print(t.commands[0])
	} else {
		fmt.Print(strings.Join(t.commands, "|"))
	}
	fmt.Println()
}

//Grid meth
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
func (g Grid) pathToValPellet(myPos Pellet) []Pellet {
	//must extract the most valuable and the shortest
	//Best are those with multiple big pac
	var paths [][]Pellet
	var shortestPath []Pellet
	for _, vp := range g.valuablePellets {
		p := g.bfs(myPos, vp)
		paths = append(paths, p)
	}
	var min = 1000
	for _, p := range paths {
		//log.Println(len(p))
		if len(p) < min {
			shortestPath = p
			min = len(p)
		}
	}
	return shortestPath
}
func (g *Grid) findValuablePath(myPos Pellet) []Pellet {
	//could do that at the beginning!! to all the grid!! to spot the highest
	//paying path!!
	var path []Pellet
	//calculate all possible paths to pellets
	//and extract the one with highest value which i will follow
	var max = 0
	var value int
	var myPath []Pellet
	for _, p := range g.pellets {
		path = g.bfs(myPos, p)
		value = getPathValue(path)
		//log.Println("inside f value", value)
		if value >= max {
			myPath = path
			max = value
		}
	}
	//log.Println("inside f", myPath)
	return myPath

}
func getPathValue(path []Pellet) int {
	var value int
	for _, p := range path {
		value += p.what
	}
	return value
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
		var t Turn
		var myScore, opponentScore int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &myScore, &opponentScore)
		t.me.score = myScore

		// visiblePacCount: all your pacs and enemy pacs in sight
		var visiblePacCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &visiblePacCount)

		//must update grid to see what's left
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
				t.me.pacs = append(t.me.pacs, Pac{Point{x, y}, pacId})
			}
			//update grid /turn?
			g.c[x][y].what = 0
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
			if value == 10 {
				g.valuablePellets = append(g.valuablePellets, g.c[x][y])
			} else {
				g.pellets = append(g.pellets, g.c[x][y])
			}

		}

		//Think : grab all the big pellets first and then just roam in cell with
		//pellet!!
		//one day should check if two pacs go for the same pellet...
		//var possPaths [][]Pellet
		var possPaths = make(map[int][][]Pellet) //==> where point is pac loc?
		if len(g.valuablePellets) > 0 {
			for _, p := range t.me.pacs {
				possPaths[p.id] = append(possPaths[p.id], g.pathToValPellet(g.c[p.x][p.y]))
			}
			for pid, path := range possPaths {
				t.commands = append(t.commands, move(pid, path[0][0].x, path[0][0].y))
			}
		} else {
			///and here...they must continue roaming? what about rand on g.pellets??
			//findvalpath does not work :'''(
			for _, p := range t.me.pacs {
				randIdx := rand.Intn(len(g.pellets))
				t.commands = append(t.commands, move(p.id, g.pellets[randIdx].x, g.pellets[randIdx].y))
			}
		}

		fmt.Fprintln(os.Stderr, possPaths)

		//output
		//fmt.Println(res) // MOVE <pacId> <x> <y>
		t.sendTurn()

		//reset
		g.valuablePellets = []Pellet{}
	}
}
