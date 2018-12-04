/* ecrire un prog qui compte les e dans une phrase*/
/* Variante sur l'exo 1.8 */

#include <stdio.h>

int main ()
{
	/*on declare*/
	char c;
	int count, nbtabs, blanks;

	/* on initialise*/
	count = nbtabs = blanks = 0;

	/*on deroule...*/
	while ( (c = getchar()) != '\n'){
		if (c == 'e')
			++count;
		else if (c == '\t')
			++nbtabs;
		else if (c == ' ')
			++blanks;
		else
			0;
	}
	printf ( "le nombre de e = %d\n", count);
	printf ("le nb de tabs = %d\n", nbtabs);
	printf ("le nb de blanc = %d\n", blanks);

	return 0;
}

