fibs :: [Integer]
fibs = 1:2:zipWith (+) fibs (tail fibs)

p2 = sum [x | x <-takeWhile (<=4000000) fibs, even x]

main = print $ p2
