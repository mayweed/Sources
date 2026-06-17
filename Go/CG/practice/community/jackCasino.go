package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var ROUNDS int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &ROUNDS)

	var CASH int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &CASH)

	for i := 0; i < ROUNDS; i++ {
		scanner.Scan()
		PLAY := scanner.Text()
		_ = PLAY // to avoid unused error

		token := strings.Fields(PLAY)
		move := token[1]
		var n1, n2 int
		if len(token) == 3 {
			n1, _ = strconv.Atoi(token[0])
			n2, _ = strconv.Atoi(token[2])
		} else {
			n1, _ = strconv.Atoi(token[0])
		}

		//ceil(a / d) = (a + d - 1) / d
		betAmount := (CASH + 3) / 4
		CASH -= betAmount

		if move == "ODD" {
			if n1%2 != 0 {
				CASH += betAmount * 2
			}
		} else if move == "EVEN" {
			if n1 != 0 && n1%2 == 0 {
				CASH += betAmount * 2
			}
		} else if move == "PLAIN" {
			if n1 == n2 {
				CASH += betAmount * 36
			}
		}
	}
	fmt.Println(CASH) // Write answer to stdout
}
