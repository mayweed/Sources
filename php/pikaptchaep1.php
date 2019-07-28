<?php
class Point {
    public $x;
    public $y;
}
class Cell extends Point{
    public $what;
    public $numPassage;
    
    /* does not work, yield NULL for all values, dont know why...
    function __construct($x,$y,$what){
        $this -> x;
        $this -> y;
        $this -> what;
    }
     */
    function countPassage(){
        //first x
        if ($grid[$this ->y][$this -> x-1 > 0] &&$grid[$this ->y][$this -> x-1 > 0]-> what ==0){
            $this -> numPassage += 1;
        }

        if ($grid[$this ->y][$this -> x+1 < $width] &&$grid[$this ->y][$this -> x+1 < $width]-> what ==0){
            $this -> numPassage += 1;
        }
        //then y
        if ($grid[$this ->y-1 > 0][$this -> x] &&$grid[$this ->y-1>0 ][$this -> x]-> what ==0){
            $this -> numPassage += 1;
        }
        if ($grid[$this ->y+1 < $height][$this -> x] &&$grid[$this ->y+1<$height ][$this -> x]-> what ==0){
            $this -> numPassage += 1;
        }
    }
}
//class Grid { ... } ???

fscanf(STDIN, "%d %d", $width, $height);

for ($y = 0; $y < $height; $y++)
{
        fscanf(STDIN, "%s", $line);
        $split = str_split($line);
        for ($x = 0; $x < $width;$x++){
            $grid[$y][$x]=new Cell();
            //test w/o construct
            $grid[$y][$x]-> y = $y;
            $grid[$y][$x]-> x = $x;
            $grid[$y][$x]-> what = $split[$x];
        }
}

$grid[0][4]->countPassage();

var_dump($grid[0][4]->numPassage,$height,$width);

for ($i = 0; $i < $height; $i++)
{
    echo("#####\n");
}
?>
