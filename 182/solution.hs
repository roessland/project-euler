main = print $ sum [e | e <- [2..phi-1], gcd e phi == 1, (gcd (e-1) (p-1) +1)*(gcd (e-1) (q-1) + 1) == 9]
    where p = 1009
          q = 3643
          phi = (p-1)*(q-1)
