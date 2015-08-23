package main

import "github.com/roessland/gopkg/primegen"
import "math"
import "fmt"

func Max(a, b int64) int64 {
    if a > b {
        return a
    } else {
        return b
    }
}

func M(p, q, N int64) int64 {
    // The maximum combination
    var m int64 = 0

    // Multiplicities
    kq := 1
    prod := p * q

    // In this there does not exist a number below N divisible by both p and q
    if prod > N {
        return 0
    }

    // Increase kq as much as possible
    for prod*q <= N {
        prod *= q
        kq += 1
    }
    // Increase kp as much as possible
    for prod*p <= N {
        prod *= p
    }
    m = prod

    // Decrease kq one by one if possible
    for kq > 1 {
        prod /= q
        kq -= 1
        // Increase kp as much as possible
        for prod*p <= N {
            prod *= p
        }
        m = Max(m, prod)
    }
    return m
}

func TestM() {
    if 96 != M(2, 3, 100) { panic("BANANA") }
    if 0 != M(2, 73, 100) { panic("WHAT") }
    if 63 != M(3, 7, 100) { panic("OY MATE") }
    if 567 != M(3, 7, 1000) { panic("YEEHAW") }
    if 75 != M(3, 5, 100) { panic("BABA") }
    if 133 != M(7, 19, 133) { panic("LADIDA") }
    if 0 != M(7, 19, 132) { panic("WHOOPS") }
}

func main() {
    TestM()
    var N int64 = 10000000
    primes := primegen.SliceFromMap(primegen.Map(N)) // 80 MB of primes
    values := make([]bool, N+1)

    sqrtN := int64(math.Sqrt(float64(N)))
    for i := int64(0); i <= sqrtN; i++ {
        p := primes[i]
        j := i + 1
        for {
            q := primes[j]
            if p * q > N {
                break
            }
            values[M(p, q, N)] = true
            j++
        }
    }

    S := int64(0)
    for index, found := range values {
        if found {
            S += int64(index)
        }
    }
    fmt.Printf("The sum of distinct values is %v\n", S)
}
