package main

func (s *State) getBfsPath(startPos, target Tile) []Tile {
	var visited = make(map[Tile]bool)
	visited[startPos] = true

	var startTile = startPos
	var queue = []Tile{startTile}

	//keep track of the preceding tile
	var parent = make(map[Tile]Tile)

	var path []Tile
	for 0 < len(queue) {
		//pop the first element/shouldnt i use container/list here?
		//t for startNode
		t := queue[0]
		queue = queue[1:]

		if t == target {
			//p like predecessors
			var p = target
			path = append(path, target)
			for parent[p] != startPos {
				path = append(path, parent[p])
				p = parent[p]
			}
			path = append(path, startPos)
			return path
		}

		//check north
		if t.pos.y-1 >= 0 && isWalkable(s.carte[t.pos.x][t.pos.y-1]) && !visited[s.carte[t.pos.x][t.pos.y-1]] {
			visited[s.carte[t.pos.x][t.pos.y-1]] = true
			parent[s.carte[t.pos.x][t.pos.y-1]] = t
			queue = append(queue, s.carte[t.pos.x][t.pos.y-1])
		}
		//check south
		if t.pos.y+1 < HEIGHT && isWalkable(s.carte[t.pos.x][t.pos.y+1]) && !visited[s.carte[t.pos.x][t.pos.y+1]] {
			visited[s.carte[t.pos.x][t.pos.y+1]] = true
			parent[s.carte[t.pos.x][t.pos.y+1]] = t
			queue = append(queue, s.carte[t.pos.x][t.pos.y+1])
		}
		//check west
		if t.pos.x-1 >= 0 && isWalkable(s.carte[t.pos.x-1][t.pos.y]) && !visited[s.carte[t.pos.x-1][t.pos.y]] {
			visited[s.carte[t.pos.x-1][t.pos.y]] = true
			parent[s.carte[t.pos.x-1][t.pos.y]] = t
			queue = append(queue, s.carte[t.pos.x-1][t.pos.y])

		}
		//check east
		if t.pos.x+1 < WIDTH && isWalkable(s.carte[t.pos.x+1][t.pos.y]) && !visited[s.carte[t.pos.x+1][t.pos.y]] {
			visited[s.carte[t.pos.x+1][t.pos.y]] = true
			parent[s.carte[t.pos.x+1][t.pos.y]] = t
			queue = append(queue, s.carte[t.pos.x+1][t.pos.y])

		}
	}
	return path
}

//the distance from startPoint to all walkable Tiles!!
//!! You can't change values associated with keys in a map, you can only reassign values.
//!! When you "fill" the map, you can't use the loop's variable, as it gets overwritten in each iteration
// see : https://stackoverflow.com/questions/42716852/how-to-update-map-values-in-go
func (s *State) calculateDist(src Tile) map[Tile]*int {
	var dist = make(map[Tile]*int)
	var path []Tile
	//should change that walkable to take into account floodfill?
	for _, target := range s.walkableTiles {
		//dont need this one
		if target == src {
			continue
		}
		path = s.getBfsPath(src, target)
		length := len(path)
		dist[target] = &length
	}
	return dist
}

func (s *State) getTargets(dist map[Tile]*int) {
	//let's find targets find the process costly (recalc the dist to all tiles etc...
	var max = 4
	//to not damage myself
	//var min = 2
	var targetTile Tile
	for k, v := range dist {
		if *v == max { //&& *v > min {
			targetTile = k
			s.targets = append(s.targets, targetTile)
		}
	}
}
