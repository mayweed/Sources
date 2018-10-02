#!/usr/bin/python

import os
import sys

"""
TODO: 
- first all the files then parse main.go
- check the import statements: should I use goimports?
- what about gofmt? https://blog.golang.org/go-fmt-your-code
- put package an import statement at the beginning of the file?
"""

# the directory should be given as an argument
filelist=os.listdir("/home/raimondeaug/scripts/Sources/Go/fb")
filelist.sort()

# check for preexisting file and delete it eventually
if "bigfile.go" in filelist:
    os.system("rm bigfile.go")

def sanitizeList(fileList):
    """
    keep only *.go files
    better use a new clean list rather than deleting elements
    """
    cleanList=[]

    for f in filelist:
        if f.endswith(".go"):
            cleanList.append(f)

    return cleanList

def scanLines(fileObject,line):
    """
    takes a file, scan all the lines
    if line found get true and exit
    else return false
    """
    #do not use open() twice!!
    for l in fileObject:
        if line==l:return False
        else: return True
    

def scanAndAdd (fileList):
    """
    First sanitize the list
    then check all lines of file and write import and package at the
    beginning
    """
    for file in fileList:
        with open(file,'r') as f:
            # w+ to create the file, works with a+
            with open("bigfile.go",'a+')as bf:
                for line in f:
                    if line.startswith("package") or line.startswith("import"):
                        #does not work...
                        #if scanLines(bf,line):
                        #    print(scanLines(bf,line),file=sys.stderr)
                        bf.write(line)


def assemble(fileList):
    #get rid of non go files
    fl=sanitizeList(fileList)

    #import and package
    scanAndAdd(fl)

    # adding what's left
    for file in fl:
        #do not include backup from cg
        if file!="allin.go":
            with open(file,'r') as f:
                # w+ to create the file, works with a+
                with open("bigfile.go",'a+')as bf:
                    bf.write("//"+file+"\n")
                    for line in f:
                        #prune files from package and import except main?
                        if line.startswith("package") or \
                            line.startswith("import") or \
                            line.startswith("//"):
                            continue
                        else:
                            bf.write(line)


assemble(filelist)

os.system("goimports bigfile.go")
