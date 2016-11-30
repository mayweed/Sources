package fb

//should use position here
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
	if etype == "SNAFFLE" {
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
}

func (s Snaffle) getSnafflePos() Position {
	pos := newPosition(s.x, s.y)
	return pos
}
