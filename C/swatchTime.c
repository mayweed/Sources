#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

typedef struct {
    int h;
    int m;
    int s;
} Time;

typedef struct {
    int h;
    int m;
    char sign;
} UtcOffset;

//convert from a given timezone to Biel tz
Time convert (Time t, UtcOffset u){ 

    Time newTime;

    // on convertit tout en sec
    int total =
        t.h * 3600 +
        t.m * 60 +
        t.s;

    int delta =
        u.h*3600 +
        u.m *60;
    if (u.sign == '-')
        delta = -delta;

    int biel = total - delta + 3600;
    biel = (biel % 86400 + 86400) % 86400;

    newTime.h = biel/3600;
    biel %= 3600;
   
    newTime.m = biel/60;
    biel %= 60;

    newTime.s = biel;

    return newTime;

}


int main()
{
    Time t;
    UtcOffset u;

    scanf("%d:%d:%d UTC%c%d:%d",
        &t.h, &t.m, &t.s,
        &u.sign,
        &u.h, &u.m);
    
    fprintf(stderr,
        "h=%d m=%d s=%d sign=%c utcH=%d utcM=%d\n",
        t.h, t.m, t.s,
        u.sign,
        u.h, u.m);

    // Write an answer using printf(). DON'T FORGET THE TRAILING \n
    printf("answer\n");

    return 0;
}
