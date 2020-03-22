package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
type Player struct {
	id         int
	currentPos Point
	hitPoints  int
	canGoWest  bool
	canGoEast  bool
	canGoNorth bool
	canGoSouth bool
}
type Action struct {
	possibleDirections []string //= ['N','S','E','W']
}

//no opp nothing, goal is roaming
type State struct {
	board string
	me    Player
}

//first keep it stateless
func (p *Player) checkDirections(pos Point, board string, visited map[int]bool) {

	if pos.x-1 > 0 && board[pos.y*WIDTH+pos.x-1] != 'x' && !visited[pos.y*WIDTH+pos.x-1] {
		p.canGoWest = true
	}
	if pos.x+1 < WIDTH && board[pos.y*WIDTH+pos.x+1] != 'x' && !visited[pos.y*WIDTH+pos.x+1] {
		p.canGoEast = true
	}
	if pos.y-1 > 0 && board[(pos.y-1)*WIDTH+pos.x] != 'x' && !visited[(pos.y-1)*WIDTH+pos.x] {
		p.canGoNorth = true
	}
	if pos.y+1 < HEIGHT && board[(pos.y+1)*WIDTH+pos.x] != 'x' && !visited[(pos.y+1)*WIDTH+pos.x] {
		p.canGoSouth = true
	}
}

//first:a simple bot that roams through the map avoiding islands
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	me := Player{}
	var width, height, myId int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)
	//did i really care about that??
	//me.id = myId
	var carte [HEIGHT][WIDTH]Tile
	var board string

	//will put that here for now?
	visited := make(map[int]bool)

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text()
		board = board + line
	}

	//init graph keeping board
	//might use byte here?
	//BIG TEST never did that before!!
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			carte[i][j] = Tile{Point{i, j}, string(board[j*WIDTH+i]), false}
		}
	}
	//my starting pos
	var startPos = Point{7, 7}
	fmt.Println(startPos.x, startPos.y)
	log.Println(carte[8][10])
	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)
		me.currentPos = Point{x, y}
		visited[y*width+x] = true

		me.checkDirections(me.currentPos, board, visited)
		log.Println(me)

		//I know...but did i grasp the logic??
		// !!! You cannot move on a cell you already visited before
		// see surface this is not a replacement for a good floodfill or sth, but...
		if me.canGoSouth {
			fmt.Println("MOVE S TORPEDO")
		}
		if !me.canGoSouth && me.canGoEast {
			fmt.Println("MOVE E TORPEDO")
		}
		if !me.canGoSouth && !me.canGoEast && me.canGoNorth {
			fmt.Println("MOVE N TORPEDO")
		}
		if !me.canGoNorth && !me.canGoEast && !me.canGoSouth && me.canGoWest {
			fmt.Println("MOVE W TORPEDO")
		}
		if !me.canGoNorth && !me.canGoEast && !me.canGoSouth && !me.canGoWest {
			fmt.Println("SURFACE")
		}
		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		scanner.Scan()
		//opponentOrders := scanner.Text()

		//reset turn player data
		//write a reset turn eventually...
		me.currentPos = Point{}
		me.canGoNorth = false
		me.canGoSouth = false
		me.canGoWest = false
		me.canGoEast = false
		me.hitPoints = 0
	}
}