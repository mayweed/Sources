//package fb

//should use position here
type Wizard struct {
	entityId   int
	entityType string
	x          int
	y          int
	vx         int
	vy         int
	state      int
    hasJustThrown bool
    grabbedSnaffle int
}

func newWizard(id int, etype string, x, y, vx, vy, state int) Wizard {
	return Wizard{
		entityId:   id,
		entityType: etype,
		x:          x,
		y:          y,
		vx:         vx,
		vy:         vy,
		state:      state,
        hasJustThrown: false
	}
}

//if dist between snaffle and wiz <=250, grabbed
//wiz radius==400, snaffle radius==150, to grab the center
//of a snaffle must be in the radius of the wiz...
//what about distEntity?
func (w Wizard) GrabbedSnaffle(snaffles []Snaffle) Snaffle{
    for _,snaffle:=range snaffles{
        if dist(snaffle.x,snaffle.y,w.x,w.y)=<250{
            snaffle.grabbedBy=wiz.entityId
            return snaffle
        }
    }
}
func (w Wizard) hasGrabbleSnaffle() bool {
	if w.state == 1 {
		return true
	} else {
		return false
	}
}

//it's either mine or the opp one
func (w Wizard) distToGoal(wizX,wizY,goalX,goalY int) int{
    return dist(w.x,w.y,goalX,goalY)
}
