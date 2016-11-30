package fb

//should use position here
type Wizard struct {
	entityId   int
	entityType string
	x          int
	y          int
	vx         int
	vy         int
	state      int
}

func newWizard(id int, etype string, x, y, vx, vy, state int) Wizard {
	if etype != "SNAFFLE" {
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
}

func (w Wizard) hasGrabbedSnaffle() bool {
	if w.state == 1 {
		return true
	} else {
		return false
	}
}
