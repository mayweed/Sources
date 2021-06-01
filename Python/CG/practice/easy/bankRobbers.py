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

# idée: un cpteur qui déroule, une fois un vault fini on le retire de la liste etc…
# une liste en cours
count = 0
#while combs not empty??
# init vaults tested
currVaults = []
for i in range(r):
    currVaults.append(combs.pop(0))

print(len(combs),file=sys.stderr)

while len(currVaults) != 0:
    count+=1
    for i in range(len(currVaults)):
        print(len(currVaults),file=sys.stderr)
        if currVaults[i] == count:
            #print(i,file=sys.stderr)
            #done with that vault…
            currVaults.pop(i)
            #time to grab the next one
            if len(combs)>0:
                currVaults.append(combs.pop(0))
       
#waiting the one with longest combs ends… 
#if len(combs) == r:
#print(max(combs))

        print(r,combs,currVaults,count,file=sys.stderr)