//copy pasted from there is no spoone
//should modify it with tron to take Point
//on a grid into account?
const (
	WIDTH=30
	HEIGHT=20
)

//should i write a grid type?
type Cell struct{
	x,y int
	visited bool
}

//a simple grid made of cells
var grid=make([][]Cell,height)
for i := 0; i < height; i++ {
    grid[i]=make([]Cell,width)
    for j:= range(grid[i]){
        grid[i][j]=Cell{i,j}
    }
}

log.Println(grid)
