//This is an attempt to retain the state of the playing field f//for this game

//dimension of the field
const (
		WIDTH  = 16001
		HEIGHT = 7501
)

//a goal is basically made of 2 poles
type Goal struct{
    pole1 Position
    pole2 Position
}

//pass the coordinates of the poles
func newGoal(p1x,p1y,p2x,p2y int){
    if myTeamId == 0{
        return Goal{
            pole1=newPosition(0,1750)
            pole2=newPosition(0,5750)
        }
    }else {
        return Goal{
            pole1=newPosition(16000,1750)
            pole2=newPosition(16000,5750)
            }
     }
}

//a playing field is made of snaffles,wiz,2 goals and a score
type PlayingField struct{
    snaffles []Snaffle
    wizards []Wizard
    myGoal Goal
    oppGoal Goal
    score int
}


