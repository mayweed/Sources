#!/usr/bin/python

import os

filelist=os.listdir("/home/raimondeaug/scripts/Sources/Go/fb")
bf=open("bigfile.txt",'r+')

for file in filelist:
    if file.endswith(".go"):
        with open(file,'r') as f:
            f.write(bf)

#should close bf
