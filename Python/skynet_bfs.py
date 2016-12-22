import sys

class Node:
    def __init__(self,id):
        self.id=id
        self.exit_gw=False
        
class Edge:  
    def __init__(self,src,dest):
        self.src=src
        self.dest=dest
    def getSrc(self):
        return self.src
    def getDest(self):
        return self.dest
        
class Graph(Edge):
    def __init__(self):
        self.nodes=[]
        self.edges={}
        self.exits=[]
    def addNode(self,node):
        if node not in self.nodes:
            self.nodes.append(node)
            self.edges[node.id]=[] 
    def addEdge(self, edge):
        self.edges[edge.getSrc()].append(edge.getDest())
        self.edges[edge.getDest()].append(edge.getSrc())    
    def bfs(self,start,end):
        pass
    
    ### HANDLERS ###
    def neighbours(self,node):
        return self.edges[node]
    def exitFlag(self,exits):
        for e in exits:
            for n in self.nodes:
                if n==e:
                    n.exit_gw=True
              
class Algos(Graph):
    '''a class to store some graph algos'''
    def bfs(self,start,end):
        queue=[]
        path=[]
        queue.append(path)
        path.append(start)
        while len(queue)!=0:
            tmpPath=queue.pop(0)
            lastNode=tmpPath[-1]
            if lastNode==end:return tmpPath
            for n in self.neighbours(lastNode):
                if n not in tmpPath  :
                    newPath=tmpPath+[n]
                    queue.append(newPath)
              
# n: the total number of nodes in the level, including the gateways
# l: the number of links
# e: the number of exit gateways
n, l, e = [int(i) for i in input().split()]
graph=Graph()

for i in range(l):
    # n1: N1 and N2 defines a link between these nodes
    n1, n2 = [int(j) for j in input().split()]
    graph.addNode(Node(n1))
    graph.addNode(Node(n2))
    graph.addEdge(Edge(n1,n2))
    
exits=[int(input()) for i in range(e)]
graph.exitFlag(exits)

# game loop
while 1:
    si = int(input())  # The index of the node on which the Skynet agent is positioned this turn
    
    # search for the shortest path to whatever exit 
    #for gw in exits:
    #    path=graph.bfs(si,exits[gw])
        
    # Write an action using print
    #print(...,file=sys.stderr)
