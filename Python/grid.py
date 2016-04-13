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
        if links[(cell.row,cell.column)]=True: return True
        else:return False
    
    def linked_cells(self):
        '''
        return the list of all cells connected to self
        '''
        linked=[]
        for cell in list(links.key()):
            if links[cell]==True: linked.append(cell)
        return linked
