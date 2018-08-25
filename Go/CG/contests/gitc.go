package main

import (
	"fmt"
	"log"
	"strings"
)

//MAP
type Factory struct {
	id         int
	owner      int
	cyborgs    int
	production int
}

type Troop struct {
	id             int
	owner          int
	from           int
	to             int
	cyborgs        int
	remainingTurns int
}
type Link struct {
	from     int
	to       int
	distance int
}

//STATE
type gameMap struct {
	factoryCount int
	linkCount    int
	links        []Link
	//factories map[int]Factory
	myFactories      []Factory
	oppFactories     []Factory
	neutralFactories []Factory
	myTroops         []Troop
	oppTroops        []Troop
}

//is there a link between f1 and f2?
func (g gameMap) linkTo(f1, f2 Factory) bool {
	for _, l := range g.links {
		if l.from == f1.id && l.to == f2.id {
			return true
		}
	}
	return false
}
func (g gameMap) facWithMaxCyb() int {
	var max, id int
	for _, f := range g.myFactories {
		if f.cyborgs > max {
			max = f.cyborgs
			id = f.id
		}
	}
	return id
}

//should g be passed as a pointer here? No need of & thereafter??
func (g *gameMap) readEntity() {
	// entityCount: the number of entities (e.g. factories and troops)
	var entityCount int
	fmt.Scan(&entityCount)

	for i := 0; i < entityCount; i++ {
		var entityId int
		var entityType string
		var arg1, arg2, arg3, arg4, arg5 int
		fmt.Scan(&entityId, &entityType, &arg1, &arg2, &arg3, &arg4, &arg5)

		if entityType == "FACTORY" && arg1 == 0 {
			g.neutralFactories = append(g.neutralFactories, Factory{entityId, arg1, arg2, arg3})
		} else if entityType == "FACTORY" && arg1 == 1 {
			g.myFactories = append(g.myFactories, Factory{entityId, arg1, arg2, arg3})
		} else if entityType == "FACTORY" && arg1 == -1 {
			g.oppFactories = append(g.oppFactories, Factory{entityId, arg1, arg2, arg3})
		}

		if entityType == "TROOP" && arg1 == 1 {
			g.myTroops = append(g.myTroops, Troop{entityId, arg1, arg2, arg3, arg4, arg5})
		} else if entityType == "TROOP" && arg1 == -1 {
			g.oppTroops = append(g.oppTroops, Troop{entityId, arg1, arg2, arg3, arg4, arg5})
		}
	}
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
	//var commands []string

	for {
		board.readEntity()

		//ALGO to get out of woods: take each of my fac with troops and move to neutral fac first
		//and then those of opp with less cyb?
		//chooseSource()
		if len(board.myFactories) == 1 {
			//  src=board.myFactories[0]
		} else {
			//choose the one with max cyborgs (arg2)
		}

		//choose destination
		//printMove and queue it!! for one or two sources?

		//send commands

		//LOGS
		log.Println(board.neutralFactories)

		// Any valid action, such as "WAIT" or "MOVE source destination cyborgs"
		fmt.Println("WAIT")
	}
}
