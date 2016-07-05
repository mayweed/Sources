package cg

import "fmt"
import "math"

type position struct{
    x,y int
}
func NewPosition(x,y int) position{
    return position{
        x:x,
        y:y,
    }
}
func (p position) Move(x,y int) (int,int){
    x=p.x+x
    y=p.y+y
    return x,y
}
func (p position) distFrom(x,y int) float64{
    dist:=math.Sqrt((float64(p.x)-float64(x))*(float64(p.x)-float64(x))+(float64(p.y)-float64(y))*(float64(p.y)-float64(y)))
    return dist
}
