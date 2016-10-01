#!/bin/bash
# $1 = nombre de bureaux
# $2 = bureau courant
for i in $(seq 0 $(($1-1))|grep -v $2); do
  echo "+ \"\$[desk.name$i]\" MoveToDesk 0 $i";
done;
