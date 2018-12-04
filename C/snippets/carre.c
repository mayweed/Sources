
#include <stdio.h>

int main()
{
	//declarations
	int line, i;

	//initialisation
	i = 0;

	//la boucle externe assure que l'on passe x fois dans la boucle interne
	while (i<5){
		//la boucle interne est executee completement a chaque passe
	     for (line= 0; line < 4; line++) //une ligne de 4 etoiles
		putchar ('*');

	    putchar ('\n');
	
	i +=1;
	}	
	return 0;

}


