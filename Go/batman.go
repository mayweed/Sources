package main

import (
    "fmt"
    "math/rand"
    )

// POSITION
type position struct{
    x,y int
}
func NewPosition(x,y int) position{
    return position{
        x:x,
        y:y,
    }
}
func bombDir(pos position) string{
    var dir string
    switch{
        case pos.y < bomb_pos.y:
            dir="UP"
        case pos.y > bomb_pos.y:
            dir="DOWN"
    }
    return dir
}

//Should use command args here:os.Args
var W,H =10,10

var bomb_pos=NewPosition(rand.Intn(W),rand.Intn(H))
var batman_pos=NewPosition(rand.Intn(W),rand.Intn(H))

func main() {
    fmt.Println("BOMBE:",bomb_pos,"BATMAN:",batman_pos)

    //takes a pos and yields "UP"/"DOWN" etc..
    //remember: top left is (0,0)
    fmt.Println(bombDir(batman_pos))
}
