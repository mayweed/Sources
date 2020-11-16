package main

import (
	"fmt"
)

type Recipe struct {
	id    int
	ings  []int
	price int
}

type Cast struct {
	id         int
	deltas     []int
	castable   bool
	repeatable bool
}

type Witch struct {
	inv   []int
	score int
}
type State struct {
	commandes []Recipe
	//must be Me et Opp
	witches  []Witch
	casts    []Cast
	oppCasts []Cast
}

func (c Cast) isCastable() bool {
	return c.castable
}
func (s State) mustRest() bool {
	var acc int
	for _, c := range s.casts {
		if c.castable {
			acc += 1
		}
	}
	if acc == len(s.casts) {
		return true
	} else {
		return false
	}
}
func (s State) findMaxPrice() (int, Recipe) {
	var max = 0
	var potMax int
	var p Recipe
	for _, potion := range s.commandes {
		if potion.price > max {
			max = potion.price
			potMax = potion.id
			p = potion
		}
	}
	return potMax, p
}

/*
func (s State) validatePotion(p Recipe) bool {
	if p.ing0+s.witches[0].inv0 >= 0 {
		if p.ing1+s.witches[0].inv1 >= 0 {
			if p.ing2+s.witches[0].inv2 >= 0 {
				if p.ing3+s.witches[0].inv3 >= 0 {
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

func (s State) canAfford(c Cast) bool {
	//must check the cast
	if s.witches[0].inv0+c.d0 < 0 || s.witches[0].inv1+c.d1 < 0 || s.witches[0].inv2+c.d2 < 0 || s.witches[0].inv3+c.d3 < 0 {
		return false
	} else {
		return true
	}
}
func (s State) checkWhatIneed(p Recipe) map[int]int {
	var needs = make(map[int]int)
	if s.witches[0].inv0+p.ing0 < 0 {
		needs[0] = int(math.Abs(float64(p.ing0)))
	}
	if s.witches[0].inv1+p.ing1 < 0 {
		needs[1] = int(math.Abs(float64(p.ing1)))
	}
	if s.witches[0].inv2+p.ing2 < 0 {
		needs[2] = int(math.Abs(float64(p.ing2)))
	}
	if s.witches[0].inv3+p.ing3 < 0 {
		needs[3] = int(math.Abs(float64(p.ing3)))
	}
	return needs
}
//it takes a recipe and gave me a sort
func (s State) pickCast(p Recipe) Cast {
	n := s.checkWhatIneed(p)
	var possCast []Cast
	for _, c := range s.casts {
		//must check inventory after??
		//if s.ing0 < 0 && s.witches[0].inv0 >0
		for i := 0; i < 4; i++ {
			if n[i] > 0 && c.d0 > 0 {
				//c can be a candidate
				possCast = append(possCast, c)
				//continue? The best would be to check for other needs and eval
				//higher the cast if multiple needs are met
			}
		}
	}
}

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
				var d []int
				d = append(d, delta0, delta1, delta2, delta3)
				s.commandes = append(s.commandes, Recipe{actionId, d, price})
			} else if actionType == "CAST" {
				var r []int
				r = append(r, delta0, delta1, delta2, delta3)
				s.casts = append(s.casts, Cast{actionId, r, castable, repeatable})
			}

		}
		for i := 0; i < 2; i++ {
			// inv0: tier-0 ingredients in inventory
			// score: amount of rupees
			var inv0, inv1, inv2, inv3, score int
			fmt.Scan(&inv0, &inv1, &inv2, &inv3, &score)
			var inv []int
			inv = append(inv, inv0, inv1, inv2, inv3)
			s.witches = append(s.witches, Witch{inv, score})
		}
		/*
			// this one rests too much and brew too late!!!
			t := s.getPotionToDeliver()
			_, pot := s.findMaxPrice()
			n := s.checkWhatIneed(pot)
			log.Println(n, pot)

			if len(t) == 0 {
				// must filter the cast!! in a func and pick the cast
				// given castable,canAfford and the possibility of délivery
				for _, c := range s.casts {
					if c.castable {
						if s.canAfford(c) {
							fmt.Println("CAST ", c.id)
						}
					} else {
						continue
					}
				}
				if s.mustRest() {
					fmt.Println("REST")
				}

			} else {
				fmt.Println("BREW ", t[0])
			}
			log.Println(t)
		*/
		fmt.Println("WAIT")
	}
}
