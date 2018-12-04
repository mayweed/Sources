#include<stdio.h>
int main()
/*pour messurer la difference entre prefixer ++ et le postfixer*/
{
    int A = 1;
    int B = 2;
    int C = 3;
    int *P1, *P2;
    P1=&A;
    P2=&C;
    //*P1=++(*P2);
    *P1=(*P2)++;
    
// pourquoi A=3 et C=4 ?
    printf ("La valeur de A est %x\n", &A); 
    printf ("La valeur de C est %d\n", C);
//test adresse
//L'adresse mémoire du pointeur P1
    printf ("La valeur de P1 est %p\n", &P1);
//L'adresse contenu par le pointeur P1 qui équivaut à l'adresse de A
//bien entendu.
    printf ("La valeur de P1 est %p\n",&(*P1));
}
//Réponse à la question:
//Tout simplement parce que ++ fonctionne de cette façon.
//
//Quand vous faites :*P1=(*P2)++;
//Pour prenez la valeur de P1 et obtenez l'adresse de A. Idem pour P2 et
//l'adresse de C. A ce moment là, la valeur de C est mise dans A
//(opérateur de de-référencement * oblige). Cela revient donc à A=C=3.
//
//L'opérateur ++ ne vient qu'après tout cela et incrémente après
//l'affection de C dans A, donc *P2 = C = 4.
//Pour obtenir A = C = 4, vous pouvez faire *P1=(*P2)++ + 1;.
//
//Si je regarde la norme C99 (document n1256), je lis au paragraphe
//6.5.2.4, point 2 :
//The result of the postfix ++ operator is the value of the operand.
//After the result is obtained, the value of the operand is incremented.
//
