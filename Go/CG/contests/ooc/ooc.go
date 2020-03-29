package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	HEIGHT = 15
	WIDTH  = 15
)

//POINT
//think about type Point in import "image"
type Point struct {
	x, y int
}

//idea where is the last torpedo pos located??
func getPosBySector(p Point) int {
	var zone int
	if p.x >= 0 && p.x <= 4 && p.y > 0 && p.y <= 4 {
		zone = 1
	}
	if p.x >= 0 && p.x <= 4 && p.y > 4 && p.y <= 9 {
		zone = 4
	}
	if p.x >= 0 && p.x <= 4 && p.y > 9 && p.y <= 14 {
		zone = 7
	}
	if p.x > 4 && p.x <= 9 && p.y > 0 && p.y <= 4 {
		zone = 2
	}
	if p.x > 4 && p.x <= 9 && p.y > 4 && p.y <= 9 {
		zone = 5
	}
	if p.x > 4 && p.x <= 9 && p.y > 9 && p.y <= 14 {
		zone = 8
	}
	if p.x > 9 && p.x <= 14 && p.y > 0 && p.y <= 4 {
		zone = 3
	}
	if p.x > 9 && p.x <= 14 && p.y > 4 && p.y <= 9 {
		zone = 6
	}
	if p.x > 9 && p.x <= 14 && p.y > 9 && p.y <= 14 {
		zone = 9
	}
	return zone
}

//TILE
//getTile() with limit checking
//to see what happen if border crossed?
//https://rosettacode.org/wiki/Bitmap/Flood_fill#Go ==> would like to see the getpx func i imagine a getTile()
// with limit checkers in a Grid Type...
type Tile struct {
	pos   Point
	what  string
	color string
}

func isWalkable(t Tile) bool {
	return t.what == "."
}

//OPP
type Opp struct {
	hitPoints       int
	oppDirection    string
	torpedoPos      []Point
	lastTorpedoZone int
	surfaceZone     int
	sonarZone       int
	dirs            string
}

func (o *Opp) parseOppOrders(orders string) {
	//sanitize orders
	st := strings.ReplaceAll(orders, "|", " ")
	//split it
	ord := strings.Split(st, " ")
	for idx, w := range ord {
		if w == "N" || w == "E" || w == "W" || w == "S" {
			o.oppDirection = w
			//a string of ALL dir followed by my opp
			//must include the silent one!!
			o.dirs = o.dirs + w
		}
		//in move n torpedo i cant see torpedo i only see it when he fires it
		if w == "TORPEDO" {
			x, _ := strconv.Atoi(ord[idx+1])
			y, _ := strconv.Atoi(ord[idx+2])
			o.torpedoPos = append(o.torpedoPos, Point{x, y})
		}
		if w == "SURFACE" {
			e, _ := strconv.Atoi(ord[idx+1])
			o.surfaceZone = e
		}
		if w == "SONAR" {
			e, _ := strconv.Atoi(ord[idx+1])
			o.sonarZone = e
		}
		if w == "SILENCE" {
			//to know the gaps
			o.dirs = o.dirs + "X"
		}
	}
}
func (o *Opp) getLastTorpZone() (int, Point) {
	var zone int
	var lastTPos Point
	if len(o.torpedoPos) > 0 {
		if len(o.torpedoPos) == 1 {
			zone = getPosBySector(o.torpedoPos[0])
			lastTPos = o.torpedoPos[0]
		} else {
			zone = getPosBySector(o.torpedoPos[len(o.torpedoPos)-1])
			lastTPos = o.torpedoPos[len(o.torpedoPos)-1]
		}
	}
	o.lastTorpedoZone = zone
	return zone, lastTPos
}

//ME
type Me struct {
	id           int
	currentPos   Tile
	hitPoints    int
	visitedTiles map[Tile]bool
	//dirs
	canGoWest  bool
	canGoEast  bool
	canGoNorth bool
	canGoSouth bool
	//Torpedo
	torpedoCooldown int
	canFireTorpedo  bool
	//Sonar
	sonarCooldown int
	canUseSonar   bool
	sonarResult   string
	//Silence
	silenceCooldown int
	canUseSilence   bool
	//Opp
	oppSurfaceHint string
	//comm to send
	commands []string
}

