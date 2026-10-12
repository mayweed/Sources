#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

int main()
{
    char bank_cards[101] = "";
    //scanf("%[^\n]", bank_cards); fgetc(stdin);
    fgets(bank_cards, sizeof(bank_cards),stdin);
    char player_cards[101] = "";
    //scanf("%[^\n]", player_cards);
    fgets(player_cards, sizeof(player_cards),stdin);

    // Write an answer using printf(). DON'T FORGET THE TRAILING \n
    // To debug: fprintf(stderr, "Debug messages...\n");

    fprintf(stderr,"%s",bank_cards);
    printf("Blackjack!\n");

    return 0;
}
