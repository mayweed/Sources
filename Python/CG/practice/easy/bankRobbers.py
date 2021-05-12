import sys
import math

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.
#lock combination
r = int(input())
v = int(input())
for i in range(v):
    c, n = [int(j) for j in input().split()]
    print(c,n)
    print((10*n)*(5**(c-n))) #10*5*5 oki??
# Write an answer using print
# To debug: print("Debug messages...", file=sys.stderr, flush=True)

print("1")
