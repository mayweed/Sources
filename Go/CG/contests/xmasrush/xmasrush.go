package main

import (
	"fmt"
	"log"
)

const (
	MAP_WIDTH = 7
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

type Grid [7][7]Tile

type Tile struct {
	direction    string
	position     Point
	hasItem      bool
	itemName     string
	itemPlayerId int
}

type Player struct {
	totalOfQuests int
	quests        []string
	tile          Tile
	position      Point
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

//this works in wood 2 I suppose, should refine by calculating which one is the
//nearest from current pos...in this case the quest is passed in argument!!
//should definitely REWRITE it!!
func (s *State) getItemTilesPos(p Player) Tile {
	var quest = s.players[0].quests[0]
	var tile Tile
	for _, tile = range s.itemTiles {
		//dont forget the player id
		if tile.itemName == quest && tile.itemPlayerId == 0 {
			break
		}
	}
	return tile
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
		s.players[i].position.x = playerX
		s.players[i].position.y = playerY
		s.players[i].tile.direction = playerTile
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
		case -1:
			s.players[0].tile.hasItem = true
		case -2:
			s.players[1].tile.hasItem = true
		default:
			s.grid[itemY][itemX].hasItem = true
			s.grid[itemY][itemX].itemName = itemName
			s.grid[itemY][itemX].itemPlayerId = itemPlayerId
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
			s.players[0].quests = append(s.players[0].quests, questItemName)
		case 1:
			s.players[1].quests = append(s.players[1].quests, questItemName)
		}
	}
}

//a bfs?
//Should yse a Tile!!or write Point -> tile
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
func (s *State) think() {
	//ternary op would be great here, to test only
	if s.turn.turnType == 0 {
		fmt.Println("PUSH 3 RIGHT") // PUSH <id> <direction> | MOVE <direction> | PASS
	} else {
		//s.bfsPath()
		var path = s.bfsPath(s.grid[s.players[0].position.y][s.players[0].position.x], s.getItemTilesPos(s.players[0]))
		if len(path) == 0 {
			//=> move in the direction if possible or pass
			fmt.Println("MOVE RIGHT")
		} else { //test purpose if len(path) < 20 {
			//go for it directly!! Indeed, try to...
			for x, p := range path {
				//that or a queue...or an index??
				if x+1 < len(path) {
					dir := p.position.printDirection(path[x+1].position)
					s.turn.directions = append(s.turn.directions, dir)
				}
			}
		}
	}
}

func main() {

	for {
		//clean state to begin with
		s := State{}
		s.read()
		s.think()
		//s.printTurn()

		//TEST LOGS
		//log.Println(s.players[0].quests)
		log.Println(s.getItemTilesPos(s.players[0]))
		log.Println(s.getNeighbours(s.grid[s.players[0].position.y][s.players[0].position.x]))
		log.Println(s.bfsPath(s.grid[s.players[0].position.y][s.players[0].position.x], s.getItemTilesPos(s.players[0])))
		log.Println(s.turn.directions)

	}
}
