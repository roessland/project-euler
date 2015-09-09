-- Runs in 20 seconds
--
-- cabal install primes
import Data.Numbers.Primes

main = do print $ head $ [n | (n,p) <- zip [1..] primes, (((p-1)^n) `mod` (p^2) + ((p+1)^n) `mod` (p^2)) `mod` (p^2) > 10^10]
