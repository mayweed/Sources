package main

import "math"

type Point struct {
	x, y int
}

func addPoint(p1, p2 Point) Point {
	var dest Point
	dest = Point{(p1.x + p2.x), (p1.y + p2.y)}
	return dest
}

func getSector(p Point) int {
	zone := math.Ceil(float64(p.x+1)/5.0) + math.Floor(float64(p.y/5.0)*3)
	return int(zone)
}