func (m *Me) isTorpCharge() bool {
	var c bool
	if m.torpedoCooldown > 0 {
		c = true
		m.canFireTorpedo = false
	} else {
		c = false
		m.canFireTorpedo = true
	}
	return c
}
func (m *Me) isSonarCharge() bool {
	var c bool
	if m.sonarCooldown > 0 {
		c = true
	} else {
		c = false
		m.canUseSonar = true
	}
	return c
}
func (m *Me) isSilenceCharge() bool {
	var c bool
	if m.silenceCooldown > 0 {
		c = true
	} else {
		c = false
		m.canUseSilence = true
	}
	return c
}
func (m *Me) move(dir string) {
	var command string
	switch dir {
	case "N":
		command = fmt.Sprintf("MOVE %s", dir)
	case "S":
		command = fmt.Sprintf("MOVE %s", dir)
	case "W":
		command = fmt.Sprintf("MOVE %s", dir)
	case "E":
		command = fmt.Sprintf("MOVE %s", dir)
	}
	//having torp charge is important but sonar?
	//no dynamic, be charged that's all
	if m.isTorpCharge() {
		command = command + " TORPEDO"
	}
	if !m.isTorpCharge() && m.isSonarCharge() {
		command = command + " SONAR"
	}
	if !m.isTorpCharge() && !m.isSonarCharge() && m.isSilenceCharge() {
		command = command + " SILENCE"
	}
	m.commands = append(m.commands, command)
}
func (m *Me) surface() {
	m.commands = append(m.commands, "SURFACE")
}
func (m *Me) sonar(sector int) {
	command := fmt.Sprintf("SONAR %d", sector)
	m.commands = append(m.commands, command)
}
func (m *Me) torpedo(tile Tile) {
	command := fmt.Sprintf("TORPEDO %d %d", tile.pos.x, tile.pos.y)
	m.commands = append(m.commands, command)
}
func (m *Me) silence(direction string, distance int) {
	command := fmt.Sprintf("SILENCE %s %d", direction, distance)
	m.commands = append(m.commands, command)
}
func (m *Me) msg(s string) {
	command := fmt.Sprintf("MSG %s", s)
	m.commands = append(m.commands, command)
}
func (m *Me) sendTurn() {
	if len(m.commands) == 1 {
		fmt.Print(m.commands[0])
	} else {
		fmt.Print(strings.Join(m.commands, "|"))
	}
	fmt.Println()
}

//STATE
type State struct {
	board         string
	carte         [HEIGHT][WIDTH]Tile
	walkableTiles []Tile
	visitedTiles  map[Point]bool
	me            Me
	opp           Opp
	targets       []Tile
}

//YannTt'as 3 mouvements possible, tu floodfill pour chaque, garde celui qui te laisse le plus de cases dispo après move
func (s *State) checkDirections(t Tile) {
	if t.pos.x-1 > 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && !s.me.visitedTiles[s.carte[t.pos.x-1][t.pos.y]] {
		s.me.canGoWest = true
	}
	if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && !s.me.visitedTiles[s.carte[t.pos.x+1][t.pos.y]] {
		s.me.canGoEast = true
	}
	if t.pos.y-1 > 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && !s.me.visitedTiles[s.carte[t.pos.x][t.pos.y-1]] {
		s.me.canGoNorth = true
	}
	if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && !s.me.visitedTiles[s.carte[t.pos.x][t.pos.y+1]] {
		s.me.canGoSouth = true
	}
}
func (s *State) floodfill(t Tile) {
	var queue []Tile
	queue = append(queue, t)

	log.Println("WNUM", len(s.walkableTiles))
	var numT int

	for len(queue) != 0 {
		var t = queue[0]
		queue = queue[1:]

		//t.color = "blue"

		//TORPEDO RANGE if the tile is blue can be used to fire torpedo
		//NOT the way to do that
		//if path := s.getBfsPath(s.me.currentPos, t); len(path) > 4 {
		//	return
		//}
		//check north
		//if it's not yet blue and is walkable, we havent visited it yet!!
		if t.pos.y-1 >= 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && s.carte[t.pos.x][t.pos.y-1].color != "blue" {
			s.carte[t.pos.x][t.pos.y-1].color = "blue" //wouldn't black be better?
			queue = append(queue, s.carte[t.pos.x][t.pos.y-1])
			log.Println(s.carte[t.pos.x][t.pos.y-1].color)
		}
		//check south
		if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && s.carte[t.pos.x][t.pos.y+1].color != "blue" {
			s.carte[t.pos.x][t.pos.y+1].color = "blue"
			queue = append(queue, s.carte[t.pos.x][t.pos.y+1])
		}
		//check west
		if t.pos.x-1 >= 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && s.carte[t.pos.x-1][t.pos.y].color != "blue" {
			s.carte[t.pos.x-1][t.pos.y].color = "blue"
			queue = append(queue, s.carte[t.pos.x-1][t.pos.y])

		}
		//check east
		if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && s.carte[t.pos.x+1][t.pos.y].color != "blue" {
			s.carte[t.pos.x+1][t.pos.y].color = "blue"
			queue = append(queue, s.carte[t.pos.x+1][t.pos.y])

		}
		numT += 1
		//log.Println(t.pos, t.color)
	}
	log.Println("FIN", numT)
}

