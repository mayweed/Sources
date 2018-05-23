package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

//FACTORY
type Factory struct {
	id         int
	//1 for me, -1 for opp, 0 neutral
	owner      int
	cyborgs    int
	production int
	//first int is id 
	distance map[int]Link
}
type Link struct{
    from int
    to int
    distance int
}
type gameMap struct{
    factories map[int]*Factory
    links []Link
    troops []Troop
    me Player
    opp Player
}
func (g gameMap) initPlayer(p Player) {
    for _,fac := range g.factories{
        if fac.owner == p.id {
		    p.factories = append(p.factories,fac)
        }
    }
}
func (g gameMap) yieldNeutralFac() []Factory{
    var neutral []Factory
    for _,fac := range g.factories{					
	    if fac.owner == 0 {
	        neutral = append(neutral,*fac )
	    }
    }
    return neutral
}
    
func (f *Factory) updateFactory(fac Factory){
    f.owner=fac.owner
	f.cyborgs=fac.cyborgs
	f.production=fac.production
}
//given a factory, how many cyb will it have in X turns
//SHOULD CHECK INCOMING/OUTCOMING TROOPS!!
func (f Factory) cybNextTurn(numTurns int) int{
    var numStart=f.cyborgs
    for i:=0;i<numTurns;i++{
       numStart+=f.production
    }
    return numStart
}
//sort interface
type byProd []Factory
func(b byProd) Len()int{return len(b)}
func(b byProd) Less(i,j int) bool{return b[i].production<b[j].production}
func (b byProd) Swap(i,j int){b[i],b[j]=b[j],b[i]}

//TROOP
type Troop struct {
	id             int
	//1 for me, -1 for opp
	owner          int
	from           int
	to             int
	cyborgs        int
	remainingTurns int

}

//PLAYER
type Player struct {
	id        int
	factories []Factory
	troops    []Troop
	score     int
}
func (p *Player) countScore(){
    p.score=0
	for _,v :=range p.factories{
	    p.score+=v.cyborgs
	}
	for _,v :=range p.troops{
	    p.score+=v.cyborgs
	}
}
//WHY num??
func (p Player) pickSourceFactory(num int, lastStart Factory) Factory {
	var startNode Factory
	//should select the one with maxnodes?
	for _, factory := range p.factories {
		if factory.cyborgs >= num && factory.id != lastStart.id {
			startNode = factory
		}
	}
	return startNode
}


//yield all opp + neutral fact with 1 or less cyborgs
func (g gameMap) zeroFactory() []Factory{
    var fact []Factory
    sort.Sort(byProd(g.opp.factories))
    for _,v := range g.opp.factories{
        if v.cyborgs <= 1{
            fact=append(fact,v)
        }
    }
    for _,v := range g.neutralFactories{
        if v.cyborgs <=1{
            fact = append(fact,v)
        }
    }
    return fact
}
//return the id of the nearest node of a given factory
//and a factory I own if possible
//func (g gameState) pickMinNode(f Factory) Factory {
//	var minDist = 20
//	var id int
//	fact := g.getFactory(f.id)
//	for _, v := range g.network.edges[f.id] {
//		if fact.owner == 1 {
//			if v[1] < minDist {
//				minDist = v[1]
//				id = v[0]
//			}
//		}
//	}
//	x := g.getFactory(id)
//	return x
//}
//Should use sort Interface to sort fact struc by production!!
//func (g gameMap) maxProdFactory() Factory {
//	var maxFact Factory
//	var max = 0
//	for _, fact := range g.opp.factories {
//		if fact.production >= max {
//			maxFact = fact
//			max = fact.production
//		}
//	}
//	for _, fac := range g.neutralFactories {
		//if fac.production >= max {
		//	maxFact = fac
		//	max = fac.production
		//}
	//}
	//return maxFact
//}
//here should return a list of factories...
//func (g gameState) pickDestFactory(startNode Factory) Factory {
    //var cybToSend int
//	var maxP = g.maxProdFactory()
//	var minD = g.pickMinNode(startNode)
//	if ok := maxP.amIowner(); !ok {
//		return maxP
//	} else if ok := minD.amIowner(); !ok {
//		return minD
//	} else {
//		for _, v := range g.opp.factories {
//		    //should take the one with the least cyb and the highest prodrate!!
//			if v.id != startNode.id && v.owner != 1 {
//				return v
//			}
//		}
//	}
//	return Factory{}
//}

