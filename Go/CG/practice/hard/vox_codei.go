package main

import (
	"bufio"
	"fmt"
	"os"
)

//A cell is a pair of coordinate + what's on it!!
type Cell struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// width: width of the firewall grid
	// height: height of the firewall grid
	var width, height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height)

	for i := 0; i < height; i++ {
		scanner.Scan()
		//mapRow := scanner.Text() // one line of the firewall grid
	}
	for {
		// rounds: number of rounds left before the end of the game
		// bombs: number of bombs left
		var rounds, bombs int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &rounds, &bombs)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("3 0") // Write action to stdout
	}
}
