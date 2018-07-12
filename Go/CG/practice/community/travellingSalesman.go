package main

import (
	"fmt"
	"log"
	"math"
)

type Point struct {
	x, y int
}

func (p Point) distanceTo(q Point) float64 {
	return math.Sqrt((float64(q.x) - float64(p.x)) + (float64(q.y) - float64(p.y)))
}

//Take the l[x] Point and calculate all distances from it?
func calculateDist(l []Point, x int) []float64 {
	dep := l[x]
	dist := []float64{}
	//should filter dep no?
	for _, p := range l {
		dist = append(dist, dep.distanceTo(p))
	}
	return dist
}

//should yield the index no to match the point?
func minInASlice(l []float64) int {
	var max float64
	var index int
	for i, d := range l {
		if d > max {
			max = d
			index = i
		}
	}
	return index
}

func main() {
	var N int
	fmt.Scan(&N)
	c := []Point{}
	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
		c = append(c, Point{X, Y})
	}
	l := calculateDist(c, 0)
	i := minInASlice(l)
	log.Println(l, c[i])
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("distance") // Write answer to stdout
}
