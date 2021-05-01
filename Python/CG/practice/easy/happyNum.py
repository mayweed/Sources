import sys
import math

def happy(n):
        return sum(int(x) ** 2 for x in str(n))

n = int(input())
for i in range(n):
    x = int(input())

    # a string map etc…
    #recursion / map + lambda?

    res = happy(x)
    attempt = 0
    while attempt < 50:
        if res == 1:
            break
        else:
            res = happy(res)
        attempt += 1

    if res == 1:
        print(x,":)")
    else:
        print(x,":(")
