//dup2 prints the count and text of lines that appear more than once
//in the input. It reads from stdin or from a list of named files.
//go run dup2.go dup2.txt dup2_bis.txt dup2_ter.txt
package main

import(
    "bufio"
    "fmt"
    "os"
    )

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    fileCount:=make(map[string]map[string]int)
    if len(files)==0 {
        countLines(os.Stdin, counts, fileCount)
    } else {
        for _,arg := range files {
            f, err := os.Open(arg)
        if err != nil {
            fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
            continue
            }
        countLines(f, counts,fileCount)
        f.Close()
        }
    }

    for line,n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n",n,line)
            for fn,fc := range fileCount[line]{
                fmt.Printf("%dx in %s\n",fc,fn)
                }
        }
    }
}

func countLines(f *os.File, counts map[string]int,fileCount map[string]map[string]int) {
    input := bufio.NewScanner(f)
    //with f.Name you got the name of the open file!!
    for input.Scan() {
        filename:=f.Name()
        text:=input.Text()
        counts[text]++
        if fileCount[text]==nil{
            fileCount[text]=make(map[string]int)
        }
        fileCount[text][filename]++
    }
    //no err handling from input.Err()
}
