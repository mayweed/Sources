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

type Tile struct {
	direction string
	position  Point
	//either me or the opp!!
	hasItem      bool
	itemName     string
	itemPlayerId int
}

type Grid [7][7]Tile

func (g Grid) getTile(x int, y int) Tile {
	var t Tile
	if isValidPos(Point{x, y}) {
		t = g[y][x]
	}
	return t
}
func (g *Grid) setTile(x int, y int, t Tile) {
	g[y][x] = t
}

func (g *Grid) pushUp(pushedTile Tile, col int) Tile {
	var maxRow = MAP_HEIGHT - 1
	var poppedTile = g.getTile(col, 0)

	for i := 0; i < maxRow; i++ {
		g.setTile(col, i, g.getTile(col, i+1))
		//getTile(col, i).move(Constants.Direction.UP);
	}
	g.setTile(col, maxRow, pushedTile)
	//getTile(col, maxRow).move(new Vector2(col, maxRow));

	return poppedTile
}

//wip...return the playerTile
//seems to sorta work
func (g *Grid) pushRight(pushedTile Tile, row int) Tile {
	var maxCol = MAP_WIDTH - 1
	var poppedTile = g.getTile(row, maxCol)

	for i := maxCol; i > 0; i-- {
		g.setTile(i, row, g.getTile(i-1, row))
		//getTile(i, row).move(Constants.Direction.RIGHT);
	}
	//you set the tile to push at the beginning left (push right
	g.setTile(0, row, pushedTile)
	//getTile(0, row).move(new Vector2(0, row));
	return poppedTile
}

func (g *Grid) pushDown(pushedTile Tile, col int) Tile {
	var maxRow = MAP_HEIGHT - 1
	var poppedTile = g.getTile(col, maxRow)

	for i := maxRow; i > 0; i-- {
		g.setTile(col, i, g.getTile(col, i-1))
		//getTile(col, i).move(Constants.Direction.DOWN);
	}
	g.setTile(col, 0, pushedTile)
	//getTile(col, 0).move(new Vector2(col, 0));
	return poppedTile
}

func (g *Grid) pushLeft(pushedTile Tile, row int) Tile {
	var maxCol = MAP_WIDTH - 1
	var poppedTile = g.getTile(0, row)
	for i := 0; i < maxCol; i++ {
		g.setTile(i, row, g.getTile(i+1, row))
		//getTile(i, row).move(Constants.Direction.LEFT);
	}
	g.setTile(maxCol, row, pushedTile)
	//getTile(maxCol, row).move(new Vector2(maxCol, row));
	return poppedTile
}

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
	//the tile where the item to quest is
	questTile Tile
	itemTiles []Tile
	//the tile in hand
	playerTile Tile
	position   Point
	//path if avail to quest item directly
	path []Tile
}

func (p *Player) getQuestTile() {
	if !p.isPlayerTileQuestTile() {
		for _, tile := range p.itemTiles {
			if tile.itemName == p.questItemName {
				p.questTile = tile
			}
		}
	}
}
func (p Player) isPlayerTileQuestTile() bool {
	if p.playerTile.itemName == p.questItemName {
		return true
	} else {
		return false
	}
}

func (p *Player) initQuestTile() {

}

type Action struct {
	actionType string
	id         int
	directions []string //better to pass them in arg?
	steps      int
}
type Turn struct {
	turnType string
	action   Action
	command  string
}

//helpers func
func (t *Turn) move() {
	var c string
	if len(t.action.directions) > 1 {
		c = strings.Join(t.action.directions, " ")
	} else {
		c = t.action.directions[0]
	}
	t.command = fmt.Sprintf("MOVE %s", c)
}
func (t *Turn) push(id int, direction string) {
	t.command = fmt.Sprintf("PUSH %d %s", id, direction)
}
func (t *Turn) pass() {
	t.command = fmt.Sprintf("PASS")
}

type State struct {
	players  [2]Player
	grid     Grid
	numItems int
	turn     Turn
	gameTurn int
}

//golang yields the ascii code not the num...it's a quickfix...
//see first answer here: https://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
func (s State) getNeighbours(t Tile) []Tile {
	var neighbours []Tile
	//to be qualified as neighbour one must be able to communicate
	//if UP is 1 on my tile, the upper tile is a neighbour if and only if down is
	//open!! Take one as a rune '1' ?
	if t.position.y-1 > 0 && t.position.y-1 < 7 {
		if t.direction[0] == '1' && s.grid[t.position.y-1][t.position.x].direction[2] == '1' {
			neighbours = append(neighbours, s.grid[t.position.y-1][t.position.x])
		}
	}
	//if RIGHT (direction 1) is ok, check left (dir 3) on the neighbouring cell
	if t.position.x+1 > 0 && t.position.x+1 < 7 {
		if t.direction[1] == '1' && s.grid[t.position.y][t.position.x+1].direction[3] == '1' {
			neighbours = append(neighbours, s.grid[t.position.y][t.position.x+1])
		}
	}
	//if DOWN is oki, the neighbouring must have up oki
	if t.position.y+1 > 0 && t.position.y+1 < 7 {
		if t.direction[2] == '1' && s.grid[t.position.y+1][t.position.x].direction[0] == '1' {
			neighbours = append(neighbours, s.grid[t.position.y+1][t.position.x])
		}
	}
	//if LEFT is oki neighbour must have right!!
	if t.position.x-1 > 0 && t.position.x-1 < 7 {
		if t.direction[3] == '1' && s.grid[t.position.y][t.position.x-1].direction[1] == '1' {
			neighbours = append(neighbours, s.grid[t.position.y][t.position.x-1])
		}
	}
	return neighbours
}

