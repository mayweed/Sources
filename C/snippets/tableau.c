// le but est de s'approprier les tabs en c
//

#include <stdio.h>

int main()
{
	int tab[10];
	int j = 0;
	int y = 0;

	//on remplit le tableau
	for ( y = 0; y < 10; y++){
	       tab[y] = j;
		j++; //c une variable temp, bien le mettre dans la boucle!!
	}

	//on imprime le contenu
	for (y = 0; y < 10; y++)
		printf ("valeur de tab i = %d\n", tab[y]);

	return 0;
}	
