import Math.NumberTheory.Primes.Factorisation

triangle n = ((n)*(n+1)) `div` 2
triangles = map triangle [1..]

main = print $ take 1 [x | x <- triangles, (tau x) > 500]
