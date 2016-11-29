package fb

type Position struct {
	x, y int
}

func newPosition(x, y int) Position {
	return Position{
		x: x,
		y: y,
	}
}
