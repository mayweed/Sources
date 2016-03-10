#!/usr/bin/python3

# A dict in dict to represent the graph
M={"A":{"B","D"},
   "B":{"A","F","G","H"},
   "C":{"E","F"},
   "D":{"A","E"},
   "E":{"C","D","F","H"},
   "F":{"B","C","E","H"},
   "G":{"B","H"},
   "H":{"B","E","F","G"}} 

# Curly braces or the set() function can be used to create sets.(PyTut)
afaire={"G"}
fait=set()

while afaire:
    temp=set()
    for s in afaire:
        temp|=M[s]
    print(temp)
    temp-=fait
    temp-=afaire
    fait|=afaire
    afaire=set(temp)
    print(temp)
