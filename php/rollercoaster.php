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

$cash = 0;

foreach ($groups as $index => $group) {
    if ($group - $places < 0) {
        $gain[$index]=$cash;
        //next group = index+1;?
        $places=$L;
    } else {
        $cash+=$group;
        $places -= $group;
    }
}

echo("$cash\n");
