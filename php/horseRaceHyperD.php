<?php

function distance ($v1, $e1,$v2,$e2){
   $result = abs($v2-$v1) + abs($e2-$e1);
   return $result;
}

fscanf(STDIN, "%d",
    $N
);
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
    //if ($i >= 1)
    //calculate dist between $t[$i] and $t[$i-1]
    //$dist = distance($t[$i]["v"],$t[$i]["e"],$t[$i-1]["v"],$t[$i-1]["e"])
}

// Write an action using echo(). DON'T FORGET THE TRAILING \n
// To debug (equivalent to var_dump): error_log(var_export($var, true));

//var_dump($t);
//echo("42\n");
echo ($t[2]["v"]);
?>
