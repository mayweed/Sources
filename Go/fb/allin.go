package main

import "fmt"
import "math"
import "log"

//CONSTS
const (
	HEIGHT     = 7501.
	WIDTH      = 16001.
)

//POINT
type Point struct {
	x, y float64
}

func newPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}
func dist(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func distEntity(wizard Wizard, snaffle Snaffle) float64 {
	return dist(wizard.x, wizard.y, snaffle.x, snaffle.y)
}

//WIZARDS
type Wizard struct {
	entityId   int
	entityType string
	vx         int
	vy         int
	state      int
	x          float64
	y          float64
}

//SNAFFLES
type Snaffle struct {
	entityId   int
	entityType string
	vx         int
	vy         int
	state      int
	x          float64
	y          float64
}

func (s Snaffle) getSnafflePos() Point {
	pos := newPoint(s.x, s.y)
	return pos
}

//COMMANDS? EVALUATOR?
//check for the closest snaffle? Should update snaffle: when one is chosen
//a second best one should be available (2 wizards!!)
//check for the nearest snaffle?
func pickNearestSnaffle(wiz Wizard, snaffles []Snaffle) Snaffle {
	var best = WIDTH
	var nearestSnaffle Snaffle
	for _, snaffle := range snaffles {
		distance := distEntity(wiz, snaffle)
		log.Println("Snaffle: ",snaffle.entityId,"Distance: ",distance)
		if distance < best {
			best = distance
			nearestSnaffle = snaffle
		}
	}
	return nearestSnaffle
}

//check for closest snaffle from oppGoal if dist to closest is < to
//nearest go for it!!
func pickClosestSnaffle(oppGoal Point, snaffles []Snaffle) Snaffle {
	var best = WIDTH
	var closestSnaffle Snaffle
	for _, snaffle := range snaffles {
		//'int' from distEntity!!
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
func command(arg string, dest Point, thrust int) {
	if arg == "move" {
		fmt.Printf("MOVE %d %d %d\n", int(dest.x), int(dest.y), thrust)
	} else if arg == "throw" {
		fmt.Printf("THROW %d %d %d\n", int(dest.x), int(dest.y), thrust)
	}
}

//MAIN
func main() {

	// myTeamId: if 0 you need to score on the right of the map, if 1 you need to score on the left
	var myTeamId int
	fmt.Scan(&myTeamId)
	var oppGoal Point
	switch myTeamId {
	case 0:
		//myGoal=newPosition(0,3750)
		oppGoal = newPoint(16000., 3750.)
	case 1:
		//myGoal=newPosition(16000,3750)
		oppGoal = newPoint(0., 3750.)
	}

	for {
		// entities: number of entities still in game
		var entities int
		fmt.Scan(&entities)
		var myWiz []Wizard
		var snaffles []Snaffle
		var oppWiz []Wizard
		for i := 0; i < entities; i++ {
			// entityType: "WIZARD", "OPPONENT_WIZARD" or "SNAFFLE" (or "BLUDGER" after first league)
			// state: 1 if the wizard is holding a Snaffle, 0 otherwise
			var entityId int
			var entityType string
			var x, y, vx, vy, state int
			fmt.Scan(&entityId, &entityType, &x, &y, &vx, &vy, &state)
			if entityType == "WIZARD" {
				myWiz = append(myWiz, Wizard{entityId, entityType,vx, vy, state, float64(x), float64(y)})
			} else if entityType == "OPPONENT_WIZARD" {
				oppWiz = append(oppWiz, Wizard{entityId, entityType, vx, vy, state, float64(x), float64(y)})
			} else if entityType == "SNAFFLE" {
				snaffles = append(snaffles, Snaffle{entityId, entityType, vx, vy, state, float64(x), float64(y)})
			    }
		    }
		    
		//pick the nearest, go for it...
		var bestSnaffle Snaffle
		for _, wiz := range myWiz {
		    if wiz.state==1 {
			    fmt.Printf("THROW %d %d 500\n",int(oppGoal.x), int(oppGoal.y))
			} else {
			    bestSnaffle = pickNearestSnaffle(wiz, snaffles)
			    //log.Println(wiz.entityId, int(wiz.x), int(wiz.y), bestSnaffle.entityId)
			    fmt.Printf("MOVE %d %d 150\n",int(bestSnaffle.x), int(bestSnaffle.y))
			}
		}
    }
}

            // i.e.: "MOVE x y thrust" or "THROW x y power"
//            fmt.Printf("MOVE 8000 3750 100\n")
