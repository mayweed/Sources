//echo with join() to compare exec time
package main

import(
    "fmt"
    "strings"
    "os"
    "time"
    )

func main(){
    start := time.Now()
    fmt.Println(strings.Join(os.Args[1:]," "))
    fmt.Printf("%2fs elapsed\n",time.Since(start).Seconds())
}
