import sys
import math
import functools

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

n = int(input())
card = 0
for i in range(n):
    card = input()

    c = str(card).replace(' ','')

    x = len(c)
    l1,l2 = [],[]
    #from right to left and -1 to get c[0]
    for i in range(x-2,-1,-2):
        l1.append(int(c[i]))
    nl1=list(map(lambda x : x*2,l1))
    #conv string pour add les nums ensuite il doit y avoir plus concis non?
    sum=0
    for num in nl1:
        for digit in str(num):
            sum+=int(digit)
    #print(sum)

    for i in range(x-1,-1,-2):
        l2.append(int(c[i]))
    nl2= functools.reduce(lambda a,b:a+b,l2)

    check = sum + nl2
    #print(c)
    #print(check)
# Write an answer using print
# To debug: print("Debug messages...", file=sys.stderr, flush=True)
    if check%10:
        print("NO")
    else:
        print("YES")