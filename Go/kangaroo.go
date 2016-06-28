package main
import "fmt"

func main(){
    x1:=0
    v1:=3
    x2:=4
    v2:=2
    var jump_X int=0
    var jump_Y int=0

for x1<100 && x2<100{
    landing_position:=x1+v1
    x1=landing_position
    jump_X+=1

    landing_position2:=x2+v2
    x2=landing_position2
    jump_Y+=1
    fmt.Println(x1,x2)
    if x1==x2 && jump_X==jump_Y{
        fmt.Println("YES")
        return
        }
    }
    fmt.Println("NO")
}

