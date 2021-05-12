import sys
import math

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

n = int(input())
card = 0
for i in range(n):
    card = input()

c = str(card).replace(' ','')


x = len(c)
l1,l2 = [],[]
for i in range(x):
    if i%2 == 0:
        l2.append(c[i])
        print("l2",i,card[i])
    else:
        l1.append(c[i])
        print("l1",i,card[i])

print(c)
print(l1,l2)
# Write an answer using print
# To debug: print("Debug messages...", file=sys.stderr, flush=True)

print("YES or NO")
