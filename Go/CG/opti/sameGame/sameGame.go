package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	for {
		for i := 0; i < 15; i++ {
			scanner.Scan()
			inputs := strings.Split(scanner.Text(), " ")
			for j := 0; j < 15; j++ {
				// color: Color of the tile
				color, _ := strconv.ParseInt(inputs[j], 10, 32)
				_ = color
			}
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println("3 6 Hello SameGame\\n:-)") // Selected tile "x y [message]".
	}
}
