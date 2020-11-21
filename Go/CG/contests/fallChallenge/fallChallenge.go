//Check the cast: do i need what it yields for the first recipe???
package main

import (
	"fmt"
	"log"
	"math"
)

const (
	ING_TYPE_COUNT = 4
)

type (
	Recipe struct {
		id    int
		ings  []int
		price int
	}

	Cast struct {
		id         int
		deltas     []int
		castable   bool
		repeatable bool
		tomeIndex  int
		taxCount   int
	}

	Witch struct {
		//what i got
		inv []int
		//what i need to deliver a given recipe
		needs []int
		score int
	}
	State struct {
		deliveries []Recipe
		me         Witch
		casts      []Cast
		opp        Witch
		oppCasts   []Cast
		tomeSpell  []Cast
	}
)

//init deliveries and casts
func (s *State) init() {
	// actionCount: the number of spells and recipes in play
	var actionCount int
	fmt.Scan(&actionCount)

	for i := 0; i < actionCount; i++ {
		var actionId int
		var actionType string
		var delta0, delta1, delta2, delta3, price, tomeIndex, taxCount int
		var castable, repeatable bool
		var _castable, _repeatable int
		fmt.Scan(&actionId, &actionType, &delta0, &delta1, &delta2, &delta3, &price, &tomeIndex, &taxCount, &_castable, &_repeatable)
		castable = _castable != 0
		repeatable = _repeatable != 0

		switch actionType {
		case "BREW":
			var d []int
			d = append(d, delta0, delta1, delta2, delta3)
			s.deliveries = append(s.deliveries, Recipe{actionId, d, price})
		case "CAST":
			var r []int
			r = append(r, delta0, delta1, delta2, delta3)
			s.casts = append(s.casts, Cast{actionId, r, castable, repeatable, tomeIndex, taxCount})
		case "OPPONENT_CAST":
			var r []int
			r = append(r, delta0, delta1, delta2, delta3)
			s.oppCasts = append(s.oppCasts, Cast{actionId, r, castable, repeatable, tomeIndex, taxCount})
		case "LEARN":
			var r []int
			r = append(r, delta0, delta1, delta2, delta3)
			s.oppCasts = append(s.tomeSpell, Cast{actionId, r, castable, repeatable, tomeIndex, taxCount})

		}
	}
}

//init inventories
func (s *State) initInventory() {
	for i := 0; i < 2; i++ {
		// inv0: tier-0 ingredients in inventory
		// score: amount of rupees
		var inv0, inv1, inv2, inv3, score int
		fmt.Scan(&inv0, &inv1, &inv2, &inv3, &score)
		var inv []int
		inv = append(inv, inv0, inv1, inv2, inv3)
		switch i {
		case 0:
			s.me = Witch{inv, []int{}, score}
		case 1:
			s.opp = Witch{inv, []int{}, score}
		}
	}
}

//the one recipe which yields max profit
//in bronze there is a bonus for the first one so…
func (s State) findMaxPrice() (int, Recipe) {
	var max = 0
	var potMax int
	var p Recipe
	for _, potion := range s.deliveries {
		if potion.price > max {
			max = potion.price
			potMax = potion.id
			p = potion
		}
	}
	return potMax, p
}

//check if a witch can deliver a given recipe
func (w Witch) canDeliver(r Recipe) bool {
	for i := 0; i < ING_TYPE_COUNT; i++ {
		if w.inv[i]+r.ings[i] < 0 {
			return false
		}
	}
	return true
}

func (s State) checkRecipe() int {
	for _, p := range s.deliveries {
		if s.me.canDeliver(p) {
			return p.id
		}
	}
	return 0 //no recipe found
}

//check if i can cast a given spell
func (w Witch) canCast(c Cast) bool {
	for i := 0; i < ING_TYPE_COUNT; i++ {
		if w.inv[i]+c.deltas[i] < 0 {
			return false
		}
	}
	return true
}

func (w *Witch) checkWhatIneed(r Recipe) {
	for i := 0; i < ING_TYPE_COUNT; i++ {
		if w.inv[i]+r.ings[i] < 0 {
			w.needs = append(w.needs, int(math.Abs(float64(r.ings[i]))))
		} else {
			w.needs = append(w.needs, 0)
		}
	}
}

//pick a given cast for a given witch in a given gamestate
func (w Witch) pickCast(s State) []Cast {
	var possCast []Cast
	var added = make(map[int]bool)
	for _, c := range s.casts {
		for i := 0; i < ING_TYPE_COUNT; i++ {
			//i need the ing, the cast can provide AND i can pay…
			if w.needs[i] > 0 && w.canCast(c) && c.castable {
				//c can be a candidate
				if !added[c.id] {
					possCast = append(possCast, c)
					added[c.id] = true
				}
			}
		}
	}
	return possCast
}

//if i chose that cast what my inv will look like?
//seems to work: can simulate what my next turn inv
//will be…
func (s State) applyCast(c Cast) []int {
	var cpState = s
	var newInv []int
	//be sure i can cast it
	if c.castable {
		for i := 0; i < ING_TYPE_COUNT; i++ {
			//apply cast
			newInv = append(newInv, cpState.me.inv[i]+c.deltas[i])
			//c.castable = true //exhausted on next turn
		}
	}
	return newInv
}

//i could apply any cast, even all, but must eval which is the
//best to choose...
func (s State) possibleInvNextTurn() [][]int {
	var allInv [][]int
	for _, c := range s.casts {
		allInv = append(allInv, s.applyCast(c))
	}
	return allInv
}
func main() {

	for {
		var s State
		s.init()
		s.initInventory()

		//so set a target then check my needs and casts accordingly
		//the idea: fulfill recipes

		//_, target := s.findMaxPrice()
		//s.me.checkWhatIneed(target)
		log.Println("N: ", s.me.needs, "I :", s.me.inv)
		p := s.me.pickCast(s)
		//log.Println("CASTS: ", s.casts, "POSS CASTS: ", p)
		//target := s.deliveries[0]

		test := s.possibleInvNextTurn()
		log.Println(test)
		if po := s.checkRecipe(); po != 0 {
			fmt.Println("BREW ", po)
		} else if len(p) > 0 {
			fmt.Println("CAST ", p[0].id)
		} else {
			fmt.Println("REST")
		}
	}
}
