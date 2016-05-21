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
         for i in in range (len L[1:]):
            #Do not forget the end of a list
            if L[i]==prev:
                count+=1
                result.append([count,i]) 
            if i!=prev:
                result.append([count,i])
                count=1
                prev=i
    return(result)

#for i in range(5):
#    print(conway([2]))

print(conway([3,2,1,1,1,1]))
#Should implement recursively
# base case: len==1 ==>1
# if len > 1:
# read the line char by char
# if one is alone of is kind: base case
# else count len and output final_count(len) i
