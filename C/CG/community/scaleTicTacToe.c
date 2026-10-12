#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

typedef enum {
    NO_WINNER,
    IN_PROGRESS,
    X_WIN,
    O_WIN,
    DRAW
} GameState;

typedef struct {
    GameState state;
    int row[13];
    int col[13];
    char symbol;
    const char* message;
} WinInfo;

WinInfo checkCol (int n, int g, char board[n][n+1]){
    //à mettre dans une func
    WinInfo info;
    info.state = NO_WINNER;
    info.symbol = ' ';

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
                info.state = X_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
               info.symbol = '|';
               info.message = "The winner is X.\n";
               return info;
            }
            else if (countO == g){
                 info.state = O_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
               info.symbol = '|';
               info.message = "The winner is O.\n";
               return info;
            }

        }
    }
    return info;
}

WinInfo checkRow (int n, int g,char board[n][n+1]){
    //à mettre dans une func
    WinInfo info;
    info.state = NO_WINNER;
    info.symbol = ' ';

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
                info.state = X_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
                info.message = "The winner is X.\n";
                info.symbol = '-';
                return info;
            }
            else if (countO == g){
                 info.state = O_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
                info.message = "The winner is O.\n";
                info.symbol = '-';
                return info;
            }

        }
    }
    return info;
}

/*
↘  dRow = +1  dCol = +1

↙  dRow = +1  dCol = -1
*/
WinInfo checkDiagTopBottom(int n, int g, char board[n][n+1]){
    //à mettre dans une func
    WinInfo info;
    info.state = NO_WINNER;
    info.symbol = ' ';
    
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
                if (player == 'X'){
                    info.state = X_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
                info.message = "The winner is X.\n";
                    info.symbol = '\\';
                    return info;
                } else if (player == 'O'){
                    info.state = O_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
                info.message = "The winner is O.\n";
                    info.symbol = '\\';
                    return info;
                }
            }
        }
    }
    return info;
}

WinInfo checkDiagBottomUp(int n, int g, char board[n][n+1]){
    //à mettre dans une func
    WinInfo info;
    info.state = NO_WINNER;
    info.symbol = ' ';
    
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
                if (player == 'X'){
                    info.state = X_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
                info.message = "The winner is X.\n";
                    info.symbol = '/';
                    return info;
                } else if (player == 'O'){
                    info.state = O_WIN;
                for (int k = 0; k < g; k++){
                    info.row[k] = winRow[k];
                    info.col[k] = winCol[k];
                }
                info.message = "The winner is O.\n";
                    info.symbol = '/';
                    return info;
                }
            }
        }
    }
    return info;
}

WinInfo checkWinner(int n, int g, char board[n][n+1])
{
   //à mettre dans une func
    WinInfo info;
    info.state = NO_WINNER;
    info.symbol = ' ';

    WinInfo c = checkCol(n, g, board);
    if (c.state != NO_WINNER){
        //modify board
        //for (int i = 0; i < g; i++) {
        //board[c.row[i]][c.col[i]] = c.symbol;
        //}
    return c;
    }

    c = checkRow(n, g, board);
    if (c.state != NO_WINNER)
        return c;

    c = checkDiagTopBottom(n, g, board);
    if (c.state != NO_WINNER)
        return c;

    c = checkDiagBottomUp(n, g, board);
    if (c.state != NO_WINNER)
        return c;

    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            if (board[i][j] == ' '){
                info.state = IN_PROGRESS;
                info.message = "The game isn't over yet!\n";
                return info;
            }
        }
    }
    info.message = "The game ended in a draw!\n";
    info.state = DRAW;
    return info;
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
    WinInfo winner = checkWinner(n,g,board);

    // on change le contenu des cases
    if (winner.state == X_WIN || winner.state == O_WIN){
        for (int i = 0; i < g; i++){
            board[winner.row[i]][winner.col[i]] = winner.symbol;
        }

    }

    // on affiche la grille
    for (int i = 0; i < n; i++){
        for (int j = 0; j < n; j++){
            printf("%c", board[i][j]);
        }
        printf("\n");
    }

    // Write an answer using printf(). DON'T FORGET THE TRAILING \n
    //fprintf(stderr, "%c\n",board[1][0]);
    printf("%s",winner.message);

    return 0;
}
