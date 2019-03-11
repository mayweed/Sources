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
	what string
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
	doughCrates []Cell
	//tools
	customerWindow Cell
	dishwasher     Cell
	choppingBoard  Cell
	oven           Cell
	//tables
	emptyTables []Cell
	dishTable   []Table
	bbTable     []Table
	icTable     []Table
	//perso
	myDish Cell
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
	players      [2]Chef
}

func (s *State) parseKitchen() {
	var kitchenLine string
	for y := 0; y < HEIGHT; y++ {
		fmt.Scan(&kitchenLine)
		kl := strings.Split(kitchenLine, "")
		for x, c := range kl {
			s.k.grid[y][x].x = x
			s.k.grid[y][x].y = y
			s.k.grid[y][x].what = c
			switch c {
			case "B":
				s.k.blueCrates = append(s.k.blueCrates, s.k.grid[y][x])
			case "I":
				s.k.iceCrates = append(s.k.iceCrates, s.k.grid[y][x])
			case "S":
				s.k.strawCrates = append(s.k.strawCrates, s.k.grid[y][x])
			case "H":
				s.k.doughCrates = append(s.k.doughCrates, s.k.grid[y][x])
			case "W":
				s.k.customerWindow = s.k.grid[y][x]
			case "D":
				s.k.dishwasher = s.k.grid[y][x]
			case "C":
				s.k.choppingBoard = s.k.grid[y][x]
			case "O":
				s.k.oven = s.k.grid[y][x]
			case "#":
				s.k.emptyTables = append(s.k.emptyTables, s.k.grid[y][x])
			}
		}
	}

}
func (s *State) parseChefs() {
	var playerX, playerY int
	var playerItem string
	fmt.Scan(&playerX, &playerY, &playerItem)
	s.players[0] = Chef{Cell{playerX, playerY, ""}, playerItem}

	var partnerX, partnerY int
	var partnerItem string
	fmt.Scan(&partnerX, &partnerY, &partnerItem)
	s.players[1] = Chef{Cell{playerX, playerY, ""}, playerItem}

}
func (s *State) parseTables() {
	// numTablesWithItems: the number of tables in the kitchen that currently hold an item
	var numTablesWithItems int
	fmt.Scan(&numTablesWithItems)

	for i := 0; i < numTablesWithItems; i++ {
		var tableX, tableY int
		var item string
		fmt.Scan(&tableX, &tableY, &item)
		t := Table{Cell{tableX, tableY, ""}, item}
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

//find an empty table to store my dish
func (s *State) findEmptyTable(c Cell) Cell {
	var t Cell
	if s.k.grid[c.y][c.x-1].what == "#" && c.x-1 > 0 {
		//this is not optimal: you should go to the straw crate, you got
		//a dish you put it on a nearby table then you do your chopping
		//things and take the dish
		s.k.myDish = s.k.grid[c.y][c.x-1]
		t = s.k.grid[c.y][c.x-1]
	} else if s.k.grid[c.y][c.x+1].what == "#" && c.x+1 < WIDTH {
		s.k.myDish = s.k.grid[c.y][c.x+1]
		t = s.k.grid[c.y][c.x+1]
	} else if s.k.grid[c.y-1][c.x].what == "#" && c.y-1 > 0 {
		s.k.myDish = s.k.grid[c.y-1][c.x]
		t = s.k.grid[c.y-1][c.x]
	} else if s.k.grid[c.y+1][c.x].what == "#" && c.y+1 < HEIGHT {
		s.k.myDish = s.k.grid[c.y+1][c.x]
		t = s.k.grid[c.y+1][c.x]
	}
	return t
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

		s.parseChefs()
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

		//should serve the customer with most award first
		//pick an order!! By default the first..
		//here i can simulate move to see what is the best crate to go first
		//take all the order and serve them, and score the best one (biggest award?)
		//write an func (s *State)executeOrder(order string){} which yields a turn
		order := s.c[0].customerItem
		myItems := s.players[0].items

		//i need to factor this code
		var res string
		if strings.Contains(order, "CROISSANT") &&
			!strings.Contains(myItems, "DOUGH") &&
			!strings.Contains(myItems, "CROISSANT") &&
			ovenContents == "NONE" {
			res = use(s.k.grid[s.k.doughCrates[0].y][s.k.doughCrates[0].x])
		} else if strings.Contains(myItems, "DOUGH") &&
			ovenContents == "NONE" {
			res = use(s.k.oven)
		} else if ovenContents == "DOUGH" {
			//sth is cooking
			res = "WAIT"
		} else if ovenContents == "CROISSANT" {
			res = use(s.k.oven)
			//	}
		} else if strings.Contains(myItems, "CROISSANT") && !strings.Contains(myItems, "DISH") {
			res = use(s.k.dishwasher)
		} else if strings.Contains(order, "CHOPPED_STRAWBERRIES") && !strings.Contains(myItems, "STRAWBERRIES") {
			//WARNING!! when you put your dish with a croissant on a table, myItem is
			//reseted to NONE. So should keep my own track of what i collected!!
			if strings.Contains(myItems, "DISH") {
				et := s.findEmptyTable(s.players[0].pos)
				res = use(et)
			} else {
				res = use(s.k.grid[s.k.strawCrates[0].y][s.k.strawCrates[0].x])
			}
		} else if strings.Contains(order, "CHOPPED_STRAWBERRIES") && !strings.Contains(myItems, "CHOPPED_STRAWBERRIES") {
			//i already picked straws, go chopping instead
			res = use(s.k.choppingBoard)
			//should i add a || with strawberries?
			//should pick my dish after if i got one
		} else if !strings.Contains(myItems, "DISH") { // &&
			//	(strings.Contains(myItems, "CROISSANT") || strings.Contains(myItems, "CHOPPED_STRAWBERRIES")) {
			res = use(s.k.dishwasher)
		} else if strings.Contains(order, "BLUEBERRIES") && !strings.Contains(myItems, "BLUEBERRIES") {
			res = use(s.k.grid[s.k.blueCrates[0].y][s.k.blueCrates[0].x])
		} else if strings.Contains(order, "ICE_CREAM") && !strings.Contains(myItems, "ICE_CREAM") {
			res = use(s.k.grid[s.k.iceCrates[0].y][s.k.iceCrates[0].x])
		} else {
			//nothing left to do just go to customer?
			res = use(s.k.customerWindow)
		}

		fmt.Println(res)

		//flush state between turns
		s.c = []Customer{}

		//LOGS
		log.Println(s.c, ovenContents, ovenTimer, "ORDER", order, "myItems", myItems)
	}
}
