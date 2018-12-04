/* Ecrire un prog qui copie son entree sur sa sortie en remplacant
 * les series de un ou plusieurs espaces par un seul espace
 */

#include<stdio.h>

int main()
{
    char c; //le char courant
    char dernierc; //le char precedent qui sert de marqueur

    while ((c = getchar()) != EOF){
        if (c !=' ' || dernierc == ' ')
            putchar(c);
    }
    dernierc = c; //bien la mettre a jour

    return 0;
}


