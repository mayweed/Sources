#!/usr/bin/python

import os

filelist=os.listdir("/home/raimondeaug/scripts/Sources/Go/fb")
bf=open("bigfile.txt",'w+')

for file in filelist:
    if file.endswith(".go"):
        with open(file,'r') as f:
            bf.write()

#should close bf
