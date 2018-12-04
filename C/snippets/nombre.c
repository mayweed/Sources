#include <stdio.h>

int main()
{
	char c;

	printf("Entrez un nombre:");
	c = getchar();

	//retenir: sans les '', cela ne marche pas cf p 43
	if (c >= '1' && c <='3')
		printf ("trouve\n");
	else
		printf ("perdu\n");

	return 0;
}
