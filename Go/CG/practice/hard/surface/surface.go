package main

import (
	"fmt"
	"log"
	"strings"
)

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	x, y int
	what string
}

/*
Flood-fill (node, target-color, replacement-color):
 1. If target-color is equal to replacement-color, return.
 2. If the color of node is not equal to target-color, return.
 3. Set the color of node to replacement-color.
 4. Perform Flood-fill (one step to the south of node, target-color, replacement-color).
    Perform Flood-fill (one step to the north of node, target-color, replacement-color).
    Perform Flood-fill (one step to the west of node, target-color, replacement-color).
    Perform Flood-fill (one step to the east of node, target-color, replacement-color).
 5. Return
*/
func floodFill(c Cell, targetColor string, replacementColor string) {
	//if it's water 'o' should mark the cell
}
func main() {
	//init map
	var width int
	fmt.Scan(&width)

	var height int
	fmt.Scan(&height)

	var grid = make([][]Cell, height)
	var row string
	for i := 0; i < height; i++ {
		fmt.Scan(&row)
		s := strings.Split(row, "")
		grid[i] = make([]Cell, width)
		for j := 0; j < width; j++ {
			grid[i][j].x = j
			grid[i][j].y = i
			grid[i][j].what = s[j]
			log.Println(grid[i][j])
		}

	}

	//number of coord to be tested
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
	}
	for i := 0; i < N; i++ {
		fmt.Println("answer") // Write answer to stdout
	}
}
