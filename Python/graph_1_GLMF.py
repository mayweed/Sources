#!/usr/bin/python3

M={"A":{"B","D"},
   "B":{"A","F","G","H"},
   "C":{"E","F"},
   "D":{"A","E"},
   "E":{"C","D","F","H"},
   "F":{"B","C","E","H"},
   "G":{"B","H"},
   "H":{"B","E","F","G"}} 
afaire={"G"}
fait=set()

while afaire:
    temp=set()
    for s in afaire:
        temp|=M[s]
    temp-=fait
    temp-=afaire
    fait|=afaire
    afaire=set(temp)
    print(temp)
