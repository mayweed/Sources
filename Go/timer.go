package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func tsToSec(ts string) int {
	var seconds int
	var min int

	ss := strings.Split(ts, ":")
	min, _ = strconv.Atoi(ss[0])
	min *= 60

	seconds, _ = strconv.Atoi(ss[1])
	seconds += min

	return seconds
}

func secToTs(sec int) string {
	var min int
	for sec > 60 {
		sec -= 60
		min++
	}
    ts:=fmt.Sprintf("%d:%02d",min,sec)
	return ts
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	var calculus float64
	var timeStamp string
	var s string

	for i := 0; i < n; i++ {
		scanner.Scan()
		timeStamp = scanner.Text()
		t := tsToSec(timeStamp)
		//log.Println(calculusToTs(t))

		var timeToStart = tsToSec(secToTs(int(calculus)))
		if t < timeToStart {
			s = secToTs(timeToStart)
			break
		}
		calculus = float64(t) - 256./(math.Exp2(float64(i)))
		//log.Println(calculus)
		log.Println("Timestamp:", timeStamp, "Time to start:", secToTs(int(calculus)))
	}

	if calculus < 0 {
		s = "0:00"
	} else {
		s = secToTs(int(calculus))
	}
	//switch here? func validateN(num int) string{ ??
	if n == 0 {
		s = "NO GAME"
	}
	if n == 7 {
		//clash immediately
		s = timeStamp
	}
	fmt.Println(s)
}
