// Write a merge sort of...
package main

import (
    "fmt"
    "math/rand"
    "time"
    )

//not right that length thing...
func merge(leftArray,rightArray []int, length int)[]int{
    var output []int
    var i,j int
    //this loop counter not right!!
    for k:=0;k<length;k++{
        if leftArray[i] < rightArray[j]{
            output[k]=leftArray[i]
            i+=1
        }else if rightArray[j] < leftArray[i]{
            output[k]=rightArray[j]
            j+=1
        }
    }
    return output
}

func mergeSort(array []int,length int)[]int{
    if len(array) < 2{
        return array
    }
    middle:=len(array)/2
    left:=mergeSort(array[:middle])
    right:=mergeSort(array[middle:])
    return merge(left,right,length)
}

func main() {
    //cf go doc example...
    r := rand.New(rand.NewSource(time.Now().UnixNano()))

    //generate an array with 10 random positive values
    var array []int
    for i:=0;i<10;i++{
        array=append(array,r.Int())
    }
    length:=len(array)
    fmt.Println(array)
    fmt.Println(mergeSort(array,length))
}
