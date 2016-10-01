#!/bin/bash

NOW=`date '+%Y%m%d'`
NAME=fvwm${NOW}
EXT=jpg

cd ${screenshots_dir}

scrot -d 5 ${NAME}.${EXT}

gqview ${NAME}.${EXT}
