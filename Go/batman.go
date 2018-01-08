package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

//not a method, it's calculator
func calculateCenter(p1, p2 Point) Point {
	return Point{
		x: (p1.x + p2.x) / 2,
		y: (p1.y + p2.y) / 2,
	}
}

//should redraw the area wrt the bombdir
//func focus()

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)

	//to keep track of where I am
	//update those structs
	var myPos = Point{X0, Y0}
	//var oldPos=Point{X0,Y0}

	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)

		//use a map litteral dir[bombDir] yields a Point struct
		//those values should be updated wrt batman pos
		//ex: if myPos.x==23 and myPos.y==33, UL={23,33}...etc etc...
		var dir = map[string]Point{
			"UR": Point{W, 0},
			"UL": Point{0, 0},
			"DR": Point{W, H},
			"DL": Point{0, H},
			//HERE!!Y or X do not move
			"U": Point{0, myPos.y},
			"D": Point{myPos.x, H},
			"R": Point{W, myPos.y},
			"L": Point{0, myPos.y},
		}

		fmt.Fprintln(os.Stderr, H, W, bombDir, myPos.x, myPos.y, dir[bombDir])

		// the location of the next window Batman should jump to.
		//fmt.Println(dir[bombDir])
		pos := calculateCenter(myPos, dir[bombDir])
		//oldPos=myPos
		fmt.Println(pos.x, pos.y)
		myPos = pos
	}
}
