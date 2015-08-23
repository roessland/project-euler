# Problem 57 - Square root convergents
# Counts the number of expansions where the amount of digits in the numerator
# is bigger than the amount of digits in the denominator.
from fractions import gcd

N = 1000
t = 1
n = 1


nums = 0

for i in xrange(1, N+1):
    # Next fraction
    t, n = 2*n + t, n + t

    # Simplify
    GCD = gcd(t, n)
    t, n = t/GCD, n/GCD

    # Check length
    if len(str(t)) > len(str(n)):
        nums += 1

print nums
