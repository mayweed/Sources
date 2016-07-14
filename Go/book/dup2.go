//dup2 prints the count and text of lines that appear more than once
//in the input. It reads from stdin or from a list of named files.
//go run dup2.go dup2.txt dup2_bis.txt 
package main

import(
    "bufio"
    "fmt"
    "os"
    )

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files)==0 {
        countLines(os.Stdin, counts)
    } else {
        for k,arg := range files {
            f, err := os.Open(arg)
        if err != nil {
            fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
            continue
            }
        fmt.Println(files[k])
        countLines(f, counts)
        f.Close()
        }
    }

    for line,n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n",n,line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    //no err handling from input.Err()
}
