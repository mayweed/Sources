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

// ---------------------- BASIC TYPES -------------------------

type Cell struct {
	x, y int // x = colonne, y = ligne
}

type Table struct {
	pos  Cell
	item string
}

type Kitchen struct {
	grid [HEIGHT][WIDTH]Cell

	blueCrates  []Cell
	iceCrates   []Cell
	strawCrates []Cell
	doughCrates []Cell

	customerWindow Cell
	dishwasher     Cell
	choppingBoard  Cell
	oven           Cell

	dishTable []Table
	bbTable   []Table
	icTable   []Table
}

type Customer struct {
	customerItem  string
	customerAward int
}

type Chef struct {
	pos   Cell
	items string
}

type State struct {
	k            Kitchen
	c            []Customer
	players      [2]Chef
	numCustomers int
}

// ---------------------- HELPERS -------------------------

func use(c Cell) string {
	return fmt.Sprintf("USE %d %d", c.x, c.y)
}

// ---------------------- PARSING -------------------------

func (s *State) parseKitchen() {
	var kitchenLine string
	for y := 0; y < HEIGHT; y++ {
		fmt.Scan(&kitchenLine)
		kl := strings.Split(kitchenLine, "")
		for x, c := range kl {

			s.k.grid[y][x] = Cell{x, y}

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
			}
		}
	}
}

func (s *State) parseChefs() {
	var x, y int
	var item string

	fmt.Scan(&x, &y, &item)
	s.players[0] = Chef{Cell{x, y}, item}

	var px, py int
	var pitem string
	fmt.Scan(&px, &py, &pitem)

	// ✅ correction du bug : mauvais playerX / item recopié
	s.players[1] = Chef{Cell{px, py}, pitem}
}

func (s *State) parseTables() {
	var num int
	fmt.Scan(&num)

	// ✅ éviter accumulation d’anciennes tables
	s.k.dishTable = nil
	s.k.bbTable = nil
	s.k.icTable = nil

	for i := 0; i < num; i++ {
		var x, y int
		var item string
		fmt.Scan(&x, &y, &item)

		t := Table{Cell{x, y}, item}

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

// ---------------------- MAIN -------------------------

func main() {
	var s State

	var numAllCustomers int
	fmt.Scan(&numAllCustomers)
	s.numCustomers = numAllCustomers

	for i := 0; i < numAllCustomers; i++ {
		var item string
		var award int
		fmt.Scan(&item, &award)
	}

	s.parseKitchen()

	for {
		var turns int
		fmt.Scan(&turns)

		s.parseChefs()
		s.parseTables()

		var ovenContents string
		var ovenTimer int
		fmt.Scan(&ovenContents, &ovenTimer)

		var numCustomers int
		fmt.Scan(&numCustomers)

		// ✅ reset clients
		s.c = nil

		for i := 0; i < numCustomers; i++ {
			var item string
			var award int
			fmt.Scan(&item, &award)
			s.c = append(s.c, Customer{item, award})
		}

		// ---------------------- LOGIQUE COOKING -------------------------

		order := s.c[0].customerItem
		myItems := s.players[0].items

		var res string

		// ------------------ CROISSANT -----------------------

		if strings.Contains(order, "CROISSANT") &&
			!strings.Contains(myItems, "DOUGH") &&
			!strings.Contains(myItems, "CROISSANT") &&
			ovenContents == "NONE" {

			if len(s.k.doughCrates) > 0 {
				c := s.k.doughCrates[0]
				res = use(c)
			} else {
				res = "WAIT"
			}

		} else if strings.Contains(myItems, "DOUGH") && ovenContents == "NONE" {
			res = use(s.k.oven)

		} else if ovenContents == "DOUGH" {
			res = "WAIT"

		} else if ovenContents == "CROISSANT" {
			res = use(s.k.oven)

		} else if strings.Contains(myItems, "CROISSANT") && !strings.Contains(myItems, "DISH") {
			res = use(s.k.dishwasher)

			// ------------------ FRAISES -----------------------

		} else if strings.Contains(order, "CHOPPED_STRAWBERRIES") &&
			!strings.Contains(myItems, "STRAWBERRIES") {

			if len(s.k.strawCrates) > 0 {
				c := s.k.strawCrates[0]
				res = use(c)
			} else {
				res = "WAIT"
			}

		} else if strings.Contains(order, "CHOPPED_STRAWBERRIES") &&
			!strings.Contains(myItems, "CHOPPED_STRAWBERRIES") {

			res = use(s.k.choppingBoard)

			// ---------------- DISH ---------------------------

		} else if !strings.Contains(myItems, "DISH") {
			res = use(s.k.dishwasher)

			// ---------------- BLUEBERRIES --------------------

		} else if strings.Contains(order, "BLUEBERRIES") &&
			!strings.Contains(myItems, "BLUEBERRIES") {

			if len(s.k.blueCrates) > 0 {
				c := s.k.blueCrates[0]
				res = use(c)
			} else {
				res = "WAIT"
			}

			// ---------------- ICE CREAM ----------------------

		} else if strings.Contains(order, "ICE_CREAM") &&
			!strings.Contains(myItems, "ICE_CREAM") {

			if len(s.k.iceCrates) > 0 {
				c := s.k.iceCrates[0]
				res = use(c)
			} else {
				res = "WAIT"
			}

			// ---------------- SERVIR -------------------------

		} else {
			res = use(s.k.customerWindow)
		}

		// ---------------------- OUTPUT -----------------------

		fmt.Println(res)

		log.Println("ORDER:", order, "myItems:", myItems, "oven:", ovenContents, ovenTimer)
	}
}
