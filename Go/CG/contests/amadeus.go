package main

import (
	"fmt"
	"log"
)

type board struct {
	planetCount int
	planets     []planet
	edge        map[int]int
}

type planet struct {
	id             int
	myUnits        int
	myTolerance    int
	otherUnits     int
	otherTolerance int
	canAssign      int
}

func (p planet) isTarget() bool {
	if p.myTolerance > 0 && p.myUnits > 0 {
		return true
	} else {
		return false
	}
}

func main() {

	var planetCount, edgeCount int
	fmt.Scan(&planetCount, &edgeCount)
	b := board{planetCount: planetCount, edge: make(map[int]int)}

	for i := 0; i < edgeCount; i++ {
		var planetA, planetB int
		fmt.Scan(&planetA, &planetB)
		b.edge[planetA] = planetB
	}
	for {
		for i := 0; i < planetCount; i++ {
			var myUnits, myTolerance, otherUnits, otherTolerance, canAssign int
			fmt.Scan(&myUnits, &myTolerance, &otherUnits, &otherTolerance, &canAssign)
			b.planets = append(b.planets, planet{i, myUnits, myTolerance, otherUnits, otherTolerance, canAssign})
		}

		log.Println(b.planetCount, b.planets)
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("1") // Write action to stdout
		fmt.Println("5")
		fmt.Println("2")
		fmt.Println("0")
		fmt.Println("3")
		fmt.Println("NONE")

		b.planets = []planet{}
	}
}
