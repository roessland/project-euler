digits :: (Integral a) => a -> [a]
digits n
    | n < 0              = error "Negative number"
    | 0 <= n && n <= 9   = [n]
    | otherwise          = digits (n `div` 10) ++ [n `mod` 10]

isPalindrome :: Integral a => a -> Bool
isPalindrome n = digits n == (reverse . digits) n

p4 = maximum [x*y | x <- [999,998..100], y <- [999, 998..100], isPalindrome (x*y)]

main = print p4
