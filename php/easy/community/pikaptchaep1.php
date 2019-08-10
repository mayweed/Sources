<?php
class Point {
    public $x;
    public $y;
}
class Cell extends Point{
    public $what;
    public $numPassage = 0;
    
    /* does not work, yield NULL for all values, dont know why...
    function __construct($x,$y,$what){
        $this -> x;
        $this -> y;
        $this -> what;
    }
  */
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

//would be better to got that as a method
for ($y = 0; $y < $height; $y++)
{
    for ($x = 0; $x < $width;$x++){
         //first x
        if ($x-1 >= 0 && $grid[$y][$x-1]-> what == "0"){
            $grid[$y][$x] -> numPassage += 1;
        }
        if ($x+1 < $width && $grid[$y][$x+1]-> what == "0"){
            $grid[$y][$x]->numPassage += 1;
        }
        //then y
        if ($y-1 >= 0 && $grid[$y-1][$x]-> what == "0"){
            $grid[$y][$x] -> numPassage += 1;
        }
        if ($y+1 < $height && $grid[$y+1][$x]-> what == "0"){
            $grid[$y][$x] -> numPassage += 1;
        }
    }
}
/*
//var_dump($grid[0][1]-> what,$grid[0][1]->numPassage);
//print grid
//should be a method!!
for ($i = 0; $i < $height; $i++)
{
    for ($j =0; $j < $width;$j++){
            echo($grid[$i][$j] -> what);
    }
    echo("\n");
}
 */
//grid output
//method here!!
for ($i = 0; $i < $height; $i++)
{
    for ($j =0; $j < $width;$j++){
        if ($grid[$i][$j] -> what == "#"){
            echo($grid[$i][$j] -> what);
        }else{
            echo($grid[$i][$j]->numPassage);
        }
    }
    echo("\n");
}
?>