//voronoi to get possible zones??
//need to keep track of the dist
//Idea: all path 4 cells from torpedoPos in the direction of opp
func (s *State) getBfsPath(startPos, target Tile) []Tile {
	var visited = make(map[Tile]bool)
	visited[startPos] = true

	var startTile = startPos
	var queue = []Tile{startTile}

	//keep track of the preceding tile
	var parent = make(map[Tile]Tile)

	var path []Tile
	for 0 < len(queue) {
		//pop the first element/shouldnt i use container/list here?
		//t for startNode
		t := queue[0]
		queue = queue[1:]

		if t == target {
			//p like predecessors
			var p = target
			path = append(path, target)
			for parent[p] != startPos {
				path = append(path, parent[p])
				p = parent[p]
			}
			path = append(path, startPos)
			return path
		}

		//check north
		if t.pos.y-1 >= 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && !visited[s.carte[t.pos.x][t.pos.y-1]] {
			visited[s.carte[t.pos.x][t.pos.y-1]] = true
			parent[s.carte[t.pos.x][t.pos.y-1]] = t
			queue = append(queue, s.carte[t.pos.x][t.pos.y-1])
		}
		//check south
		if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && !visited[s.carte[t.pos.x][t.pos.y+1]] {
			visited[s.carte[t.pos.x][t.pos.y+1]] = true
			parent[s.carte[t.pos.x][t.pos.y+1]] = t
			queue = append(queue, s.carte[t.pos.x][t.pos.y+1])
		}
		//check west
		if t.pos.x-1 >= 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && !visited[s.carte[t.pos.x-1][t.pos.y]] {
			visited[s.carte[t.pos.x-1][t.pos.y]] = true
			parent[s.carte[t.pos.x-1][t.pos.y]] = t
			queue = append(queue, s.carte[t.pos.x-1][t.pos.y])

		}
		//check east
		if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && !visited[s.carte[t.pos.x+1][t.pos.y]] {
			visited[s.carte[t.pos.x+1][t.pos.y]] = true
			parent[s.carte[t.pos.x+1][t.pos.y]] = t
			queue = append(queue, s.carte[t.pos.x+1][t.pos.y])

		}
	}
	return path
}

//the distance from startPoint to all walkable Tiles!!
//!! You can't change values associated with keys in a map, you can only reassign values.
//!! When you "fill" the map, you can't use the loop's variable, as it gets overwritten in each iteration
// see : https://stackoverflow.com/questions/42716852/how-to-update-map-values-in-go
func (s *State) calculateDist(src Tile) map[Tile]*int {
	var dist = make(map[Tile]*int)
	var path []Tile
	//should change that walkable to take into account floodfill?
	for _, target := range s.walkableTiles {
		//dont need this one
		if target == src {
			continue
		}
		path = s.getBfsPath(src, target)
		length := len(path)
		dist[target] = &length
	}
	return dist
}

