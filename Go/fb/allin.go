package main

import (
	"fmt"
	"math"
	//"log"
)

//POINT
type point struct {
	x, y float64
}

func newPoint(x, y float64) point {
	return point{
		x: x,
		y: y,
	}
}
func dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func distEntity(wiz, snaf entity) float64 {
	return dist(wiz.x, wiz.y, snaf.x, snaf.y)
}

type entity struct {
	entityId   int
	entityType string
	vx         int
	vy         int
	state      int
	x          float64
	y          float64
}

type playground struct {
	width    float64
	height   float64
	myGoal   point
	oppGoal  point
	myWiz    []entity
	oppWiz   []entity
	snaffles []entity
	bludgers []entity
}

func snaffleCarried(wiz entity, snaffles []entity) entity {
	//why cant I just return inside if ? got a no return error??
	var carried entity
	for _, snaf := range snaffles {
		if (snaf.x == wiz.x) && (snaf.y == wiz.y) {
			carried = snaf
			break
		}
	}
	return carried
}

//check for the nearest snaffle?
func (p playground) pickNearestSnaffle(wiz entity, snaffles []entity) entity {
	var best = p.width
	var nearestSnaffle entity
	for _, snaffle := range snaffles {
		distance := distEntity(wiz, snaffle)
		//log.Println("Snaffle: ",snaffle.entityId,"Distance: ",distance)
		if distance < best {
			best = distance
			nearestSnaffle = snaffle
		}
	}
	return nearestSnaffle
}

//check for closest snaffle from oppGoal if dist to closest is < to
//nearest go for it!!
func (p playground) pickClosestSnaffle(oppGoal point, snaffles []entity) entity {
	var best = p.width
	var closestSnaffle entity
	for _, snaffle := range snaffles {
		distance := dist(oppGoal.x, snaffle.x, oppGoal.y, snaffle.y)
		if distance < best {
			best = distance
			closestSnaffle = snaffle
		}
	}
	return closestSnaffle
}

//move to somewhere not right:(0 <= thrust <= 150, 0 <= power <= 500)
//should I use sprintf and yield a string?
func command(arg string, dest point, thrust int) {
	if arg == "move" {
		fmt.Printf("MOVE %d %d %d\n", int(dest.x), int(dest.y), thrust)
	} else if arg == "throw" {
		fmt.Printf("THROW %d %d %d\n", int(dest.x), int(dest.y), thrust)
	}
}

//MAIN
func main() {
	//playground
	pg := playground{
		width:  16001.,
		height: 7501.,
	}
	// myTeamId: if 0 you need to score on the right of the map, if 1 you need to score on the left
	var myTeamId int
	fmt.Scan(&myTeamId)

	switch myTeamId {
	case 0:
		pg.myGoal = newPoint(0., 3750.)
		pg.oppGoal = newPoint(16000., 3750.)
	case 1:
		pg.myGoal = newPoint(16000., 3750.)
		pg.oppGoal = newPoint(0., 3750.)
	}

	for {
		var myScore, myMagic int
		fmt.Scan(&myScore, &myMagic)

		var opponentScore, opponentMagic int
		fmt.Scan(&opponentScore, &opponentMagic)

		// entities: number of entities still in game
		var entities int
		fmt.Scan(&entities)

		for i := 0; i < entities; i++ {
			// entityType: "WIZARD", "OPPONENT_WIZARD" or "SNAFFLE" (or "BLUDGER" after first league)
			// state: 1 if the wizard is holding a Snaffle, 0 otherwise
			var entityId int
			var entityType string
			var x, y, vx, vy, state int
			fmt.Scan(&entityId, &entityType, &x, &y, &vx, &vy, &state)
			if entityType == "WIZARD" {
				pg.myWiz = append(pg.myWiz, entity{entityId, entityType, vx, vy, state, float64(x), float64(y)})
			} else if entityType == "OPPONENT_WIZARD" {
				pg.oppWiz = append(pg.oppWiz, entity{entityId, entityType, vx, vy, state, float64(x), float64(y)})
			} else if entityType == "SNAFFLE" {
				pg.snaffles = append(pg.snaffles, entity{entityId, entityType, vx, vy, state, float64(x), float64(y)})
			} else if entityType == "BLUDGER" {
				pg.bludgers = append(pg.bludgers, entity{entityId, entityType, vx, vy, state, float64(x), float64(y)})
			}
		}

		//pick the nearest, go for it...
		var bestSnaffle entity
		for _, wiz := range pg.myWiz {
			if wiz.state == 1 {
				fmt.Printf("THROW %d %d 500\n", int(pg.oppGoal.x), int(pg.oppGoal.y))
			} else {
				bestSnaffle = pg.pickNearestSnaffle(wiz, pg.snaffles)
				//log.Println(wiz.entityId, int(wiz.x), int(wiz.y), bestSnaffle.entityId)
				fmt.Printf("MOVE %d %d 150\n", int(bestSnaffle.x), int(bestSnaffle.y))
			}

		}
		pg.myWiz = []entity{}
		pg.oppWiz = []entity{}
		pg.bludgers = []entity{}
	}
}
