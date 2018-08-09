package main

import (
	"fmt"
	"log"
	"strings"
)

//CARD
type Card struct {
	cardNumber, instanceId, location, cardType, cost, attack, defense int
	abilities                                                         string
	myHealthChange, opponentHealthChange, cardDraw                    int
}

//PLAYERS
// PlayerInfo stores some data about a player
type PlayerInfo struct {
	health int
	mana   int
	deck   int
	runes  int
	decks  []Card
}

// Players contains info about all players
type Players struct {
	me    PlayerInfo
	enemy PlayerInfo
}

func readPlayers() *Players {

	p := &Players{}

	var playerHealth, playerMana, playerDeck, playerRune int

	fmt.Scan(&playerHealth, &playerMana, &playerDeck, &playerRune)
	me := PlayerInfo{health: playerHealth, mana: playerMana, deck: playerDeck, runes: playerRune}
	p.me = me

	fmt.Scan(&playerHealth, &playerMana, &playerDeck, &playerRune)
	enemy := PlayerInfo{health: playerHealth, mana: playerMana, deck: playerDeck, runes: playerRune}
	p.enemy = enemy

	return p
}

//COMMANDS
func cmdAttack(myCardID, enemyCardID int) string {
	return fmt.Sprintf("ATTACK %d %d", myCardID, enemyCardID)
}

func cmdSummon(cardID int) string {
	return fmt.Sprintf("SUMMON %d", cardID)
}
func sendCommands(commands []string) {
	cmd := "PASS"
	if len(commands) == 0 {
		log.Println("List of commands is empty, PASS will be sent")
	} else {
		cmd = strings.Join(commands, ";")
	}
	fmt.Println(cmd)
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
		players := readPlayers()

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
		log.Println(players)
	}
}
