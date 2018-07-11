import sys
import math

n = int(input())

# https://www.daniweb.com/programming/software-development/code/216880/check-if-a-number-is-a-prime-number-python
#Thx to him for that func!!
def isPrime(n):
    '''check if integer n is a prime'''
    # make sure n is a positive integer
    n = abs(int(n))
    # 0 and 1 are not primes
    if n < 2:
        return False
    # 2 is the only even prime number
    if n == 2: 
        return True    
    # all other even numbers are not primes
    if not n & 1: 
        return False
    # range starts with 3 and only needs to go up the squareroot of n
    # for all odd numbers
    for x in range(3, int(n**0.5)+1, 2):
        if n % x == 0:
            return False
    return True
        
def isCarNum(n):
    a=498081
    return a**n%n == a%n
        
# Write an action using print
print(isPrime(9), file=sys.stderr)

#Carmichael num are NOT prime
if isPrime(n):
    print("NO")
elif not isPrime(n):
    if isCarNum(n):
        print("YES")
    else:
        print("NO")
