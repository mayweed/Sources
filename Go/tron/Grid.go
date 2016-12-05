//copy pasted from there is no spoone
//should modify it with tron to take Point
//on a grid into account?
var grid=make([][]string,height)
    for i := 0; i < height; i++ {
        scanner.Scan()
        inputs:=strings.Split(scanner.Text(),"")
        grid[i]=make([]string,width)
        for j:= range(grid[i]){
            //no need of that, could compare strings...
            //x,_:= strconv.Atoi(inputs[j])
            grid[i][j]=inputs[j]
        }
    }
    log.Println(grid)