func (s *State) read() {
	var turnType int
	fmt.Scan(&turnType)
	if turnType == 0 {
		s.turn.turnType = "PUSH"
	} else {
		s.turn.turnType = "MOVE"
	}

	//cf GameBoard => sendMapToPlayer()
	for y := 0; y < MAP_HEIGHT; y++ {
		for x := 0; x < MAP_WIDTH; x++ {
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
		//special case
		switch itemX {
		case -1:
			s.players[0].playerTile.hasItem = true
			s.players[0].playerTile.position = Point{-1, -1}
			s.players[0].playerTile.itemName = itemName
			s.players[0].playerTile.itemPlayerId = itemPlayerId
		case -2:
			s.players[1].playerTile.hasItem = true
			s.players[1].playerTile.position = Point{-2, -2}
			s.players[1].playerTile.itemName = itemName
			s.players[1].playerTile.itemPlayerId = itemPlayerId
		default:
			s.grid[itemY][itemX].hasItem = true
			s.grid[itemY][itemX].itemName = itemName
			s.grid[itemY][itemX].itemPlayerId = itemPlayerId
		}

		//-1 /-1 change wrt whom his the opponent!!
		if itemPlayerId == 0 && itemX != -1 && itemX != -2 {
			s.players[0].itemTiles = append(s.players[0].itemTiles, s.grid[itemY][itemX])
		} else if itemPlayerId == 1 && itemX != -2 && itemX != -1 {
			s.players[1].itemTiles = append(s.players[1].itemTiles, s.grid[itemY][itemX])
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

		//goal: return all possibles paths to choose from!!
		lastTile := path[len(path)-1]
		if lastTile == questTile {
			return path
			//if i do that, will it ever return when goal reached??
		} //else if lastTile.

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
		s.turn.action.steps = len(path)
		for x, p := range path {
			if x+1 < len(path) {
				dir := p.position.printDirection(path[x+1].position)
				s.turn.action.directions = append(s.turn.action.directions, dir)
			}
		}
	}

}

// Should simulate map no? What happens if i push here? is it good or not?
//handle the case where i got the itemTile as player tile
//This is useless heuristics should use pushUP etc test all the possibilités and eval
//the best!!
func (s *State) printTurn() { // string {
	if s.turn.turnType == "PUSH" {
		//tile IS playerTile, bug if i tried to push?
		if s.players[0].questTile.position.x == -1 || s.players[0].questTile.position.x == -2 {
			//push at the beginning or end to grab item next turn!!
			if s.players[0].position.x == 0 {
				s.turn.push(s.players[0].position.y, "LEFT")
			} else if s.players[0].position.x == MAP_WIDTH-1 {
				s.turn.push(s.players[0].position.y, "RIGHT")
			} else {
				//cas général
				s.turn.push(s.players[1].position.y, "LEFT")
			}
		} else if s.players[0].questTile.position.y == -1 || s.players[0].questTile.position.y == -2 {
			if s.players[0].position.y == 0 {
				s.turn.push(s.players[0].position.x, "UP")
			} else if s.players[0].position.y == MAP_HEIGHT-1 {
				s.turn.push(s.players[0].position.x, "DOWN")
			} else {
				//ET pour ce que n'est pas aux 2 bouts???
				s.turn.push(s.players[0].position.x, "UP")
			}
		} else {

			if s.players[0].questTile.position.y != s.players[0].position.y {
				if s.players[0].questTile.position.y < s.players[0].position.y {
					s.turn.push(s.players[0].questTile.position.y, "RIGHT")
				}
				if s.players[0].questTile.position.y > s.players[0].position.y {
					s.turn.push(s.players[0].questTile.position.y, "LEFT")
					//got invalid input? see item init in turn??
				}
			} else if s.players[0].questTile.position.x != s.players[0].position.x {
				if s.players[0].questTile.position.x < s.players[0].position.x {
					s.turn.push(s.players[0].questTile.position.x, "UP")
				}
				if s.players[0].questTile.position.x > s.players[0].position.x {
					s.turn.push(s.players[0].questTile.position.x, "DOWN")
				}
			}
		}
	}

	if s.turn.turnType == "MOVE" {
		if len(s.turn.action.directions) == 0 {
			//to begin with then should go for half path near the quest?
			s.turn.pass()
		} else {
			s.turn.move()
		}
	}
	fmt.Println(s.turn.command)
}

func (s *State) think() {
	if s.players[0].isPlayerTileQuestTile() {
		s.players[0].questTile = s.players[0].playerTile
	} else {
		//init quest tile
		s.players[0].getQuestTile()
		s.players[1].getQuestTile()
	}

	if !s.players[0].isPlayerTileQuestTile() {
		s.players[0].path = s.bfsPath(s.grid[s.players[0].position.y][s.players[0].position.x], s.players[0].questTile)
	}

	if !s.players[1].isPlayerTileQuestTile() {
		s.players[1].path = s.bfsPath(s.grid[s.players[1].position.y][s.players[1].position.x], s.players[1].questTile)
	}
}

func main() {
	for {
		//clean state to begin with
		s := State{}
		s.read()
		s.think()
		s.printTurn()

		log.Println(s.turn.command)
		log.Println(s.players[0].itemTiles)
		//TEST LOGS
		log.Println(s.players[0].questItemName, "in", s.players[0].questTile)
		//a problem in my push
		log.Println(s.players[0].position)

		//only print them when needed
		//should go elsewhere...stringer??
		for id, p := range s.players {
			if len(p.path) > 0 {
				log.Println(id, p.path)
			}
		}
	}
}
