package main

import (
	//	"fmt"
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
	case "1":
		switch c[1] {
		case "0":
			val = 10
			suit = c[2]
		}
	default:
		val, _ = strconv.Atoi(c[0])
	}
	//parse things correctly
	if val == 10 {
		suit = c[2]
	} else {
		suit = c[1]
	}

	return Card{val, suit}
}

func main() {
	// n: the number of cards for player 1
	//var n int
	//fmt.Scan(&n)

	/*
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
	*/
	//Should extend capacity of each queue to (m+n)
	var queueCardp1 = []Card{{5, "C"}, {3, "D"}, {2, "C"}, {7, "D"}, {8, "C"},
		{7, "S"}, {5, "D"}, {5, "H"}, {6, "D"}, {5, "S"}, {4, "D"}, {6, "H"}, {6, "S"},
		{3, "C"}, {3, "S"}, {7, "C"}, {4, "S"}, {4, "H"}, {7, "H"}, {4, "C"}, {2, "H"},
		{6, "C"}, {8, "D"}, {3, "H"}, {2, "D"}, {2, "S"}}

	var queueCardp2 = []Card{{14, "C"}, {9, "H"}, {13, "H"}, {13, "C"}, {13, "D"},
		{13, "S"}, {1, "0"}, {1, "0"}, {9, "S"}, {12, "D"}, {11, "S"}, {1, "0"},
		{8, "S"}, {12, "H"}, {11, "D"}, {14, "D"}, {11, "C"}, {14, "S"}, {12, "S"},
		{14, "H"}, {11, "H"}, {1, "0"}, {9, "C"}, {8, "H"}, {12, "C"}, {9, "D"}}

	var turn int
	// Do not handle war!!
	log.Println("BEFORE: ", queueCardp1, queueCardp2)
	//oki this is queue!! I should dequeue til empty!! or nil for a
	//slice
	for {
		if queueCardp1[0].value > queueCardp2[0].value {
			queueCardp1 = append(queueCardp1[1:], queueCardp1[0], queueCardp2[0])
			queueCardp2 = queueCardp2[1:]
			log.Println("QP1", queueCardp1, queueCardp2)
		}
		if queueCardp2[0].value > queueCardp1[0].value {
			queueCardp2 = append(queueCardp2[1:], queueCardp1[0], queueCardp2[0])
			queueCardp1 = queueCardp1[1:]
			log.Println("QP2", queueCardp1, queueCardp2)
		}
		//		/* WAR
		if queueCardp1[0].value == queueCardp2[0].value {
			//first player wins the war
			if queueCardp1[4].value > queueCardp2[4].value {
				queueCardp1 = append(queueCardp1[4:], queueCardp1[0:4]...)
				queueCardp1 = append(queueCardp1[4:], queueCardp2[0:4]...)
				queueCardp2 = queueCardp2[4:]
			}
			//second player wins the war
			if queueCardp2[4].value > queueCardp1[4].value {
				queueCardp2 = append(queueCardp2[4:], queueCardp1[0:4]...)
				queueCardp2 = append(queueCardp2[4:], queueCardp2[0:4]...)
				queueCardp1 = queueCardp1[4:]
			}
			//		       if queueCardp1[4].value == queueCardp2[4].value
			//chained battles
		}
		//		*/
		if len(queueCardp1) == 0 || len(queueCardp2) == 0 {
			break
		}
		turn += 1
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
