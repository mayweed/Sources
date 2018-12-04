// Verify that getchar() != EOF is 0 or 1
// A retenir: le shortcut sous Linux pour EOF, c'est Ctrl+d
#include<stdio.h>

int main()
{
	int c;

	c = (getchar() != EOF);

	printf ("%d\n", c);

	return 0;
}
