package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ------------------------------------------------------------
// CORE CONSTANTS
// ------------------------------------------------------------
const (
	HEIGHT = 15
	WIDTH  = 15
)

// ------------------------------------------------------------
// BASIC TYPES
// ------------------------------------------------------------

type Point struct {
	x, y int
}

type Tile struct {
	pos   Point
	what  string
	color string
}

type Me struct {
	id              int
	currentPos      Tile
	hitPoints       int
	canGoNorth      bool
	canGoSouth      bool
	canGoWest       bool
	canGoEast       bool
	torpedoCooldown int
	oppSurfaceHint  string
}

type Opp struct {
	hitPoints    int
	oppDirection string
	torpedoPos   []Point
	enemyZone    int
}

// ------------------------------------------------------------
// TURN & COMMANDS
// ------------------------------------------------------------

type Turn struct {
	commands []string
}

func (t *Turn) move(dir string, charge bool) {
	cmd := fmt.Sprintf("MOVE %s", dir)
	if charge {
		cmd += " TORPEDO"
	}
	t.commands = append(t.commands, cmd)
}

func (t *Turn) surface() {
	t.commands = append(t.commands, "SURFACE")
}

func (t *Turn) torpedo(tile Tile) {
	cmd := fmt.Sprintf("TORPEDO %d %d", tile.pos.x, tile.pos.y)
	t.commands = append(t.commands, cmd)
}

func (t *Turn) msg(s string) {
	t.commands = append(t.commands, fmt.Sprintf("MSG %s", s))
}

func (t *Turn) sendTurn() {
	if len(t.commands) == 1 {
		fmt.Print(t.commands[0])
	} else {
		fmt.Print(strings.Join(t.commands, "|"))
	}
	fmt.Println()
}

// ------------------------------------------------------------
// STATE & GAME LOGIC HOLDER
// ------------------------------------------------------------

type State struct {
	board         string
	carte         [HEIGHT][WIDTH]Tile
	walkableTiles []Tile
	visitedTiles  map[Point]bool
	me            Me
	opp           Opp
	t             Turn
}

// ------------------------------------------------------------
// GRID INITIALIZATION
// ------------------------------------------------------------

// Checks whether a tile is walkable (NOTE: includes space!)
func isWalkable(t Tile) bool {
	return t.what == "." || t.what == " "
}

// Assign enemy zone based on position
func (s *State) getPosBySector(p Point) {
	switch {
	case p.x >= 0 && p.x <= 4 && p.y >= 0 && p.y <= 4:
		s.opp.enemyZone = 1
	case p.x >= 0 && p.x <= 4 && p.y > 4 && p.y <= 9:
		s.opp.enemyZone = 4
	case p.x >= 0 && p.x <= 4 && p.y > 9 && p.y <= 14:
		s.opp.enemyZone = 7

	case p.x > 4 && p.x <= 9 && p.y >= 0 && p.y <= 4:
		s.opp.enemyZone = 2
	case p.x > 4 && p.x <= 9 && p.y > 4 && p.y <= 9:
		s.opp.enemyZone = 5
	case p.x > 4 && p.x <= 9 && p.y > 9 && p.y <= 14:
		s.opp.enemyZone = 8

	case p.x > 9 && p.x <= 14 && p.y >= 0 && p.y <= 4:
		s.opp.enemyZone = 3
	case p.x > 9 && p.x <= 14 && p.y > 4 && p.y <= 9:
		s.opp.enemyZone = 6
	case p.x > 9 && p.x <= 14 && p.y > 9 && p.y <= 14:
		s.opp.enemyZone = 9
	}
}

// ------------------------------------------------------------
// PARSING OPPONENT ORDERS
// ------------------------------------------------------------

func (s *State) parseOppOrders(orders string) {
	clean := strings.ReplaceAll(orders, "|", " ")
	parts := strings.Split(clean, " ")

	for idx, w := range parts {

		// direction of opponent
		if w == "N" || w == "S" || w == "E" || w == "W" {
			s.opp.oppDirection = w
		}

		// torpedo firing
		if w == "TORPEDO" {
			x, _ := strconv.Atoi(parts[idx+1])
			y, _ := strconv.Atoi(parts[idx+2])
			s.opp.torpedoPos = append(s.opp.torpedoPos, Point{x, y})
		}

		// SURFACE sectors hint
		if w == "SURFACE" {
			sector, _ := strconv.Atoi(parts[idx+1])
			s.opp.enemyZone = sector

			// simple hint mapping
			switch sector {
			case 1:
				s.me.oppSurfaceHint = "NW"
			case 2:
				s.me.oppSurfaceHint = "N"
			case 3:
				s.me.oppSurfaceHint = "NE"
			case 4:
				s.me.oppSurfaceHint = "CW"
			case 5:
				s.me.oppSurfaceHint = "C"
			case 6:
				s.me.oppSurfaceHint = "CE"
			case 7:
				s.me.oppSurfaceHint = "SW"
			case 8:
				s.me.oppSurfaceHint = "S"
			case 9:
				s.me.oppSurfaceHint = "SE"
			}
		}
	}
}

