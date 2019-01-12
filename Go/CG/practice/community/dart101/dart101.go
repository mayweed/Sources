package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Player struct {
	id     int
	name   string
	shoots []string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	var players = make([]Player, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		p := scanner.Text()
		players[i].id = i
		players[i].name = p
	}
	for i := 0; i < N; i++ {
		scanner.Scan()
		s := scanner.Text()
		s2 := strings.Split(s, " ")
		players[i].shoots = s2
		//for score := range players[i].shoots{
		//  var result int
		//	s3 := strconv.Atoi(score)
		//  // then cases if not X result +=s3
		//}
	}
	log.Println(players[0])
}
