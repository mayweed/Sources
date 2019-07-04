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

    $t[$i]=
        array(
            "v" =>$V,
            "e" => $E
            );
            
    //naive solution
   //try to compare the min velocity/eleg with the 2d min one
   //you sort them in ascending order...
   foreach ($t as $x){
       foreach($t as $y){
           //avoid computing dist between same points
           if ($x == $y) continue;
           
           //you calculate the distance between each
           $d=distance ($x["v"],$x["e"],$y["v"],$y["e"]);
           
           if ($d < $min) $min=$d;
       }
   }
}

// Write an action using echo(). DON'T FORGET THE TRAILING \n
echo("$min\n");
?>
