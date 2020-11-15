package main

import (
	"fmt"
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
	id         int
	d1         int
	d2         int
	d3         int
	d4         int
	castable   bool
	repeatable bool
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

func (s State) validatePotion(p Potion) bool {
	if p.ing1+s.witches[0].inv0 >= 0 {
		if p.ing2+s.witches[0].inv1 >= 0 {
			if p.ing3+s.witches[0].inv2 >= 0 {
				if p.ing4+s.witches[0].inv3 >= 0 {
					return true
				}
			}
		}
	}
	return false
}

func (s State) getPotionToDeliver() []int {
	var ids []int
	for _, p := range s.commandes {
		if s.validatePotion(p) {
			ids = append(ids, p.id)
		}
	}
	return ids
}

func (s State) canAfford(c Sort) bool {
	//must check the cast
	if s.witches[0].inv0+c.d1 < 0 || s.witches[0].inv1+c.d2 < 0 || s.witches[0].inv2+c.d3 < 0 || s.witches[0].inv3+c.d4 < 0 {
		return false
	} else {
		return true
	}
}

/*
//if i chose that cast what my inv will look like?
func (s State) applyCast(c Sort) {
	var cpState = s
	if c.castable {
		if c.d1 < 0 {
			//apply cast
			cpState.witches[0].inv0 + c.d1
		} else if c.d1 > 0 {
			cpState.witches[0].inv0 + c.d1
		}
	}
}
*/
/*
func (s State) think() {
	var _, target = s.findMaxPrice()
		if target.ing2 == 0 {
		//check for a cast
		for _, c := range cpState.casts {

		}
	}
}
*/
func main() {

	for {
		var s State

		// actionCount: the number of spells and recipes in play
		var actionCount int
		fmt.Scan(&actionCount)

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
			var castable, repeatable bool
			var _castable, _repeatable int
			fmt.Scan(&actionId, &actionType, &delta0, &delta1, &delta2, &delta3, &price, &tomeIndex, &taxCount, &_castable, &_repeatable)
			castable = _castable != 0
			repeatable = _repeatable != 0

			//init sorts et potions
			if actionType == "BREW" {
				s.commandes = append(s.commandes, Potion{id: actionId, ing1: delta0, ing2: delta1, ing3: delta2, ing4: delta3, price: price})
			} else if actionType == "CAST" {
				s.casts = append(s.casts, Sort{actionId, delta0, delta1, delta2, delta3, castable, repeatable})
			}

		}
		for i := 0; i < 2; i++ {
			// inv0: tier-0 ingredients in inventory
			// score: amount of rupees
			var inv0, inv1, inv2, inv3, score int
			fmt.Scan(&inv0, &inv1, &inv2, &inv3, &score)
			s.witches = append(s.witches, Witch{inv0, inv1, inv2, inv3, score})
		}

		// this one rests too much and brew too late!!!
		t := s.getPotionToDeliver()

		if len(t) == 0 {
			// must filter the cast!! in a func and pick the cast
			// given castable,canAfford and the possibility of délivery
			for _, c := range s.casts {
				if c.castable {
					if s.canAfford(c) {
						fmt.Println("CAST ", c.id)
						continue
					}
				} else {
					fmt.Println("REST")
				}

			}
		} else {
			fmt.Println("BREW ", t[0])
		}
	}
}
