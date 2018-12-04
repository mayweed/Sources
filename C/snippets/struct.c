#include <stdio.h>

// cf
// http://gradot.wordpress.com/2012/09/10/initialisation-tableaux-et-structures-en-c/

struct s
{
     int a;
     int b;
     int c;
     int d;
};

struct s sg;    // contenu vaut zero
int tg[5];      // idem

int main(void)
{
   struct s s1 = {12, .c=17};      // pas de warning
   struct s s2 = {12, 15};         // warning : missing initializer
   struct s s3 = {12, 17, 19, 21}; // pas de warning

   int t1[5] = {1,2,3};            // les deux derniers elements valent 0
   int t2[ ] = {1,2,3};            // t2 a une longueur de 3

   struct s s4;                    // contenu indefini
   int t3[5];                      // idem

//    int t4[];
//    // Genere : "error: array size
//    missing in 't4'"

    printf ("mon test fait %d\n", s1.a + s1.b);
}
