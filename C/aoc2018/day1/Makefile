CC=clang
CFLAGS=-ggdb3 -Ispl/include -O0 -std=c99 -Wall -Werror -Wno-unused-variable

part1: part1.c 
	$(CC) -ggdb3 -Ispl/include -O0 -std=c99 -Wall -Werror -Wno-unused-variable -o part1 part1.c -Lspl/lib -lm

part2: part2.c 
	$(CC) -ggdb3 -Ispl/include -O0 -std=c99 -Wall -Werror -Wno-unused-variable -o part2 part2.c -Lspl/lib -lm

clean:
	rm -f *.o 
