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
	id         int
	currentPos Point
	hitPoints  int
	canGoWest  bool
	canGoEast  bool
	canGoNorth bool
	canGoSouth bool
}

//first keep it stateless
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
		log.Println(idx, w)
		//in move n torpedo my opponent cant see i m recharging my torpedo so...
		if w == "TORPEDO" {
			x, _ := strconv.Atoi(ord[idx+1])
			y, _ := strconv.Atoi(ord[idx+2])
			o.torpedoPos = append(o.torpedoPos, Point{x, y})
		}
	}
}

//question is: how to triangulate opp pos from his torpedoes?? BFS??

type Action struct {
	possibleDirections []string //= ['N','S','E','W']
}

//no opp nothing, goal is roaming
type State struct {
	board string
	carte [HEIGHT][WIDTH]Tile
	me    Me
	opp   Opp
}

//first:a simple bot that roams through the map avoiding islands
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	//me := Me{}
	//opp := Opp{}
	var s State
	var width, height, myId int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)
	//did i really care about that??
	s.me.id = myId
	//var carte [HEIGHT][WIDTH]Tile
	//var board string

	//will put that here for now?
	visited := make(map[int]bool)

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		s.board = s.board + line
	}

	//init graph keeping board
	//might use byte here?
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
		visited[y*width+x] = true

		s.me.checkDirections(s.me.currentPos, s.board, visited)
		//log.Println(me)

		//I know...but did i grasp the logic??
		// !!! You cannot move on a cell you already visited before
		// see surface this is not a replacement for a good floodfill or sth, but...
		if s.me.canGoSouth {
			fmt.Println("MOVE S TORPEDO")
		}
		if !s.me.canGoSouth && s.me.canGoEast {
			fmt.Println("MOVE E TORPEDO")
		}
		if !s.me.canGoSouth && !s.me.canGoEast && s.me.canGoNorth {
			fmt.Println("MOVE N TORPEDO")
		}
		if !s.me.canGoNorth && !s.me.canGoEast && !s.me.canGoSouth && s.me.canGoWest {
			fmt.Println("MOVE W TORPEDO")
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
		log.Println(s.opp.oppDirection, s.opp.torpedoPos)

		//log.Println(torpedoCooldown, sonarResult)

		//reset turn player data
		//write a reset turn eventually...
		s.me.currentPos = Point{}
		s.me.canGoNorth = false
		s.me.canGoSouth = false
		s.me.canGoWest = false
		s.me.canGoEast = false
		s.me.hitPoints = 0
	}
}
