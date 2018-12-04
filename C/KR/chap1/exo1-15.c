//exo1-15
// Ecrire une fonction convert avec un arg fahr
#include<stdio.h>

// f like fahrenheit, c like celsius
// Could use float, but be consistent ;)
int convert (int f)
{
        int c;
        c = 5 * (f - 32) / 9;
        return c;
}

int main(void)
{
	int fahr, celsius;
	int lower, upper, step;

	lower = 0;
	upper =300;
	step = 20;
	printf ("Table de conversion fahrenheit/Celsius\n");

	fahr = lower;
	while (fahr <= upper){
        celsius = convert (fahr);
		printf ("%d\t%d\n", fahr, celsius);
		fahr += step;
	}
	return 0;
}
