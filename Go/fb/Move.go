//package fb

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
func command(arg string, dest Position, thrust int) {
	if arg == "move" {
		fmt.Printf("MOVE %d %d %d\n", dest.x, dest.y, thrust)
	} else if arg == "throw" {
		fmt.Printf("THROW %d %d %d\n", dest.x, dest.y, thrust)
	}
}
