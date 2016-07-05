package main

import "fmt"
//import "os"
import "math"

type Point struct {
    x,y float64
}
func (p Point) distance(from_p Point) float64{
      return math.Sqrt(((p.x-from_p.x)*(p.x-from_p.x)) +((p.y-from_p.y)*(p.y-from_p.y)))
}
func NewPosition(x,y float64) Point{
    return Point{
        x:x,
        y:y,
    }
}

type Pod struct{
    pos Point
    angle float64
}
func NewPod(position Point,nextCheckpointAngle float64) Pod{
    return Pod{
        pos:position,
        angle:nextCheckpointAngle,
    }
}
func (p Pod) getAngle(targetPoint Point) float64{
    var d float64
    d=p.pos.distance(targetPoint)

    var dx float64
    dx=(targetPoint.x - p.pos.x)/d
    var dy float64
    dy=(targetPoint.y-p.pos.y)/d

    var angle float64
    angle=math.Acos(dx)*180.0 / math.Pi

    if dy < 0{
        angle=360.0-angle
    }
    return angle
}
func (p Pod) diffAngle(targetPoint Point) float64{
    a:=p.getAngle(targetPoint)
    var left,right float64
    if p.angle <=a{
        right=a-p.angle
    }else{
        right=360.0 - p.angle+a
    }
    if p.angle >=a{
        left=p.angle-a
    }else{
        left=p.angle+360.0-a
    }

    if right < left{
        return right
    }else{
        return -left
    }
}
func (p Pod) rotate(targetPoint Point) {
    a:=p.diffAngle(targetPoint)

    p.angle+=a

    if p.angle >= 360.0{
        p.angle=p.angle -360.0
    }else if p.angle < 0.0{
        p.angle+=360.0
    }
}

func main() {
    for {
        // nextCheckpointX: x position of the next check point
        // nextCheckpointY: y position of the next check point
        // nextCheckpointDist: distance to the next checkpoint
        // nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
        var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
        fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)

        var opponentX, opponentY int
        fmt.Scan(&opponentX, &opponentY)


        // fmt.Fprintln(os.Stderr, "Debug messages...")

        // You have to output the target position
        // followed by the power (0 <= thrust <= 100)
        // i.e.: "x y thrust"
        var speed int
        if nextCheckpointAngle > 90 || nextCheckpointAngle < -90{
            speed=0
        }else{
            speed=100
        }
        //if nextCheckpointDist < 600{
        fmt.Printf("%d %d %d\n", nextCheckpointX, nextCheckpointY,speed)
        //}else{
        //    fmt.Printf("%d %d %d\n", nextCheckpointX, nextCheckpointY,speed)
        //}

    }
}
