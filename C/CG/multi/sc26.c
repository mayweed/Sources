#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

typedef struct {
    int id;
    int player;
    int x;
    int y;
    int movement_speed;
    int carry_capacity;
    int harvest_power;
    int chop_power;
    int carry_plum;
    int carry_lemon;
    int carry_apple;
    int carry_banana;
    int carry_iron;
    int carry_wood;
}Troll;

enum fruitType {
    PLUM,
    LEMON,
    APPLE,
    BANANA,
};

typedef struct{
    int x;
    int y;
    int size;
    int health;
    int fruits;
    int cooldown;
}Ftree;

typedef enum {
    MOVE,
    HARVEST,
    DROP,
    WAIT,
    MSG,
}actionType;

typedef struct{
    actionType type;
    int trollId;
}Action;

int main()
{
    int width;
    int height;
    scanf("%d%d", &width, &height); fgetc(stdin);
    for (int i = 0; i < height; i++) {
        char line[width + 2]; // = "";
        //scanf("%[^\n]", line); fgetc(stdin);
        fgets(line, sizeof line, stdin);
    }

    // game loop
    while (1) {
        for (int i = 0; i < 2; i++) {
            int plum;
            int lemon;
            int apple;
            int banana;
            int iron;
            int wood;
            scanf("%d%d%d%d%d%d", &plum, &lemon, &apple, &banana, &iron, &wood);
        }
        int trees_count;
        scanf("%d", &trees_count);
        for (int i = 0; i < trees_count; i++) {
            char type[7] = "";
            int x;
            int y;
            int size;
            int health;
            int fruits;
            int cooldown;
            scanf("%s%d%d%d%d%d%d", type, &x, &y, &size, &health, &fruits, &cooldown);
        }
        int trolls_count;
        scanf("%d", &trolls_count);
        for (int i = 0; i < trolls_count; i++) {
            int id;
            int player;
            int x;
            int y;
            int movement_speed;
            int carry_capacity;
            int harvest_power;
            int chop_power;
            int carry_plum;
            int carry_lemon;
            int carry_apple;
            int carry_banana;
            int carry_iron;
            int carry_wood;
            scanf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d", &id, &player, &x, &y, &movement_speed, &carry_capacity, &harvest_power, &chop_power, &carry_plum, &carry_lemon, &carry_apple, &carry_banana, &carry_iron, &carry_wood);
        }

        // Write an action using printf(). DON'T FORGET THE TRAILING \n
        // To debug: fprintf(stderr, "Debug messages...\n");


        // valid actions:
        // MOVE <id> <x> <y>
        // HARVEST <id> - when you are on the same cell as a tree
        // DROP <id> - when you are next to your shack and carry items
        printf("MOVE 0 7 7\n");
    }

    return 0;
}
