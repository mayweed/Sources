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

type Point struct {
	x, y int
}

//idea where is the last torpedo pos located??
//pff cant return in if directly...(no return etc..)
//TOO SIMPLE should think rewrite!!
//voronoi to get possible zones??
func getDirFromPoint(myPos Point, torpedoPos Point) string {
	var s string
	if myPos.x > torpedoPos.x && myPos.y == torpedoPos.y {
		s = "W"
	}
	if myPos.x < torpedoPos.x && myPos.y == torpedoPos.y {
		s = "E"
	}
	if myPos.x == torpedoPos.x && myPos.y > torpedoPos.y {
		s = "N"
	}
	if myPos.x == torpedoPos.x && myPos.y < torpedoPos.y {
		s = "S"
	}
	return s
}

//a graph might help?
type Tile struct {
	pos  Point
	what string
	//floodfill
	color string
}

func isWalkable(t Tile) bool {
	return t.what == "."
}

//should i add ,... in arg list??
func (s *State) numOfEdges(t Tile) int {
	//could be useful
	return len(s.getNeighbours(t))
}

type Me struct {
	id              int
	currentPos      Tile
	hitPoints       int
	canGoWest       bool
	canGoEast       bool
	canGoNorth      bool
	canGoSouth      bool
	torpedoCooldown int
	oppSurfaceHint  string
}

type Opp struct {
	hitPoints    int
	oppDirection string
	torpedoPos   []Point
	enemyZone    int
}

func (s *State) parseOppOrders(orders string) {
	//sanitize orders
	st := strings.ReplaceAll(orders, "|", " ")
	//split it
	ord := strings.Split(st, " ")
	for idx, w := range ord {
		if w == "N" || w == "E" || w == "W" || w == "S" {
			s.opp.oppDirection = w
		}
		//in move n torpedo my opponent cant see i m recharging my torpedo so...
		if w == "TORPEDO" {
			x, _ := strconv.Atoi(ord[idx+1])
			y, _ := strconv.Atoi(ord[idx+2])
			s.opp.torpedoPos = append(s.opp.torpedoPos, Point{x, y})
		}
		//important hint here, must find a way to parse it sector by sector
		//the idea: if the enemy is in the SW zone and I head south,
		//next i will go S and w and s etc..
		//must find sth better
		if w == "SURFACE" {
			e, _ := strconv.Atoi(ord[idx+1])
			s.opp.enemyZone = e
			switch s.opp.enemyZone {
			case 1:
				s.me.oppSurfaceHint = "NW"
			case 2:
				s.me.oppSurfaceHint = "N"
			case 3:
				s.me.oppSurfaceHint = "NE"
			case 4:
				s.me.oppSurfaceHint = "CW" //C like center
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

//question is: how to triangulate opp pos from his torpedoes?? BFS??
//WRITE A GUESS FUNC TO GUESS ENEMY POS FROM HINTS!!

type Turn struct {
	commands []string
}

//Helpers for turn
//false = no charge, true = charge. if multiple arm could be string?
func (t *Turn) move(dir string, c bool) {
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
	if c {
		command = command + " TORPEDO"
	}
	t.commands = append(t.commands, command)
}
func (t *Turn) surface() {
	t.commands = append(t.commands, "SURFACE")
}
func (t *Turn) torpedo(tile Tile) {
	command := fmt.Sprintf("TORPEDO %d %d", tile.pos.x, tile.pos.y)
	t.commands = append(t.commands, command)
}
func (t *Turn) msg(s string) {
	command := fmt.Sprintf("MSG %s", s)
	t.commands = append(t.commands, command)
}

//if and only if commands > 1
func (t *Turn) sendTurn() {
	if len(t.commands) == 1 {
		fmt.Print(t.commands[0])
	} else {
		fmt.Print(strings.Join(t.commands, "|"))
	}
	fmt.Println()
}

//no opp nothing, goal is roaming
type State struct {
	board         string
	carte         [HEIGHT][WIDTH]Tile
	walkableTiles []Tile
	visitedTiles  map[Point]bool
	me            Me
	opp           Opp
	t             Turn
}

//simple helper
func (s *State) hasBeenVisited(t Tile) bool {
	return s.visitedTiles[t.pos]
}

//YannTt'as 3 mouvements possible, tu floodfill pour chaque, garde celui qui te laisse le plus de cases dispo après move
func (s *State) checkDirections(t Tile) {
	if t.pos.x-1 > 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && !s.hasBeenVisited(s.carte[t.pos.x-1][t.pos.y]) {
		s.me.canGoWest = true
	}
	if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && !s.hasBeenVisited(s.carte[t.pos.x+1][t.pos.y]) {
		s.me.canGoEast = true
	}
	if t.pos.y-1 > 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && !s.hasBeenVisited(s.carte[t.pos.x][t.pos.y-1]) {
		s.me.canGoNorth = true
	}
	if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && !s.hasBeenVisited(s.carte[t.pos.x][t.pos.y+1]) {
		s.me.canGoSouth = true
	}
}

//for ANY given walkable tile, yield its walkable valid tile!!
func (s *State) getNeighbours(t Tile) []Tile {
	var neighbours []Tile
	//NOTE should update the way i handle error/exception (same for test!!)
	// here's what i need https://blog.golang.org/error-handling-and-go
	if isWalkable(t) {
		//check north
		if t.pos.y-1 > 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) {
			neighbours = append(neighbours, s.carte[t.pos.x][t.pos.y-1])
		}
		//check south
		if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) {
			neighbours = append(neighbours, s.carte[t.pos.x][t.pos.y+1])
		}
		//check west
		if t.pos.x-1 > 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) {
			neighbours = append(neighbours, s.carte[t.pos.x-1][t.pos.y])
		}
		//check east
		if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) {
			neighbours = append(neighbours, s.carte[t.pos.x+1][t.pos.y])
		}
	}
	return neighbours
}

