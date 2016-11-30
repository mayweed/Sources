package fb

// TODO:
//- wiz should be considered separately: one moves the snaffle and tries to intercept those thrown
//by opponent, an other way keeps throwing to score!!
//Find best move
//here: loop on wizard pick a snaffle and move to it?
//Needs two lines for each wiz considered separately!!
//TODO:one will go for the closest snaffle to his position
//the other one for the closest to oppGoal.
//should I make a map of possible moves and evaluate them ??

//check for the nearest snaffle?
func pickNearestSnaffle(wiz Wizard,snaffles []Snaffle) Snaffle{
    var best=WIDTH
	var nearestSnaffle Snaffle
    for _,snaffle :=range snaffles{
        distance:=distEntity(wiz,snaffle)
        if distance < best{
			best=distance
			nearestSnaffle=snaffle
		}
	}
	return nearestSnaffle
}

//check for closest snaffle from oppGoal
func pickClosestSnaffle(oppGoal Position,snaffles []Snaffle) Snaffle{
	var best=WIDTH
	var closestSnaffle Snaffle
    for _,snaffle :=range snaffles{
        distance:=dist(oppGoal.x,snaffle.x,oppGoal.y,snaffle.y)
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

//check wiz to find best moves??
//a func that yields a map
func findBestMove(myWiz []Wizard) map[Wizard]string{
	var choices= make(map[Wizard]string) //a map with a wiz and a tag for action??
	var bestSnaffle Entity
	for _,wiz := range myWiz{
	//state is often 0, two wiz same direction...
	if wiz.state==0{
		//no snaffle
		bestSnaffle=pickSnaffle(wiz,snaffle)
		destination:=newPosition(bestSnaffle.x,bestSnaffle.y)
		//command("move",destination,100)
		//break
		choices[wiz]=
			}else if wiz.hasGrabbedSnaffle(){
				command("throw",oppGoal,400)
			}
		}
