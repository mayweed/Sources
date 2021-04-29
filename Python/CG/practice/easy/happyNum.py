import sys
import math

n = int(input())
for i in range(n):
    x = int(input())
    res = 0
    #combien de tentatives? 50? et on checke le dernier chiffre du dernier rés?
    while x > 0:
        if x%10 > 0:
            y = x%10
            res += (y**2) #does not work if not y but x%10?
        else:
            res += x ** 2
        x= x//10
    print(res,file=sys.stderr)
    #print(str(x) + :)) ou l’inverse…

# if res == 1 it’s okay else go for another pass on res this time
#if res