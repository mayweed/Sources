package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	MAP_WIDTH  = 7
	MAP_HEIGHT = 7
)

type Point struct {
	x, y int
}

//P start Point, p2 where he wants to go
func (p Point) printDirection(p2 Point) string {
	var dir string
	//return directly output "missing return at the end of func"...
	if p2.x > p.x {
		dir = "RIGHT"
	}
	if p2.x < p.x {
		dir = "LEFT"
	}
	if p2.y < p.y {
		dir = "UP"
	}
	if p2.y > p.y {
		dir = "DOWN"
	}
	return dir
}
func isValidPos(p Point) bool {
	return p.x >= 0 && p.y >= 0 &&
		p.x < MAP_WIDTH && p.y < MAP_HEIGHT
}

type Grid [7][7]Tile

type Tile struct {
	direction    string
	position     Point
	hasItem      bool
	itemName     string
	itemPlayerId int
}

func (g Grid) getTile(x int, y int) Tile {
	var t Tile
	if isValidPos(Point{x, y}) {
		t = g[x][y]
	}
	return t
}
func (g *Grid) setTile(x int, y int, t Tile) {
	g[x][y] = t
}

//wip...return the playerTile
func (g *Grid) pushRight(pushedTile Tile, row int) Tile {
	var maxCol = MAP_WIDTH - 1
	var poppedTile = g.getTile(row, maxCol)

	for i := maxCol; i > 0; i-- {
		g.setTile(i, row, g.getTile(i-1, row))
		//getTile(i, row).move(Constants.Direction.RIGHT);
	}
	//setTile(0, row, pushedTile);
	//getTile(0, row).move(new Vector2(0, row));
	return poppedTile
}

/*
Grid Example/Debug purpose
1110 0101 1001 1010 1101 0101 0111
1101 1010 1111 1010 1001 1010 1011
0110 1001 1001 0011 0111 0111 1001
0110 0110 1101 1011 1101 1110 1111
0111 1101 1101 1100 0110 0110 0111
1110 1010 0110 1010 1111 1010 1001
1001 0101 0111 1010 0110 0101 1001
*/

//shouldn't i use a buffer instead of sprintf??
func (g *Grid) printGrid() {
	var row string
	//http://xahlee.info/golang/golang_rune.html
	//was of great help to print rune
	var printDir = make(map[string]rune)
	printDir["0101"] = '═'
	printDir["1010"] = '║'
	printDir["0110"] = '╔'
	printDir["0011"] = '╗'
	printDir["1100"] = '╚'
	printDir["1001"] = '╝'
	printDir["1110"] = '╠'
	printDir["1011"] = '╣'
	printDir["0111"] = '╦'
	printDir["1101"] = '╩'
	printDir["1111"] = '╬'
	for y := 0; y < MAP_HEIGHT; y++ {
		for x := 0; x < MAP_WIDTH; x++ {
			//%c → the character represented by the corresponding Unicode code point
			row += fmt.Sprintf("%c ", printDir[g[y][x].direction])
		}
		row += fmt.Sprintf("\n")
	}
	log.Println(row)
}

/*
//To simulate you take the grid
//Parse the action,
// apply it
// give back the future grid
//better : you pass a turn!!
func (s State) simulateGrid(t Turn){
	g:=s.grid
	if t.turntype == "PUSH"{
		switch a.pushDirection{
			case "RIGHT": pushRight(grid)
			//if it's right/left you offset x

			//up/down you offset y
			...
		}
	}
}
*/
type Player struct {
	totalOfQuests int
	//At most one is revealed
	questItemName string
	//the item tile for a given player
	itemTile Tile
	//the tile in hand
	playerTile Tile
	position   Point
}

type Turn struct {
	turnType   int
	directions []string
	//action Action
}

type State struct {
	players   [2]Player
	grid      Grid
	numItems  int
	itemTiles []Tile
	turn      Turn
}

//golang yields the ascii code not the num...it's a quickfix...
//see first answer here: https://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
//UGLY CODE, will be non understandable in 6 weeks from now
func (s State) getNeighbours(t Tile) []Tile {
	var neighbours []Tile
	//check borders!!
	//to be qualified as neighbour one must be able to communicate
	//if UP is 1 on my tile, the upper tile is a neighbour if and only if down is
	//open!!
	if t.position.y-1 > 0 && t.position.y-1 < 7 {
		//if isValidPos(t.position.y - 1) { ==> look at that later, that func needs an
		//overhaul
		if t.direction[0] == 49 && s.grid[t.position.y-1][t.position.x].direction[2] == 49 {
			neighbours = append(neighbours, s.grid[t.position.y-1][t.position.x])
		}
	}
	//if RIGHT (direction 1) is ok, check left (dir 3) on the neighbouring cell
	if t.position.x+1 > 0 && t.position.x+1 < 7 {
		if t.direction[1] == 49 && s.grid[t.position.y][t.position.x+1].direction[3] == 49 {
			neighbours = append(neighbours, s.grid[t.position.y][t.position.x+1])
		}
	}
	//if DOWN is oki, the neighbouring must have up oki
	if t.position.y+1 > 0 && t.position.y+1 < 7 {
		if t.direction[2] == 49 && s.grid[t.position.y+1][t.position.x].direction[0] == 49 {
			neighbours = append(neighbours, s.grid[t.position.y+1][t.position.x])
		}
	}
	//if LEFT is oki neighbour must have right!!
	if t.position.x-1 > 0 && t.position.x-1 < 7 {
		if t.direction[3] == 49 && s.grid[t.position.y][t.position.x-1].direction[1] == 49 {
			neighbours = append(neighbours, s.grid[t.position.y][t.position.x-1])
		}
	}
	return neighbours
}

