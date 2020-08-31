package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkCols(g [][]int64) {
	//inc y, not x

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var grid = make([][]int64, 9)

	for i := 0; i < 9; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		grid[i] = make([]int64, 9)
		for j := 0; j < 9; j++ {
			n, _ := strconv.ParseInt(inputs[j], 10, 32)
			grid[i][j] = n
		}
	}
	fmt.Println("true")

}
