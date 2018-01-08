package main

import (
	"fmt"
	"log"
	"strings"
)

//calculer le plus court chemin avec un DFS
//l'enregister dans edgeTo[]
//depiler edgeTo[] et construire la grille d'output
const (
	LEFT  string = "<"
	RIGHT string = ">"
	UP    string = "^"
	DOWN  string = "v"
)

/*
var LEFT = point{-1, 0}
var RIGHT = point{1, 0}
var DOWN = point{0, 1}
var UP = point{1, 0}
*/
type point struct {
	x, y int
}
type node struct {
	point
	value   int
	isHole  bool
	isWater bool
}

func main() {
	var W, H int
	fmt.Scan(&W, &H)

	var grid = make([][]string, H)
	for i := 0; i < H; i++ {
		var row string
		fmt.Scan(&row)
		inputs := strings.Split(row, "")
		grid[i] = make([]string, W)
		for j := range grid[i] {
			//here a switch to init node
			grid[i][j] = inputs[j]
		}
	}
	log.Println(H, W, grid)
	//TEST 1 to test
	//fmt.Println(">.")
}
