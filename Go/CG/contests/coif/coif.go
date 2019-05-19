package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}
type Player struct {
	gold   int
	income int
	units  int
}
type Building struct {
	location Point
	owner    int
	btype    int
}
type Map struct {
	location  Point
	state     string
	buildings Building
}

func main() {
	var numberMineSpots int
	fmt.Scan(&numberMineSpots)

	for i := 0; i < numberMineSpots; i++ {
		var x, y int
		fmt.Scan(&x, &y)
	}
	for {
		var gold int
		fmt.Scan(&gold)

		var income int
		fmt.Scan(&income)

		var opponentGold int
		fmt.Scan(&opponentGold)

		var opponentIncome int
		fmt.Scan(&opponentIncome)

		for i := 0; i < 12; i++ {
			var line string
			fmt.Scan(&line)
			strings.Split(line, "")
		}
		var buildingCount int
		fmt.Scan(&buildingCount)

		for i := 0; i < buildingCount; i++ {
			var owner, buildingType, x, y int
			fmt.Scan(&owner, &buildingType, &x, &y)
		}
		var unitCount int
		fmt.Scan(&unitCount)

		for i := 0; i < unitCount; i++ {
			var owner, unitId, level, x, y int
			fmt.Scan(&owner, &unitId, &level, &x, &y)
		}
		fmt.Println("WAIT") // Write action to stdout
	}
}
