package main

import (
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
	//strings are made of bytes in Go
	c := strings.Split(card, "")

	switch c[0] {
	case "J":
		val = 11
	case "Q":
		val = 12
	case "K":
		val = 13
	case "A":
		val = 14
	default:
		val, _ = strconv.Atoi(c[0])
	}
	suit = c[1]

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
	for i := 0; i < m; i++ {
		// cardp2: the m cards of player 2
		var cardp2 string
		fmt.Scan(&cardp2)
		c := ParseCard(cardp2)
		queueCardp2 = append(queueCardp2, c)

	}
	// Do not handle war!!
	log.Println("BEFORE: ", queueCardp1, queueCardp2)
	for {
		if queueCardp1[0].value > queueCardp2[0].value {
			queueCardp1 = append(queueCardp1[1:], queueCardp1[0], queueCardp2[0])
			queueCardp2 = queueCardp2[1:]
		}
		if queueCardp2[0].value > queueCardp1[0].value {
			queueCardp2 = append(queueCardp2[1:], queueCardp1[0], queueCardp2[0])
			queueCardp1 = queueCardp1[1:]
		}
		/* WAR
		   if queueCardp1[0].value == queueCardp2[0].value{
		       //first player wins the war
		       if queueCardp1[4].value > queueCardp2[4].value{
		           queueCardp1=append(queueCardp1[4:],queueCardp1[0:3],queueCardp2[0:3])
		           queueCardp2=queueCardp2[4:]
		       }
		       //second player wins the war
		       if queueCardp2[4].value > queueCardp1[4].value{
		           queueCardp2=append(queueCardp2[4:],queueCardp1[0:3],queueCardp2[0:3])
		           queueCardp1=queueCardp1[4:]
		       }
		       if queueCardp1[4].value == queueCardp2[4].value
		       //chained battles
		*/
		if len(queueCardp1) == 0 || len(queueCardp2) == 0 {
			break
		}
	}
	log.Println("AFTER: ", queueCardp1, queueCardp2)
	//fmt.Println("PAT")// Write answer to stdout
}
