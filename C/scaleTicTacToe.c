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

bool checkRow (int n, int g,char board[n][n+1]){
    int countX = 0;
    int countO = 0;

    //store cell’s index
    int winRow[g]; 
    int winCol[g] ;

    for (int i = 0; i < n; i++){
        for (int j = 0; j <n; j++){
            if (board[i][j] == 'X') {
                countX ++;
                countO = 0;
                winRow[countX-1] = i;
                winCol[countX-1]=j;
            } else if (board[i][j] == 'O'){
                countX = 0;
                countO +=1;
                winRow[countO-1] = i;
                winCol[countO-1]=j;
            }
            if (countX == g){
                printf("We have a winner X");
            }
            else if (countO == g){
                printf("O got it");
            }

        }
    }
}

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
