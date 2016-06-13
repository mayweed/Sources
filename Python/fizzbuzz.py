#!/usr/bin/python

for num in range(100):
    if num%3==0:
        if num%5==0:
            print(num,"fizzbuzz")
        else:            
            print(num,"fizz")
    elif num%5==0:
        print(num,"buzz")

