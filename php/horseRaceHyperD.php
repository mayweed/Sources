<?php

function distance ($v1, $e1,$v2,$e2){
   $result = abs($v2-$v1) + abs($e2-$e1);
   return $result;
}

fscanf(STDIN, "%d",
    $N
);

$min= 10000000;

for ($i = 0; $i < $N; $i++)
{
    fscanf(STDIN, "%d %d",
        $V,
        $E
    );
    //or $t[]
    $t[$i]=
        array(
            "v" =>$V,
            "e" => $E
            );
            
   foreach ($t as $x){
       foreach($t as $y){
           //avoid computing dist between same points
           if ($x == $y) continue;
           
           //you calculate the distance between each
           $d=distance ($x["v"],$x["e"],$y["v"],$y["e"]);
           //echo ("$d\n");
           
           if ($d < $min) $min=$d;
       }
   }
    
    //algo wrong: you must compare each element: $t[0] to $t[1] -> $t[9]
    //$t[1] to all elements above and the one below etc...
    //this is not so simple...
    //if ($i >= 1){
        //calculate dist between $t[$i] and $t[$i-1]   
     //   $dist =  distance($t[$i]["v"],$t[$i]["e"],$t[$i-1]["v"],$t[$i-1]["e"]);
        
        //echo("$dist\n");
        //echo("$min\n");
    //}
        
    //put dist in an array and sort the array in ascending order (arr[0]= answer)
    
}

// Write an action using echo(). DON'T FORGET THE TRAILING \n
// To debug (equivalent to var_dump): error_log(var_export($var, true));

//var_dump($t);
//echo("42\n");

echo("$min\n");
//echo ($t[2]["v"]);
?>