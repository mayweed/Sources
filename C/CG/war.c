#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define CARD_MAX 57

//Some values/cant start with a digit
//une struct? avec deux mbres: le char* et un entier (pour ><?)
const char* value= '2345678910JQKA';

// Let's play with struct
// Should retrieve those values correctly from input...
struct card{
    char* value;
    char* suit;
}card_p1,card_p2;

//that should be a struct of arrays
struct hand{
    //struct card
    char* card //initialize that with cardp??
    int number;
};

int main()
{
    char cardp1[4];
    char cardp2[4];
    /*
    make that a struct array:
    handp[1] should be (a pointer to?) one card (value+suit)
    so you create struct hand{struct card} handp!,handp2;
    then initialize handp1[0] to the first card (OKI HOW in the hell??)
    */
    char handp1[n];
    
    int n; // the number of cards for player 1
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        char cardp1[4]; // the n cards of player 1
        scanf("%s", cardp1);
   //     handp1[i]=cardp1;
    }
    
    int m; // the number of cards for player 2
    scanf("%d", &m);
    for (int i = 0; i < m; i++) {
        char cardp2[4]; // the m cards of player 2
        scanf("%s", cardp2);
    }
    
    // Write an action using printf(). DON'T FORGET THE TRAILING \n
    //fprintf(stderr,"%s %s",cardp1,cardp2);

    //printf("PAT\n");

    return 0;
}
