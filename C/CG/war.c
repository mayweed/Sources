#include <stdio.h>
#include <stdlib.h>

/* A card token is at most three characters (for example, 10H) plus '\0'. */
struct card {
    char text[4];
};

struct hand {
    struct card *cards;
    size_t count;
};

static int read_hand(struct hand *hand)
{
    int count;

    if (scanf("%d", &count) != 1 || count < 0) {
        return 0;
    }

    hand->count = (size_t)count;
    hand->cards = calloc(hand->count, sizeof *hand->cards);
    if (hand->count != 0 && hand->cards == NULL) {
        return 0;
    }

    for (size_t i = 0; i < hand->count; ++i) {
        if (scanf("%3s", hand->cards[i].text) != 1) {
            free(hand->cards);
            hand->cards = NULL;
            hand->count = 0;
            return 0;
        }
    }

    return 1;
}

int main(void)
{
    struct hand player1 = { NULL, 0 };
    struct hand player2 = { NULL, 0 };

    if (!read_hand(&player1) || !read_hand(&player2)) {
        fprintf(stderr, "Invalid or incomplete hand input\n");
        free(player1.cards);
        free(player2.cards);
        return EXIT_FAILURE;
    }

    /* The hands are now available as player1.cards and player2.cards. */

    free(player1.cards);
    free(player2.cards);
    return EXIT_SUCCESS;
}
