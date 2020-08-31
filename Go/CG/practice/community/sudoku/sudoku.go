package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkCols(g [][]int64) bool {
	//inc y, not x
	for x := 0; x < 9; x++ {
		var checkCols = make(map[int64]int)
		for y := 0; y < 9; y++ {
			checkCols[g[y][x]] += 1
			if checkCols[g[y][x]] > 1 {
				return false
			}
		}
	}
	return true

}
func checkRows(g [][]int64) bool {
	//inc x, not y
	for y := 0; y < 9; y++ {
		var checkCols = make(map[int64]int)
		for j := 0; j < 9; j++ {
			checkCols[g[y][j]] += 1
			if checkCols[g[y][j]] > 1 {
				return false
			}
		}
	}
	return true

}
func checkSubGrid(g [][]int64) bool {
	//first three

	checkDigits := make(map[int64]int)
	for y := 0; y <= 2; y++ {
		for x := 0; x <= 2; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}

		}
	}
	checkDigits = make(map[int64]int)
	for y := 0; y <= 2; y++ {

		for x := 3; x <= 5; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}

		}
	}
	checkDigits = make(map[int64]int)
	for y := 0; y <= 2; y++ {

		for x := 6; x <= 8; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}

		}
	}

	//next three
	checkDigits = make(map[int64]int)
	for y := 3; y <= 5; y++ {

		for x := 0; x <= 2; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}

		}
	}
	checkDigits = make(map[int64]int)
	for y := 3; y <= 5; y++ {

		for x := 3; x <= 5; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}

		}
	}
	checkDigits = make(map[int64]int)
	for y := 3; y <= 5; y++ {

		for x := 6; x <= 8; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}
		}
	}
	//last three
	checkDigits = make(map[int64]int)
	for y := 6; y <= 8; y++ {

		for x := 0; x <= 2; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}
		}
	}
	checkDigits = make(map[int64]int)
	for y := 6; y <= 8; y++ {

		for x := 3; x <= 5; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}
		}
	}
	checkDigits = make(map[int64]int)
	for y := 6; y <= 8; y++ {

		for x := 6; x <= 8; x++ {
			checkDigits[g[y][x]] += 1
			if checkDigits[g[y][x]] > 1 {
				return false
			}

		}
	}
	return true
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

	log.Println(checkSubGrid(grid))
	if checkCols(grid) && checkRows(grid) && checkSubGrid(grid) {

		fmt.Println("true")
	} else {

		fmt.Println("false")
	}

}
