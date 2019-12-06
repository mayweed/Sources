package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	WIDTH      = 13
	HEIGHT     = 11
	BOMB_RANGE = 3
)

//STATE:description of the state of the game at the turn T
type Point struct {
	x, y int
}

//distance in a grid idea voronoi
func (p Point) manhattanDist(p2 Point) float64 {
	var dx = p2.x - p.x
	var dy = p2.y - p.y
	return math.Abs(float64(dx)) + math.Abs(float64(dy))
}

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	position  Point
	hasMe     bool
	hasPlayer bool
	hasCrate  bool
	hasBomb   bool
	isEmpty   bool
}

type Player struct {
	position       Point
	id             int
	numOfBombsLeft int
	rangeOfBombs   int
}

type Bomb struct {
	position  Point
	ownerId   int
	countdown int
	expRange  int
}

type State struct {
	me    Player
	board [HEIGHT][WIDTH]Cell
	//should convert grid into a bitfield...but what about players?
	grid    string
	players []Player
	bombs   []Bomb
	crates  []Point
}

func (s *State) getCellFromXY(x, y int) byte {
	return s.grid[y*WIDTH+x]
}
func (s *State) getCellFromPoint(p Point) byte {
	return s.grid[p.y*WIDTH+p.x]
}
func (s *State) cratesAround(c Cell) int {
	var numCrates int
	//for any given free cell let's see how many crates are in range
	for _, crate := range s.crates {
		if c.position.x == crate.x && math.Round(math.Abs(float64(c.position.y-crate.y))) <= BOMB_RANGE ||
			c.position.y == crate.y && math.Round(math.Abs(float64(c.position.x-crate.x))) <= BOMB_RANGE {
			numCrates += 1
		}
	}
	return numCrates
}

//TURN and Action
//a turn is an action + a destination
type Turn struct {
	c Cell
	//evalScore float64
}

//should make it generic, so that it works for every player
//idea: simulate bomb explosion
func (s *State) applyTurn(t Turn) {
}

//should be in action type
func move(c Cell) string {
	s := fmt.Sprintf("MOVE %d %d", c.position.x, c.position.y)
	return s
}
func bomb(c Cell) string {
	s := fmt.Sprintf("BOMB %d %d", c.position.x, c.position.y)
	return s
}

func (s *State) think() string {
	//first select a batch of random possible move
	//evaluate them: is this enough far from any given bomb? is this close to
	// a foe (bombs will become lethal)? AND is there any crates around?

	var cells []Cell
	//so first move to a random cell
	//select 10 cells check cratesAround
	for i := 0; i < 10; i++ {
		x := rand.Intn(12)
		y := rand.Intn(10)
		if s.board[y][x].isEmpty {
			cells = append(cells, s.board[y][x])
		}
	}
	//very light eval, should take into account the range of others players bomb (and
	//mine too, watch out not be killed by my own bombs!!)
	var max int
	var cell Cell
	for _, c := range cells {
		num := s.cratesAround(c)
		//log.Println(num, c.position)
		if num > max {
			max = num
			cell = c
		}
	}
	//log.Println(cell)
	var res string
	//attempt to understand
	//OOPS you go first and bomb then!!
	//so calculate a path with most crates and bomb at a given time?
	if s.me.numOfBombsLeft > 0 {
		res = bomb(cell)
	} else {
		res = move(cell)
	}

	//wait til param1 of bombPlaced == 0
	return res
}

func main() {
	rand.Seed(time.Now().Unix())
	var s State

	for {

		//read Grid
		var width, height, myId int
		fmt.Scan(&width, &height, &myId)

		for y := 0; y < height; y++ {
			var row string
			fmt.Scan(&row)
			s.grid += row
			//here you write a parseGrid linear fashion
			for x := 0; x < width; x++ {
				//dont know why got randomly "index out of range"???
				//log.Println(x, len(row))
				if x == len(row) {
					break
				}
				if row[x] == '.' {
					s.board[y][x].position = Point{x, y}
					s.board[y][x].isEmpty = true
				} else {
					s.board[y][x].position = Point{x, y}
					s.board[y][x].hasCrate = true
					//just need their positions on the grid, no more no less
					s.crates = append(s.crates, Point{x, y})
				}
			}
		}

		log.Println(s.getCellFromXY(5, 7))
		//read Entities
		var entities int
		fmt.Scan(&entities)

		for i := 0; i < entities; i++ {
			var entityType, owner, x, y, param1, param2 int
			fmt.Scan(&entityType, &owner, &x, &y, &param1, &param2)

			if owner == myId {
				s.me.position = Point{x, y}
				s.me.id = owner
				s.me.numOfBombsLeft = param1
				s.me.rangeOfBombs = param2
				s.board[y][x].hasMe = true
				s.board[y][x].isEmpty = false
			} else {
				switch entityType {
				case 0:
					s.board[y][x].hasPlayer = true
					s.board[y][x].isEmpty = false
					s.players = append(s.players, Player{position: Point{x, y}, id: owner, numOfBombsLeft: param1, rangeOfBombs: param2})
				case 1:
					s.board[y][x].hasBomb = true
					s.board[y][x].isEmpty = false
					s.bombs = append(s.bombs, Bomb{position: Point{x, y}, ownerId: owner, countdown: param1, expRange: param2})
				}
			}

		}

		res := s.think()
		fmt.Println(res)
		//MOVE test
		//should list possible moves+simulate where to leave bombs to get
		//more boxes destroy
		//fmt.Println("MOVE 10 10") // Write action to stdout

		//LOGS
		//log.Println(s.me.numOfBombsLeft)
		//s.grid = "" //reset state
	}
}
