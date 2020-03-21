package main

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var width, height, myId int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width, &height, &myId)

	for i := 0; i < height; i++ {
		scanner.Scan()
		//line := scanner.Text()
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("7 7") // Write action to stdout
	for {
		var x, y, myLife, oppLife, torpedoCooldown, sonarCooldown, silenceCooldown, mineCooldown int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &x, &y, &myLife, &oppLife, &torpedoCooldown, &sonarCooldown, &silenceCooldown, &mineCooldown)
		var sonarResult string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &sonarResult)

		scanner.Scan()
		//opponentOrders := scanner.Text()

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("MOVE N TORPEDO") // Write action to stdout
	}
}