// ------------------------------------------------------------
// GRID NEIGHBOURS & MOVEMENT HELPERS
// ------------------------------------------------------------

// Returns all valid neighbour tiles (N/S/E/W)
func (s *State) getNeighbours(t Tile) []Tile {
	neighbours := []Tile{}

	x := t.pos.x
	y := t.pos.y

	// North
	if y-1 >= 0 && isWalkable(s.carte[x][y-1]) {
		neighbours = append(neighbours, s.carte[x][y-1])
	}
	// South
	if y+1 < HEIGHT && isWalkable(s.carte[x][y+1]) {
		neighbours = append(neighbours, s.carte[x][y+1])
	}
	// West
	if x-1 >= 0 && isWalkable(s.carte[x-1][y]) {
		neighbours = append(neighbours, s.carte[x-1][y])
	}
	// East
	if x+1 < WIDTH && isWalkable(s.carte[x+1][y]) {
		neighbours = append(neighbours, s.carte[x+1][y])
	}

	return neighbours
}

// ------------------------------------------------------------
// BFS (Breadth-First Search) — CORRIGÉ
// ------------------------------------------------------------
// Retourne un vrai chemin minimal entre deux tiles.
// Si aucun chemin, retourne nil.

func (s *State) BFS(start Tile, goal Tile) []Tile {

	if start.pos == goal.pos {
		return []Tile{start}
	}

	visited := make(map[Point]bool)
	parent := make(map[Point]Point)

	queue := []Tile{}
	queue = append(queue, start)
	visited[start.pos] = true

	found := false

	// ---- BFS principal ----
	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		// Objectif atteint ?
		if current.pos == goal.pos {
			found = true
			break
		}

		// Explore les voisins
		for _, nei := range s.getNeighbours(current) {
			if !visited[nei.pos] {
				visited[nei.pos] = true
				parent[nei.pos] = current.pos
				queue = append(queue, nei)
			}
		}
	}

	// ---- Pas de chemin ----
	if !found {
		return nil
	}

	// ---- Reconstruction du chemin ----
	path := []Tile{}
	curr := goal.pos

	for curr != start.pos {
		path = append(path, s.carte[curr.x][curr.y])
		curr = parent[curr]
	}
	path = append(path, start)

	// ---- Reverse path (pour faire start → goal) ----
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}

	return path
}

// ------------------------------------------------------------
// FLOODFILL (OPTIONNEL POUR AMÉLIORER IA)
// ------------------------------------------------------------
// Donne un score de "liberté" d’une zone pour éviter de s’enfermer.

func (s *State) floodCount(start Tile) int {
	visited := make(map[Point]bool)
	stack := []Tile{start}
	visited[start.pos] = true

	count := 0

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		for _, nei := range s.getNeighbours(cur) {
			if !visited[nei.pos] {
				visited[nei.pos] = true
				stack = append(stack, nei)
			}
		}
	}

	return count
}

// ------------------------------------------------------------
// MOVEMENT LOGIC (CORRIGÉE ET STABLE)
// ------------------------------------------------------------

// Check if the current tile has already been visited
func (s *State) hasBeenVisited(t Tile) bool {
	return s.visitedTiles[t.pos]
}

// Reset all movement directions before computing new ones
func (s *State) resetDirections() {
	s.me.canGoNorth = false
	s.me.canGoSouth = false
	s.me.canGoWest = false
	s.me.canGoEast = false
}

// Evaluate possible movement directions from your current tile
func (s *State) checkDirections(t Tile) {

	x := t.pos.x
	y := t.pos.y

	// NORTH
	if y-1 >= 0 && isWalkable(s.carte[x][y-1]) && !s.hasBeenVisited(s.carte[x][y-1]) {
		s.me.canGoNorth = true
	}

	// SOUTH
	if y+1 < HEIGHT && isWalkable(s.carte[x][y+1]) && !s.hasBeenVisited(s.carte[x][y+1]) {
		s.me.canGoSouth = true
	}

	// WEST
	if x-1 >= 0 && isWalkable(s.carte[x-1][y]) && !s.hasBeenVisited(s.carte[x-1][y]) {
		s.me.canGoWest = true
	}

	// EAST
	if x+1 < WIDTH && isWalkable(s.carte[x+1][y]) && !s.hasBeenVisited(s.carte[x+1][y]) {
		s.me.canGoEast = true
	}
}

// ------------------------------------------------------------
// MOVEMENT DECISION (PRINCIPALE)
// ------------------------------------------------------------

