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

	var gridState = true //bool
	for i := 0; i < 9; i++ {
		var checkRows = make(map[int64]int)
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		for j := 0; j < 9; j++ {
			n, _ := strconv.ParseInt(inputs[j], 10, 32)
			checkRows[n] += 1
			if checkRows[n] > 1 {
				gridState = false
				break
			}
		}
		if !gridState {
			break
		}
	}

	if gridState {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
