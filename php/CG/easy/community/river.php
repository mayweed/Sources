<?php
function river($num) {
	$sum=0;
	$n = $num;
	while ($n > 0) {
		$sum += $n % 10;
		$n = $n / 10;
	}
	$num += $sum;
	return $num;
}

fscanf(STDIN, "%d", $r1 );
fscanf(STDIN, "%d", $r2 );

while ($r1 != $r2){
    if ($r1 < $r2){
        $r1=river($r1);
    }else{
        $r2=river($r2);
    }
}
echo("$r1\n");
?>
