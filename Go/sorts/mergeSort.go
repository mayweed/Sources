// Write a merge sort top-down here
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Merge 2 sorted sequences left and right
func merge(leftArray, rightArray []int) []int {
	var output []int
	var i, j int
	var count int

	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			output = append(output, leftArray[i])
			i += 1
			count += 1
		} else if rightArray[j] < leftArray[i] {
			output = append(output, rightArray[j])
			j += 1
			count += 1
		}
	}

	//remainding items...
	for i < len(leftArray) {
		output = append(output, leftArray[i])
		i += 1
	}
	for j < len(rightArray) {
		output = append(output, rightArray[j])
		j += 1
	}
	return output
}

// Sort by merge two sequences
func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	middle := len(array) / 2
	left := mergeSort(array[:middle])
	right := mergeSort(array[middle:])
	return merge(left, right)
}

func main() {
	//input file
	file, err := os.Open("IntegerArray.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//file is Reader interface!!(it implements Read!!)
	scanner := bufio.NewScanner(file)

	var array []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		array = append(array, num)
	}

	log.Println(array)
	start := time.Now()
	fmt.Println(mergeSort(array))
	end := time.Now()
	fmt.Println("delta:", end.Sub(start))
}
