package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

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

//Action+Turn should be mixed together??
//Then a func that create a moves slice string
//and send it. And GET RID of actionType
type Action struct {
	actionType  string
	from        int
	to          int
	cyborgCount int
}

//should be append to moves in Turn should add ';'
func (a Action) printAction() string {
	var s string
	if a.actionType == "move" {
		s = fmt.Sprintf("MOVE %d %d %d", a.from, a.to, a.cyborgCount)
	} else {
		//if no move, just wait?
		s = fmt.Sprintf("WAIT")
	}
	return s
}

type Turn struct {
	//encapsulate Action in string (sprintf)
	moves []string
}

func (t Turn) sendCommands() {
	cmd := "WAIT"
	if len(t.moves) == 0 {
		//log.Println("List of commands is empty, WAIT will be sent")
	} else {
		cmd = strings.Join(t.moves, ";")
	}
	fmt.Println(cmd)
}

type Player struct {
	factories []Factory
	troops    []Troop
	turn      Turn
	score     int
}

func (p Player) facWithMaxCyb() int {
	var max, id int
	for _, f := range p.factories {
		if f.cyborgs > max {
			max = f.cyborgs
			id = f.id
		}
	}
	return id
}
func (p *Player) calcScore(){
    var total int
    for _,f := range p.factories{
        total+=f.cyborgs
    }
    for _,t := range p.troops{
        total+=t.cyborgs
    }
    p.score=total
}

type State struct {
	factoryCount     int
	linkCount        int
	links            []Link
	neutralFactories []Factory
	me               Player
	opp              Player
}
//This is ugly can't i put that in an anon func i pass on?
func (s *State) sortLinksByDist(){
    sort.Slice(s.links, func(i, j int) bool { return s.links[i].distance< s.links[j].distance})
}
func (s *State) sortFacByProd(){
    sort.Slice(s.neutralFactories, func(i, j int) bool { return s.neutralFactories[i].production< s.neutralFactories[j].production})
    sort.Slice(s.opp.factories, func(i, j int) bool { return s.opp.factories[i].production< s.opp.factories[j].production})
    sort.Slice(s.me.factories, func(i, j int) bool { return s.me.factories[i].production< s.me.factories[j].production})
}
func (s *State) readMap() {
	// factoryCount: the number of factories
	var factoryCount int
	fmt.Scan(&factoryCount)
	s.factoryCount = factoryCount

	// linkCount: the number of links between factories
	var linkCount int
	fmt.Scan(&linkCount)
	s.linkCount = linkCount

	for i := 0; i < linkCount; i++ {
		var factory1, factory2, distance int
		fmt.Scan(&factory1, &factory2, &distance)
		s.links = append(s.links, Link{factory1, factory2, distance})
	}
	s.sortLinksByDist()
}

func (s *State) readEntity() {
	// entityCount: the number of entities (e.g. factories and troops)
	var entityCount int
	fmt.Scan(&entityCount)

	for i := 0; i < entityCount; i++ {
		var entityId int
		var entityType string
		var arg1, arg2, arg3, arg4, arg5 int
		fmt.Scan(&entityId, &entityType, &arg1, &arg2, &arg3, &arg4, &arg5)

		if entityType == "FACTORY" && arg1 == 0 {
			s.neutralFactories = append(s.neutralFactories, Factory{entityId, arg1, arg2, arg3})
		} else if entityType == "FACTORY" && arg1 == 1 {
			s.me.factories = append(s.me.factories, Factory{entityId, arg1, arg2, arg3})
		} else if entityType == "FACTORY" && arg1 == -1 {
			s.opp.factories = append(s.opp.factories, Factory{entityId, arg1, arg2, arg3})
		}

		if entityType == "TROOP" && arg1 == 1 {
			s.me.troops = append(s.me.troops, Troop{entityId, arg1, arg2, arg3, arg4, arg5})
		} else if entityType == "TROOP" && arg1 == -1 {
			s.opp.troops = append(s.opp.troops, Troop{entityId, arg1, arg2, arg3, arg4, arg5})
		}
	}
	s.me.calcScore()
	s.opp.calcScore()
	s.sortFacByProd()
	
	log.Println(s.me.factories)
}

func (s *State) clearState() {
	s.me.factories = []Factory{}
	s.me.troops = []Troop{}
	s.me.turn.moves = []string{}
	s.neutralFactories = []Factory{}
	s.opp.factories = []Factory{}
}

//is there a link between f1 and f2?
func (s State) linkTo(f1, f2 Factory) bool {
	for _, l := range s.links {
		if l.from == f1.id && l.to == f2.id {
			return true
		}
	}
	return false
}

//ALGO to get out of woods: take each of my fac with troops and move to neutral fac first
//and then those of opp with less cyb?
//could list all possible acttions and choose first those factories with
//highest prod rate?
//must define a better cyb count
//must take the time to build factories. Sending all those cybs wear out
//factories...
func (s *State) think() {
	for _, src := range s.me.factories {
		//lame and does not work
		if src.cyborgs < 5 {
			continue
		}
		if len(s.neutralFactories) != 0 {
			for _, dest := range s.neutralFactories {
				s.me.turn.moves = append(s.me.turn.moves, Action{"move", src.id, dest.id, 1}.printAction())
			}
		} else {
			for _, dest := range s.opp.factories {
				s.me.turn.moves = append(s.me.turn.moves, Action{"move", src.id, dest.id, 1}.printAction())
			}
		}
		//log.Println(s.me.score)
	}
}

func main() {
	board := State{}
	board.readMap()
	for {
		board.readEntity()
		board.think()
		board.me.turn.sendCommands()
		board.clearState()
	}
}
