package main

import (
	"fmt"
	"log"
)

type graph struct {
	factoryCount int
	linkCount    int
	//a slice of int slices(dest + distance)
	edges map[int][][]int

	factories []factory
	troops    []troop
}
type factory struct {
	id         int
	cyborgs    int
	production int
	//arg1 to know
	ownership int
}

type troop struct {
	from           int
	to             int
	cyborgs        int
	remainingTurns int
	ownership      int
}

func mv(from, to, cyb int) string {
	s := fmt.Sprintf("MOVE %d %d %d\n", from, to, cyb)
	return s
}

/*
func (g graph) sendTroop()(from,to, numCyb int){
    for _,v := range g.factories{
        if v.ownership==1{
*/
/*
strat: check my factories, check opponent fact, divide my
num of cyb/opp fact,send troops to it
Need a real game state with lots of intell
*/

func main() {
	// factoryCount: the number of factories
	var factoryCount int
	fmt.Scan(&factoryCount)

	// linkCount: the number of links between factories
	var linkCount int
	fmt.Scan(&linkCount)

	network := graph{
		factoryCount: factoryCount,
		linkCount:    linkCount,
		edges:        make(map[int][][]int),
	}

	for i := 0; i < linkCount; i++ {
		var factory1, factory2, distance int
		fmt.Scan(&factory1, &factory2, &distance)
		//not directed, should i use a edge struct?
		network.edges[factory1] = append(network.edges[factory1], []int{factory2, distance})
		network.edges[factory2] = append(network.edges[factory2], []int{factory1, distance})
	}
	for {
		// entityCount: the number of entities (e.g. factories and troops)
		var entityCount int
		fmt.Scan(&entityCount)

		var myFactories []factory
		var oppFactories []factory
		var neutralFactories []factory

		for i := 0; i < entityCount; i++ {
			var entityId int
			var entityType string
			var arg1, arg2, arg3, arg4, arg5 int
			fmt.Scan(&entityId, &entityType, &arg1, &arg2, &arg3, &arg4, &arg5)
			switch entityType {
			case "FACTORY":
				network.factories = append(network.factories, factory{entityId, arg2, arg3, arg1})
				if arg1 == 1 {
					fac := factory{entityId, arg2, arg3, arg1}
					myFactories = append(myFactories, fac)
				} else if arg1 == -1 {
					fac := factory{entityId, arg2, arg3, arg1}
					oppFactories = append(oppFactories, fac)
				} else if arg1 == 0 {
					fac := factory{entityId, arg2, arg3, arg1}
					neutralFactories = append(neutralFactories, fac)
				}
			case "TROOP":
				network.troops = append(network.troops, troop{arg2, arg3, arg4, arg5, arg1})
			}
		}
		var startNode = myFactories[0]
		myFactories = myFactories[1:]

		//DOES NOT WORK AS INTENDED especially:
		//always sending to the same node!!
		// + should use WAIT!!

		//keep track of where I previously sent troops
		var lastSendTroopsNode = startNode.id

		//should put that in a list: never send twice to the same factory??
		//except maybe to take it?
		var min = network.pickMinNode(startNode)
		var s string

		if min == lastSendTroopsNode {
			//pick anothernode,let's say a node with 0 cyb and neutral
			s = mv(startNode.id, neutralFactories[0].id, 2)
			//TEST!!
			fmt.Println("Bind!!")
			lastSendTroopsNode = neutralFactories[0].id
		} else {
			s = mv(startNode.id, min, 2)
			lastSendTroopsNode = min
		}

		//log.Println(min,lastSendTroopsNode,neutralFactories[0].id,network.edges)
		// Any valid action, such as "WAIT" or "MOVE source destination cyborgs"
		fmt.Printf("%s", s)
	}
}
