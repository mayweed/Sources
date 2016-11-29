package main

import "fmt"
//import "os"

func main() {
    const (
		HEIGHT=7501
		WIDTH=16001
    )
    // myTeamId: if 0 you need to score on the right of the map, if 1 you need to score on the left
    var myTeamId int
    fmt.Scan(&myTeamId)
	var oppGoal Position
	switch myTeamId{
	case 0:
		//don't need that for the moment right?
		//myGoal=newPosition(0,3750)
		oppGoal=newPosition(16000,3750)
		case 1:
		//myGoal=newPosition(16000,3750)
		oppGoal=newPosition(0,3750)
	}
        
    for {
        // entities: number of entities still in game
        var entities int
        fmt.Scan(&entities)
		var myWiz []Entity
		var snaffle []Entity
		var oppWiz []Entity
        for i := 0; i < entities; i++ {
            // entityType: "WIZARD", "OPPONENT_WIZARD" or "SNAFFLE" (or "BLUDGER" after first league)
            // state: 1 if the wizard is holding a Snaffle, 0 otherwise
            var entityId int
            var entityType string
            var x, y, vx, vy, state int
			fmt.Scan(&entityId, &entityType, &x, &y, &vx, &vy, &state)
			if entityType=="WIZARD"{
				myWiz=append(myWiz,newEntity(entityId, entityType, x, y, vx, vy, state))
			} else if entityType=="OPPONENT_WIZARD"{
				oppWiz=append(oppWiz,newEntity(entityId, entityType, x, y, vx, vy, state))
			} else if entityType=="SNAFFLE"{
				snaffle=append(snaffle,newEntity(entityId, entityType, x, y, vx, vy, state))
			}
	}
				
	//check for the closest snaffle?
	func pickSnaffle(wizard Entity,snaffles []Entity) Entity{
	    var best=WIDTH
		var closestSnaffle Entity
	    for _,snaffle :=range snaffles{
	        distance:=distEntity(wizard,snaffle)
	        if distance < best{
				best=distance
				closestSnaffle=snaffle
			}
		}
		return closestSnaffle
	}
	//move to somewhere not right:(0 <= thrust <= 150, 0 <= power <= 500)
	func command (arg string,dest Position,thrust int) string{
		if arg=="move"{
			fmt.Printf("MOVE" +" dest.x"+" dest.y"+" thrust\n")
		}else if arg=="throw"{
			fmt.Printf("THROW"+" dest.x"+" dest.y"+" thrust\n")
		}
	}
	//Find best move
		//here: loop on wizard pick a snaffle and move to it?
		//should not include command in the loop!!!
		//Needs two lines for each wiz considered separately!!
	var bestSnaffle Entity
		for _,wiz := range myWiz{
			if wiz.state==0{
				//no snaffle
				bestSnaffle=pickSnaffle(wiz,snaffle)
				destination:=newPosition(bestSnaffle.x,bestSnaffle.y)
				command("move",destination,100)
				//break
			}else if wiz.state==1{
				command("throw",oppGoal,400)
			}
		}
	
            // Edit this line to indicate the action for each wizard 
            // i.e.: "MOVE x y thrust" or "THROW x y power"
            //fmt.Printf("MOVE 8000 3750 100\n")
       
    }
}
