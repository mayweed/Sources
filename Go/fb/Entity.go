package fb

//should use position here
type Entity struct{
	entityId:int,
	//either wiz,opponent wiz or snaffle
	entityType:string,
	x:int,
	y:int,
	vx:int,
	vy:int,
	state:int,
}

func newEntity (id int,etype string,x,y,vx,vy,state int) Entity{
	return Entity{
		entityId:id,
		entityType:etype,
		x:x,
		y:y,
		vx:vx,
		vy:vy,
		state:state,
	}
}

func (e Entity) getType() string{
	return e.entityType
}
