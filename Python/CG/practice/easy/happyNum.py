import sys
import math

n = int(input())
for i in range(n):
    x = int(input())
    res = 0
    attempt = 0
    #combie0n de tentatives? 50? et on checke le dernier chiffre du dernier rés?
    # a string map etc…
    #recursion / map + lambda?
    while attempt < 50:
        print(attempt,x)
        while x > 0:
            if x%10 > 0:
                y = x%10
                res += (y**2) #does xot work if not y but x%10?
            else:
                res += x ** 2
            x= x//10
        attempt+=1

    #print(res,file=sys.stderr)
    #print(str(x) + :)) ou l’inverse…
