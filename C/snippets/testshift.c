#include<stdio.h>

//recupere sur DVP, ecrit par Régis Portalez (cf http://bit.ly/R0hPhk)
// First place to avoid previous implicit declaration warning
// Need to SEE the results of the bit op.
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

// comment passer ces 2 args dans la fonction interne?
int testshift (int n, int p){
 	int i; //nb que l'on veut décaler
	printf ("i est égal à %i\t", ~i);
	affichebin(~i);
	putchar ('\n');

	i= ~i << p; // on décale à partir de p.
	printf ("i est égal à %i\t", i);
	affichebin (i);
	putchar ('\n');
}


int main (void){
     testshift (0,4);
     return 0;
}
