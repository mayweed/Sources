import sys
import math

r = int(input())
v = int(input())

#lock combination
combs = []

for i in range(v):
    c, n = [int(j) for j in input().split()]
    time =((10**n)*(5**(c-n))) 
    combs.append(time)
    print(time)

print(combs)