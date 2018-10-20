package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	DAMAGE_DURATION = 5
)

type Link struct {
	from     int
	to       int
	distance int
}

type Factory struct {
	id             int
	owner          int
	unitCount      int
	productionRate int
	disabled       int
}
type Troop struct {
	id             int
	owner          int
	from           int
	to             int
	unitCount      int
	remainingTurns int
}
type Bomb struct {
	owner          int
	from           int
	to             int
	remainingTurns int
}

type Action struct {
	actionType  string
	from        int
	to          int
	cyborgCount int
	msg         string
}

type Player struct {
	factories []Factory
	troops    []Troop
	bombs     []Bomb
	action    Action
	turn      []string
	score     int
}

func (p Player) facWithMaxCyb() int {
	var max, id int
	for _, f := range p.factories {
		if f.unitCount > max {
			max = f.unitCount
			id = f.id
		}
	}
	return id
}
func (p *Player) getScore() {
	var total int
	for _, f := range p.factories {
		total += f.unitCount
	}
	for _, t := range p.troops {
		total += t.unitCount
	}
	p.score = total
}
func (p *Player) move(from, to, cyborgCount int) {
	p.action.actionType = "MOVE"
	p.action.from = from
	p.action.to = to
	p.action.cyborgCount = cyborgCount
	p.turn = append(p.turn, fmt.Sprintf("MOVE %d %d %d", p.action.from, p.action.to, p.action.cyborgCount))
}
func (p *Player) bomb(from, to int) {
	p.action.actionType = "BOMB"
	p.action.from = from
	p.action.to = to
	p.turn = append(p.turn, fmt.Sprintf("BOMB %d %d", p.action.from, p.action.to))
}
func (p *Player) wait() {
	p.action.actionType = "WAIT"
	p.turn = append(p.turn, fmt.Sprintf("WAIT"))
}
func (p *Player) msg(msg string) {
	p.action.actionType = "MSG"
	p.action.msg = msg
	p.turn = append(p.turn, fmt.Sprintf("MSG %s", p.action.msg))
}
func (p *Player) sendCommands() {
	if len(p.turn) == 0 {
		p.wait()
		p.msg("sending wait, turn empty")
	}
	cmd := strings.Join(p.turn, ";")
	fmt.Println(cmd)
}

type State struct {
	factoryCount     int
	linkCount        int
	links            []Link
	neutralFactories []Factory
	bombs            []Bomb
	turn             int
	me               Player
	opp              Player
}

//This is ugly can't i put that in an anon func i pass on?
func (s *State) sortLinksByDist() {
	sort.Slice(s.links, func(i, j int) bool { return s.links[i].distance < s.links[j].distance })
}
func (s *State) sortFacByProd() {
	sort.Slice(s.neutralFactories, func(i, j int) bool {
		return s.neutralFactories[i].productionRate < s.neutralFactories[j].productionRate
	})
	sort.Slice(s.opp.factories, func(i, j int) bool { return s.opp.factories[i].productionRate < s.opp.factories[j].productionRate })
	sort.Slice(s.me.factories, func(i, j int) bool { return s.me.factories[i].productionRate < s.me.factories[j].productionRate })
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
		switch entityType {
		case "FACTORY":
			switch arg1 {
			case 0:
				s.neutralFactories = append(s.neutralFactories, Factory{id: entityId, owner: arg1, unitCount: arg2, productionRate: arg3})
			case 1:
				s.me.factories = append(s.me.factories, Factory{id: entityId, owner: arg1, unitCount: arg2, productionRate: arg3})
			case -1:
				s.opp.factories = append(s.opp.factories, Factory{id: entityId, owner: arg1, unitCount: arg2, productionRate: arg3})
			}
		case "TROOP":
			switch arg1 {
			case 0:
				s.me.troops = append(s.me.troops, Troop{entityId, arg1, arg2, arg3, arg4, arg5})
			case -1:
				s.opp.troops = append(s.opp.troops, Troop{entityId, arg1, arg2, arg3, arg4, arg5})
			}
		case "BOMB":
			switch arg1 {
			case 1:
				s.me.bombs = append(s.me.bombs, Bomb{arg1, arg2, arg3, arg4})
			case -1:
				s.opp.bombs = append(s.opp.bombs, Bomb{arg1, arg2, arg3, arg4})
			}
		}
	}
	s.me.getScore()
	s.opp.getScore()
	s.sortFacByProd()
}

func (s *State) clearState() {
	s.me.factories = []Factory{}
	s.me.troops = []Troop{}
	s.me.turn = []string{}
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
	var needTroopsFactories []Factory
	for _, src := range s.me.factories {
		//I should send troops to the factories!!
		if src.unitCount < 5 {
			needTroopsFactories = append(needTroopsFactories, src)
			continue
		}
		if len(s.neutralFactories) != 0 {
			for _, dest := range s.neutralFactories {
				s.me.move(src.id, dest.id, 1)
			}
		} else {
			for _, dest := range s.opp.factories {
				s.me.move(src.id, dest.id, 1)
			}
		}
	}
	//test to help does not work no? should think strat?
	for _, dest := range needTroopsFactories {
		for _, src := range s.me.factories {
			if src.unitCount > 5 {
				s.me.move(src.id, dest.id, 1)
			}
		}
	}
}

func main() {
	board := State{}
	board.readMap()
	for {
		board.readEntity()
		board.think()
		board.me.sendCommands()
		board.clearState()
	}
}
