import sys
 
class Graph:
    def __init__(self):
        self.nodes=[]
        self.edges={}
        self.exits=[]
    def neighbours(self,node):
        return self.edges[node]
        
class Node(Graph):
    self.exit_gw=false
    def __init__(self,name):
        self.name=name
    def addNode(self,node):
        if node not in self.nodes:
            self.nodes.append(node)
            self.edges[node]=[]    
            
class Edge(Graph):    
    def __init__(self,src,dest):
        self.src=src
        self.dest=dest
    def addEdge(self, edge):
        self.edges[edge.getSrc()].append(edge.getDest())
        self.edges[edge.getDest()].append(edge.getSrc())    
              
              
# n: the total number of nodes in the level, including the gateways
# l: the number of links
# e: the number of exit gateways
n, l, e = [int(i) for i in input().split()]
graph=Graph()

for i in range(l):
    # n1: N1 and N2 defines a link between these nodes
    n1, n2 = [int(j) for j in input().split()]
    graph.addNode(n1)
    graph.addNode(n2)
    graph.addEdge(Edge(n1,n2))
#graph.test()    
exits=[int(input()) for i in range(e)]

def bfs(graph,start,end):
    queue=[]
    path=[]
    queue.append(path)
    path.append(start)
    while len(queue)!=0:
        tmpPath=queue.pop(0)
        lastNode=tmpPath[-1]
        if lastNode==end:return tmpPath
        for n in graph.neighbours(lastNode):
            if n not in tmpPath  :
                newPath=tmpPath+[n]
                queue.append(newPath)
# game loop
while 1:
    si = int(input())  # The index of the node on which the Skynet agent is positioned this turn
    
    # search for the shortest path to whatever exit 
    for gw in exits:
        path=bfs(graph,si,exits[gw])
        
    # Write an action using print
    print(n,l,exits[0],path,si,si_neighbours,file=sys.stderr)
