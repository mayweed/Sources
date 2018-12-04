#include<stdio.h>

//recupere sur DVP, ecrit par Régis Portalez (cf http://bit.ly/R0hPhk)
// First place to avoid previous implicit declaration warning
int affichebin(unsigned n)
{
	unsigned bit = 0 ;
	unsigned mask = 1 ;
	int i;
	for (i = 0 ; i < 16 ; ++i)
	{
		bit = (n & mask) >> i ;
		printf("%d", bit) ;
		mask <<= 1 ;
	}
}

int testand (int n){
      int i;
      for (i = 0; i < 8; i++){
	//formatting properly cf p.13 K&R
	// le compilo est pas très content mais ça marche mon espace!
	printf ("n est égal à %4.i\ ", n >> i);
        affichebin (n >> i);
        putchar ('\n');
        }
      return n;
}


int main (void){
        //guess value just test; first with 2 & n <<, second with 256 et n >>
        // Goal is to observe the diff in binary
	testand (256);
        return 0;


}

