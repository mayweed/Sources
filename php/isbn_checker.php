<?php
fscanf(STDIN, "%d", $N);

for ($i = 0; $i < $N; $i++)
{
    $ISBN = stream_get_line(STDIN, 20 + 1, "\n");
}
var_dump($ISBN);

function checkIsbn10($isbn){
    $checkDigit = $isbn%10;
    $checkNum = $isbn/10;
}

function checkIsbn13($isbn){
    $checkDigit = $isbn%10;
    //$checkNum = $isbn/10;
    echo("$checkDigit\n");
}


//check length and content
//if length incorrect or content not digit => increment invalid add the isbn to the
//output list
if (strlen($ISBN) == 10 || strlen($ISBN) ==13){// || is_numeric((int)$ISBN)){
    $invalid +=1;
    if (strlen($ISBN)==10){
        checkIsbn10($ISBN);
    }
    if (strlen($ISBN)==13){
        checkIsbn13($ISBN);
    }
}

// Write an action using echo(). DON'T FORGET THE TRAILING \n
// To debug (equivalent to var_dump): error_log(var_export($var, true));
echo("answer\n");
?>
