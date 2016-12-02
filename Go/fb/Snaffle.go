//package fb

//should use position here
type Snaffle struct {
	entityId   int
	entityType string
	x          int
	y          int
	vx         int
	vy         int
    state   int
    thrownBy   int
    grabbedBy int
}

func newSnaffle(id int, etype string, x, y, vx, vy, state int) Snaffle {
		return Snaffle{
			entityId:   id,
			entityType: etype,
			x:          x,
			y:          y,
			vx:         vx,
			vy:         vy,
			state:      state,
            //thrownBy:   0
		}
}

func (s Snaffle) getSnafflePos() Position {
	pos := newPosition(s.x, s.y)
	return pos
}
