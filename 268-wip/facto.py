def gcd(*numbers):
    """Return the greatest common divisor of the given integers"""
    from fractions import gcd
    return reduce(gcd, numbers)

# Least common multiple is not in standard libraries? It's in gmpy, but this is simple enough:

def lcm(*numbers):
    """Return lowest common multiple."""
    def lcm(a, b):
        return (a * b) // gcd(a, b)
    return reduce(lcm, numbers, 1)

N = 100
P1 = 2*3
P2 = 2*5
P3 = 2*7
P4 = 3*5
P5 = 3*7
P6 = 5*7

#N0 = 0
#N1 = N0 + (N-1)//P1
#N2 = N1 + (N-1)//P2 - (N-1)/lcm(P1, P2)
#N3 = N2 + (N-1)//P3 - (N-1)/lcm(P1, P3) - (N-1)/lcm(P2, P3)
#N4 = N3 + (N-1)//P4 - (N-1)/lcm(P1, P4) - (N-1)/lcm(P2, P4) - (N-1)/lcm(P3, P4)
#N5 = N4 + (N-1)//P5 - (N-1)/lcm(P1, P5) - (N-1)/lcm(P2, P5) - (N-1)/lcm(P3, P5) - (N-1)/lcm(P4, P5)
#N6 = N5 + (N-1)//P6 - (N-1)/lcm(P1, P6) - (N-1)/lcm(P2, P6) - (N-1)/lcm(P3, P6) - (N-1)/lcm(P4, P6) - (N-1)/lcm(P5, P6)
print - ((N-1)/lcm(P1, P2))
print - ((N-1)/lcm(P1, P2))
print - ((N-1)/lcm(P1, P3) + (N-1)/lcm(P2, P3))
print - ((N-1)/lcm(P1, P4) + (N-1)/lcm(P2, P4) + (N-1)/lcm(P3, P4))
print - ((N-1)/lcm(P1, P5) + (N-1)/lcm(P2, P5) + (N-1)/lcm(P3, P5) + (N-1)/lcm(P4, P5))
print - ((N-1)/lcm(P1, P6) + (N-1)/lcm(P2, P6) + (N-1)/lcm(P3, P6) + (N-1)/lcm(P4, P6) + (N-1)/lcm(P5, P6))
