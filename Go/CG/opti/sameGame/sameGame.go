package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y  int
	color int64
}

func outputTile(p Point) {
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	for {
		//15 is width AND height
		var board = make([][]Point, 15)

		for y := 14; y >= 0; y-- {
			scanner.Scan()
			inputs := strings.Split(scanner.Text(), " ")
			board[y] = make([]Point, 15)
			for x := 0; x < 15; x++ {
				c, _ := strconv.ParseInt(inputs[x], 10, 32)
				board[y][x] = Point{y, x, c}
			}
		}
		log.Println(board)
		fmt.Println("3 6 Hello SameGame\\n:-)") // Selected tile "x y [message]".
	}
}
