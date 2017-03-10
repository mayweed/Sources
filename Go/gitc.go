package main

import (
	"fmt"
	"log"
)

type factory struct {
	id         int
	cyborgs    int
	production int
	owner      int
}
type troop struct {
	id             int
	from           int
	to             int
	cyborgs        int
	remainingTurns int
	owner          int
}
type graph struct {
	factoryCount int
	linkCount    int
	//a int from a slice of ints(factory2 + distance)
	edges map[int][][]int
}

//IDEA: develop a game state + a tree of possible moves
type player struct {
	id        int
	factories []factory
	troops    []troop
	score     int
	cybTroop  int
	lastMove  move
}
type move struct {
	from    factory
	to      factory
	cyborgs int
}
type gameState struct {
	//map of the game
	network graph
	//opponent
	opponent player
	//neutral
	neutralFactories []factory
	//game
	numOfTurns int
	//moves
	possibleMoves []move
}

//COMMAND HELPER should implement WAIT especially
//at endgame if I have 0 cybs in my last fact should
//wait til endgame. //+chain command with ;
func mv(from, to, cyb int) string {
	s := fmt.Sprintf("MOVE %d %d %d\n", from, to, cyb)
	return s
}

//helper func
func amIowner(f factory) bool {
	if f.owner == 1 {
		return true
	} else {
		return false
	}
}

//GRAPH METHODS:should be attached either to player or to gamestate!!
func (g gameState) getFactory(id int) factory {
	for _, fac := range g.opponent.factories {
		if fac.id == id {
			return fac
		}
	}
	for _, fact := range g.neutralFactories {
		if fact.id == id {
			return fact
		}
	}
	return factory{}
}

//zero factory
func (g gameState) zeroFactory() []factory {
	//yield all opp + neutral fact with 1 or less cyborgs
	//TODO: ordered them by production rate!!
	var fact []factory
	for _, v := range g.opponent.factories {
		if v.cyborgs <= 1 {
			fact = append(fact, v)
		}
	}
	for _, v := range g.neutralFactories {
		if v.cyborgs <= 1 {
			fact = append(fact, v)
		}
	}
	return fact
}

//should REWORK that...
func (p player) pickSourceFactory(num int, lastStart factory) factory {
	var startNode factory
	//should select the one with maxnodes?
	for _, factory := range p.factories {
		if factory.cyborgs >= num && factory.id != lastStart.id {
			startNode = factory
		}
	}
	return startNode
}

//return the id of the nearest node of a given factory
//and a factory I own if possible
func (g gameState) pickMinNode(f factory) factory {
	var minDist = 20
	var id int
	fact := g.getFactory(f.id)
	for _, v := range g.network.edges[f.id] {
		if fact.owner == 1 {
			if v[1] < minDist {
				minDist = v[1]
				id = v[0]
			}
		}
	}
	x := g.getFactory(id)
	return x
}

//Should use sort Interface to sort fact struc by production!!
func (g gameState) maxProdFactory() factory {
	var maxFact factory
	var max = 0
	for _, fact := range g.opponent.factories {
		if fact.production >= max {
			maxFact = fact
			max = fact.production
		}
	}
	for _, fac := range g.neutralFactories {
		if fac.production >= max {
			maxFact = fac
			max = fac.production
		}
	}
	return maxFact
}

//should I pas player as arg here?
func (g gameState) pickDestFactory(startNode factory) factory {
	var maxP = g.maxProdFactory()
	log.Println(maxP)
	var minD = g.pickMinNode(startNode)
	log.Println(minD)
	if ok := amIowner(maxP); !ok {
		return maxP
	} else if ok := amIowner(minD); !ok {
		return minD
	} else {
		for _, v := range g.opponent.factories {
			//should take the one with the least cyb and the highest prodrate!!
			if v.id != startNode.id && v.owner != 1 {
				return v
			}
		}
	}
	return factory{}
}

/* TODO
//should check a factory node
//should have a way to assess score
//a string like "WAIT" it's just checking next turn
//to see if last fact remaining will be overtaken
//in that case "WAIT" instead of Can't send a troop from a factory you don't control (0)
func (g gameState) checkEndGame() string{
    if len(facts)==1{
TODO    */

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

	eval := gameState{
		network: network,
		opponent: player{
			id:        -1,
			factories: []factory{},
		},
		neutralFactories: []factory{},
	}

	for i := 0; i < linkCount; i++ {
		var factory1, factory2, distance int
		fmt.Scan(&factory1, &factory2, &distance)
		//not directed SHOULD BE al list of lists... to handle multiple edges
		//or list of edges
		network.edges[factory1] = append(network.edges[factory1], []int{factory2, distance})
		network.edges[factory2] = append(network.edges[factory2], []int{factory1, distance})
	}

	//player
	var me = player{
		id:        1,
		factories: []factory{},
	}

	for {
		// entityCount: the number of entities (e.g. factories and troops)
		var entityCount int
		fmt.Scan(&entityCount)

		eval.numOfTurns += 1

		//use sort Interface to sort factory by production, highest first
		for i := 0; i < entityCount; i++ {
			var entityId int
			var entityType string
			var arg1, arg2, arg3, arg4, arg5 int
			fmt.Scan(&entityId, &entityType, &arg1, &arg2, &arg3, &arg4, &arg5)
			switch entityType {
			case "FACTORY":
				if arg1 == me.id {
					fac := factory{entityId, arg2, arg3, arg1}
					me.factories = append(me.factories, fac)
				} else if arg1 == eval.opponent.id {
					fac := factory{entityId, arg2, arg3, arg1}
					eval.opponent.factories = append(eval.opponent.factories, fac)
				} else if arg1 == 0 {
					fac := factory{entityId, arg2, arg3, arg1}
					eval.neutralFactories = append(eval.neutralFactories, fac)
				}
			case "TROOP":
				if arg1 == me.id {
					t := troop{entityId, arg2, arg3, arg4, arg5, arg1}
					me.troops = append(me.troops, t)
				} else if arg1 == eval.opponent.id {
					t := troop{entityId, arg2, arg3, arg4, arg5, arg1}
					eval.opponent.troops = append(eval.opponent.troops, t)
				}
			}
		}
		var s string
		//should write a func to compute the best num of cyb to send+ choose my startNode wrt that..
		//Rule: you overtake a factory if you send more cyb than it will have when you arrived.
		//EX: send one cyb to *all* nodes with cyb==0
		var num = 3

		//problem here with the second arg. Should keep the second arg cf eval for last stat?
		var startNode = me.pickSourceFactory(num, factory{})
		var dest = eval.pickDestFactory(startNode)

		//should modify mv to chain commands with ;
		s = mv(startNode.id, dest.id, num)
		me.lastMove = move{startNode, dest, num}
		//This one happens too: Can't send a troop from a factory you don't control (3)
		//ex: try to send last cyb from a node the foe will capture next turn
		fmt.Printf("%s", s)

		eval.numOfTurns += 1

		//SCORE
		for _, v := range me.factories {
			me.score += v.cyborgs
		}
		//troop??
		for _, v := range me.troops {
			me.cybTroop += v.cyborgs
		}
		log.Println(me.score, me.cybTroop)

		me.factories = []factory{}
		eval.opponent.factories = []factory{}
		eval.neutralFactories = []factory{}
		me.score = 0
		me.cybTroop = 0
	}

}
