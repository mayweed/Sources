#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

typedef struct{
    //int x;
    //int y;
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
        for (int j = 0; j < w; j++){
            board[i][j].contenu = line[j];
           // board[i][j].x = j;
            //board[i][j].y = i;
            board[i][j].count = 0;
        }
    }
    fprintf(stderr,"%c\n",board[1][0].contenu);

    // track my pos
    int x,y = 0;
    Cell myPos;
    
    if (x >= 0 && x < w && y >= 0 && y < h) {
        switch (board[y][x].contenu){
            case '>': //tant que contenu différent de # si # regarder en haut et en bas en fonction de side
                myPos = board[y][x+1];
                myPos.count +=1;
                break;
            case '<':
                myPos = board[y][x-1];
                myPos.count +=1;
                break;
            case 'v':
                myPos = board[y+1][x];
                myPos.count +=1;
                break;
            case '^': 
                myPos = board[y-1][x];
                myPos.count +=1;
                break;
        }
    }
    char side[2] = "";
    scanf("%s", side);

    fprintf(stderr, "%s",side);

    for (int i = 0; i < h; i++) {
        for (int j = 0; j < w; j++){
            if (board[i][j].contenu == '#'){
                printf("%c",board[i][j]);
            } else{
                printf("%i",board[i][j].count);
            }
        }
        printf("\n");
    }

    return 0;
}
