<?php
fscanf(STDIN, "%d", $N);

$invalid=0;

for ($i = 0; $i < $N; $i++)
{
    $ISBN = stream_get_line(STDIN, 20 + 1, "\n");
    if (strlen($ISBN) == 10 || strlen($ISBN) ==13){
        if (strlen($ISBN)==10){
            if (checkIsbn10($ISBN)){
                $invalid+=1;
                $invalids[]=$ISBN;
            }
        }
        if (strlen($ISBN)==13){
            if(checkIsbn13($ISBN)){
                $invalid+=1;
                $invalids[]=$ISBN;
            }
        }
    }else{
        $invalid+=1;
        $invalids[]=$ISBN;
    }

}
echo("$invalid invalid:\n");
foreach ($invalids as $elt){
    echo("$elt\n");
}

function checkIsbn10($isbn){
    //first extract the checkDigit
    $checkDigit = $isbn[strlen($isbn)-1];
    if ($checkDigit == "X") { $checkDigit = 10;}

    //sum the isbn
    $sum = 0;
    $index=2;
    for ($i=strlen($isbn)-2;$i>=0;$i--){
        $sum+=$isbn[$i]*$index;
        $index+=1;
    }

    //check it
    if (($sum%11 + $checkDigit)%11){
        //there is a remainder!!
        //echo("INVALID $isbn\n");
        return true;
    }//else{
        //no remainder
        //echo("VALID $isbn\n");
      //  return false;
    //}

}

function checkIsbn13($isbn){
    //first extract the checkDigit
    $checkDigit = $isbn[strlen($isbn)-1];
    if ($checkDigit == "X") { $checkDigit = 10;}

    //sum the isbn
    $sum = 0;
    for ($i=strlen($isbn)-2;$i>=0;$i--){
        if ($i % 2){
            $index=1;
        }else{
            $index=3;
        }
        $sum+=$isbn[$i]*$index;
    }

    //check it
    if (($sum%10 + $checkDigit)%10){
        //there is a remainder!!
        //echo("INVALID $isbn\n");
        return true;
    }//else{
        //no remainder
        //echo("VALID $isbn\n");
    //}
}

function printOutput(){
}
?>
