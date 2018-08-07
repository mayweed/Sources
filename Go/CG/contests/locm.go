package main

import (
	"fmt"
	"log"
)

type Card struct {
	cardNumber, instanceId, location, cardType, cost, attack, defense int
	abilities                                                         string
	myHealthChange, opponentHealthChange, cardDraw                    int
}

//during the draft phase choose the card with max attack.
//func pickDraftCards
func main() {

	var turn int
	for {

		for i := 0; i < 2; i++ {
			var playerHealth, playerMana, playerDeck, playerRune int
			fmt.Scan(&playerHealth, &playerMana, &playerDeck, &playerRune)
		}
		var opponentHand int
		fmt.Scan(&opponentHand)

		var cardCount int
		fmt.Scan(&cardCount)

		var cards []Card
		var myDeck []Card

		var idCardMax int
		var max int

		for i := 0; i < cardCount; i++ {
			var cardNumber, instanceId, location, cardType, cost, attack, defense int
			var abilities string
			var myHealthChange, opponentHealthChange, cardDraw int
			fmt.Scan(&cardNumber, &instanceId, &location, &cardType, &cost, &attack, &defense, &abilities, &myHealthChange, &opponentHealthChange, &cardDraw)

			cards = append(cards, Card{cardNumber, instanceId, location, cardType, cost, attack, defense, abilities, myHealthChange, opponentHealthChange, cardDraw})
			if attack > max {
				max = attack
				//0 1 or 2
				idCardMax = i
			}
		}
		myDeck = append(myDeck, cards[idCardMax])
		turn += 1
		log.Println(idCardMax, turn)
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		if turn < 30 {
			fmt.Println("PICK", idCardMax) // Write action to stdout
		} else {
			//should devise a proper strat!!
			fmt.Println("PASS")
		}

	}
}
