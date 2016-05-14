package main

import ("fmt" 
        "strings"
        )
var s string="hello,you"

func main(){
    t:=strings.Split(s,",")
    fmt.Printf("%s\n",strings.Split(s,","))
    fmt.Printf("%s\n",t[0])
}

