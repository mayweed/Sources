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

func main() {
	var N int
	fmt.Scan(&N)
	c := []Point{}
	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
		c = append(c, Point{X, Y})
	}
	log.Println(calculateDist(c, 0))
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("distance") // Write answer to stdout
}
