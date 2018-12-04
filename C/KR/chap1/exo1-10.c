/* Exo 1.10*/
/* Un prog qui copie son input sur son output en remplacant
 * un backspace par \b, une tab par \t etc..*/

#include <stdio.h>

#define IN 1 
#define OUT 0

int main ()
{
    int etat;
	char c;

    etat = OUT;

while ((c = getchar()) != EOF){
    if (c == '\t'){
       putchar ('\\'); 
       putchar ('t');
       etat  = IN;
    }
    else if (c == '\\'){
        putchar ('\\');
        putchar ('\\');
        etat = IN;
    }
    else if (c == ' '){ //does not work with a real \b...
        putchar ('\\');
        putchar ('b');
        etat = IN;
    }
    else 
        etat = OUT;

    if ( etat == OUT)
        putchar (c);
}
    return 0;
}
