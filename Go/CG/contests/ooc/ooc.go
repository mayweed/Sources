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

//a graph might help?
type Tile struct {
	pos     Point
	what    string
	visited bool
}

//check that a tile is not off limit + walkable
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
	currentPos      Point
	hitPoints       int
	canGoWest       bool
	canGoEast       bool
	canGoNorth      bool
	canGoSouth      bool
	torpedoCooldown int
	oppSurfaceHint  string
}

func (s *State) checkDirections(pos Point, visited map[int]bool) {
	//TODO should add a check on num of edges!! if num of edges walkable <2 dont fuckin
	//go!! REFACTOR!!
	//if pos.x-1 > 0 && isWalkable(s.carte[pos.x-1][pos.y]) && s.numOfEdges(s.carte[pos.x-1][pos.y]) >=2
	//&& !visited[pos.y*WIDTH+pos.x-1] {
	//must used the bool from Tile also...
	if pos.x-1 > 0 && s.board[pos.y*WIDTH+pos.x-1] != 'x' && !visited[pos.y*WIDTH+pos.x-1] {
		s.me.canGoWest = true
	}
	if pos.x+1 < WIDTH && s.board[pos.y*WIDTH+pos.x+1] != 'x' && !visited[pos.y*WIDTH+pos.x+1] {
		s.me.canGoEast = true
	}
	if pos.y-1 > 0 && s.board[(pos.y-1)*WIDTH+pos.x] != 'x' && !visited[(pos.y-1)*WIDTH+pos.x] {
		s.me.canGoNorth = true
	}
	if pos.y+1 < HEIGHT && s.board[(pos.y+1)*WIDTH+pos.x] != 'x' && !visited[(pos.y+1)*WIDTH+pos.x] {
		s.me.canGoSouth = true
	}
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
func move(dir string, c bool) string {
	var s string
	switch dir {
	case "N":
		s = fmt.Sprintf("MOVE N")
	case "S":
		s = fmt.Sprintf("MOVE S")
	case "W":
		s = fmt.Sprintf("MOVE W")
	case "E":
		s = fmt.Sprintf("MOVE E")
	}
	if c {
		s = s + " TORPEDO"
	}
	return s
}
func surface() string {
	return fmt.Sprintf("SURFACE")
}
func torpedo(p Point) string {
	return fmt.Sprintf("TORPEDO %d %d", p.x, p.y)
}
func msg(s string) string {
	return fmt.Sprintf("MSG %s", s)
}

//if and only if commands > 1
func sendTurn(commands []string) string {
	if len(commands) == 1 {
		return commands[0]
	} else {
		return strings.Join(commands, "|")
	}
}

//no opp nothing, goal is roaming
type State struct {
	board         string
	carte         [HEIGHT][WIDTH]Tile
	walkableTiles []Tile
	me            Me
	opp           Opp
	t             Turn
}

//for ANY given walkable tile, yield its walkable valid tile!!
func (s *State) getNeighbours(t Tile) []Tile {
	var neighbours []Tile
	//be sure it's walkable we're searching from...!!
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

/*
//a bfs? taken from xmasrush one...
//need to keep track of the dist
func (s *State) bfsPath(playerTilePos Tile) []Tile {
	var visited = make(map[Tile]bool)
	visited[playerTilePos] = true

	var startTile = []Tile{playerTilePos}
	var queue = [][]Tile{startTile}

	for 0 < len(queue) {
		//pop the first element
		path := queue[0]
		queue = queue[1:]

		lastTile := path[len(path)-1]
		if lastTile == questTile {
			return path
		}

		for _, tile := range s.getNeighbours(lastTile) {
			var newPath = path
			if !visited[tile] {
				visited[tile] = true
				newPath = append(newPath, tile)
				queue = append(queue, newPath)
			}
		}
	}
	return []Tile{}
}
*/
//first:a simple bot that roams through the map avoiding islands
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

	//will put that here for now?
	visited := make(map[int]bool)

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		s.board = s.board + line
	}

	//init graph keeping board
	//BIG TEST never did that before!!
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			s.carte[i][j] = Tile{Point{i, j}, string(s.board[j*WIDTH+i]), false}
			//get a list of walkable cells to choose randomly a starting point
			if s.carte[i][j].what == "." {
				s.walkableTiles = append(s.walkableTiles, s.carte[i][j])
			}
		}
	}

	//my starting pos
	//must be in sector 5, center. Must find a walkable tile in there
	//idea : from all center tiles filter the walkable choose randomly one
	var startPos = s.walkableTiles[rand.Intn(len(s.walkableTiles))]
	fmt.Println(startPos.pos.x, startPos.pos.y)
	log.Println(startPos) //to know it...

	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)
		s.me.currentPos = Point{x, y}
		s.me.hitPoints = myLife
		s.opp.hitPoints = oppLife
		s.me.torpedoCooldown = torpedoCooldown
		visited[y*width+x] = true

		s.checkDirections(s.me.currentPos, visited)

		var c bool
		if s.me.torpedoCooldown <= 3 {
			c = true
		} else {
			c = false
		}
		//TEST, now should fire in the range!! and should chase the goose!!
		if s.me.torpedoCooldown == 0 {
			s.t.commands = append(s.t.commands, (torpedo(Point{3, 5})))
		}
		//I know...but did i grasp the logic??
		// !!! You cannot move on a cell you already visited before
		// see surface this is not a replacement for a good floodfill or sth, but...
		if s.me.canGoSouth {
			s.t.commands = append(s.t.commands, move("S", c))
		}
		if !s.me.canGoSouth && s.me.canGoEast {
			s.t.commands = append(s.t.commands, move("E", c))
		}
		if !s.me.canGoSouth && !s.me.canGoEast && s.me.canGoNorth {
			s.t.commands = append(s.t.commands, move("N", c))
		}
		if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && s.me.canGoWest {
			s.t.commands = append(s.t.commands, move("W", c))
		}
		if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && !s.me.canGoWest {
			fmt.Println("SURFACE")
			//should reset visited
			for c, _ := range visited {
				visited[c] = false
			}
		}

		//must write the  command chain!!
		//if torpedoCooldown == 0 : FIRE!!!
		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		scanner.Scan()
		opponentOrders := scanner.Text()
		s.parseOppOrders(opponentOrders)
		//should yield 0 4/0 6/1 5
		//MUST update the way i make test!!
		log.Println(s.numOfEdges(s.carte[0][5]))
		log.Println(s.opp.enemyZone)

		res := sendTurn(s.t.commands)
		fmt.Println(res)
		//reset turn player data
		//write a reset turn eventually...
		s.me.currentPos = Point{}
		s.me.canGoNorth = false
		s.me.canGoSouth = false
		s.me.canGoWest = false
		s.me.canGoEast = false
		s.me.hitPoints = 0
		s.t.commands = []string{}
	}
}
