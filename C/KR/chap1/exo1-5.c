//exo1-5
//
#include<stdio.h>

int main()
{
	int fahr, celsius;
	int lower, upper, step;

	lower = 0;
	upper =300;
	step = 20;
	printf ("Table de conversion fahrenheit/Celsius\n");

	fahr = upper;
	while (fahr >= lower){
		celsius = 5 * (fahr-32) / 9;
		printf ("%d\t%d\n", fahr, celsius);
		fahr -= step;
	}
	return 0;
}
