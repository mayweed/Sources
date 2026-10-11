#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

//parse input? je dois avoir h,m,s et utc h et m
void parseInput(char* rawtime){
    char* time = strtok(rawtime,':');
    fprintf(stderr,time);
    return time;
}

//convert from a given timezone to Biel tz
//float convert (int h, in)

int main()
{
    char rawtime[21] = "";
    scanf("%[^\n]", rawtime);

    fprintf(stderr,"%s",rawtime);
    // Write an answer using printf(). DON'T FORGET THE TRAILING \n
    // To debug: fprintf(stderr, "Debug messages...\n");

    printf("answer\n");

    return 0;
}
