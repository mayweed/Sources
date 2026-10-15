#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

typedef struct{
    char contenu;
    int count; // le nb de passage
} Cell;

// d for the directions (0 v,1 >,2 ^,3 < )
int left(int d)  { return (d + 3) % 4; }
int right(int d) { return (d + 1) % 4; }
int back(int d)  { return (d + 2) % 4; }

int main()
{
    int w;
    int h;
    scanf("%d%d", &w, &h);

    Cell board[h][w];
    int startX = 0;
    int startY = 0;

    for (int i = 0; i < h; i++) {
        char line[256];
        scanf("%s", line);
        for (int j = 0; j < w; j++){
            board[i][j].contenu = line[j];
            board[i][j].count = 0;
            if (strchr("^v<>", board[i][j].contenu)){
                startX = j;
                startY = i;
            }
        }
    }
    char side;
    scanf(" %c", &side);
    fprintf(stderr,"%c",side);

    //directions | char dirs[4] = {'v', '>', '^', '<'};
    int dx[4] = {0, 1, 0, -1};  // down, right, up, left
    int dy[4] = {1, 0, -1, 0};

    // track my pos
    int x = startX;
    int y = startY;
    char startDir = board[startY][startX].contenu;
    int dir=0;

    switch(startDir){
        case 'v': dir = 0; break;
        case '>': dir = 1; break;
        case '^': dir = 2; break;
        case '<': dir = 3; break;
    }

    bool firstStep = true;

    while (true) {

        
        board[y][x].count++;
        int order[4];
    if (side == 'L') {
            order[0] = left(dir);
            order[1] = dir;
            order[2] = right(dir);
            order[3] = back(dir);
        } else {
            order[0] = right(dir);
            order[1] = dir;
            order[2] = left(dir);
            order[3] = back(dir);
        }
        // try directions in order
        bool moved = false;
        for (int i = 0; i < 4; i++) {
            int nd = order[i];
            int nx = x + dx[nd];
            int ny = y + dy[nd];
            //fprintf(stderr,"try nd=%d -> (%d,%d)",nd, nx, ny);     
   
            if (nx >= 0 && nx < w && ny >= 0 && ny < h &&
                board[ny][nx].contenu != '#') {
                x = nx;
                y = ny;
                dir = nd;
                moved = true; 
                break;
            }
        }
        if (!moved) {
                board[y][x].count--;
                break;
                }

        // stop condition (loop detection or out of bounds)
       if (x == startX && y == startY && !firstStep){ 
           break;
        } 
        firstStep = false;

    }
    
    for (int i = 0; i < h; i++) {
        for (int j = 0; j < w; j++){
            if (board[i][j].contenu == '#'){
                printf("%c",board[i][j].contenu);
            } else{
                printf("%i",board[i][j].count);
            }
        }
        printf("\n");
        
    }

    return 0;
}
