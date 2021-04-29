package main

import "fmt"

type Graph struct {
	nodes   []int
	adjList map[int]int
}

func main() {
	// n: the number of relationships of influence
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		// x: a relationship of influence between two people (x influences y)
		var x, y int
		fmt.Scan(&x, &y)
	}
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	// The number of people involved in the longest succession of influences
	fmt.Println("2")
}