func (s *State) getTargets(dist map[Tile]*int) {
	//let's find targets find the process costly (recalc the dist to all tiles etc...
	var max = 4
	//to not damage myself
	//var min = 2
	var targetTile Tile
	for k, v := range dist {
		if *v == max { //&& *v > min {
			targetTile = k
			s.targets = append(s.targets, targetTile)
		}
	}
}
func getBestMove(s State) {
	//it takes state as an arg, clone it to sim?
	//determine which direction is best wrt the num of wtiles left and the opp
	//location (indeed where we think he might be thx to the info from surface, sonar
	//or torpedo)
	//should be there in the end
	//s.checkDirections(m.currentPos)
	//determine what commands we must send so that everything gets charged + choose
	//between torp/sonar/silence?
}

//question is: how to triangulate opp pos from his torpedoes?? BFS??
//WRITE A GUESS FUNC TO GUESS ENEMY POS FROM HINTS!!
//I know...but did i grasp the logic??
func (s *State) woodMoves() {
	if s.me.canGoSouth {
		s.me.move("S")
	}
	if !s.me.canGoSouth && s.me.canGoEast {
		s.me.move("E")
	}
	if !s.me.canGoSouth && !s.me.canGoEast && s.me.canGoNorth {
		s.me.move("N")
	}
	if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && s.me.canGoWest {
		s.me.move("W")
	}
	//TEST
	//if i am round the torp zone sonar to see what happens
	if getPosBySector(s.me.currentPos.pos) == s.opp.lastTorpedoZone {
		//s.me.sonar(s.opp.lastTorpedoZone) //useless should be able to use sonar in the neighhouring zones
		_, p := s.opp.getLastTorpZone()
		s.me.torpedo(s.carte[p.x][p.y])
	}
	/*
		//Torpedo spam to TEST
		//pb with canFireTorpedo: not enougght charge??
		if !s.me.isTorpCharge() {
			t := s.targets[rand.Intn(len(s.targets))]
			s.me.torpedo(t)
		}
	*/

	//one direction is possible but cell has already been visited!! so surface
	//must neutralize torpedo here
	if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && !s.me.canGoWest {
		s.me.surface()
		//should reset visited
		for c, _ := range s.me.visitedTiles {
			s.me.visitedTiles[c] = false
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	//random num generator
	rand.Seed(time.Now().Unix())

	var s State
	var width, height, myId int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)
	//did i really care about that??
	s.me.id = myId

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		s.board = s.board + line
	}

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			//red by default
			s.carte[i][j] = Tile{Point{i, j}, string(s.board[j*WIDTH+i]), "red"}
			if s.carte[i][j].what == "." {
				//green like "you could go there"
				s.carte[i][j].color = "green"
				s.walkableTiles = append(s.walkableTiles, s.carte[i][j])
			}
		}
	}

	//my starting pos
	var startPos = s.walkableTiles[rand.Intn(len(s.walkableTiles))]
	//	var startPos = s.carte[7][8]
	fmt.Println(startPos.pos.x, startPos.pos.y)

	//Dont use it yet so...
	//s.floodfill(startPos)
	s.me.visitedTiles = make(map[Tile]bool)

	var turn int
	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)
		//should write a New func for State no?
		s.me.currentPos = s.carte[x][y]
		s.me.visitedTiles[s.me.currentPos] = true
		s.me.hitPoints = myLife
		s.me.torpedoCooldown = torpedoCooldown
		s.me.sonarCooldown = sonarCooldown
		s.me.silenceCooldown = silenceCooldown

		s.opp.hitPoints = oppLife

		//TEST TARGET
		dist := s.calculateDist(s.me.currentPos)
		s.getTargets(dist)
		log.Println(s.targets, len(s.targets))

		s.checkDirections(s.me.currentPos)
		s.woodMoves()

		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)
		s.me.sonarResult = sonarResult

		scanner.Scan()
		opponentOrders := scanner.Text()
		s.opp.parseOppOrders(opponentOrders)

		s.me.sendTurn()

		//reset turn player data
		//write a reset turn eventually...
		s.me.currentPos = Tile{}
		s.me.canGoNorth = false
		s.me.canGoSouth = false
		s.me.canGoWest = false
		s.me.canGoEast = false
		s.me.commands = []string{}
		s.targets = []Tile{}
		turn += 1
	}
}
