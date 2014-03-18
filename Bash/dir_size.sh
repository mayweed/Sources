#!/bin/bash

# Calcul la taile d'un répertoire a partir de l'argument donné

du -mcah $1 | tail -1 | cut -f1
