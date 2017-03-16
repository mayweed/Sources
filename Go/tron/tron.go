package main

import (
    "fmt"
    "log"
    )

type point struct{
	x,y int
}
//func (p point)comparePoint
//equals less etc??
type player struct{
    id int
    position point
    lastPos point
}
//This one should have a string meth
//to got traces + an update meth to take into
//account moves
type board struct{
    width int //=30
    height int //=20
    //list of points
    cells [][]point
}
//should be inited with w h and
func (b board) initBoard (width,height int,start ...point{})board{
    //a simple grid made of cells
    var i,j int
    var grid=make([][]point,height)
    for i = 0; i < height; i++ {
        grid[i]=make([]point,width)
        for j= range(grid[i]){
            grid[i][j]=point{i,j}
        }
    }
    //should MARK start point here!!
    return board{
        width:w,
        height:h,
        cells:grid,
    }
}
type gameState struct{
    //based on id: player[0] gives first player etc...
    players []player
    //should update board with trace, either with a lettre or sth
    //where a player has played (or the id??)
    //board grid
    numOfTurns int
}

func main() {
    //is this the right struct for this?
    var actions=make(map[string][]int)
    actions["LEFT"]=[]int{-1,0}
    actions["RIGHT"]=[]int{1,0}
    actions["UP"]=[]int{0,-1}
    actions["DOWN"]=[]int{0,1}

    g:= gameState{}

    for {
        // N: total number of players (2 to 4).
        // P: your player number (0 to 3).
        var N, P int
        fmt.Scan(&N, &P)

        for i := 0; i < N; i++ {
            // X0: starting X coordinate of lightcycle (or -1)
            // Y0: starting Y coordinate of lightcycle (or -1)
            // X1: starting X coordinate of lightcycle (can be the same as X0 if you play before this player)
            // Y1: starting Y coordinate of lightcycle (can be the same as Y0 if you play before this player)
            var X0, Y0, X1, Y1 int
            fmt.Scan(&X0, &Y0, &X1, &Y1)
            log.Println(X0,Y0,X1,Y1)
        }

        g.numOfTurns+=1
        log.Println(g.numOfTurns)
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        fmt.Println("LEFT") // A single line with UP, DOWN, LEFT or RIGHT
    }
}
