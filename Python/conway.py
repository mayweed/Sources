#!/usr/bin/python

def conway(L):
    """
    takes a string:say-as-you-read
    """
    count=1
    result=" "
    prev=L[0]        
    for i in L[1:]:
        #skip the spaces
        if i==" ":continue
        else:
            if i!=prev:
                result+=str(count)+" "+prev+" "
                count=1
                prev=i
            elif i==prev:
                count+=1

    # At the end of the string... or if string==1
    result+=str(count)+" "+prev+" "

    return(result)

#for i in range(5):
#    print(conway("2"))

print(conway("3"))