//a bfs? taken from xmasrush one...
//need to keep track of the dist
//Idea: all path 4 cells from torpedoPos in the direction of opp
func (s *State) getBfsPath(playerTilePos Tile) {
	var visited = make(map[Tile]bool)
	visited[playerTilePos] = true

	var startTile = playerTilePos
	var queue = []Tile{startTile}

	//keep track of the preceding tile
	var parent = make(map[Tile]Tile)

	for 0 < len(queue) {
		//pop the first element/shouldnt i use container/list here?
		startNode := queue[0]
		queue = queue[1:]

		for _, adjTile := range s.getNeighbours(startNode) {
			if !visited[adjTile] {
				visited[adjTile] = true
				parent[adjTile] = startNode
				queue = append(queue, adjTile)
			}
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

	//init graph keeping board
	//BIG TEST never did that before!!
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			//red by default
			s.carte[i][j] = Tile{Point{i, j}, string(s.board[j*WIDTH+i]), "red"}
			//get a list of walkable cells to choose randomly a starting point
			if s.carte[i][j].what == "." {
				//green like "you could go there"
				s.carte[i][j].color = "green"
				s.walkableTiles = append(s.walkableTiles, s.carte[i][j])
			}
		}
	}

	//my starting pos
	var startPos = s.walkableTiles[rand.Intn(len(s.walkableTiles))]
	fmt.Println(startPos.pos.x, startPos.pos.y)
	//fmt.Println("7 8")
	//log.Println(startPos) //to know it...

	s.visitedTiles = make(map[Point]bool)

	var turn int
	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)

		s.me.currentPos = s.carte[x][y]
		s.visitedTiles[s.me.currentPos.pos] = true
		s.me.hitPoints = myLife
		s.opp.hitPoints = oppLife
		s.me.torpedoCooldown = torpedoCooldown

		var c bool
		if s.me.torpedoCooldown <= 3 {
			c = true
		} else {
			c = false
		}
		/*
			//TEST, now should fire in the range!! and should chase the goose!!
			if s.me.torpedoCooldown == 0 {
				//this is shitty com'on!!
				s.t.commands = append(s.t.commands, (torpedo(s.carte[s.me.currentPos.pos.x+2][s.me.currentPos.pos.y])))
			}
			//TEST
			var dir string
			if len(s.opp.torpedoPos) != 0 {
				dir = getDirFromPoint(s.me.currentPos.pos, s.opp.torpedoPos[0])
			}
		*/
		//I know...but did i grasp the logic??
		s.checkDirections(s.me.currentPos)
		if s.me.canGoSouth {
			s.t.move("S", c)
		}
		if !s.me.canGoSouth && s.me.canGoEast {
			s.t.move("E", c)
		}
		if !s.me.canGoSouth && !s.me.canGoEast && s.me.canGoNorth {
			s.t.move("N", c)
		}
		if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && s.me.canGoWest {
			s.t.move("W", c)
		}
		//one direction is possible but cell has already been visited!! so surface
		if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && !s.me.canGoWest {
			s.t.surface()
			//should reset visited
			for c, _ := range s.visitedTiles {
				s.visitedTiles[c] = false
			}
		}

		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		scanner.Scan()
		opponentOrders := scanner.Text()
		s.parseOppOrders(opponentOrders)

		//TEST
		//s.floodfill(startPos)
		//log.Println(s.carte[startPos.pos.x-1][startPos.pos.y].color)
		log.Println(s.t.commands)
		log.Println("N: ", s.me.canGoNorth, "S: ", s.me.canGoSouth, "W: ", s.me.canGoWest, "E: ", s.me.canGoEast)
		//log.Println(s.opp.torpedoPos)

		//TEST
		s.getBfsPath(s.me.currentPos)

		//fmt.Println(s.sendTurn())
		//s.sendTurn()
		fmt.Println(s.t.commands[0])
		//reset turn player data
		//write a reset turn eventually...
		s.me.currentPos = Tile{}
		s.me.canGoNorth = false
		s.me.canGoSouth = false
		s.me.canGoWest = false
		s.me.canGoEast = false
		s.t.commands = []string{}
		turn += 1
	}
}
