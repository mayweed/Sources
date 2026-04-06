#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

typedef struct{
    int x;
    int y;
    char contenu;
    int count; // le nb de passage
} Cell;

int main()
{
    int w;
    int h;
    scanf("%d%d", &w, &h);

    Cell board[h][w];

    for (int i = 0; i < h; i++) {
        char line[256] = "";
        scanf("%s", line);
        for (int x = 0; x < w; x++){
            board[i][x].contenu = line[x];
            board[i][x].x = x;
            board[i][x].y = i;
        }
    }
    fprintf(stderr,"%c\n",board[1][0].contenu);
    
 
    char startPos = board[0][0].contenu;

//    if (x >= 0 && x < w && y >= 0 && y < h) {
        switch (startPos){
            case '>': //tant que contenu différent de # si # regarder en haut et en bas en fonction de side
                x = x + 1;
                break;
            case '<':
                x = x - 1;
                break;
            case 'v':
                y = y + 1;
                break;
            case '^': 
                y = y - 1;
                break;
        }
    //}
    char side[2] = "";
    scanf("%s", side);

    fprintf(stderr, "%s",side);

    for (int i = 0; i < h; i++) {

        // Write an action using printf(). DON'T FORGET THE TRAILING \n
        // To debug: fprintf(stderr, "Debug messages...\n");

        printf("#####\n");
    }

    return 0;
}
