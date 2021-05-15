package main

import "container/list"

//see https://cybernetist.com/2019/03/09/breadth-first-search-using-go-standard-library/
//cant use tile as map key: it has a slice!!
func floodfillByDir(t Tile) {
	var visited = make(map[Tile]bool)
	visited[t] = true

	var queue = list.New()
	queue.PushBack(t)

	for queue.Len() > 0 {
		s := queue.Front()
		//type switch
		for _, neigh := range s.Value.(Tile).neighbours {
			if !visited[neigh] {
				visited[neigh] = true
				queue.PushBack(neigh)
			}
		}
		queue.Remove(s)
	}
}
