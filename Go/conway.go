package main

import "fmt"
import "strings"
import "strconv"

func conway(L string) int{
    var count int=1
    var result string=""
    var Lsplit []string
    Lsplit=strings.Split(L," ")
    prev:=Lsplit[0]
    for i:=1; i<len(Lsplit);i++{
        if Lsplit[i]==" "{
            continue
        }else{
            if Lsplit[i]!=prev{
                result+=strconv.Itoa(count)+" "+prev+" "
                count=1
                prev=Lsplit[i]
            }else if Lsplit[i]==prev{
                count+=1
            }
        }
    }
    result+=strconv.Itoa(count)+" "+prev

    return result
}


func main() {
    var R string
    //fmt.Scan(&R)
    R="1"

    var L int
    //fmt.Scan(&L)
    L=2

    var res string
    for x:=0;x<L-1;x++{
        res=conway(R)
        R=res
    }
    fmt.Println(R)
    // fmt.Fprintln(os.Stderr, "Debug messages...")
}
