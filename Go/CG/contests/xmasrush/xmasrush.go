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
	direction    string
	position     Point
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
}

/*
//To simulate you take the grid
//Parse the action,
// apply it
// give back the future grid
//better : you pass a turn!!
// between each turn you must reset the state!! you push right on the original grid
// not on the test one!! and for EACH row
func (s State) simulateGrid(t Turn){
	g:=s.grid
	if t.turn.actionType == "PUSH"{
	for c,p := range s.players{
	for row:=0;row < MAP_WIDTH;row++{
		switch a.directions[0]{
			case "RIGHT":
			g.pushRight(s.players[c].playerTile,row)
			s.reset()
			//if it's right/left you offset x

			//up/down you offset y
			...
		}
	}
	}
}
*/
type Player struct {
	totalOfQuests int
	//At most one is revealed
	quests []string
	//the tile where the item to quest is
	questTiles []Tile
	itemTiles  []Tile
	//the tile in hand
	playerTile        Tile
	playerTileisQuest bool
	position          Point
	//path if avail to quest item directly
	path      [][]Tile
	otherPath [][]Tile
}

func (p *Player) getQuestTile() {
	for _, tile := range p.itemTiles {
		for _, quest := range p.quests {
			if tile.position.x == -1 || tile.position.x == -2 {
				p.playerTileisQuest = true
				//continue //do not add a playerTile to quest!!
			}
			if tile.itemName == quest {
				p.questTiles = append(p.questTiles, tile)
			}
		}
	}
}

/*
func (p Player) isPlayerTileQuestTile() (bool, int) {
	for index, quest := range p.quests {
		if p.playerTile.itemName == quest {
			return true, index
		}
	}
	return false, -1
}
*/
type Action struct {
	actionType string
	id         int
	directions [][]string //better to pass them in arg?
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
	for index := range t.action.directions {
		c = strings.Join(t.action.directions[index], " ")
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
			s.players[0].quests = append(s.players[0].quests, questItemName)
		case 1:
			s.players[1].quests = append(s.players[0].quests, questItemName)
		}
	}
	s.players[0].getQuestTile()
	//s.players[1].getQuestTile()

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
		//all the tile of the map?
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

func (s *State) getDirectionsFromPath(path [][]Tile) []string {
	var dirs []string
	if len(path) != 0 {
		for index, _ := range path {
			if len(path[index]) > 0 {
				s.turn.action.steps = len(path[index])
				for x, _ := range path[index] {
					if x+1 < len(path[index]) {
						dir := path[index][x].position.printDirection(path[index][x+1].position)
						dirs = append(dirs, dir)
					}
				}
			}
		}
	}
	return dirs

}

