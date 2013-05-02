# Solves problem 32 - Pandigital products
import itertools

def magic(numbers):
    """Convert list of digits to integer"""
    return int(''.join([ "%d"%x for x in numbers]))

prods = set()

#perm = [3,9,1,8,6,7,2,5,4]

previousperm = 0
for perm in itertools.permutations([1,2,3,4,5,6,7,8,9]):
    if previousperm != perm[1]:
        print perm[1]
        previousperm = perm[1]

    for multpos in range(1,8):
        for eqpos in range(multpos+1, 8):
            multiplicand = magic(perm[0:multpos])
            multiplier = magic(perm[multpos:eqpos])
            product = magic(perm[eqpos:])
            #print multiplicand, multiplier, product
            if multiplicand * multiplier == product:
                prods.add(product)

print sum(prods)
    
