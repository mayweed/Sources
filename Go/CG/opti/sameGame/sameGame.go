package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Grid [][]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var board Grid

	for {
		for i := 0; i < 15; i++ {
			scanner.Scan()
			inputs := scanner.Text() //strings.Split(scanner.Text(), " ")
			for j := 0; j < 15; j++ {
				// color: Color of the tile
				board[i][j] = strconv.ParseInt(inputs[i*width+j])
				_ = color
			}
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("3 6 Hello SameGame\\n:-)") // Selected tile "x y [message]".
	}
}
