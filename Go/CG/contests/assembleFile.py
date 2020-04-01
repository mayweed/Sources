#!/usr/bin/python

import os
import sys

"""
TODO: 
- first all the files then parse main.go
- check the import statements: should I use goimports?
- what about gofmt? https://blog.golang.org/go-fmt-your-code
"""

# must check that it's directory + add usage() 
filelist=os.listdir(sys.argv[1])
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

def assemble(fileList):
    #get rid of non go files
    fl=sanitizeList(fileList)

    for file in fl:
            with open(file,'r') as f:
                with open("bigfile.go",'a+')as bf:
                    bf.write("//"+file+ "crafted with love by assembleFile\n")
                    #add a package in our big file
                    bf.write("package main")
                    for line in f:
                        if line.startswith("package") :
                            continue
                        elif line.startswith("import"):
                            # import group
                            if line.endswith("("):
                                #tant que la ligne ne commence pas par ) ignorer
                                for line in f:
                                    continue
                                    if line.startswith(")")
                                        break
                            elif:
                                #import by line
                                continue
                        else:
                            bf.write(line)


assemble(filelist)

os.system("goimports bigfile.go")
