#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <time.h>
#include <stdbool.h>

typedef struct {
    int64_t x;
    int64_t y;
} xy_t;

typedef enum {
    CMD_MOVE,
    CMD_BOMB
} ctype;

typedef struct {
    ctype type;
    xy_t xy;
} cmd_t;

typedef struct {
    int owner;
    int count;
    int range;
}bomb_t;

typedef struct{
    int nature; // seul param 1 est vraiment utile non?
}obj_t;

typedef struct{
    int id;
    xy_t pos;
    int nbBomb;
    int range;
}player_t;

void print_action(cmd_t *cmd) {
    if (cmd->type == CMD_MOVE) {
        printf("MOVE %lld %lld\n", cmd->xy.x, cmd->xy.y);
    } else if (cmd->type == CMD_BOMB) {
        printf("BOMB %lld %lld\n", cmd->xy.x, cmd->xy.y);
    }
}

char get_board(char *board, int w, int x, int y) {
    return board[y * w + x];
}

bool check_board(char *board, int w, int x, int y, char c) {
    return get_board(board, w, x, y) == c;
}

bool check_board_range(char *board, int w, int h, int x, int y, char c) {
    if (x < 0 || x >= w || y < 0 || y >= h) {
        return false;
    }
    return check_board(board, w, x, y, c);
}

bool bomb_in_range(int x1, int y1, int x2, int y2) {
    return abs(x1 - x2) <= 1 && abs(y1 - y2) <= 1;
}

bool danger_here(char *board, int w, int h, int x, int y, int bomb_x, int bomb_y, int range) {
    if (x == bomb_x && abs(y - bomb_y) <= range) return true;
    if (y == bomb_y && abs(x - bomb_x) <= range) return true;
    return false;
}

int count_boxes_hit(char *board, int w, int h, int x, int y, int range) {
    int hit = 0;

    // Directions: up, down, left, right
    int dx[4] = {0, 0, -1, 1};
    int dy[4] = {-1, 1, 0, 0};

    for (int d = 0; d < 4; d++) {
        for (int step = 1; step <= range; step++) {
            int nx = x + dx[d] * step;
            int ny = y + dy[d] * step;

            // Hors du plateau
            if (nx < 0 || nx >= w || ny < 0 || ny >= h)
                break;

            char cell = get_board(board, w, nx, ny);

            // Mur indestructible
            //if (cell == '#')
            //    break;

            // Caisse destructible
            if (cell == '0') {
                hit++;
                break;  // On s'arrête dans cette direction
            }
        }
    }
    return hit;
}

xy_t think (char *board, int w, int h){
    // choose random target
    xy_t moves [4];
    for (int i =0; i<4; i++){
        moves[i].x=rand() % w;
        moves[i].y=rand() % h;
        fprintf(stderr, "moves proposés : %d %d\n",moves[i].x, moves[i].y);
    }

    int best = 0;
    xy_t myMove = moves[0];

    for (int a = 0; a < 3; a++){
        int t = count_boxes_hit(board,w,h, moves[a].x,moves[a].y,3);
        if (t > best){
            best = t;
            myMove= moves[a];
            fprintf(stderr,"test nove %d %d\n",moves[a].x, moves[a].y);
        }
    }
    fprintf(stderr, "best : %d %d %d\n", best, myMove.x, myMove.y);
    return myMove;
}

int main() {
    int w, h, my_id;
    scanf("%d%d%d", &w, &h, &my_id);

    srand(time(NULL));


    bool hasBomb = false;
    int turn = 0;

    char board[h * w];

    bomb_t myBomb;
    player_t me;
    obj_t objet;

    xy_t gameMove = { -1, -1 };
    bool hasObjective = false;

    // game loop
    while (1) {
        //   read board
        for (int i = 0; i < h; i++) {
            char row[w + 1];
            scanf("%s", row);
            for (int n = 0; n < w; n++) {
                board[i * w + n] = row[n];
            }
        }

        int entities;
        scanf("%d", &entities);

        for (int i = 0; i < entities; i++) {
            int entity_type, owner, x, y, param_1, param_2;
            scanf("%d%d%d%d%d%d",&entity_type, &owner, &x, &y, &param_1, &param_2);

            if (owner == my_id){
                switch (entity_type) {
                case 2: // objet
                    objet = (obj_t){param_1};
                    break;

                case 1: // Bombe
                    myBomb = (bomb_t){ owner, param_1, param_2 };
                    break;

                case 0: // Player
                    me = (player_t){ owner, {x, y}, param_1, param_2 };
                    break;

                default:
                    break;
                }
            }
        }
        fprintf(stderr,"me : %d %d %d\n", me.nbBomb, me.pos.x, me.pos.y);

        if (!hasObjective){
            gameMove = think(board,w,h);
            fprintf(stderr,"test : %d %d\n",gameMove.x,gameMove.y);
            hasObjective = true;
        }
        if (me.pos.x == gameMove.x && me.pos.y == gameMove.y){
            hasObjective = false;
        }

       fprintf(stderr,"move : %d %d %b\n",gameMove.x, gameMove.y,hasObjective);


        int hits = count_boxes_hit(board, w, h, me.pos.x, me.pos.y, me.range);

            if (me.nbBomb > 0 && hits > 0) {
                printf("BOMB %d %d\n", me.pos.x, me.pos.y);
                continue;
            }

        // 4. Déplacement vers l’objectif
        printf("MOVE %d %d\n", gameMove.x, gameMove.y);

    turn++;
    }
    return 0;
}

/*
tour n :
    analyser toutes les bombes autour
    SI en danger → move vers position safe

    sinon :
        calculer combien de caisses je peux casser depuis ma position
        SI >= 1 et j’ai une sortie →
            poser bombe
        sinon →
            chercher caisse la plus proche (BFS)
            move vers elle


*/
