/*Compter les mots d'une entree en C*/

#include <stdio.h>

int main()
{
	char c;
	int i, nw, nt;

	nw = nt = 0;

	while ((c = getchar()) != '\n'){
		if ( c == '\t')
			++nt;
		else if ( c == ' ')
			++nw;
		}
	/*printf("le nombre de mots est %d et de tabs %d\n", nw, nt);*/
	for (i=0; i < nw; i++)
		putchar ('|') && putchar ('\n');
	return 0;
}
	
