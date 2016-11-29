package fb

import "math"

func dist(x1, y1, x2, y2 int) float64 {
	dist := math.Sqrt((float64(x1)-float64(x2))*(float64(x1)-float64(x2)) + (float64(y1)-float64(y2))*(float64(y1)-float64(y2)))
	return dist
}

func distEntity(wizard, snaffle Entity) int {
	distance := dist(wizard.x, snaffle.x, wizard.y, snaffle.y)
	return int(distance)
}
