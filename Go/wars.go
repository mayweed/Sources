package main

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Card struct {
	value int
	suit  string
}

func ParseCard(card string) Card {
	var val int
	var suit string
	c := strings.Split(card, "")

	//first char
	switch c[0] {
	case "J":
		val = 11
	case "Q":
		val = 12
	case "K":
		val = 13
	case "A":
		val = 14
	case "1":
		val = 10
	default:
		val, _ = strconv.Atoi(c[0])
	}

	//second char
	if val == 10 {
		suit = c[2]
	} else {
		suit = c[1]
	}

	return Card{val, suit}
}

func main() {
	// n: the number of cards for player 1
	var n int
	fmt.Scan(&n)

	var deckP1 []Card
	for i := 0; i < n; i++ {
		// cardp1: the n cards of player 1
		var cardp1 string
		fmt.Scan(&cardp1)
		c := ParseCard(cardp1)
		deckP1 = append(deckP1, c)
	}
	// m: the number of cards for player 2
	var m int
	fmt.Scan(&m)

	var queueCardp2 []Card
	var my_deck list.List

	for i := 0; i < m; i++ {
		// cardp2: the m cards of player 2
		var cardp2 string
		fmt.Scan(&cardp2)
		//log.Println(cardp2)
		c := ParseCard(cardp2)
		queueCardp2 = append(queueCardp2, c)
		my_deck.PushBack(c)

	}

	// Do not handle war!!
	var turn int

	//to dump cards in a war:
	// -dump the equal cards + the 3 cards if next cards in decks == again just od it again
	var warBufferP1 []Card
	var warBufferP2 []Card
	//prevent declare and not used thing...
	//see:https://golang.org/doc/effective_go.html#blank
	//cond compil??or on the fly when needed???
	_ = warBufferP1
	_ = warBufferP2

	log.Println("BEFORE: ", deckP1, queueCardp2)
	for {
		if len(deckP1) == 0 || len(queueCardp2) == 0 {
			break
		} else {
			if deckP1[0].value > queueCardp2[0].value {
				deckP1 = append(deckP1[1:], deckP1[0], queueCardp2[0])
				queueCardp2 = queueCardp2[1:]
				//log.Println(deckP1)
			} else if queueCardp2[0].value > deckP1[0].value {
				queueCardp2 = append(queueCardp2[1:], deckP1[0], queueCardp2[0])
				deckP1 = deckP1[1:]
				//log.Println(queueCardp2)
			}
			turn += 1
		}
		/* WAR
		   if deckP1[pop].value == queueCardp2[pop].value{
		   then they take an other card
		   if one is > to the other, the player takes all four cards
		   at the back of his deck: his cards first
		   if it's equal, then players must take the the card+3 to determine
		   who wins etc...
		*/

	}
	log.Println("AFTER: ", deckP1, queueCardp2)

	var player int
	if len(deckP1) == 0 {
		player = 2
	} else {
		player = 1
	}
	fmt.Printf("%d %d\n", player, turn)
	//fmt.Println("PAT")// Write answer to stdout
}
