package main

import (
	"fmt"
	"log"
)

type Potion struct {
	id    int
	ing1  int
	ing2  int
	ing3  int
	ing4  int
	price int
}

type Sort struct {
	id       int
	d1       int
	d2       int
	d3       int
	d4       int
	castable int
}

type Witch struct {
	inv0  int
	inv1  int
	inv2  int
	inv3  int
	score int
}
type State struct {
	commandes []Potion
	witches   []Witch
	casts     []Sort
	oppCasts  []Sort
}

func (s State) findMaxPrice() (int, Potion) {
	var max = 0
	var potMax int
	var p Potion
	for _, potion := range s.commandes {
		if potion.price > max {
			max = potion.price
			potMax = potion.id
			p = potion
		}
	}
	return potMax, p
}
func (s State) think() {
	var _, target = s.findMaxPrice()
	var cpState = s
	if target.ing2 == 0 {
		//check for a cast
		for _, c := range cpState.casts {
			if c.castable == 1 {
				if c.d1 == -1 {
					//apply cast
					cpState.witches[0].inv0 - 1
					//me.invX+1
					//fmt.Println("CAST ", c.id)
				}
			}

		}
	}
}
func main() {

	for {
		// actionCount: the number of spells and recipes in play
		var actionCount int
		fmt.Scan(&actionCount)

		var s State

		for i := 0; i < actionCount; i++ {
			// actionId: the unique ID of this spell or recipe
			// actionType: in the first league: BREW; later: CAST, OPPONENT_CAST, LEARN, BREW
			// delta0: tier-0 ingredient change
			// delta1: tier-1 ingredient change
			// delta2: tier-2 ingredient change
			// delta3: tier-3 ingredient change
			// price: the price in rupees if this is a potion
			// tomeIndex: in the first two leagues: always 0; later: the index in the tome if this is a tome spell, equal to the read-ahead tax
			// taxCount: in the first two leagues: always 0; later: the amount of taxed tier-0 ingredients you gain from learning this spell
			// castable: in the first league: always 0; later: 1 if this is a castable player spell
			// repeatable: for the first two leagues: always 0; later: 1 if this is a repeatable player spell
			var actionId int
			var actionType string
			var delta0, delta1, delta2, delta3, price, tomeIndex, taxCount int
			//var castable, repeatable bool
			var _castable, _repeatable int
			fmt.Scan(&actionId, &actionType, &delta0, &delta1, &delta2, &delta3, &price, &tomeIndex, &taxCount, &_castable, &_repeatable)

			//idée basique: trouver les ing qui manquent pour max et voir si
			//je peux jeter un sort pour les avoir
			if actionType == "BREW" {
				s.commandes = append(s.commandes, Potion{id: actionId, ing1: delta0, ing2: delta1, ing3: delta2, ing4: delta3, price: price})
			} else if actionType == "CAST" {
				s.casts = append(s.casts, Sort{actionId, delta0, delta1, delta2, delta3, _castable})
			}

		}
		for i := 0; i < 2; i++ {
			// inv0: tier-0 ingredients in inventory
			// score: amount of rupees
			var inv0, inv1, inv2, inv3, score int
			fmt.Scan(&inv0, &inv1, &inv2, &inv3, &score)
			s.witches = append(s.witches, Witch{inv0, inv1, inv2, inv3, score})
		}

		log.Println(s.commandes, s.witches[0])
		log.Println(s.casts)
		//potMax, _ := s.findMaxPrice()
		// in the first league: BREW <id> | WAIT; later: BREW <id> | CAST <id> [<times>] | LEARN <id> | REST | WAIT
		//cast while you can
		for _, c := range s.casts {
		}

		//fmt.Println("BREW ", potMax)
	}
}
