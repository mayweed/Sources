package main

import "log"

//put that here, cant make it work dont get stubborn time to move on!!
func (s *State) floodfill(t Tile, depth float64) map[string]float64 {
	var queue []Tile
	queue = append(queue, t)

	var areaN, areaS, areaW, areaE float64
	var ffd = make(map[string]float64)
	log.Println("BEGIN FOR: ", s.me.currentPos)
	for len(queue) != 0 {
		var t = queue[0]
		queue = queue[1:]

		if areaN == depth || areaS == depth || areaW == depth || areaE == depth {
			//best ffDir = max des 4
			ffd["N"] = areaN
			ffd["S"] = areaS
			ffd["E"] = areaE
			ffd["W"] = areaW
			return ffd
		}
		//!!! cant check south if i go north!! i just came from south!!
		//check north
		if t.pos.y-1 >= 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && !s.me.visitedTiles[s.carte[t.pos.x][t.pos.y-1]] {
			//s.carte[t.pos.x][t.pos.y-1].color = "blue" //wouldn't black be better?
			queue = append(queue, s.carte[t.pos.x][t.pos.y-1])
			log.Println("N cells visited: ", s.carte[t.pos.x][t.pos.y-1])
			areaN += 1.0
		}
		//check south
		if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && !s.me.visitedTiles[s.carte[t.pos.x][t.pos.y+1]] {
			//s.carte[t.pos.x][t.pos.y+1].color = "blue"
			queue = append(queue, s.carte[t.pos.x][t.pos.y+1])
			areaS += 1.0
			log.Println("S cells visited: ", s.carte[t.pos.x][t.pos.y+1])
		}
		//check west
		if t.pos.x-1 >= 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && !s.me.visitedTiles[s.carte[t.pos.x-1][t.pos.y]] {
			//s.carte[t.pos.x-1][t.pos.y].color = "blue"
			queue = append(queue, s.carte[t.pos.x-1][t.pos.y])
			areaW += 1.0
			log.Println("W cells", s.carte[t.pos.x-1][t.pos.y])

		}
		//check east
		if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && !s.me.visitedTiles[s.carte[t.pos.x+1][t.pos.y]] {
			//s.carte[t.pos.x+1][t.pos.y].color = "blue"
			queue = append(queue, s.carte[t.pos.x+1][t.pos.y])
			areaE += 1.0
			log.Println("EAST cells visited: ", s.carte[t.pos.x+1][t.pos.y])

		}
	}
	return nil
}
