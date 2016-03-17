#!/usr/bin/python

def conway(L):
    """
    takes a list:say-as-you-read
    """
    count=1
    result=[]
    if len(L)==1:result=[1,L[0]]
    else:
         prev=L[0]        
         for i in L[1:]:
            #Do not forget the end of a list
            if i==prev:
                count+=1
                result.append([count,i]) 
            if i!=prev:
                result.append([count,i]) 
                count=1
    return(result)

print(conway([2,1]))
