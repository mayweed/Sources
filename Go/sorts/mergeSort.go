// Write a merge sort top-down here
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Merge 2 sorted sequences left and right
func merge(leftArray, rightArray []int) []int {
	var output []int
	var i, j int

	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] < rightArray[j] {
			output = append(output, leftArray[i])
			i += 1
		} else if rightArray[j] < leftArray[i] {
			output = append(output, rightArray[j])
			j += 1
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
	//cf go doc example...
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//generate an array with 10 random positive values
	var array []int
	for i := 0; i < 10; i++ {
		array = append(array, r.Int())
	}
	log.Println(array)
	start := time.Now()
	fmt.Println(mergeSort(array))
	end := time.Now()
	fmt.Println("delta:", end.Sub(start))
}
