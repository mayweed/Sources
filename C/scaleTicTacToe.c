#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

/*
bool checkRow(...);
bool checkColumn(...);
bool checkDiag(...);
bool checkWinner(...);

je checke d’abord s’il y a un winner, 
ensuite je regarde s’il reste des espaces libres, 
si oui ça n’est pas fini 
sinon c’est un draw, c’est bien ça?
*/

typedef enum {
    //IN_PROGRESS,
    NO_WINNER,
    X_WIN,
    O_WIN
    //DRAW
} GameState;

GameState checkCol (int n, int g, char board[n][n+1]){
    //store cell’s index
    int winRow[g]; 
    int winCol[g] ;

    for (int j = 0; j < n; j++){
        int countX = 0;
        int countO = 0;
        for (int i = 0; i <n; i++){
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
            } else {
                countX = 0;
                countO = 0;
            }
            if (countX == g){
                return X_WIN;
            }
            else if (countO == g){
                return O_WIN;
            }

        }
    }
    return NO_WINNER;
}

GameState checkRow (int n, int g,char board[n][n+1]){
    //store cell’s index
    int winRow[g]; 
    int winCol[g] ;

    for (int i = 0; i < n; i++){
        int countX = 0;
        int countO = 0;
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
            } else {  //case vide
                countX = 0;
                countO = 0;
            }
            if (countX == g){
                return X_WIN;
            }
            else if (countO == g){
                return O_WIN;
            }

        }
    }
    return NO_WINNER;
}

/*
↘  dRow = +1  dCol = +1

↙  dRow = +1  dCol = -1
*/
GameState checkDiagTopBottom(int n, int g, char board[n][n+1]){
    //store cell’s index
    int winRow[g]; 
    int winCol[g] ;

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {

            if (board[i][j] == ' ')
                continue;

            char player = board[i][j];
            int countPlayer = 1;
            int k = 1;
            winRow[0] = i;
            winCol[0]=j;

            while (i+k < n &&
                j+k < n &&
                board[i+k][j+k] == player) {
                     winRow[countPlayer] = i+k;
                    winCol[countPlayer]= j+k;
                    countPlayer++;
                    k++;
            }

            if (countPlayer >= g) {
                if (player == 'X')
                    return X_WIN;
                else
                    return O_WIN;
            }
        }
    }
    return NO_WINNER;
}

GameState checkDiagBottomUp(int n, int g, char board[n][n+1]){
    //store cell’s index
    int winRow[g]; 
    int winCol[g] ;

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            
            if (board[i][j] == ' ')
                continue;

            char player = board[i][j];
            int countPlayer = 1;
            int k = 1;
            winRow[0] = i;
            winCol[0]=j;

            while (i+k < n &&
                j-k >= 0 &&
                j-k < n &&
                board[i+k][j-k] == player) {
                    winRow[countPlayer] = i+k;
                    winCol[countPlayer]= j-k;
                    countPlayer++;
                    k++;
            }

            if (countPlayer >= g) {
                if (player == 'X')
                    return X_WIN;
                else
                    return O_WIN;
            }
        }
    }
    return NO_WINNER;
}

int main()
{
    int n;
    int g;
    scanf("%d%d", &n, &g); fgetc(stdin);
    char board[n][n+1];

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