// Should simulate map no? What happens if i push here? is it good or not?
//handle the case where i got the itemTile as player tile
//This is useless heuristics should use pushUP etc test all the possibilités and eval
//the best!!
func (s *State) printTurn() { // string {
	if s.turn.turnType == "PUSH" {
		//it should be one of the quest tiles, not necessarily the first!!
		//if s.players[0].questTiles[0].position.x == -1 || s.players[0].questTiles[0].position.x == -2 {
		//i write pushUp etc..to simu and i end up with spaghetti code!!
		if s.players[0].playerTileisQuest {
			//push at the beginning or end to grab item next turn!!
			if s.players[0].position.x == 0 {
				s.turn.push(s.players[0].position.y, "LEFT")
			} else if s.players[0].position.x == MAP_WIDTH-1 {
				s.turn.push(s.players[0].position.y, "RIGHT")
			} else {
				//cas général
				s.turn.push(s.players[1].position.y, "LEFT")
			}
		} else if s.players[0].questTiles[0].position.y == -1 || s.players[0].questTiles[0].position.y == -2 {
			if s.players[0].position.y == 0 {
				s.turn.push(s.players[0].position.x, "UP")
			} else if s.players[0].position.y == MAP_HEIGHT-1 {
				s.turn.push(s.players[0].position.x, "DOWN")
			} else {
				//ET pour ce que n'est pas aux 2 bouts???
				s.turn.push(s.players[0].position.x, "UP")
			}
		} else {

			if s.players[0].questTiles[0].position.y != s.players[0].position.y {
				if s.players[0].questTiles[0].position.y < s.players[0].position.y {
					s.turn.push(s.players[0].questTiles[0].position.y, "RIGHT")
				}
				if s.players[0].questTiles[0].position.y > s.players[0].position.y {
					s.turn.push(s.players[0].questTiles[0].position.y, "LEFT")
					//got invalid input? see item init in turn??
				}
			} else if s.players[0].questTiles[0].position.x != s.players[0].position.x {
				if s.players[0].questTiles[0].position.x < s.players[0].position.x {
					s.turn.push(s.players[0].questTiles[0].position.x, "UP")
				}
				if s.players[0].questTiles[0].position.x > s.players[0].position.x {
					s.turn.push(s.players[0].questTiles[0].position.x, "DOWN")
				}
			}
		}
	}

	if s.turn.turnType == "MOVE" {
		//should be player path here, not directions!!
		// MUST handle the case with multiple path to choose from!!
		var nonEmptyPath bool
		for index := range s.players[0].path {
			if len(s.players[0].path[index]) > 0 {
				s.turn.action.directions = append(s.turn.action.directions, s.getDirectionsFromPath(s.players[0].path))
				nonEmptyPath = true
			}
		}
		/*
			//try to make it move..
			//use only if I have nothing, really not accurate (distance? Voronoi?)
			//not that good at first (like going out without compass!!)
			if !nonEmptyPath {
				for y := 0; y < MAP_HEIGHT; y++ {
					for x := 0; x < MAP_WIDTH; x++ {
						if x == s.players[0].position.x && y == s.players[0].position.y {
							continue
						} else {
							p := s.bfsPath(s.grid[s.players[0].position.y][s.players[0].position.x], s.grid[y][x])
							if len(p) > 0 && len(p) < 20 {
								s.players[0].otherPath = append(s.players[0].otherPath, p)
							}
						}
					}
					//s.players[0].otherPath = [][]Tile{}
				}
				for i := range s.players[0].otherPath {
					//should compare the quest!!should have a voronoi of the quest here!!
					if s.players[0].otherPath[i][len(s.players[0].otherPath[i])-1].position.y > s.players[0].position.y {
						s.turn.action.directions = append(s.turn.action.directions, s.getDirectionsFromPath(s.players[0].otherPath))
						nonEmptyPath = true
					}
				}
			}
		*/
		if nonEmptyPath {
			s.turn.move()
		} else {
			s.turn.pass()
		}

	}
	fmt.Println(s.turn.command)
}

func (s *State) think() {
	for _, t := range s.players[0].questTiles {
		if t.position.x == -1 || t.position.y == -2 {
			continue
		} else {
			//they can be multiple paths to evaluate!! Would have to work on that!!
			s.players[0].path = append(s.players[0].path, s.bfsPath(s.grid[s.players[0].position.y][s.players[0].position.x], t))
		}
	}

	//if !s.players[1].isPlayerTileQuestTile() {
	//s.players[1].path = s.bfsPath(s.grid[s.players[1].position.y][s.players[1].position.x], s.players[1].questTile)
	//}
}

func main() {
	for {
		//clean state to begin with
		s := State{}
		s.read()
		s.think()
		s.printTurn()

		log.Println(s.turn.turnType)
		//log.Println(s.players[0].questTiles)
		//log.Println(s.players[0].playerTile)
		log.Println(s.players[0].otherPath)
		//log.Println(s.players[0].questTiles)
		//TEST LOGS
		//log.Println(s.players[0].quests)
		//a problem in my push
		//log.Println(s.players[0].position)

	}
}
