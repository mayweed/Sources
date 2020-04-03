#!/usr/bin/python

import os 
import os.path
import sys

"""
TODO: 
- does not work if it's not called in the directory!! Cant find files otherwise!!
  Must fix that!!
- wtf is wrong with goimports?
"""

# must check that it's directory + add usage() 
#filelist=os.listdir(sys.argv[1])
filelist= os.listdir("/home/guillaume/scripts/Sources/Go/CG/contests/ooc")
#why does it bug if i delete that?
filelist.sort()

# check for preexisting file and delete it eventually
#what about pathlib? cf https://linuxize.com/post/python-check-if-file-exists/
if os.path.isfile("cgfile.go"):
    os.system("rm cgfile.go")

#if i go by file i dont need that. Would be great though to have choice
#by args between dir and files...
def sanitizeList(fileList):
    """
    keep only *.go files in the directory
    better use a new clean list rather than deleting elements
    """
    cleanList=[]

    for f in filelist:
        if f.endswith(".go"):
            cleanList.append(f)

    return cleanList

def main():
    #get rid of non go files
    fl=sanitizeList(filelist)

    with open("cgfile.go",'a+')as bf:
        bf.write("//CG File crafted with love by assembleFile\n")
        bf.write("package main\n\n")
        for file in fl:
            inImport = True 
            with open(file,'r') as f:
                for line in f:
                    # 2 cases: either global import() or in line import
                    if inImport:
                        #remove all the lines in each file til the end of import
                        if line.startswith(")") :
                            inImport=False
                        #the beginning of a file after inline imports
                        elif line.startswith("func") or line.startswith("const") \
                                or line.startswith("//"):
                            inImport = False
                            #keep those ones
                            bf.write(line)
                        else:
                            continue
                    else:
                        bf.write(line)

#do NOT forget to add export PATH=$PATH:$GOPATH/bin in .bashrc
os.system("goimports cgfile.go")

if __name__ == "__main__":
        main()
