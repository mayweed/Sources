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
    def __str__(self):
        for node in range(len(self.nodes)):
            print("{0}:{1}".format(node,g.neighbours(node)))


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

#print(g)

def printPath(path):
    """
    print path, path is a list of nodes
    """
    result=''
    for node in path:
        result+=str(node)
        if node != path[-1]:
            #append only if not the last...
            result+='->'
    return result

# TEST
#print(printPath([1,2,3]))

def bfs(graph,start,end):
    """
    write a bfs algo
    graph is a Graph object
    start, end are node's name
    """
    #structure used to put nodes to visit
    queue=[]
    #my path
    path=[]
    #init the structure
    path.append(start)
    queue.append(path)

    #you loop on the queue where you put your nodes
    while len(queue) != 0:
        tmpPath=queue.pop(0)
        lastNode=tmpPath[len(tmpPath)-1]
        #we're done
        if lastNode == end: return tmpPath
        #else we must run through children
        for node in graph.neighbours(lastNode):
            if node not in tmpPath:
                newPath=tmpPath+[node]
                queue.append(newPath)

print(printPath(bfs(g,0,5)))
