#!/bin/bash
# $1 = plus petit layer
# $2 = plus grand layer
# $3 = layer actuel
echo "+ \"Par défaut\" Pick Layer default"
for i in $(seq $1 $2); do
  echo -n "+ \"Layer $i";
  case $i in
    2) echo -n " (bas)";;
    4) echo -n " (normal)";;
    6) echo -n " (haut)";; 
  esac;
  if [[ $i == $3 ]]; then
    echo " (actuel)\" Nop";
  else
    echo "\" Pick Layer 0 $i";
  fi;
done;

