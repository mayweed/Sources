package main

import "fmt"

//import "os"

func main() {
	const (
		HEIGHT = 7501
		WIDTH  = 16001
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
	//SHOULD MOVE THAT ELSEWHERE (move.go?)
	//check wiz to find best moves??
	//a func that yields a map
	//func findBestMove(myWiz []Wizard) map[Wizard]string{
	//	var choices= make(map[Wizard]string) //a map with a wiz and a tag for action??
	var bestSnaffle Snaffle
	var oldDestination Position
	for _,wiz := range myWiz{
	//state is often 0, two wiz same direction...
	if wiz.state==0{
		//no snaffle 
		bestSnaffle=pickNearestSnaffle(wiz,snaffle)
		destination:=newPosition(bestSnaffle.x,bestSnaffle.y)
		if destination==oldDestination{
			//change destination for the second one...
			closestSnaffle=pickClosestSnaffle(oppGoal,snaffles)
		}
		oldDestination=destination
		//command("move",destination,100)
		//break
		//choices[wiz]=
		}else if wiz.hasGrabbedSnaffle(){
			command("throw",oppGoal,400)
		}
	}
}
