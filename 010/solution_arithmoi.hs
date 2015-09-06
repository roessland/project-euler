import Math.NumberTheory.Primes.Sieve
main = print $ sum $ takeWhile (<2000000) primes
