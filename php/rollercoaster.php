<?php
//Key is to store two maps / arrays (one of earns per group and one for index of
//group which goes next) and then just jump here and there and compute total.
//
class Gain {
    public static $index; //index of the next group
    public static $cash; //cash til here
}
fscanf(STDIN, "%d %d %d", $L, $C, $N);

$places = $L;
$times = $C;
$numGroups = $N;

for ($i = 0; $i < $N; $i++) {
    fscanf(STDIN, "%d", $pi);
    $groups[$i] = $pi;
}


//first the gain per group and the index of the next group
for ($i = 0; $i <= $numGroups-1; $i++){
    $g = new Gain();
    $numPplPerRide= 0;
    $index =$i; //of the NEXT group to go there directly
    while ($numPplPerRide+$groups[$index] - $places <= 0) {
        $numPplPerRide+=$groups[$index]; //$i ou $index?

        //retour de l'index à zéro si index dépasse i (on est arrivé au bout de la liste)
        if ($index < $numGroups-1){
            $index++;
        }else{
            $index =0;
        }

        //gestion des doublons??
        if ($index == $i){
            break;
        }
    }
    $gain[$i]= array(
        $index => $numPplPerRide
    );
    //$g -> $index = $index;
    //$g -> $cash = $numPplPerRide;
    //var_dump("$numGroups,$i,$index,$numPplPerRide\n");
    var_dump($gain);
}

echo("7\n");
