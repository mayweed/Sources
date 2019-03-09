package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	WIDTH  = 11
	HEIGHT = 7
)

type Cell struct {
	x, y int
	//what string
}

type Kitchen struct {
	grid           [HEIGHT][WIDTH]Cell
	blueCrates     []Cell
	iceCrates      []Cell
	customerWindow Cell
	dishwasher     Cell
}

type State struct {
	k Kitchen
}

func (s *State) parseKitchen() {
	var kitchenLine string
	for y := 0; y < HEIGHT; y++ {
		fmt.Scan(&kitchenLine)
		kl := strings.Split(kitchenLine, "")
		for x, c := range kl {
			s.k.grid[y][x].x = x
			s.k.grid[y][x].y = y
			//k.grid[y][x].what = c
			switch c {
			case "B":
				s.k.blueCrates = append(s.k.blueCrates, s.k.grid[y][x])
			case "I":
				s.k.iceCrates = append(s.k.iceCrates, s.k.grid[y][x])
			case "W":
				s.k.customerWindow = s.k.grid[y][x]
			case "D":
				s.k.dishwasher = s.k.grid[y][x]
			}
		}
	}

}
func main() {
	var s State

	var numAllCustomers int
	fmt.Scan(&numAllCustomers)

	for i := 0; i < numAllCustomers; i++ {
		// customerItem: the food the customer is waiting for
		// customerAward: the number of points awarded for delivering the food
		var customerItem string
		var customerAward int
		fmt.Scan(&customerItem, &customerAward)
	}

	s.parseKitchen()
	log.Println(s.k.iceCrates)

	for {
		var turnsRemaining int
		fmt.Scan(&turnsRemaining)

		var playerX, playerY int
		var playerItem string
		fmt.Scan(&playerX, &playerY, &playerItem)

		var partnerX, partnerY int
		var partnerItem string
		fmt.Scan(&partnerX, &partnerY, &partnerItem)

		// numTablesWithItems: the number of tables in the kitchen that currently hold an item
		var numTablesWithItems int
		fmt.Scan(&numTablesWithItems)

		for i := 0; i < numTablesWithItems; i++ {
			var tableX, tableY int
			var item string
			fmt.Scan(&tableX, &tableY, &item)
		}
		// ovenContents: ignore until wood 1 league
		var ovenContents string
		var ovenTimer int
		fmt.Scan(&ovenContents, &ovenTimer)

		// numCustomers: the number of customers currently waiting for food
		var numCustomers int
		fmt.Scan(&numCustomers)

		for i := 0; i < numCustomers; i++ {
			var customerItem string
			var customerAward int
			fmt.Scan(&customerItem, &customerAward)
		}

		fmt.Println("WAIT")
	}
}
