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
			s.casts = append(s.casts, Cast{actionId, r, castable, repeatable})
		case "OPPONENT_CAST":
			var r []int
			r = append(r, delta0, delta1, delta2, delta3)
			s.casts = append(s.oppCasts, Cast{actionId, r, castable, repeatable})

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

//are all spells exhausted? do i need rest?
func (s State) mustRest() bool {
	for _, c := range s.casts {
		if c.castable {
			return false
		}
	}
	return true
}

//the one recipe which yields max profit
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

/*
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
		s.init()
		s.initInventory()

		_, target := s.findMaxPrice()
		s.me.checkWhatIneed(target)
		log.Println("T: ", target, "N: ", s.me.needs, "I :", s.me.inv)
		/*
			so set a target then check my needs and casts accordingly
			the idea: fulfill recipes
				// this one rests too much and brew too late!!!

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
