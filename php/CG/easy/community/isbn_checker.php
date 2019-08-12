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
$s="";
foreach ($invalids as $elt){
    $s.= "$elt";
    $s.= "\n";
}
$s=trim($s,"\n");
echo("$s"); //HERE pb: Found: 978043551907XEnd of line (\n) | Expected: 978043551907XNothing


function checkIsbn10($isbn){
    //first extract the checkDigit
    $checkDigit = $isbn[strlen($isbn)-1];
    if ($checkDigit == "X") { $checkDigit = 10;}

    //sum the isbn
    $sum = 0;
    $index=2;
    for ($i=strlen($isbn)-2;$i>=0;$i--){
        if ($isbn[$i]=="X"){
            $invalids[]=$isbn;
            break;
        }

        $sum+=$isbn[$i]*$index;
        $index+=1;
    }

    //check it
    if (($sum%11 == 0) && $checkDigit !=0 ){
        return true;
    }

    if (($sum%11 + $checkDigit)%11){
        //there is a remainder!!
        return true;
    }
}

function checkIsbn13($isbn){
    //first extract the checkDigit
    $checkDigit = $isbn[strlen($isbn)-1];
    if ($checkDigit == "X") { $checkDigit = 10;}

    //sum the isbn
    $sum = 0;
    for ($i=strlen($isbn)-2;$i>=0;$i--){
        //check for false char...really not good...but pff...
        if ($isbn[$i]=="X"){
            $invalids[]=$isbn;
            break;
        }
        if ($i % 2){
            $index=3;
        }else{
            $index=1;
        }
        $sum+=$isbn[$i]*$index;
    }

    //check it
    if (($sum%10 == 0) && $checkDigit !=0 ){
        return true;
    }
    if (($sum%10 + $checkDigit)%10 != 0){
        //there is a remainder!!
        return true;
    }
}
?>
