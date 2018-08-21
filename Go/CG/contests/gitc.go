package main

import (
	"fmt"
	"log"
	"strings"
)

type Factory struct {
	id         int
	cyborgs    int
	production int
	owner      int
}
type Troop struct {
	id             int
	from           int
	to             int
	cyborgs        int
	remainingTurns int
	owner          int
}
type Link struct {
	from     int
	to       int
	distance int
}
type gameMap struct {
	factoryCount int
	linkCount    int
	factories    map[int]Factory
	troops       []Troop
	links        []Link
}

//COMMANDS
func cmdMove(source, destination, cyborgCount int) string {
	return fmt.Sprintf("MOVE %d %d %d", source, destination, cyborgCount)
}

func cmdWait(cardID int) string {
	return fmt.Sprintf("WAIT")
}
func sendCommands(commands []string) {
	cmd := "PASS"
	if len(commands) == 0 {
		log.Println("List of commands is empty, PASS will be sent")
	} else {
		cmd = strings.Join(commands, ";")
	}
	fmt.Println(cmd)
}

func main() {

	board := gameMap{}

	//put in a initMap() func
	// factoryCount: the number of factories
	var factoryCount int
	fmt.Scan(&factoryCount)
	board.factoryCount = factoryCount

	// linkCount: the number of links between factories
	var linkCount int
	fmt.Scan(&linkCount)
	board.linkCount = linkCount

	for i := 0; i < linkCount; i++ {
		var factory1, factory2, distance int
		fmt.Scan(&factory1, &factory2, &distance)
		board.links = append(board.links, Link{factory1, factory2, distance})
		//et vice versa?
	}
	var commands []string

	for {
		//put in a readEntity func
		// entityCount: the number of entities (e.g. factories and troops)
		var entityCount int
		fmt.Scan(&entityCount)

		for i := 0; i < entityCount; i++ {
			var entityId int
			var entityType string
			var arg1, arg2, arg3, arg4, arg5 int
			fmt.Scan(&entityId, &entityType, &arg1, &arg2, &arg3, &arg4, &arg5)
			//switch entityType {
			//  case "FACTORY":
			//}

		}

		//LOGS
		log.Println(board)

		//ALGO to get out of woods: take each of my fac with troops and move to neutral fac first
		//and then those of opp with less cyb?
		// Any valid action, such as "WAIT" or "MOVE source destination cyborgs"
		fmt.Println("WAIT")
	}
}
