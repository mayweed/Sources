#!/usr/bin/python

"""
GRAPH test etc...
"""
class Node:
    def __init__(self,name):
        self.name=name
    def getName(self):
        return self.name
    def __str__(self):
        return str(self.name)
          
class Edge:    
    def __init__(self,src,dest):
        self.src=src
        self.dest=dest
    def getSrc(self):
        return self.src
    def getDest(self):
        return self.dest
    def __str__(self):
        return self.src + '->' + self.dest

class Graph:
    def __init__(self):
        self.nodes=[]
        #a dict of edges: K:anodes
        self.edges={}
    def node(self):
        return self.nodes
    def addNode(self,node):
        if node not in self.nodes:
            self.nodes.append(node)
            self.edges[node]=[]
    def addEdge(self, edge):
        #Oops!!  This is no digraph!!
        #oki not clean...?
        self.edges[edge.getSrc()].append(edge.getDest())
        self.edges[edge.getDest()].append(edge.getSrc())
    def isConnected(self,node1,node2):
        if node2 in self.edges[node1]: return True
        else: return False
    def neighbours(self,node):
        return self.edges[node]

#Empty graph
g=Graph()

#Add nodes
for node in range(6):
    g.addNode(node)

#Add edges
g.addEdge(Edge(0,1))
g.addEdge(Edge(0,2))
g.addEdge(Edge(1,2))
g.addEdge(Edge(1,3))
g.addEdge(Edge(2,3))
g.addEdge(Edge(2,4))
g.addEdge(Edge(3,4))
g.addEdge(Edge(3,5))

#Look for final result
for node in range(6):
    #pretty print please method here?
    print("{0}:{1}".format(node,g.neighbours(node)))
