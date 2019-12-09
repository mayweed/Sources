package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	//first open file
	file, err := os.Open("inputD1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read each line add the result /3(rounded) -2
	scanner := bufio.NewScanner(file)
	var res int
	for scanner.Scan() { // internally, it advances token based on sperator
		num, _ := strconv.Atoi(scanner.Text())
		div := math.Floor(float64(num) / 3)
		res += int(div) - 2
	}
	fmt.Println(res)
}
