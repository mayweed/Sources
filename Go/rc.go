package main

import (
	"fmt"
	"log"
)

func keepAdding(participants, L int, ppl []int) {
	var count int
	var p = ppl[0] + participants
	count += 1
	if p > L {
		p -= participants
		count -= 1
		ppl = append(ppl[count:], ppl[count])
		break //exit the loop
	} else {
		keepAdding(participants, L, ppl)
	}
	//should return sth: queue + rides??
}

func main() {
	var L, C, N int
	fmt.Scan(&L, &C, &N)

	var ppl []int
	for i := 0; i < N; i++ {
		var Pi int
		fmt.Scan(&Pi)
		ppl = append(ppl, Pi)
	}
	log.Println("Places:", L, "Times:", C, "Groups:", N, "Initial Queue:", ppl)

	//for each ride keep track of how many ppl goes in
	var rides = make(map[int]int, C)

	// ride n°2=ride[0]+ride[1]. Idea you wanna know
	//earned dirhams for a particular ride , call ride[X]
	var dirhamPerRide = make(map[int]int)

	var count_dirhams int
	var participants = ppl[0]
	ppl = append(ppl[1:], ppl[0])

	for i := 0; i < C; i++ {
		//ride not complete
		//must have a counter to know at which index I stopped
		for participants < L {
			var count int
			var p = ppl[0] + participants
			count += 1
			//one step to far
			if p > L {
				p -= participants
				count -= 1
				ppl = append(ppl[count:], ppl[count])
				break //exit the loop
			} //else if p < L{
			//keepAdding(participants, queue)
			//WHAT IS THE BASE CASE?? iif >

		}
		if participants == L {
			ppl = append(ppl[1:], ppl[0])
		}
		rides[i] = participants
		count_dirhams += participants
		participants = ppl[0]
		log.Println(rides)

	}
	fmt.Println(count_dirhams)
	//Test case 2 for ref
	//fmt.Println(3935)

}
