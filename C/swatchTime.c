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
/*

*/
int main()
{
    //char rawtime[21] = "";
    //fgets(rawtime, sizeof(rawtime), stdin);

    int h, m, s;
    int utcH, utcM;
    char sign;

    //avec rawtime +fgets() possible d’utiliser sscanf(rawtime,...)
    scanf("%d:%d:%d UTC%c%d:%d",
        &h, &m, &s,
        &sign,
        &utcH, &utcM);
        
    //fprintf(stderr,"[%s]\n",rawtime);
    fprintf(stderr,
        "h=%d m=%d s=%d sign=%c utcH=%d utcM=%d\n",
        h, m, s,
        sign,
        utcH, utcM);

    // Write an answer using printf(). DON'T FORGET THE TRAILING \n
    // To debug: fprintf(stderr, "Debug messages...\n");

    printf("answer\n");

    return 0;
}
