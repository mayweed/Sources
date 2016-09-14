// Write a merge sort of...
package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	//cf go doc example...
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//generate an array with 10 random positive values
	var array []int
	for i := 0; i < 10; i++ {
		array = append(array, r.Intn(10))
	}
	log.Println(array)

	for j := 1; j < len(array); j++ {
		var key = array[j]
		var i = j - 1
		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i = i - 1
		}
		array[i+1] = key
	}
	log.Println(array)
}
