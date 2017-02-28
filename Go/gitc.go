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
		//IDEA: take all v[0] of my startNode put theme in a queue
		//dequeue and send packets...
		//var min =network.pickMinNode(startNode)
		var s string
		var num = 3

		//TODO factorize plz
		for k, _ := range network.edges {
			if k != startNode.id {
				queue = append(queue, k)
			}
		}

		dest := queue[0]
		//oki new errors:
		//Can't send a troop to the factory it is issued from (0)

		if startNode.cyborgs > num {
			s = mv(startNode.id, dest, num)
		} else {
			//pick another starting node
			//should check that it has cyborgs?
			//the other way round: bfs? neighboring nodes of startNode?
			//check that the nearest node has cyborgs and use it as new base?
			//SHOULD BE ABLE TO REBASE
			//SHOULD KEEP TRACK OF CYBIORGS NUM with production rate. Should not send
			//more cyborgs than produce?
			startNode = queue[len(queue)]
		}
		//log.Println(min,lastSendTroopsNode,neutralFactories[0].id,network.edges)
		// Any valid action, such as "WAIT" or "MOVE source destination cyborgs"
		//fmt.Sprintf("MOVE" startNode network.edges[startNode][0] prod)
		fmt.Printf("%s", s)

		//put nodes at end
		queue = append(queue[1:], dest)
	}
}
