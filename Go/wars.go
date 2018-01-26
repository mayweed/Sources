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

	switch c[0] {
	case "J":
		val = 11
		suit = c[1]
	case "Q":
		val = 12
		suit = c[1]
	case "K":
		val = 13
		suit = c[1]
	case "A":
		val = 14
		suit = c[1]
	case "1":
		val = 10
		suit = c[2]
	default:
		val, _ = strconv.Atoi(c[0])
		suit = c[1]
	}

	return Card{val, suit}
}

func main() {
	// n: the number of cards for player 1
	var n int
	fmt.Scan(&n)

	var queueCardp1 []Card
	for i := 0; i < n; i++ {
		// cardp1: the n cards of player 1
		var cardp1 string
		fmt.Scan(&cardp1)
		c := ParseCard(cardp1)
		queueCardp1 = append(queueCardp1, c)
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
	log.Println("BEFORE: ", queueCardp1, queueCardp2)
	for {
		if len(queueCardp1) == 0 || len(queueCardp2) == 0 {
			break
		} else {
			if queueCardp1[0].value > queueCardp2[0].value {
				queueCardp1 = append(queueCardp1[1:], queueCardp1[0], queueCardp2[0])
				queueCardp2 = queueCardp2[1:]
				//log.Println(queueCardp1)
			} else if queueCardp2[0].value > queueCardp1[0].value {
				queueCardp2 = append(queueCardp2[1:], queueCardp1[0], queueCardp2[0])
				queueCardp1 = queueCardp1[1:]
				//log.Println(queueCardp2)
			}
			turn += 1
		}
		/* WAR
		   if queueCardp1[pop].value == queueCardp2[pop].value{
		   then they take an other card
		   if one is > to the other, the player takes all four cards
		   at the back of his deck: his cards first
		   if it's equal, then players must take the the card+3 to determine
		   who wins etc...
		*/

	}
	log.Println("AFTER: ", queueCardp1, queueCardp2)

	var player int
	if len(queueCardp1) == 0 {
		player = 2
	} else {
		player = 1
	}
	fmt.Printf("%d %d\n", player, turn)
	//fmt.Println("PAT")// Write answer to stdout
}
