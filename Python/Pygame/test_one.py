#!/usr/bin/python

import pygame

pygame.init()

#Tuple packing!!
size=width,height=320,420
black=0,0,0
red=255,0,0
screen=pygame.display.set_mode(size)

#RGB chart: http://www.discoveryplayground.com/computer-programming-for-kids/rgb-colors/
#while 1:
#make it stick!! + event!!
screen.fill(black)
pygame.draw.line(screen,red,(10,10),(100,100),4)
