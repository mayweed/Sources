package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	HEIGHT = 15
	WIDTH  = 15
)

type Point struct {
	x, y int
}

func (p *Point) isValid() bool {
	return p.x > 0 && p.x < WIDTH && p.y > 0 && p.y < HEIGHT
}

//a graph might help?
type Tile struct {
	pos     Point
	what    string
	visited bool
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
}

func (m *Me) checkDirections(pos Point, board string, visited map[int]bool) {

	if pos.x-1 > 0 && board[pos.y*WIDTH+pos.x-1] != 'x' && !visited[pos.y*WIDTH+pos.x-1] {
		m.canGoWest = true
	}
	if pos.x+1 < WIDTH && board[pos.y*WIDTH+pos.x+1] != 'x' && !visited[pos.y*WIDTH+pos.x+1] {
		m.canGoEast = true
	}
	if pos.y-1 > 0 && board[(pos.y-1)*WIDTH+pos.x] != 'x' && !visited[(pos.y-1)*WIDTH+pos.x] {
		m.canGoNorth = true
	}
	if pos.y+1 < HEIGHT && board[(pos.y+1)*WIDTH+pos.x] != 'x' && !visited[(pos.y+1)*WIDTH+pos.x] {
		m.canGoSouth = true
	}
}

type Opp struct {
	hitPoints    int
	oppDirection string
	torpedoPos   []Point
}

func (o *Opp) parseOppOrders(orders string) {
	//sanitize orders
	s := strings.ReplaceAll(orders, "|", " ")
	//split it
	ord := strings.Split(s, " ")
	for idx, w := range ord {
		if w == "N" || w == "E" || w == "W" || w == "S" {
			o.oppDirection = w
		}
		//in move n torpedo my opponent cant see i m recharging my torpedo so...
		if w == "TORPEDO" {
			x, _ := strconv.Atoi(ord[idx+1])
			y, _ := strconv.Atoi(ord[idx+2])
			o.torpedoPos = append(o.torpedoPos, Point{x, y})
		}
	}
}

//question is: how to triangulate opp pos from his torpedoes?? BFS??

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
	board string
	carte [HEIGHT][WIDTH]Tile
	me    Me
	opp   Opp
	t     Turn
}

//first:a simple bot that roams through the map avoiding islands
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

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
		}
	}

	//my starting pos
	var startPos = Point{7, 7}
	fmt.Println(startPos.x, startPos.y)
	//log.Println(carte[8][10])

	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)
		s.me.currentPos = Point{x, y}
		s.me.hitPoints = myLife
		s.opp.hitPoints = oppLife
		s.me.torpedoCooldown = torpedoCooldown
		visited[y*width+x] = true

		s.me.checkDirections(s.me.currentPos, s.board, visited)

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
		s.opp.parseOppOrders(opponentOrders)
		log.Println(s.me.torpedoCooldown)

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
