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

//first simply pick the one with max attack
func draftPhase() string {
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
	str := fmt.Sprintf("PICK %d\n", idCardMax)
	return str //should update a state instead of yielding it,myDeck
}

func main() {
	var turn int
	for {

		for i := 0; i < 2; i++ {
			var playerHealth, playerMana, playerDeck, playerRune int
			fmt.Scan(&playerHealth, &playerMana, &playerDeck, &playerRune)
		}
		var opponentHand int
		fmt.Scan(&opponentHand)

		if turn < 30 {
			s := draftPhase()
			fmt.Printf("%s", s) // Write action to stdout
		} else {
			var cardCount int
			fmt.Scan(&cardCount)

			for i := 0; i < cardCount; i++ {
				var cardNumber, instanceId, location, cardType, cost, attack, defense int
				var abilities string
				var myHealthChange, opponentHealthChange, cardDraw int
				fmt.Scan(&cardNumber, &instanceId, &location, &cardType, &cost, &attack, &defense, &abilities, &myHealthChange, &opponentHealthChange, &cardDraw)
			}
			fmt.Println("PASS")
		}
		turn += 1
		log.Println(turn)
	}
}
