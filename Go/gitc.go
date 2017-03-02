package main

import (
	"fmt"
	"log"
)

type graph struct {
	factoryCount int
	linkCount    int
	//a int from a slice of ints(factory2 + distance)
	edges map[int][][]int

	factories []factory
	troops    []troop
}

type factory struct {
	id         int
	cyborgs    int
	production int
	owner      int
}

//helper func
func amIowner(f factory) bool {
	if f.owner == 1 {
		return true
	} else {
		return false
	}
}

//slice thing does not work
func (g graph) maxProdFactory(f []factory) (factory, []factory) {
	var maxFact factory
	//cf go blog slice
	//newSlice := make([]int, len(slice), 2*cap(slice))
	var maxFacts = make([]factory, len(g.factories))
	var max = 0
	for _, fact := range f {
		if fact.production >= max {
			maxFact = fact
			max = fact.production
			//I grow the slice by one elt (cf cap)
			//maxFacts=maxFacts[0:len(maxFacts)+1]
			//make room at the front
			copy(maxFacts[1:], maxFacts[:len(maxFacts)-1])
			//assign front elt
			maxFacts[0] = fact
		} else if fact.production <= max {
			maxFacts = append(maxFacts, fact)
		}
	}
	return maxFact, maxFacts
}

type troop struct {
	from           int
	to             int
	cyborgs        int
	remainingTurns int
	owner          int
}

//OKI this is a test never done that before
//IDEA: develop a game state + a tree of possible moves
type gameState struct {
	//our sources
	myCurrentBase factory
	oppBase       factory
	//our destinations
	myLastDest  factory
	oppLastDest factory
	//our troops
	myTroops  []troop
	oppTroops []troop

	numOfTurns int
	score      int

	possibleMoves []move
}
type move struct {
	from    factory
	to      factory
	cyborgs int
}

//GRAPH METHODS
func (g graph) getFactory(id int) factory {
	for _, fac := range g.factories {
		if fac.id == id {
			return fac
		}
	}
	return factory{}
}

func (g graph) getFactQueue(startNode factory) []factory {
	var queue []factory
	for k, _ := range g.edges {
		if k != startNode.id {
			queue = append(queue, g.getFactory(k))
		}
	}
	return queue
}

func (g graph) pickSourceFactory(queue []factory, num int, lastStart factory) (node factory) {
	var startNode factory
	//should select the one with maxnodes?
	for _, factory := range g.factories {
		//should own the factory
		if factory.owner == 1 {
			if factory.cyborgs >= num && factory.id != lastStart.id {
				startNode = factory
			}
		}
	}
	return startNode
}

//func (g graph) pickDestFactory(facts []factory,startNode factory) factory{
//First minNode + maxProd of opponent?
//for _,v := range (g.factories){

//return the id of the nearest node of a given factory
//and a factory I own if possible
func (g graph) pickMinNode(f factory) factory {
	var minDist = 20
	var id int
	var idSlice []int
	fact := g.getFactory(f.id)
	for _, v := range g.edges[f.id] {
		if fact.owner == 1 {
			if v[1] < minDist {
				minDist = v[1]
				id = v[0]
				idSlice = append(idSlice, id)
			}
		}
	}
	f := g.getFactory(id)
	return f
}

//SCORE
func (g graph) baseScore() (x, y int) {
	var myScore = 0
	var oppScore = 0
	for _, v := range g.factories {
		switch v.owner {
		case 1:
			myScore = v.cyborgs
		case -1:
			oppScore = v.cyborgs
		}
	}
	return myScore, oppScore
}

//oki does not work
func (g graph) countTroops(myScore, oppScore *int) {
	for _, troop := range g.troops {
		switch troop.owner {
		case 1:
			*(myScore) = troop.cyborgs
		case -1:
			*(oppScore) = troop.cyborgs
		}
	}
}

//COMMAND HELPER
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
	eval := gameState{}

	for i := 0; i < linkCount; i++ {
		var factory1, factory2, distance int
		fmt.Scan(&factory1, &factory2, &distance)
		//not directed SHOULD BE al list of lists... to handle multiple edges
		//or list of edges
		network.edges[factory1] = append(network.edges[factory1], []int{factory2, distance})
		network.edges[factory2] = append(network.edges[factory2], []int{factory1, distance})
	}

	//enqueue nodes
	//so two queues : one with dest factory and the second one with send
	var queue []factory
	var myScore int
	var oppScore int

	for {
		// entityCount: the number of entities (e.g. factories and troops)
		var entityCount int
		fmt.Scan(&entityCount)

		var myFactories []factory
		var oppFactories []factory
		var neutralFactories []factory

		myScore, oppScore = network.baseScore()
		//test count troops
		var Score = 0
		var opp = 0

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

		m, n := network.maxProdFactory(oppFactories)
		log.Println("MAXPROD", m, n)

		//BFS like non? should i use myFactories or queue?
		//Think: myFactories those are my base start. Should attack oppFactories
		for len(myFactories) > 0 {
			var s string
			var num = 3
			var startNode = network.pickSourceFactory(myFactories, num, factory{})
			eval.myCurrentBase = startNode
			//log.Println(myFactories,oppFactories)
			queue = network.getFactQueue(startNode)
			log.Println(queue)

			//WHY in the hell use an int here?
			//HERE I should pick either a neutral or an opponent fact!!
			dest := queue[0].id

			///THIS IS checked in pick source??
			if startNode.cyborgs >= num {
				s = mv(startNode.id, dest, num)
			} else if startNode.cyborgs < num {
				startNode = network.pickSourceFactory(myFactories, num, startNode)
				s = mv(startNode.id, dest, num)
				log.Println("I was here", startNode.id)
			}
			//oki new errors:
			//Can't send a troop to the factory it is issued from (0)
			//OOPS!! src and dest MUST be different!!
			//This one happens too: Can't send a troop from a factory you don't control (3)
			fmt.Printf("%s", s)

			//WHY TWO QUEUES??
			queue = queue[1:]
			myFactories = myFactories[1:]
		}
		eval.numOfTurns += 1

		//SCORE, does not work...
		network.countTroops(&Score, &opp)
		//log.Println(Score,opp)
		//log.Println(myScore,oppScore)

	}
	//should be reset at the end of each turn
	myScore = 0
	oppScore = 0

}
