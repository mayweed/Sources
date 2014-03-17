#!/usr/bin/perl

use warnings;

# le compteur est la clé
# s'il est pas incrementé ou décrémenté
# la boucle affiche tjs la meme chose...
$x=5; 

while ($x >= 1){
   for ($j=0;$j < $x;$j++){
   print "*";
}
print "\n";
$x-=1;
}  
