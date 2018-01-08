//echo prog from chap1 The Go Prog language
package main

import (
    "fmt"
    "os"
    "time"
    )

func main() {
    start:=time.Now()
    //exo 1.1
    fmt.Println("Name of the prog is",os.Args[0])
    for k,arg := range(os.Args[1:]){
        //exo 1.2
        fmt.Println(k,arg)
    }
    fmt.Printf("%2fs elapsed\n",time.Since(start).Seconds())
}
