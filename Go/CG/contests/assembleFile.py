#!/usr/bin/python

import os
import sys

"""
TODO: 
- does not work if it's not called in the directory!! Cant find files otherwise!!
  Must fix that!!
- for each and every file it imports package + import!!
- check the import statements: should I use goimports?
"""

# must check that it's directory + add usage() 
#filelist=os.listdir(sys.argv[1])
filelist=os.listdir("/home/guillaume/scripts/Sources/Go/CG/contests/ooc")
filelist.sort()

# check for preexisting file and delete it eventually
if "bigfile.go" in filelist:
    os.system("rm bigfile.go")

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
    for f in fl:
        sys.stderr.write(f)

    with open("bigfile.go",'a+')as bf:
        bf.write("//File crafted with love by assembleFile\n")
        #add a package in our big file
        bf.write("package main")

        for file in fl:
            with open(file,'r') as f:
                for line in f:
                    if line.startswith("package"): # or \
                    #line.startswith(" \"")or line.startswith(")") :
                            continue
                    # import group does not work :'(
                    elif line.startswith("import"):
                        #if line.endswith("("):
                            #tant que la ligne ne commence pas par ) ignorer
                            #for line in f:
                                #continue
                            if line.startswith(")"):
                                break
                            #else:
                                #import by line
                            #    continue
                    else:
                        bf.write(line)

#os.system("goimports bigfile.go")
if __name__ == "__main__":
        main()
