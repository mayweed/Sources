#!/usr/bin/python

class Cell(object):
    def __init__(self,row,column):
        self.row=row
        self.column=column
        # a hashtable of links:keep track of which neighboring links are joined
        # by a passage
        # key is a Cell tuple, value is true if connected false otherwise
        links={}

    def link(self,cell):
        #should check here for boundaries of cell??
        self.links[(cell.row,cell.column)]=True

    def unlink(self,cell):
        self.links[(cell.row,cell.column)]=False

    def is_cell_linked(self,cell):
        '''
        is cell linked with self?
        '''
        if links[(cell.row,cell.column)]==True: return True
        else:return False
    
    def linked_cells(self):
        '''
        return the list of all cells connected to self
        '''
        linked=[]
        for cell in list(links.key()):
            if links[cell]==True: linked.append(cell)
        return linked

# simpler but not method like thing..
#grid=grid[row][col]
#while r < self.row:
#    for col in range(self.column):
#        grid[r][c]=Cell(r,col)

class Grid:
    #wouldn't it be more clear to write that as a 2D with contains cell objects?
    def __init__(self,row,column):
        self.row=row
        self.column=column
        #each row of cell objects in a list
        self.grid=[[Cell(r,col) for r in range(self.row)] for col in range(self.column)] 

    # http://stackoverflow.com/questions/9884132/what-exactly-are-pythons-iterator-iterable-and-iteration-protocols    
    # cf first example (with the four methods!!):http://stackoverflow.com/questions/19151/how-to-make-class-iterable
    def __iter__(self):
        #init the grid with Cell objects
        r=0
        while r < self.row:
            for col in range(self.column):
                c=Cell(r,col)
                yield c
            r+=1
    
    def get_num_cells(self):
        return self.row*self.column

    def access_cell(self,row,col):
        return self.grid[row][col]
