package main

import "fmt"
import "os"
import "bufio"
//import "strings"
//import "strconv"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var startPoint string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&startPoint)

    var endPoint string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&endPoint)

    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)

    for i := 0; i < N; i++ {
        scanner.Scan()
        //stopName := scanner.Text()
    }
    var M int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&M)

    for i := 0; i < M; i++ {
        scanner.Scan()
        //route := scanner.Text()
    }

    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("IMPOSSIBLE")// Write answer to stdout
}
