//package fb
import "math"

//make it with float64.Even if float is approximated, better
//than truncate int...?
type Point struct {
	x, y float64
}

func newPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

//Basic dist func
func dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

//dist from a point to coordinates
func (p Point) dist2(x2,y2 float64) float64{
    return dist(p.x, p.y, x2, y2)
}

//last step:point to point
func (p Point) distP2P(p1 Point) float64{
    return dist(p.x,p.y,p1.x,p1.y)
}
