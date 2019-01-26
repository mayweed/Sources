package main

import (
	"fmt"
	//"log"
)

func countNum(num int) int {
	//base case
	if num < 10 {
		return num
	} else {
		return num%10 + countNum(num/10)
	}
}

//for a num give the next one in line
//create an anonymous func?? to got river?
func river(num int) int {
	return num + countNum(num)
}

func main() {
	var r1 int
	fmt.Scan(&r1)

	var r2 int
	fmt.Scan(&r2)

	//no sets in golang...
	var x = r1
	var y = r2
	for {
		x = river(x)
		if x == y {
			break
		}
	}
	fmt.Println(x)
}
