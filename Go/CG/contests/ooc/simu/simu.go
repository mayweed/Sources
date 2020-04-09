package main

import (
	"log"
	"math/rand"
	"time"
)

/*
// sim///
//s.directions = []Direction{Direction{-1, 0, "W"}, Direction{0, -1, "N"}, Direction{+1, 0, "E"}, Direction{0, +1, "S"}}
func estimateTrajectory(s State, numTurns int) {
	//simply play the game of trajectory right?
	cpState := s //a copy of the state
	//copy of a board to test
	for i := 0; i < numTurns; i++ {
		if cpState.me.canGoEast {
			//MUST find ja simpler way to do that: getTile!!! and a valid one!! with no shit
			//if y+1 etc
			var nextPos Point
			nextPos.x += s.directions[2][0]
			nextPos.y += s.directions[2][1]

			//updateMap
			cpState.carte[nextPos.x][nextPos.y].what = "M"
		}
		//if next pos is a deadstreet do not go there!!
		//if nextPos.pos.x+1 >0 &&
	}
}
*/
func main() {
	//random num generator
	rand.Seed(time.Now().Unix())

	//init Grid
	var board = "xx..xxx............xxx.............xxx.....xx.....xxx.....xx.....xxx.......................................................xx.............xx...............xxx............xxx......x.....xxx......x......xx......................"
	var g Grid
	g.NewGrid(board)
	g.getWalkableTiles()

	//Players
	var p [2]Player //p0 is me right

	//Get a random starting pos on grid for each
	for i := range p {
		p[i].currentPos = g.wtiles[rand.Intn(len(g.wtiles))]
	}

	var turn int

	//game loop in itself
	for turn < 1 {

		//update grid manually, next step factorize and make it move
		g.carte[p[0].currentPos.x][p[0].currentPos.y].what = "M"
		g.carte[p[1].currentPos.x][p[1].currentPos.y].what = "O"

		log.Println(p[0].currentPos)
		g.printGrid()
		g.getTileNeighbours(&p[0].currentPos)
		var visitedTiles = make(map[Tile]bool)
		visitedTiles[p[0].currentPos] = true
		log.Println(p[0].currentPos.neighbours)
		turn += 1 //inc turn
	}
}
