package main

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var n int
    fmt.Scan(&n)
    
    for i := 0; i < n; i++ {
        var name, parent string
        var birth int
        var death, religion, gender string
        fmt.Scan(&name, &parent, &birth, &death, &religion, &gender)
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("orDeroFsucceSsion")// Write answer to stdout
}