package cg

import (
	"fmt"
	"math"
)

//Should add lots of thing here:add/move/sub
//inspiration:https://gist.github.com/mortoray/0826a58d06fc7f06ac6ddf1df56aecfc

type Point struct {
	x, y float64
}

func At(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}
func (p Point) Move(to Point) Point {
	x = p.x + to.x
	y = p.y + to.y
	return Point{x, y}
}
func (p Point) distFrom(destination Point) float64 {
	dist := math.Sqrt(((p.x)-(destination.x))*((p.x)-(destination.x)) + ((p.y)-(destination.y))*((p.y)-(destination.y)))
	return dist
}
