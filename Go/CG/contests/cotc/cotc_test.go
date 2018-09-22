package main

import "testing"

func TestNeighbour(t *testing.T) {
	var orientation = 2
	var initialPos = Point{2, 4}
	p := initialPos.neighbour(orientation)
	t.Errorf("Initial Position x : %d y:%d New position x : %d y: %d\n", initialPos.x, initialPos.y, p.x, p.y)
}
