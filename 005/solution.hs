gcd' a 1 = a
gcd' 1 b = b
gcd' a b = gcd' b (a `mod` b)

lcm' a b = a * (b `quot` (gcd a b))

myFoldl :: (b -> a -> b) -> b -> [a] -> b
myFoldl _ i [] = i
myFoldl f i (x:xs) = myFoldl f (f i x) xs

p5 = myFoldl lcm 1 [1..20]

main = print p5
