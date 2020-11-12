package main

import (
	"fmt"
	"log"
)

type Action struct {
	id    int
	price int
}

func main() {

	for {
		// actionCount: the number of spells and recipes in play
		var actionCount int
		fmt.Scan(&actionCount)

		var max = 0
		var potMax int

		var ids []int
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
			if max < price {
				max = price
				potMax = actionId
			}
			//castable = _castable != 0
			//repeatable = _repeatable != 0
			ids = append(ids, actionId)
			log.Println(potMax, "max", max)
		}
		for i := 0; i < 2; i++ {
			// inv0: tier-0 ingredients in inventory
			// score: amount of rupees
			var inv0, inv1, inv2, inv3, score int
			fmt.Scan(&inv0, &inv1, &inv2, &inv3, &score)
		}

		log.Println(ids)
		//randomIndex := rand.Intn(len(ids))
		// in the first league: BREW <id> | WAIT; later: BREW <id> | CAST <id> [<times>] | LEARN <id> | REST | WAIT
		fmt.Println("BREW ", potMax) //ids[randomIndex])
	}
}
