package main

import (
	"fmt"
)

func countNum(num int) int {
	//base case
	if num < 10 {
		return num
	} else {
		return num%10 + countNum(num/10)
	}
}

func river(num int) int {
	return num + countNum(num)
}

/*
func river(num int) int {
	var sum int
	var n = num
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	num += sum
	return num
}
*/
func main() {
	var r1 int
	fmt.Scan(&r1)

	var r2 int
	fmt.Scan(&r2)

	var x = r1
	var y = r2

	var seenX = make(map[int]bool)
	var seenY = make(map[int]bool)

	seenX[x] = true
	seenY[y] = true

	var res int

	for {
		seenX[river(x)] = true
		seenY[river(y)] = true

		if seenX[river(y)] {
			res = river(y)
			break
		} else if seenY[river(x)] {
			res = river(x)
			break
		}
		x = river(x)
		y = river(y)
	}
	fmt.Println(res)
}
