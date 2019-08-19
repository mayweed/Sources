#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define MAX_LENGTH 21

typedef struct node{
    int value;
    struct node *next;
}node;

int main()
{
    int N;
    scanf("%d", &N);
    
    char telephone[MAX_LENGTH];
    char *nums[N];
    for (int i = 0; i < N; i++) 
    {
        scanf("%s", telephone);
		nums[N]=&telephone;
    }

    for (int i=0;i<MAX_LENGTH-1;i++)
    {
        node *num = malloc(sizeof(int));   
        num->value=atoi(&telephone[i]);
    }

    //fprintf(stderr, "Debug messages...\n");

    // The number of elements (referencing a number) stored in the structure.
    printf("number\n");

    return 0;
}