func (s *State) possibleDir() {

	// 1) Reset direction booleans
	s.resetDirections()

	// 2) Recompute possible moves
	s.checkDirections(s.me.currentPos)

	// 3) Torpedo charge allowed ?
	charge := (s.me.torpedoCooldown <= 0)

	// 4) Movement priority rule:
	//    SOUTH → EAST → NORTH → WEST
	moved := false

	if s.me.canGoSouth {
		s.t.move("S", charge)
		moved = true
	} else if s.me.canGoEast {
		s.t.move("E", charge)
		moved = true
	} else if s.me.canGoNorth {
		s.t.move("N", charge)
		moved = true
	} else if s.me.canGoWest {
		s.t.move("W", charge)
		moved = true
	}

	// 5) If stuck → SURFACE
	if !moved {
		s.t.surface()
		// Reset visitedTiles entirely
		s.visitedTiles = make(map[Point]bool)
	}
}

// ------------------------------------------------------------
// TORPEDO HANDLING (CORRIGÉ ET FACULTATIF)
// ------------------------------------------------------------

// Firing torpedo if enemy hinted in a sector
func (s *State) tryTorpedo() Tile {
	// Very naive logic: shoot south if possible
	x := s.me.currentPos.pos.x
	y := s.me.currentPos.pos.y

	// TORPEDO RANGE = 4
	if y+4 < HEIGHT && isWalkable(s.carte[x][y+4]) {
		return s.carte[x][y+4]
	}

	return Tile{}
}

// ------------------------------------------------------------
// ------------------------------------------------------------
// MAIN LOOP — PARTIE FINALE CORRIGÉE
// ------------------------------------------------------------

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1_000_000), 1_000_000)

	rand.Seed(time.Now().UnixNano())

	var s State

	// --------------------------------------------------------
	// READ MAP SIZE + MY ID
	// --------------------------------------------------------
	var width, height, myId int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)
	s.me.id = myId

	// --------------------------------------------------------
	// READ FULL MAP
	// --------------------------------------------------------
	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		s.board += line
	}

	// --------------------------------------------------------
	// INIT GRID (CORRIGÉ ET PROPRE)
	// --------------------------------------------------------
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {

			tileChar := s.board[y*WIDTH+x]

			s.carte[x][y] = Tile{
				pos:   Point{x, y},
				what:  string(tileChar),
				color: "red",
			}

			// Register walkable tiles
			if isWalkable(s.carte[x][y]) {
				s.carte[x][y].color = "green"
				s.walkableTiles = append(s.walkableTiles, s.carte[x][y])
			}
		}
	}

	// --------------------------------------------------------
	// CHOOSE RANDOM START POSITION (AS REQUIRED BY GAME RULES)
	// --------------------------------------------------------
	start := s.walkableTiles[rand.Intn(len(s.walkableTiles))]
	fmt.Println(start.pos.x, start.pos.y)

	// Init visited
	s.visitedTiles = make(map[Point]bool)

	// --------------------------------------------------------
	// GAME LOOP
	// --------------------------------------------------------
	for {
		var x, y, myLife, oppLife int
		var torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int

		scanner.Scan()
		fmt.Sscan(scanner.Text(),
			&x, &y,
			&myLife, &oppLife,
			&torpedoCooldown, &sonarCooldown,
			&silenceCooldown, &mineCooldown)

		// ----------------------------------------------------
		// UPDATE SELF STATE
		// ----------------------------------------------------
		s.me.currentPos = s.carte[x][y]
		s.visitedTiles[s.me.currentPos.pos] = true

		s.me.hitPoints = myLife
		s.opp.hitPoints = oppLife

		s.me.torpedoCooldown = torpedoCooldown

		// Reset torpedo info
		s.opp.torpedoPos = []Point{}

		// ----------------------------------------------------
		// READ SONAR RESULT (UNUSED IN WOOD LEAGUE)
		// ----------------------------------------------------
		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		// ----------------------------------------------------
		// PARSE OPPONENT ORDERS
		// ----------------------------------------------------
		scanner.Scan()
		opponentOrders := scanner.Text()
		s.parseOppOrders(opponentOrders)

		// ----------------------------------------------------
		// APPLY MAIN MOVEMENT LOGIC
		// ----------------------------------------------------
		s.possibleDir()

		// ----------------------------------------------------
		// SEND COMMANDS
		// ----------------------------------------------------
		s.t.sendTurn()

		// ----------------------------------------------------
		// TURN RESET
		// ----------------------------------------------------
		s.me.currentPos = Tile{}
		s.me.canGoNorth = false
		s.me.canGoSouth = false
		s.me.canGoWest = false
		s.me.canGoEast = false

		s.t.commands = []string{}
	}
}
