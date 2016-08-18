// Write a merge sort of...
package main

import (
    "fmt"
    "math/rand"
    "time"
    )

//func merge(){}

func main() {
    //cf go doc example...
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    //generate an array with 10 random positive values
    var array []int
    for i:=0;i<10;i++{
        array=append(array,r.Int())
    }
    fmt.Println(array)
}
