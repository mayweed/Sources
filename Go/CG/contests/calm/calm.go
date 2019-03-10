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

type Table struct {
	pos  Cell
	item string
}
type Kitchen struct {
	grid [HEIGHT][WIDTH]Cell
	//crates
	blueCrates  []Cell
	iceCrates   []Cell
	strawCrates []Cell
	//tools
	customerWindow Cell
	dishwasher     Cell
	choppingBoard  Cell
	//tables
	dishTable []Table
	bbTable   []Table
	icTable   []Table
}
type Customer struct {
	customerItem  string
	customerAward int
}
type Chef struct {
	pos Cell
	//playeritem?
	items string
}
type State struct {
	k            Kitchen
	numCustomers int
	c            []Customer
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
			case "S":
				s.k.strawCrates = append(s.k.strawCrates, s.k.grid[y][x])
			case "W":
				s.k.customerWindow = s.k.grid[y][x]
			case "D":
				s.k.dishwasher = s.k.grid[y][x]
			case "C":
				s.k.choppingBoard = s.k.grid[y][x]
			}
		}
	}

}
func (s *State) parseTables() {
	// numTablesWithItems: the number of tables in the kitchen that currently hold an item
	var numTablesWithItems int
	fmt.Scan(&numTablesWithItems)

	for i := 0; i < numTablesWithItems; i++ {
		var tableX, tableY int
		var item string
		fmt.Scan(&tableX, &tableY, &item)
		t := Table{Cell{tableX, tableY}, item}
		switch item {
		case "DISH":
			s.k.dishTable = append(s.k.dishTable, t)
		case "BLUEBERRIES":
			s.k.bbTable = append(s.k.bbTable, t)
		case "ICE_CREAM":
			s.k.icTable = append(s.k.icTable, t)
		}
	}
}

//first action func
func use(c Cell) string {
	s := fmt.Sprintf("USE %d %d", c.x, c.y)
	return s
}

func main() {
	var s State

	var numAllCustomers int
	fmt.Scan(&numAllCustomers)
	s.numCustomers = numAllCustomers

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
		//should write a parseTurn() maybe a turn?
		var turnsRemaining int
		fmt.Scan(&turnsRemaining)

		var playerX, playerY int
		var playerItem string
		fmt.Scan(&playerX, &playerY, &playerItem)

		var partnerX, partnerY int
		var partnerItem string
		fmt.Scan(&partnerX, &partnerY, &partnerItem)

		s.parseTables()

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
			customer := Customer{customerItem, customerAward}
			s.c = append(s.c, customer)
		}

		//first get a dish
		//should split item to know what i got in hand
		//parsePlayerItem??
		var res string
		if playerItem == "NONE" {
			res = use(s.k.dishwasher)
		} else if playerItem == "DISH" {
			//no pos here?
			res = use(s.k.grid[s.k.blueCrates[0].y][s.k.blueCrates[0].x])
		} else if strings.Contains(playerItem, "BLUEBERRIES") {
			res = use(s.k.grid[s.k.iceCrates[0].y][s.k.iceCrates[0].x])
		}
		//validate the plate
		if strings.Contains(playerItem, "DISH") &&
			strings.Contains(playerItem, "ICE_CREAM") &&
			strings.Contains(playerItem, "BLUEBERRIES") {
			res = use(s.k.customerWindow)
		}

		fmt.Println(res)

		//flush state between turns
		s.c = []Customer{}

		//LOGS
		//log.Println(s.k.bbTable)
	}
}
