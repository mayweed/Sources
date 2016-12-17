package main

import "fmt"
import "math"

//import "log"

//CONSTS
const (
	HEIGHT     = 7501.
	WIDTH      = 16001.
	MAX_POWER  = 500
	MAX_THRUST = 150
)

//POSITION
type Point struct {
	x, y float64
}

func newPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

//WIZARDS
type Wizard struct {
	entityId   int
	entityType string
	x          float64
	y          float64
	vx         int
	vy         int
	state      int
}

func newWizard(id, vx, vy, state int, etype string, x, y float64) Wizard {
	return Wizard{
		entityId:   id,
		entityType: etype,
		x:          x,
		y:          y,
		vx:         vx,
		vy:         vy,
		state:      state,
	}
}

func (w Wizard) hasGrabbedSnaffle() bool {
	if w.state == 1 {
		return true
	} else {
		return false
	}
}

//SNAFFLES
type Snaffle struct {
	entityId   int
	entityType string
	x          float64
	y          float64
	vx         int
	vy         int
	state      int
}

func newSnaffle(id, vx, vy, state int, etype string, x, y float64) Snaffle {
	return Snaffle{
		entityId:   id,
		entityType: etype,
		x:          x,
		y:          y,
		vx:         vx,
		vy:         vy,
		state:      state,
	}
}

func (s Snaffle) getSnafflePos() Point {
	pos := newPoint(s.x, s.y)
	return pos
}

//UTILS
func dist(x1, y1, x2, y2 float64) float64 {
	dist := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	return dist
}

func distEntity(wizard Wizard, snaffle Snaffle) float64 {
	distance := dist(wizard.x, snaffle.x, wizard.y, snaffle.y)
	return distance
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
		if distance < best {
			best = distance
			nearestSnaffle = snaffle
		}
	}
	return nearestSnaffle
}

//check for closest snaffle from oppGoal
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
				myWiz = append(myWiz, newWizard(entityId, vx, vy, state, entityType, float64(x), float64(y)))
			} else if entityType == "OPPONENT_WIZARD" {
				oppWiz = append(oppWiz, newWizard(entityId, vx, vy, state, entityType, float64(x), float64(y)))
			} else if entityType == "SNAFFLE" {
				snaffles = append(snaffles, newSnaffle(entityId, vx, vy, state, entityType, float64(x), float64(y)))
			}
		}
		//Find best move
		//here: loop on wizard pick a snaffle and move to it?
		//should not include command in the loop!!!
		//Needs two lines for each wiz considered separately!!
		//SHOULD MOVE THAT ELSEWHERE (move.go?)
		//check wiz to find best moves??
		//a func that yields a map
		//func findBestMove(myWiz []Wizard) map[Wizard]string{
		//	var choices= make(map[Wizard]string) //a map with a wiz and a tag for action??
		//var closestSnaffle Snaffle
		var destination Point
		for _, wiz := range myWiz {
			var bestSnaffle Snaffle
			//state is often 0, two wiz same direction...
			if wiz.hasGrabbedSnaffle() {
				command("throw", oppGoal, MAX_POWER)
			} else {
				//no snaffle
				bestSnaffle = pickNearestSnaffle(wiz, snaffles)
				//from the codingame cast!!
				//bestSnaffle=snaffles[i%len(myWiz)]
				destination = newPoint(bestSnaffle.x, bestSnaffle.y)
				command("move", destination, MAX_THRUST)
			}
		}
	}
}
