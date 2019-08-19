#include <stdlib.h>
#include <stdio.h>
#include <string.h>

long long river(long long num) {
	long long sum=0;
	long long n = num;
	while (n > 0) {
		sum += n % 10;
		n = n / 10;
	}
	num += sum;
	return num;
}

int main()
{
    long long r1;
    scanf("%lld", &r1);
    long long r2;
    scanf("%lld", &r2);

    while (r1 != r2){
        if (r1 < r2){
            r1=river(r1);
        }else{
            r2=river(r2);    
        }
    }

    printf("%lld\n",r1);

    return 0;
}
