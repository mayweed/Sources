#!/usr/bin/python3

# Set operators:
# &: letters in both operands
# ^: letters either op but NOT in both
# -: letters in the left op but NOT in the right op
# |: letters in either op or both

# A dict in dict to represent the graph
M={"A":{"B","D"},
   "B":{"A","F","G","H"},
   "C":{"E","F"},
   "D":{"A","E"},
   "E":{"C","D","F","H"},
   "F":{"B","C","E","H"},
   "G":{"B","H"},
   "H":{"B","E","F","G"}} 

print(M["A"])
# Curly braces or the set() function can be used to create sets.
#(PyTut 5.4)
afaire={"G"}
# Note: to create an empty set you have to use set(), not {}
# (PyTut 5.4)
fait=set()

while afaire:
    temp=set()
    for s in afaire:
        temp|=M[s]
    temp-=fait
    # - in set: letters in temp but NOT in afaire
    temp-=afaire
    # | in set: letters in either fait or afaire
    fait|=afaire
    afaire=set(temp)
    print(temp)
