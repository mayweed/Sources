use strict;
use warnings;
#use diagnostics;
use 5.20.1;

select(STDOUT); $| = 1; # DO NOT REMOVE

chomp(my $n = <STDIN>);

my $stairs=0;

# Idea is there find a recursive solution for all cases!!
#what is the base case?
#for 2 steps (for 10 it's four)
for my $num (1..$n){
    my $remainder=$n-$num;
    if ($num+$remainder==$n && $num < $remainder){
        $stairs+=1;
    }
}

#for 3 steps (for ten it's 3)
#my $remainder_bis=1;
#for my $num (1..$n){
   # $remainder_bis=$n-($num+1);
    #print STDERR "$remainder_bis\n";
    ## cant work with 5 of course...
    #if (($num +1)+$remainder_bis==$n && ($num+1) < $remainder_bis){
    #    $stairs+=1;
    #}
#}

print "$stairs\n";
