#include<stdio.h>

int main(){

char c;
int len=0;

while (c=getchar() != '\n')
  len ++;

printf ("la longueur est %d\n", len);

int i;
for (i=0; i <= len; i++){
//histo horizontal
  printf ("-");
//histo vertical
//printf ("|\n");
}
// serait bien de pouvoir traiter plusieurs mots séparés par un espace.
printf ("\n");
return 0;
}
