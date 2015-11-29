#!/usr/bin/bash
# Place yourself in the right directory
# Then wildcard_adb.sh /your/path/*

adb shell ls $1 | tr '\r' ' ' | xargs -n1 adb pull 