func (s *State) read() {
	var turnType int
	fmt.Scan(&turnType)
	s.turn.turnType = turnType

	//cf GameBoard => sendMapToPlayer()
	for y := 0; y < 7; y++ {
		for x := 0; x < 7; x++ {
			var row string
			fmt.Scan(&row)
			s.grid[y][x].direction = row
			s.grid[y][x].position.x = x
			s.grid[y][x].position.y = y
		}
	}

	for i := 0; i < 2; i++ {
		// numPlayerCards: the total number of quests for a player (hidden and revealed)
		var numPlayerCards, playerX, playerY int
		var playerTile string
		fmt.Scan(&numPlayerCards, &playerX, &playerY, &playerTile)
		s.players[i].totalOfQuests = numPlayerCards
		s.players[i].position = Point{x: playerX, y: playerY}
		s.players[i].playerTile.direction = playerTile
	}

	// numItems: the total number of items available on board and on player tiles
	var numItems int
	fmt.Scan(&numItems)
	s.numItems = numItems

	for i := 0; i < numItems; i++ {
		var itemName string
		var itemX, itemY, itemPlayerId int
		fmt.Scan(&itemName, &itemX, &itemY, &itemPlayerId)
		switch itemX {
		//special case
		case -1:
			s.players[0].playerTile.hasItem = true
		case -2:
			s.players[1].playerTile.hasItem = true
		default:
			s.grid[itemY][itemX].hasItem = true
			s.grid[itemY][itemX].itemName = itemName
			s.grid[itemY][itemX].itemPlayerId = itemPlayerId
			//simply save the item tile
			switch itemPlayerId {
			case 0:
				s.players[0].itemTile = s.grid[itemY][itemX]
			case 1:
				s.players[1].itemTile = s.grid[itemY][itemX]
			}
			s.itemTiles = append(s.itemTiles, s.grid[itemY][itemX])
		}
	}

	// numQuests: the total number of revealed quests for both players
	var numQuests int
	fmt.Scan(&numQuests)

	for i := 0; i < numQuests; i++ {
		var questItemName string
		var questPlayerId int
		fmt.Scan(&questItemName, &questPlayerId)
		switch questPlayerId {
		case 0:
			s.players[0].questItemName = questItemName
		case 1:
			s.players[1].questItemName = questItemName
		}
	}
}

//a bfs?
func (s *State) bfsPath(playerTile, questTile Tile) []Tile {
	var visited = make(map[Tile]bool)
	visited[playerTile] = true

	var startTile = []Tile{playerTile}
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

func (s *State) getDirectionsFromPath(path []Tile) {
	if len(path) != 0 {
		//TODO:count steps
		for x, p := range path {
			if x+1 < len(path) {
				dir := p.position.printDirection(path[x+1].position)
				s.turn.directions = append(s.turn.directions, dir)
			}
		}
	}

}

// BUG: after i grab a quest if turn is push i got invalid input!!
// it seems like next quest is revealed in move mode not in push one ?? cf src to
// check...
// Should simulate map no? What happens if i push here? is it good or not?
func (s *State) printTurn() string {
	var command string
	if s.turn.turnType == 0 {
		if s.players[0].itemTile.position.y < s.players[0].position.y {
			command = fmt.Sprintf("PUSH %d %s", s.players[0].position.y, "RIGHT")
		} else if s.players[0].itemTile.position.y > s.players[0].position.y {
			command = fmt.Sprintf("PUSH %d %s", s.players[0].position.y, "LEFT")
			//got invalid input? see item init in turn??
		} else {
			command = fmt.Sprintf("PUSH %d %s", s.players[1].position.y, "RIGHT")
		}

		if s.players[0].itemTile.position.x < s.players[0].position.x {
			command = fmt.Sprintf("PUSH %d %s", s.players[0].position.x, "UP")
		} else if s.players[0].itemTile.position.x > s.players[0].position.x {
			command = fmt.Sprintf("PUSH %d %s", s.players[0].position.x, "DOWN")
		} else {
			command = fmt.Sprintf("PUSH %d %s", s.players[1].position.x, "DOWN")
		}
	} else {
		if len(s.turn.directions) == 0 {
			//to begin with then should go for half path near the quest?
			command = "PASS"
		} else {
			s := strings.Join(s.turn.directions, " ")
			command = fmt.Sprintf("MOVE %s", s)
		}
	}
	return command
}
func (s *State) think() {
	//push strat
	// first i could simple bfsPath the first player, if there is a direct route to
	// his quest, move the tile..
	var oppPath = s.bfsPath(s.grid[s.players[1].position.y][s.players[1].position.x], s.players[1].itemTile)
	log.Println("oppPath:", oppPath)
	//if len(s.players[1].turn.directions)
	//move strat
	var path = s.bfsPath(s.grid[s.players[0].position.y][s.players[0].position.x], s.players[0].itemTile)
	log.Println("path:", path)
}

func main() {

	for {
		//clean state to begin with
		s := State{}
		s.read()
		s.think()
		//should improve that!!
		comm := s.printTurn()
		fmt.Println(comm)

		//TEST LOGS
		s.grid.printGrid()
		log.Println("MY quest: ", s.players[0].questItemName, "located", s.players[0].itemTile)
		log.Println("OPP quest: ", s.players[1].questItemName, "located", s.players[1].itemTile)
		log.Println(s.bfsPath(s.grid[s.players[0].position.y][s.players[0].position.x], s.players[0].itemTile))
		log.Println(s.turn.directions)

	}
}
