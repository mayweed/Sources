#include<stdio.h>
#include<stdlib.h>

int main(void){

    FILE *input=fopen("input.txt","r");
    int total=0;
    int change=0;

    if(input){
        //BUG!! https://faq.cprogramming.com/cgi-bin/smartfaq.cgi?id=1043284351&answer=1046476070
        //while (!feof(input)){
        while(fscanf(input,"%i",&change) == 1){
            total+=change;
            //fprintf(stderr,"%d\n",total);
        }
        fclose(input);
    }

    printf("%d\n",total);
    return 0;
}
