package main

// sim///
//s.directions = []Direction{Direction{-1, 0, "W"}, Direction{0, -1, "N"}, Direction{+1, 0, "E"}, Direction{0, +1, "S"}}
func estimateTrajectory(s State, numTurns int) {
	//simply play the game of trajectory right?
	cpState := s //a copy of the state
	//copy of a board to test
	cpState.board = " xx..xxx............xxx.............xxx.....xx.....xxx.....xx.....xxx.......................................................xx.............xx...............xxx............xxx......x.....xxx......x......xx......................"
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