//COMMANDS 
func mv(from, to, cyb int) string {
	s := fmt.Sprintf("MOVE %d %d %d\n", from, to, cyb)
	return s
}
//format multiple orders
func enqueueJoin (moveOrders []string) string{
    var s []string
    for _, order := range moveOrders{
         s=append(s,order)
    }
    return strings.Join(s,";")
}

//MAIN
func main() {
    
	me := Player{id:1, factories:[]Factory{}}
	opp:=Player{id: -1, factories: []Factory{}}
    board := gameMap{factories:make(map[int]*Factory),links:[]Link,troops:[]Troop}
    
	// factoryCount: the number of factories
	var factoryCount int
	fmt.Scan(&factoryCount)

	// linkCount: the number of links between factories
	var linkCount int
	fmt.Scan(&linkCount)

	//move orders in a queue?
	//moveOrders := []string
	
	for i := 0; i < linkCount; i++ {
	    var factory1, factory2, distance int
		fmt.Scan(&factory1, &factory2, &distance)
		log.Println(factory1, factory2, distance)
        board.links=append(board.links,Link{from:factory1,to:factory2,distance:distance})
	}

	for {
		// entityCount: the number of entities (e.g. factories and troops)
		var entityCount int
		fmt.Scan(&entityCount)

        //use sort Interface to sort factory by production, highest first
		for i := 0; i < entityCount; i++ {
			var entityId int
			var entityType string
			var arg1, arg2, arg3, arg4, arg5 int
			fmt.Scan(&entityId, &entityType, &arg1, &arg2, &arg3, &arg4, &arg5)
			switch entityType {
			case "FACTORY":
			    board.factories[entityId].updateFactory(Factory{owner:arg1,cyborgs:arg2,production:arg3})
			    board.initPlayer(me)
			    board.initPlayer(opp)
			case "TROOP":
			//same for troops
			    //board.factories[entityId].updateFactory(Factory{owner:arg1,cyborgs:arg2,production:arg3})
				if arg1 == me.id {
					me.troops = append(me.troops, Troop{id:entityId,owner:arg1,from:arg2, to:arg3, cyborgs:arg4, remainingTurns:arg5})
				} else if arg1 == eval.opp.id {
					opp.troops = append(opp.troops,Troop{id:entityId,owner:arg1,from:arg2, to:arg3, cyborgs:arg4, remainingTurns:arg5})
				}
			}
		}
		
		//you tied each link to the right fact?
		//for _,link :=range eval.links{
		    //if link.from == eval.factories.id{
		    //    eval.factories.links=append(eval.factories.links,link)
		    //}
		//}
		
		//LOGS
		//for id,factory := range eval.factories{
		//    log.Println(id,factory)
		//}
		//log.Println(eval.factories[1].distance[7])

		//NAIVE STRATEGY
		//Take all the nodes that are mine, calculate to send troops to all remaining nodes
		//write a findTarget func for each node of mine :)
		//ALGO:
		/*
		for _,myFac := range me.factories{
		    -check they are no sending troops already (arg2)
		    -if not find a suitable targets: either neutral or opp that I could afford
		    -queue the command
		   */
		 //targets []Factory
		 //for _,fac := range board.factories{
		     //it's mine
		     //if fac.owner==1{
		         //selectTarget() yet to write
		         //targets=append(targets,fac.selectTarget())

		//OLD CODE
		//var num = 3
		//sort.Sort(byProd(eval.neutralFactories))
		//var startNode = me.pickSourceFactory(num, Factory{})
		//var dest = eval.pickDestFactory(startNode)
        //me.currentMove=move{startNode,dest,num}
		//s := mv(startNode.id, dest.id, num)
		//fmt.Printf("%s", s)
		fmt.Printf("WAIT")


		eval.numOfTurns += 1
		(&me).countScore()
		(&eval.opp).countScore()

		//clear the buckets
		me.factories=[]Factory{}
		me.troops=[]Troop{}
		eval.opp.factories=[]Factory{}
		eval.opp.troops=[]Troop{}
		eval.neutralFactories=[]Factory{}
	    //me.lastMove = me.currentMove
	}
}
