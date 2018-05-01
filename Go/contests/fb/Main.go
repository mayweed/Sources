package main

import "fmt"

func main() {
	const (
		HEIGHT = 7501
		WIDTH  = 16001
		MAX_THRUST = 150
		MAX_POWER=500
	)
	// myTeamId: if 0 you need to score on the right of the map, if 1 you need to score on the left
	var myTeamId int
	fmt.Scan(&myTeamId)
	var oppGoal Position
	switch myTeamId {
	case 0:
		//don't need that for the moment right?
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
		var snaffle []Snaffle
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
				snaffle = append(snaffle, newSnaffle(entityId, entityType, x, y, vx, vy, state))
			}
		}
		var bestSnaffle Snaffle
		var destination Position
		var oldDestination Position
		for _, wiz := range myWiz {
			//Pick a Snaffle
			bestSnaffle = pickNearestSnaffle(wiz, snaffles)
			dist:=distAtSnaffle(wiz,snaffles)
			if wiz.hasGrabbedSnaffle() {
				command("throw", oppGoal, MAX_POWER)
			}else{
				//wiz grabs no snaffle
				destination = newPosition(bestSnaffle.x, bestSnaffle.y)
				if destination==oldDestination{
					//should pick a second snaffle here...
				command("move", destination, MAX_THRUST)
				oldDestination=destination
			}
		}
	}
}