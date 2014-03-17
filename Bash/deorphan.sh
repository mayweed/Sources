#!/bin/bash

for d in `pacman -Qdt`; do pacman -R $d; done
