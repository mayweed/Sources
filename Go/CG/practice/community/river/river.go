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

//is the last elt of list 2 in list 1?
func In(n int, l []int) int {
	var m int
	for _, e := range l {
		if e == n {
			m = e
			break
		}
	}
	return m
}

func main() {
	var r1 int
	fmt.Scan(&r1)

	var r2 int
	fmt.Scan(&r2)

	//no sets in golang...
	var x = r1
	var y = r2
	var riverX []int
	var riverY []int
	riverX = append(riverX, r1)
	riverY = append(riverY, r2)
	var res int
	for {
		riverX = append(riverX, river(x))
		riverY = append(riverY, river(y))

		if ok := In(riverX[len(riverX)-1], riverY); ok != 0 {
			res = ok
			break
		} else if ok := In(riverY[len(riverY)-1], riverX); ok != 0 {
			res = ok
			break
		}
		x = riverX[len(riverX)-1]
		y = riverY[len(riverY)-1]
		//log.Println(riverX,riverY)
	}
	fmt.Println(res)
}
