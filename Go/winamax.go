package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
    )

type Card struct{
    value int
    suit string
}

func ParseCard(card string) Card{
    var val int
    var suit string
    //strings are made of bytes in Go
    c:=strings.Split(card,"")
    
    switch c[0]{
    case "J":
        val=11
    case "Q":
        val=12
    case "K":
        val=13
    case "A":
        val=14
    default:
        val,_=strconv.Atoi(c[0])
    }
    suit=c[1]
    
    return Card{val,suit}
}
        

func main() {
    // n: the number of cards for player 1
    var n int
    fmt.Scan(&n)
    
    var queueCardp1 []Card
    for i := 0; i < n; i++ {
        // cardp1: the n cards of player 1
        var cardp1 string
        fmt.Scan(&cardp1)
        c:=ParseCard(cardp1)
        queueCardp1=append(queueCardp1,c)
    }
    // m: the number of cards for player 2
    var m int
    fmt.Scan(&m)
    
    var queueCardp2 []Card
    for i := 0; i < m; i++ {
        // cardp2: the m cards of player 2
        var cardp2 string
        fmt.Scan(&cardp2)
        c:=ParseCard(cardp2)
        queueCardp2=append(queueCardp2,c)
        
    }
    log.Println(queueCardp1,queueCardp2)
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println("PAT")// Write answer to stdout
}
