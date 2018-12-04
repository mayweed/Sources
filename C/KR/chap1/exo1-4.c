// Ecrire un table de conversion Celsius -> fahrenheit
// Si celsius = (5/9) * (fahr - 32)
// alors fahr = 9/5 * celsius + 32

#include<stdio.h>

int main()
{
    float fahr, celsius;
    int mini, maxi, intervalle;

    mini = 0; 
    maxi = 300; 
    intervalle = 20;

    celsius = mini;
    printf("Table de conversion Celsius/Fahrenheit\n");

    while (celsius <= maxi){
        fahr = (9.0/5.0) * celsius + 32.0;
        printf("%3.0f %6.1f\n", celsius, fahr);
        celsius = celsius + intervalle;
    }
}

