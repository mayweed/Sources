#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

/*
bool checkRow(...);
bool checkColumn(...);
bool checkDiag(...);
bool checkWinner(...);
*/

int main()
{
    int n;
    int g;
    scanf("%d%d", &n, &g); fgetc(stdin);
    char board[n][n];

    for (int i = 0; i < n; i++) {
        char row[n + 1];
        scanf("%[^\n]", row); fgetc(stdin);
        for (int j = 0; j <n; j++){
            board[i][j]=row[j];
        }
    }

    // Write an answer using printf(). DON'T FORGET THE TRAILING \n
    //fprintf(stderr, "%c\n",board[1][0]);

    printf("Answer\n");

    return 0;
}
