// pour comprendre pourquoi l'on peut faire ip = array
//
#include<stdio.h>

int main ()
{
int array[5], i, *ip;

	for(i = 0; i < 5; i++)
       		array[i] = i;

	ip = &array[0]; // ou p = array; cf K&R p.99
	// "Since the name of an array is a synonym for the location of the initial
	// element"...
	printf("%d\n", *(ip + 3));
	return 0;
}
