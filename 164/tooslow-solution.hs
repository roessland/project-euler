nextDigits :: Integer -> Integer -> [Integer]
nextDigits a b = [0..(9-a-b)]

starts = [(a,b) | a <- [1..9], b <- [0..9], a+b < 10]

count 0 a b = 1
count n a b = sum [count (n-1) b c | c <- nextDigits a b]

--18
main = do print $ sum [count 6 (fst d) (snd d) | d <- starts]
