package main

import (
	"fmt"
)

/*
func countNum(num int) int {
	//base case
	if num < 10 {
		return num
	} else {
		return num%10 + countNum(num/10)
	}
}
*/

func river(num int) int {
	var sum int
	var n = num
	for n > 0 {
		sum += n % 10
		n = n / 10
		if n == 0 {
			break
		}
	}
	num += sum
	return num
}

/*
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
*/
func main() {
	var r1 int
	fmt.Scan(&r1)

	var r2 int
	fmt.Scan(&r2)

	var x = r1
	var y = r2

	//var riverX []int
	//var riverY []int

	//riverX = append(riverX, r1)
	//riverY = append(riverY, r2)

	var seenX = make(map[int]bool)
	var seenY = make(map[int]bool)

	seenX[x] = true
	seenY[y] = true

	var res int

	for {
		//riverX = append(riverX, river(x))
		//riverY = append(riverY, river(y))

		seenX[river(x)] = true
		seenY[river(y)] = true

		if seenX[river(y)] {
			//if ok := In(riverX[len(riverX)-1], riverY); ok != 0 {
			res = river(y)
			break
			//} else if ok := In(riverY[len(riverY)-1], riverX); ok != 0 {
		} else if seenY[river(x)] {
			res = river(x)
			break
		}
		x = river(x)
		y = river(y)
		//x = riverX[len(riverX)-1]
		//y = riverY[len(riverY)-1]
		//log.Println(riverX, riverY, seenX, seenY)
	}
	fmt.Println(res)
}
