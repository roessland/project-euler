main = print $ head
    [a*b*c | a <- [1..500],
             b <- [1..500],
             let c = 1000 - a - b,
             a^2 + b^2 == c^2]
