package main

//put that here, cant make it work dont get stubborn time to move on!!
func (s *State) floodfill(t Tile, depth int) int {
	var queue []Tile
	queue = append(queue, t)
	var area int

	var visited = make(map[Tile]bool)

	for len(queue) != 0 {
		var t = queue[0]
		queue = queue[1:]

		//!!! cant check south if i go north!! i just came from south!!
		//check north
		if t.pos.y-1 >= 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && !visited[s.carte[t.pos.x][t.pos.y-1]] {
			queue = append(queue, s.carte[t.pos.x][t.pos.y-1])
			visited[s.carte[t.pos.x][t.pos.y-1]] = true
			area += 1
		}
		//check south
		if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && !visited[s.carte[t.pos.x][t.pos.y+1]] {
			queue = append(queue, s.carte[t.pos.x][t.pos.y+1])
			visited[s.carte[t.pos.x][t.pos.y+1]] = true
			area += 1
		}
		//check west
		if t.pos.x-1 >= 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && !visited[s.carte[t.pos.x-1][t.pos.y]] {
			queue = append(queue, s.carte[t.pos.x-1][t.pos.y])
			visited[s.carte[t.pos.x-1][t.pos.y]] = true
			area += 1

		}
		//check east
		if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && !visited[s.carte[t.pos.x+1][t.pos.y]] {
			queue = append(queue, s.carte[t.pos.x+1][t.pos.y])
			visited[s.carte[t.pos.x+1][t.pos.y]] = true
			area += 1

		}
	}
	return 0
}
