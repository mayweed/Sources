package main

import (
	"fmt"
	"log"
	"math"
)

//POINT
type Point struct {
	x, y float64
}

func (p Point) distance(from_p Point) float64 {
	return math.Sqrt(((p.x - from_p.x) * (p.x - from_p.x)) + ((p.y - from_p.y) * (p.y - from_p.y)))
}

//POD
type Pod struct {
	pos                 Point
	nextCheckpoint      Point
	nextCheckpointDist  float64
	nextCheckpointAngle float64
}

func main() {
	for {
		// nextCheckpointX: x position of the next check point
		// nextCheckpointY: y position of the next check point
		// nextCheckpointDist: distance to the next checkpoint
		// nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
		var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle float64 //int
		fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)
		//myPod := Pod{Point{x, y},Point{nextCheckpointX, nextCheckpointY},nextCheckpointDist, nextCheckpointAngle}

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)

		// You have to output the target position
		// followed by the power (0 <= thrust <= 100)
		// i.e.: "x y thrust"
		var speed int
		if nextCheckpointAngle > 90 || nextCheckpointAngle < -90 {
			speed = 0
		} else {
			speed = 100
		}
		fmt.Printf("%d %d %d\n", nextCheckpointX, nextCheckpointY, speed)

		log.Println(nextCheckpointAngle)

	}
}
