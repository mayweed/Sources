package main

import (
    "fmt"
    "time"
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
//p is batman pos
func (p position) bombDir(pos position) string{
    var dir string
    switch{
        case p.y < pos.y:
            dir="UP"
        case p.y > pos.y:
            dir="DOWN"
    }
    return dir
}

//Should use command args here:os.Args
var W,H =10,10


func main() {
	// The default number generator is deterministic, so it'll
	// produce the same sequence of numbers each time by default.
	// To produce varying sequences, give it a seed that changes.
	// Note that this is not safe to use for random numbers you
	// intend to be secret, use `crypto/rand` for those.
    // cf https://play.golang.org/p/ZdFpbahgC1
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var bomb_pos=NewPosition(r1.Intn(W),r1.Intn(H))
	var batman_pos=NewPosition(r1.Intn(W),r1.Intn(H))

    fmt.Println("BOMBE:",bomb_pos,"BATMAN:",batman_pos)

    //takes a pos and yields "UP"/"DOWN" etc..
    //remember: top left is (0,0)
    fmt.Println(batman_pos.bombDir(bomb_pos))
}
