package main

import (
	"fmt"
	"log"
)

type graph struct {
	factoryCount int
	linkCount    int
	//a int from a slice of ints(factory2 + distance)
	edges     map[int][][]int
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

type troop struct {
	from           int
	to             int
	cyborgs        int
	remainingTurns int
	owner          int
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

//oki does not work
func (g graph) countScore() (x, y int) {
	var myScore = 0
	var oppScore = 0
	for _, v := range g.factories {
		switch v.owner {
		case 1:
			myScore += v.cyborgs
		case -1:
			oppScore += v.cyborgs
		}
	}
	return myScore, oppScore
}

func (g graph) pickAnotherFactory(queue []factory, doneQueue []int, num int, lastStart factory) (q1 []factory, q2 []int, node factory) {
	//oki so: take factory from a done queue check cyborgs num
	//and promote it to start node
	var startNode factory
	//should select the one with maxnodes?
	var min = num * 2
	for _, factory := range g.factories {
		//should own the factory
		if factory.owner == 1 {
			if factory.id == doneQueue[0] {
				//num == number of sent cyborgs
				if factory.cyborgs >= min && factory.owner == 1 {
					startNode = factory
				}
			}
		}
	}
	//in that case should clean queue of the new start node
	for _, v := range queue {
		if v.id == startNode.id {
			continue
		} else {
			queue = append(queue, v)
		}
	}
	return queue, doneQueue, startNode
}

//return the id of the nearest node of a given factory
//and a factory I own if possible
//BETTER:order those slices in an id slice with nearest first
func (g graph) pickMinNode(f factory) int {
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
		//log.Println(idSlice)
	}
	return id
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
	var sendQueue []int

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

		//BFS like non? should i use myFactories or queue?
		for len(myFactories) > 0 {
			var startNode = myFactories[0]
			//log.Println("STARTNODE FIRST",startNode.id,"NUM",startNode.cyborgs)
			log.Println(myFactories, oppFactories)

			var s string
			var num = 3

			queue = network.getFactQueue(startNode)

			//WHY in the hell use an int here?
			dest := queue[0].id
			//c trop nul ça...tss
			//if startNode.id==dest{dest=queue[1].id}
			//should clean queue instead...
			if startNode.cyborgs > num {
				if startNode.id == dest {
					dest = queue[1].id
				}
				s = mv(startNode.id, dest, num)
				sendQueue = append(sendQueue, dest)
			} else if startNode.cyborgs <= num {
				//the other way round: bfs? neighboring nodes of startNode?
				queue, sendQueue, startNode = network.pickAnotherFactory(queue, sendQueue, num, startNode)
				if startNode.id == dest {
					dest = network.pickMinNode(startNode)
				}
				s = mv(startNode.id, dest, num)
				log.Println("I was here", startNode.id)
			}
			//oki new errors:
			//Can't send a troop to the factory it is issued from (0)
			//OOPS!! src and dest MUST be different!!
			fmt.Printf("%s", s)
			//log.Println(network.countScore())
			//put nodes at end
			//queue=append(queue[1:],dest)
			//WHY TWO QUEUES??
			queue = queue[1:]
			myFactories = myFactories[1:]
		}
		//fmt.Printf("WAIT\n")
	}
}
