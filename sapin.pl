#!/usr/bin/perl

use warnings;

# le compteur est la clé
# s'il est pas incrementé ou décrémenté
# la boucle affiche tjs la meme chose...
$x=1; 

while ($x <= 5){
   for ($j=0;$j < $x;$j++){
   print "*";
}
print "\n";
$x+=1;
}

# seconde partie
$y=5; 

while ($y >= 1){
   for ($j=0;$j < $y;$j++){
   print "*";
}
print "\n";
$y-=1;
}  
