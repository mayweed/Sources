package main

import (
	"fmt"
	"log"
	"sort"
)

//MAP
type graph struct {
	factoryCount int
	linkCount    int
	//a int from a slice of ints(factory2 + distance)
	edges map[int][][]int
}
type factory struct {
	id         int
	cyborgs    int
	production int
	owner      int
}

//helper func
func (f factory) amIowner() bool {
	if f.owner == 1 {
		return true
	} else {
		return false
	}
}

//sort interface
type byProd []factory

func (b byProd) Len() int           { return len(b) }
func (b byProd) Less(i, j int) bool { return b[i].production < b[j].production }
func (b byProd) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type troop struct {
	id             int
	from           int
	to             int
	cyborgs        int
	remainingTurns int
	owner          int
}

//PLAYER
type player struct {
	id          int
	factories   []factory
	troops      []troop
	score       int
	lastMove    move
	currentMove move
}

func (p *player) countScore() {
	p.score = 0
	for _, v := range p.factories {
		p.score += v.cyborgs
	}
	for _, v := range p.troops {
		p.score += v.cyborgs
	}
}
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

//GAMESTATE
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
	sort.Sort(byProd(g.opponent.factories))
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

//MOVE
type move struct {
	from    factory
	to      factory
	cyborgs int
}

//COMMAND HELPER should implement WAIT especially
//at endgame if I have 0 cybs in my last fact should
//wait til endgame. //+chain command with ;
//should be attach to a move object
func mv(from, to, cyb int) string {
	//here should be m.from etc...in the end
	s := fmt.Sprintf("MOVE %d %d %d\n", from, to, cyb)
	return s
}
func (m move) String() string {
	s := fmt.Sprintf("MOVE %d to %d with %d cyborgs\n", m.from, m.to, m.cyborgs)
	return s
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
func (g gameState) pickDestFactory(startNode factory) factory {
	//var cybToSend int
	var maxP = g.maxProdFactory()
	var minD = g.pickMinNode(startNode)

	/*//test
	for _,v := range g.opponent.troops{
	    if v.id==maxP.id{
	        cybToSend=v.cyborgs+1
	    }
	}*/

	if ok := maxP.amIowner(); !ok {
		return maxP
	} else if ok := minD.amIowner(); !ok {
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
func (g gameState) checkEndGame() bool{
    if len(facts)==1{
func (g gameState) selectMove() move{
   for _,v := range g.possibleMove{
       ...
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

		sort.Sort(byProd(eval.opponent.factories))
		log.Println(eval.opponent.factories)
		//problem here with the second arg. Should keep the second arg cf eval for last stat?
		var startNode = me.pickSourceFactory(num, factory{})
		var dest = eval.pickDestFactory(startNode)
		me.currentMove = move{startNode, dest, num}
		//should modify mv to chain commands with ;
		s = mv(startNode.id, dest.id, num)
		//This one happens too: Can't send a troop from a factory you don't control (3)
		//ex: try to send last cyb from a node the foe will capture next turn
		fmt.Printf("%s", s)

		eval.numOfTurns += 1

		(&me).countScore()
		(&eval.opponent).countScore()
		log.Println(me.score, eval.opponent.score, eval.numOfTurns, me.currentMove, me.lastMove)

		//THIS is a must, without that keeps appending to the same list
		//again and again..
		me.factories = []factory{}
		me.troops = []troop{}
		eval.opponent.factories = []factory{}
		eval.opponent.troops = []troop{}
		eval.neutralFactories = []factory{}
		me.lastMove = me.currentMove
	}

}
