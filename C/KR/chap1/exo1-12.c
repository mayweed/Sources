#include<stdio.h>
/* We postulate that words get separated either by a space or by a tab*/
int main (void){
    int c;

    while ((c=getchar()) != EOF){
            putchar (c);

        if (c == ' ' || c=='\t'){
            printf ("\n");
        }
    }
    return 0;
    }
