#!/bin/bash

#Usage: ./version paquet

pacman -Qi $1 | grep -i version
