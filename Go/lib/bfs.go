package lib

type graph struct {
	nodes []int
	edges map[int][]int
}

//simple bfs
func (g graph) bfs(startNode, endNode int) []int {
	var visited = make(map[int]bool)
	visited[startNode] = true

	var node = []int{startNode}
	var queue = [][]int{node}

	for 0 < len(queue) {
		//pop the first element
		path := queue[0]
		queue = queue[1:]

		lastNode := path[len(path)-1]
		if lastNode == endNode {
			return path
		}

		for _, w := range g.edges[lastNode] {
			var newPath = path
			if !visited[w] {
				visited[w] = true
				newPath = append(newPath, w)
				queue = append(queue, newPath)
			}
		}
	}
	return []int{}
}
