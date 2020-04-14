package main

//put that here, cant make it work dont get stubborn time to move on!!
//https://rosettacode.org/wiki/Bitmap/Flood_fill#Go
/*
function fill(array, x, y)
  if !array[x, y]
    array[x, y] = true
    if y > 0: fill(array, x, y-1)
    if x > 0: fill(array, x-1, y)
    if x < array.width-1: fill(array, x+1, y)
    if y < array.height-1: fill(array, x, y+1)
*/
func floodfill(x int, y int) {
	var visited = make(map[Tile]bool)
if !visited[s.carte[t.pos.x][t.pos.y]]{

	//!!! cant check south if i go north!! i just came from south!!
	//check north
	if t.pos.y >= 0 {
		floodfill(grid,x,y-1)
		//queue = append(queue, s.carte[t.pos.x][t.pos.y-1])
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
