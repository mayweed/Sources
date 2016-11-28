package main

import "fmt"
//import "os"

//to output a command either move or throw
const (
    MOVE = "MOVE"
    THROW ="THROW"
    )
type command struct{
    ctype string
    destination position
    power int
}

func main() {
    const (
		HEIGHT=7501
		WIDTH=16001
    )
    // myTeamId: if 0 you need to score on the right of the map, if 1 you need to score on the left
    var myTeamId int
    fmt.Scan(&myTeamId)
	var myGoal Position
	switch myTeamId{
		case 0:
			 myGoal=NewPosition(0,3750)
		case 1:
			 myGoal=NewPosition(16000,3750)
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
			    
        //for i := 0; i < 2; i++ {
            
            // fmt.Fprintln(os.Stderr, "Debug messages...")
            
            // Edit this line to indicate the action for each wizard (0 <= thrust <= 150, 0 <= power <= 500)
            // i.e.: "MOVE x y thrust" or "THROW x y power"
            //fmt.Printf("MOVE 8000 3750 100\n")
        //}
    }
}
