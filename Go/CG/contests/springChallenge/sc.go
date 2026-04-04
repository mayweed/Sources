package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Pellet struct {
	Point
	what int
}

type Pac struct {
	Point
	id              int
	possibleActions Action
}

type Action struct {
	canGoUp    bool
	canGoDown  bool
	canGoLeft  bool
	canGoRight bool
}

type Player struct {
	score int
	pacs  map[int]*Pac
}

type Grid struct {
	h, w            int
	c               [][]Pellet
	valuablePellets []Pellet
	pellets         []Pellet
}

type Turn struct {
	g        *Grid
	me       Player
	commands []string
}

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

// --- GRID INITIALISATION ---------------------------------------------------

func (g *Grid) NewGrid(c string) {
	g.c = make([][]Pellet, g.w)
	for x := 0; x < g.w; x++ {
		g.c[x] = make([]Pellet, g.h)
		for y := 0; y < g.h; y++ {
			if c[(y*g.w)+x] == '#' {
				g.c[x][y] = Pellet{Point{x, y}, -1}
			} else {
				g.c[x][y] = Pellet{Point{x, y}, 0}
			}
		}
	}
}

// --- NEIGHBOURS ------------------------------------------------------------

func (g Grid) getNeighbours(p Pellet) []Pellet {
	res := []Pellet{}
	if p.x+1 < g.w && g.c[p.x+1][p.y].what != -1 {
		res = append(res, g.c[p.x+1][p.y])
	}
	if p.x-1 >= 0 && g.c[p.x-1][p.y].what != -1 {
		res = append(res, g.c[p.x-1][p.y])
	}
	if p.y+1 < g.h && g.c[p.x][p.y+1].what != -1 {
		res = append(res, g.c[p.x][p.y+1])
	}
	if p.y-1 >= 0 && g.c[p.x][p.y-1].what != -1 {
		res = append(res, g.c[p.x][p.y-1])
	}
	return res
}

func (g Grid) bfs(startPoint, endPoint Pellet) []Pellet {
	visited := make(map[Pellet]bool)
	pred := make(map[Pellet]Pellet)
	queue := []Pellet{}

	visited[startPoint] = true
	queue = append(queue, startPoint)

	found := false

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if s == endPoint {
			found = true
			break
		}

		for _, n := range g.getNeighbours(s) {
			if !visited[n] {
				visited[n] = true
				pred[n] = s
				queue = append(queue, n)
			}
		}
	}

	if !found {
		return nil
	}

	path := []Pellet{}
	curr := endPoint
	path = append(path, curr)

	for curr != startPoint {
		curr = pred[curr]
		path = append(path, curr)
	}

	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}

	return path
}

// --- STRAT : TROUVER LE PLUS COURT CHEMIN VERS UN GROS PELLET --------------

func (g Grid) pathToValPellet(myPos Pellet) []Pellet {
	if len(g.valuablePellets) == 0 {
		return nil
	}

	var best []Pellet
	min := 10000

	for _, vp := range g.valuablePellets {
		p := g.bfs(myPos, vp)
		if len(p) > 0 && len(p) < min {
			min = len(p)
			best = p
		}
	}

	return best
}

// --- PAC POSSIBLE MOVES -----------------------------------------------------

func (g Grid) getPossibleMoves(p *Pac) {
	p.possibleActions = Action{}

	if p.x+1 < g.w && g.c[p.x+1][p.y].what != -1 {
		p.possibleActions.canGoRight = true
	}
	if p.x-1 >= 0 && g.c[p.x-1][p.y].what != -1 {
		p.possibleActions.canGoLeft = true
	}
	if p.y+1 < g.h && g.c[p.x][p.y+1].what != -1 {
		p.possibleActions.canGoDown = true
	}
	if p.y-1 >= 0 && g.c[p.x][p.y-1].what != -1 {
		p.possibleActions.canGoUp = true
	}
}

// --- MAIN LOOP --------------------------------------------------------------

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1_000_000), 1_000_000)

	g := &Grid{}
	var width, height int

	sc.Scan()
	fmt.Sscan(sc.Text(), &width, &height)
	g.w = width
	g.h = height

	var line string
	for i := 0; i < height; i++ {
		sc.Scan()
		line += sc.Text()
	}

	g.NewGrid(line)

	for {
		var t Turn
		t.g = g

		var myScore, oppScore int
		sc.Scan()
		fmt.Sscan(sc.Text(), &myScore, &oppScore)
		t.me.score = myScore

		var visiblePacCount int
		sc.Scan()
		fmt.Sscan(sc.Text(), &visiblePacCount)

		t.me.pacs = make(map[int]*Pac)

		// Reset all pellets to 0 (we refill after)
		for x := 0; x < g.w; x++ {
			for y := 0; y < g.h; y++ {
				if g.c[x][y].what > 0 {
					g.c[x][y].what = 0
				}
			}
		}

		// All pacs (mine + enemy)
		for i := 0; i < visiblePacCount; i++ {
			var pacId, mine, x, y int
			var typeId string
			var speedTurnsLeft, abilityCooldown int
			sc.Scan()
			fmt.Sscan(sc.Text(), &pacId, &mine, &x, &y, &typeId, &speedTurnsLeft, &abilityCooldown)

			if mine == 1 {
				t.me.pacs[pacId] = &Pac{Point{x, y}, pacId, Action{}}
			}

			g.c[x][y].what = 0
		}

		// Pellets visibles
		var visiblePelletCount int
		sc.Scan()
		fmt.Sscan(sc.Text(), &visiblePelletCount)

		g.valuablePellets = []Pellet{}
		g.pellets = []Pellet{}

		for i := 0; i < visiblePelletCount; i++ {
			var x, y, value int
			sc.Scan()
			fmt.Sscan(sc.Text(), &x, &y, &value)

			g.c[x][y].what = value
			if value == 10 {
				g.valuablePellets = append(g.valuablePellets, g.c[x][y])
			} else {
				g.pellets = append(g.pellets, g.c[x][y])
			}
		}

		// STRAT : big pellets first
		possPaths := map[int][]Pellet{}

		if len(g.valuablePellets) > 0 {
			for pid, p := range t.me.pacs {
				myCell := g.c[p.x][p.y]
				possPaths[pid] = g.pathToValPellet(myCell)
			}

			for pid, path := range possPaths {
				if path != nil && len(path) > 1 {
					next := path[1]
					t.commands = append(t.commands, move(pid, next.x, next.y))
				} else {
					p := t.me.pacs[pid]
					t.commands = append(t.commands, move(pid, p.x, p.y))
				}
			}

		} else {

			for _, p := range t.me.pacs {
				g.getPossibleMoves(p)
				if p.possibleActions.canGoRight {
					t.commands = append(t.commands, move(p.id, p.x+1, p.y))
				} else if p.possibleActions.canGoLeft {
					t.commands = append(t.commands, move(p.id, p.x-1, p.y))
				} else if p.possibleActions.canGoUp {
					t.commands = append(t.commands, move(p.id, p.x, p.y-1))
				} else if p.possibleActions.canGoDown {
					t.commands = append(t.commands, move(p.id, p.x, p.y+1))
				}
			}
		}

		t.sendTurn()
	}
}
