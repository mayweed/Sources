<?php
//Key is to store two maps / arrays (one of earns per group and one for index of
//group which goes next) and then just jump here and there and compute total.
//
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
    $gain[$i] = $numPplPerRide;
    $nextGp[$i] = $index;
}

$total = 0;
$currentGp =0;
for ($j =0;$j <= $times-1;$j++){
    //gain of the first group to ride
    $total+=$gain[$currentGp];
    //get the index of the next group...
    //ex: $nextGp[0]==1 so go to gain[1] etc...
    //very astute a heap!!!
    $currentGp=$nextGp[$currentGp];
}

echo("$total\n");
