package main

import "fmt"
import "math"
import "math/rand"
import "os"

// POSITION
type position struct{
    x,y int
}
func NewPosition(x,y int) position{
    return position{
        x:x,
        y:y,
    }
}
func (p position) Move(x,y int) (int,int){
    x=p.x+x
    y=p.y+y
    return x,y
}
func (p position) distFrom(x,y int) float64{
    dist:=math.Sqrt((float64(p.x)-float64(x))*(float64(p.x)-float64(x)) + (float64(p.y)-float64(y))*(float64(p.y)-float64(y)))
    return dist
}
func HomePosition(myTeamId int) position{
    var home position
    switch myTeamId{
        case 0:home=NewPosition(0,0)
        case 1:home=NewPosition(16000,9000)
    }
    return home
}

//ENTITY
type entity struct{
    entityId,entityType,state,value int
    pos position
}
func NewEntity(entityId,entityType,state,value int, pos position) entity{
    return entity{
        entityId : entityId,
        pos:position{},
        entityType:entityType,
        state:state,
        value:value,
    }
}


func main() {
    // bustersPerPlayer: the amount of busters you control
    var bustersPerPlayer int
    fmt.Scan(&bustersPerPlayer)

    // ghostCount: the amount of ghosts on the map
    var ghostCount int
    fmt.Scan(&ghostCount)

    // myTeamId: if this is 0, your base is on the top left of the map, if it is one, on the bottom right
    var myTeamId int
    fmt.Scan(&myTeamId)


    for {
        // entities: the number of busters and ghosts visible to you
        var entities int
        fmt.Scan(&entities)
        var entityId, x, y, entityType, state, value int
        var ghosts_list []entity
        var my_busters_list []entity
        var foe_busters_list []entity

        for i := 0; i < entities; i++ {
            // entityId: buster id or ghost id
            // y: position of this buster / ghost
            // entityType: the team id if it is a buster, -1 if it is a ghost.
            // state: For busters: 0=idle, 1=carrying a ghost.
            // value: For busters: Ghost id being carried. For ghosts: number of busters attempting to trap this ghost.

            fmt.Scan(&entityId, &x, &y, &entityType, &state, &value)
            if entityType==myTeamId{
                b:=NewEntity(entityId,entityType,state,value,NewPosition(x,y))
                my_busters_list=append(my_busters_list,b)
            }
            if entityType==-1{
                ghosts_list=append(ghosts_list,NewEntity(entityId,entityType,state,value,NewPosition(x,y)))
            }
            if entityType!=myTeamId && entityType!=-1{
                b:=NewEntity(entityId,entityType,state,value,NewPosition(x,y))
                foe_busters_list=append(foe_busters_list,b)
            }

        }
        //make them move to see the game
        //if len(ghosts_list)!=0 should go and catch them!!
        for i:=0;i<len(my_busters_list);i++{
            if u:=len(ghosts_list);u!=0{
                //should iterate through ghosts_lists(pop?)
                //then the buster should be blocked in "transport_state" bool?
                fmt.Printf("BUST %d\n",ghosts_list[0].entityId)
                u-=1
                continue
            }else{
                fmt.Printf("MOVE %d %d\n",rand.Intn(16000),rand.Intn(9000))
            }
        }

        fmt.Fprintln(os.Stderr,myTeamId,bustersPerPlayer,ghostCount,entities,len(my_busters_list),len(ghosts_list))
        //fmt.Printf("MOVE 8000 4500\n") // MOVE x y | BUST id | RELEASE
    }
}
