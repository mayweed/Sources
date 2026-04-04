package main

import (
	"fmt"
	"sort"
	"strings"
)

const DAMAGE_DURATION = 5

// ------------------------------------------------------------
// DATA STRUCTURES
// ------------------------------------------------------------

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

// Action builder
type Player struct {
	factories []Factory
	troops    []Troop
	bombs     []Bomb
	turn      []string
	score     int
}

// ------------------------------------------------------------
// PLAYER METHODS
// ------------------------------------------------------------

func (p Player) facWithMaxCyb() int {
	max := -1
	id := -1
	for _, f := range p.factories {
		if f.unitCount > max {
			max = f.unitCount
			id = f.id
		}
	}
	return id
}

func (p *Player) getScore() {
	total := 0
	for _, f := range p.factories {
		total += f.unitCount
	}
	for _, t := range p.troops {
		total += t.unitCount
	}
	p.score = total
}

func (p *Player) move(from, to, cyborgCount int) {
	p.turn = append(p.turn, fmt.Sprintf("MOVE %d %d %d", from, to, cyborgCount))
}

func (p *Player) bomb(from, to int) {
	p.turn = append(p.turn, fmt.Sprintf("BOMB %d %d", from, to))
}

func (p *Player) wait() {
	p.turn = append(p.turn, "WAIT")
}

func (p *Player) msg(msg string) {
	p.turn = append(p.turn, fmt.Sprintf("MSG %s", msg))
}

func (p *Player) sendCommands() {
	// Si rien n'a été ajouté → WAIT
	if len(p.turn) == 0 {
		p.wait()
	}
	fmt.Println(strings.Join(p.turn, ";"))
}

// ------------------------------------------------------------
// GAME STATE
// ------------------------------------------------------------

type State struct {
	factoryCount     int
	linkCount        int
	links            []Link
	neutralFactories []Factory
	me               Player
	opp              Player
}

// ------------------------------------------------------------
// INITIALIZATION
// ------------------------------------------------------------

func (s *State) readMap() {
	fmt.Scan(&s.factoryCount)
	fmt.Scan(&s.linkCount)

	s.links = make([]Link, 0, s.linkCount)

	for i := 0; i < s.linkCount; i++ {
		var f1, f2, dist int
		fmt.Scan(&f1, &f2, &dist)
		s.links = append(s.links, Link{f1, f2, dist})
	}

	// Tri des liens par distance (utile pour heuristiques)
	sort.Slice(s.links, func(i, j int) bool {
		return s.links[i].distance < s.links[j].distance
	})
}

// ------------------------------------------------------------
// PARSE ENTITIES EACH TURN
// ------------------------------------------------------------

func (s *State) readEntity() {
	var entityCount int
	fmt.Scan(&entityCount)

	// ✅ Très important : reset à chaque tour
	s.me.factories = nil
	s.me.troops = nil
	s.me.bombs = nil

	s.opp.factories = nil
	s.opp.troops = nil
	s.opp.bombs = nil

	s.neutralFactories = nil

	for i := 0; i < entityCount; i++ {
		var id int
		var t string
		var a1, a2, a3, a4, a5 int

		fmt.Scan(&id, &t, &a1, &a2, &a3, &a4, &a5)

		switch t {

		case "FACTORY":
			f := Factory{id, a1, a2, a3, a4}
			switch a1 {
			case 0:
				s.neutralFactories = append(s.neutralFactories, f)
			case 1:
				s.me.factories = append(s.me.factories, f)
			case -1:
				s.opp.factories = append(s.opp.factories, f)
			}

		case "TROOP":
			tr := Troop{id, a1, a2, a3, a4, a5}
			if a1 == 1 {
				s.me.troops = append(s.me.troops, tr)
			} else if a1 == -1 {
				s.opp.troops = append(s.opp.troops, tr)
			}

		case "BOMB":
			b := Bomb{a1, a2, a3, a4}
			if a1 == 1 {
				s.me.bombs = append(s.me.bombs, b)
			} else if a1 == -1 {
				s.opp.bombs = append(s.opp.bombs, b)
			}
		}
	}

	s.me.getScore()
	s.opp.getScore()

	// tri des factories par production
	sort.Slice(s.me.factories, func(i, j int) bool {
		return s.me.factories[i].productionRate > s.me.factories[j].productionRate
	})
	sort.Slice(s.opp.factories, func(i, j int) bool {
		return s.opp.factories[i].productionRate > s.opp.factories[j].productionRate
	})
	sort.Slice(s.neutralFactories, func(i, j int) bool {
		return s.neutralFactories[i].productionRate > s.neutralFactories[j].productionRate
	})
}

// ------------------------------------------------------------
// HELPER : existe-t-il un lien direct ?
// ------------------------------------------------------------

func (s State) linkTo(f1, f2 Factory) bool {
	for _, l := range s.links {
		if (l.from == f1.id && l.to == f2.id) || (l.to == f1.id && l.from == f2.id) {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------
// THINK (simple version)
// ------------------------------------------------------------

func (s *State) think() {

	// Prend l'usine la plus forte
	if len(s.me.factories) == 0 {
		s.me.wait()
		return
	}

	// Exemple simple :
	// 1. Capturer neutres
	for _, src := range s.me.factories {
		if src.unitCount <= 1 {
			continue
		}

		for _, nf := range s.neutralFactories {
			if s.linkTo(src, nf) {
				s.me.move(src.id, nf.id, 1)
			}
		}
	}

	// 2. Sinon attaquer l’ennemi
	for _, src := range s.me.factories {
		if src.unitCount > 1 {
			for _, of := range s.opp.factories {
				if s.linkTo(src, of) {
					s.me.move(src.id, of.id, 1)
				}
			}
		}
	}

	// Si aucune action → WAIT
	if len(s.me.turn) == 0 {
		s.me.wait()
	}
}

// ------------------------------------------------------------
// MAIN
// ------------------------------------------------------------

func main() {
	var game State
	game.readMap()

	for {
		game.readEntity()
		game.think()
		game.me.sendCommands()
	}
}
