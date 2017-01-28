package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"log"
	"math"
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

func calculusToTs(sec int) string {
	var min int
	for sec > 60 {
		sec -= 60
		min++
	}
	ts := fmt.Sprintf("%d:%d", min, sec)
	return ts
}

//2 particular cases are missing in this code:
// - timer starts game
// - room filled
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	var calculus float64
	for i := 0; i < n; i++ {
		scanner.Scan()
		timeStamp := scanner.Text()

		t := tsToSec(timeStamp)
		calculus = float64(t) - 256./(math.Exp2(float64(n-1)))

		//log.Println(timeStamp,tsToSec(timeStamp),calculusToTs(int(calculus)),seconds)
	}

	var s string
	if calculus < 0 {
		s = "0:00"
	} else {
		s = calculusToTs(int(calculus))
	}
	if n == 0 {
		s = "NO GAME"
	}
	fmt.Println(s) // Write answer to stdout
}
