package main

import (
	"fmt"
	"log"
)

/**
 * 001010101000001111
 **/

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var B string
		var C int
		fmt.Scan(&B, &C)
		log.Println(B, string(C))
	}
	var S string
	fmt.Scan(&S)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("abracadabra") // Write answer to stdout
}
