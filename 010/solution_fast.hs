-- merge infinite sorted lists of distinct integers
-- duplicates may exist across lists, but not within lists
merge :: Ord a => [a] -> [a] -> [a]
merge xs@(x:xt) ys@(y:yt) =
    case compare x y of
        LT -> x : merge xt ys
        EQ -> x : merge xt yt
        GT -> y : merge xs yt

-- diff returns the first list minus the elements in the second list
diff :: Ord a => [a] -> [a] -> [a]
diff xs@(x:xt) ys@(y:yt) =
    case compare x y of
        LT -> x : diff xt ys
        EQ -> diff xt yt
        GT -> diff xs yt

primes, nonprimes :: [Integer]
primes    = [2, 3, 5] ++ (diff [7, 9..] nonprimes)
nonprimes = foldr1 f $ map g $ tail primes
    where
        f (x:xt) ys = x : (merge xt ys)
        g p         = [n * p | n <- [p, p + 2 ..]]

main = print $ sum $ takeWhile (<2000000) primes
