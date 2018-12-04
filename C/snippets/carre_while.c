/*fun*/
//il faut une boucle externe qui passe x fois la boucle interne 
//TODO: ca me met une ligne et 4 \n, pourquoi ne passe t on qu'une fois dans
//second while?

#include <stdio.h>

int main()
{
	int line, y;

	line = y = 0;

	//la boucle externe assure que l'on passe x fois dans la boucle interne
	while (line < 5){
		//la boucle interne est executee completement a chaque passe
	     while (y<20){ //cinq ligne de 4 etoiles
		putchar ('*');
		y +=1;
		//ca marche, ca fait travailler sur break mais c pas elegant!!
		if (y == 4 || y == 8 || y == 12 || y == 16)
		    break;
	     }
	putchar ('\n');
	line +=1;
	}

	return 0;

}


