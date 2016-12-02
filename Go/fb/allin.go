package main

import "fmt"
import "math"
import "log"

//CONSTS
const (
	HEIGHT = 7501
	WIDTH  = 16001
)

//POSITION
type Position struct {
	x, y int
}

func newPosition(x, y int) Position {
	return Position{
		x: x,
		y: y,
	}
}

//WIZARDS
type Wizard struct {
	entityId      int
	entityType    string
	x             int
	y             int
	vx            int
	vy            int
	state         int
	hasJustThrown bool
}

func newWizard(id int, etype string, x, y, vx, vy, state int) Wizard {
	return Wizard{
		entityId:      id,
		entityType:    etype,
		x:             x,
		y:             y,
		vx:            vx,
		vy:            vy,
		state:         state,
		hasJustThrown: false,
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
	x          int
	y          int
	vx         int
	vy         int
	state      int
}

func newSnaffle(id int, etype string, x, y, vx, vy, state int) Snaffle {
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

func (s Snaffle) getSnafflePos() Position {
	pos := newPosition(s.x, s.y)
	return pos
}

//UTILS
func dist(x1, y1, x2, y2 int) int {
	dist := math.Sqrt((float64(x1)-float64(x2))*(float64(x1)-float64(x2)) + (float64(y1)-float64(y2))*(float64(y1)-float64(y2)))
	return int(dist)
}

func distEntity(wizard Wizard, snaffle Snaffle) int {
	distance := dist(wizard.x, snaffle.x, wizard.y, snaffle.y)
	return int(distance)
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
func pickClosestSnaffle(oppGoal Position, snaffles []Snaffle) Snaffle {
	var best = WIDTH
	var closestSnaffle Snaffle
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
func command(arg string, dest Position, thrust int) {
	if arg == "move" {
		fmt.Printf("MOVE %d %d %d\n", dest.x, dest.y, thrust)
	} else if arg == "throw" {
		fmt.Printf("THROW %d %d %d\n", dest.x, dest.y, thrust)
	}
}

//MAIN
func main() {

	// myTeamId: if 0 you need to score on the right of the map, if 1 you need to score on the left
	var myTeamId int
	fmt.Scan(&myTeamId)
	var oppGoal Position
	switch myTeamId {
	case 0:
		//myGoal=newPosition(0,3750)
		oppGoal = newPosition(16000, 3750)
	case 1:
		//myGoal=newPosition(16000,3750)
		oppGoal = newPosition(0, 3750)
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
				myWiz = append(myWiz, newWizard(entityId, entityType, x, y, vx, vy, state))
			} else if entityType == "OPPONENT_WIZARD" {
				oppWiz = append(oppWiz, newWizard(entityId, entityType, x, y, vx, vy, state))
			} else if entityType == "SNAFFLE" {
				snaffles = append(snaffles, newSnaffle(entityId, entityType, x, y, vx, vy, state))
			}
		}
		//Find best move
		//SHOULD MOVE THAT ELSEWHERE (move.go?)
		//check wiz to find best moves??
		//a func that yields a map
		//func findBestMove(myWiz []Wizard) map[Wizard]string{
		//	var choices= make(map[Wizard]string) //a map with a wiz and a tag for action??
		var bestSnaffle Snaffle
		var closestSnaffle Snaffle
		var wizPos Position
		var oldWizPos Position
		var destination Position
		for _, wiz := range myWiz {
			//state is often 0, two wiz same direction...
			wizPos = newPosition(wiz.x, wiz.y)
			if wiz.state == 0 {
				//no snaffle
				bestSnaffle = pickNearestSnaffle(wiz, snaffles)
				destination = newPosition(bestSnaffle.x, bestSnaffle.y)
				//This loop is USELESS!!!
				if oldWizPos == wizPos {
					log.Println(oldWizPos, wizPos)
					//change destination for the second one...
					closestSnaffle = pickClosestSnaffle(oppGoal, snaffles)
					destination = newPosition(closestSnaffle.x, closestSnaffle.y)
				}
				//oldDestination=destination
				command("move", destination, 120)
			} else if wiz.hasGrabbedSnaffle() {
				command("throw", oppGoal, 500)
				wiz.hasJustThrown = true
				//if a wiz has just thrown must pursue the ball to score!!
				//should mark the snaffle and f*ckin run after it to trhow it max!!
			}
			oldWizPos = newPosition(wiz.x, wiz.y)
			log.Println(wiz.hasJustThrown)
		}
	}
}
