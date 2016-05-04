#!/usr/bin/python

"""
This is codingame lib: you reuse the same basis structure regularly 
I'll put mine here to copy/paste when needed!!
I'll probably polish them to make it more generic...
"""

"""
GRAPH
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
        if node2 in self.edges[node1]:
            return True
        else:
            return False
    def lst_neighbours(self,node):
        return self.edges[node]
