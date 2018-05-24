#!/usr/bin/python

def conway(L):
    """
    takes a int:say-as-you-read
    """
    count=1
    result=""
    L=L.split()
    prev=L[0]

    for i in L[1:]:
        #skip the spaces
        if i==" ":continue
        else:
            if i!=prev:
                result+=str(count)+" "+str(prev)+" "
                count=1
                prev=i
            elif i==prev:
                count+=1

    # At the end of the string... or if string==1
    result+=str(count)+" "+str(prev)

    return(result)

#for i in range(5):
#    print(conway("2"))

#print(conway("1 1 25"))
R="25"
L=9
for x in range(L):
    res=conway(R)
    R=res
    print(res)
