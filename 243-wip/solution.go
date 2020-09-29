// Problem 243.
// The basic observation here is that
//
//     R(d) = EulerPhi(d)/(d-1)
//
// where EulerPhi(d) is the number of numbers less than d which are relatively
// prime to d. In other words, keep incrementing d and evaluating R(d) until
// the result is found.
package main

import (
    "fmt"
)

func GCD(r0, r1 int) int {
    if r1 > r0 {
        r0, r1 = r1, r0
    }

    for r1 != 0 {
        r0, r1 = r1, r0 % r1
    }

    return r0
}

func TestGCD() {
    if GCD(453242343, 442143147) != 3 {
        panic("GCD fail")
    }
}

func EulerPhi(n int) int {
    phi := 0
    for d := 1; d < n; d = d+1 {
        if GCD(n, d) == 1 {
            phi++
        }
    }
    return phi
}

func TestEulerPhi() {
    if EulerPhi(60) != 16 {
        panic("phi fail!")
    }
}

func R(n int) float64 {
    return float64(EulerPhi(n)) / float64(n-1)
}


func main() {
    TestGCD()
    TestEulerPhi()

    upper := float64(15499)/float64(94744)
    almost := float64(19499)/float64(94744)
    for n := 2; ; n = n+1 {
        resilience := R(n)
        fmt.Printf("%v ", n)
        if resilience < almost {
            fmt.Printf("That was close %v\n", n)
        }
        if resilience < upper {
            fmt.Printf("The number is %v!\n", n)
            break
        }
    }
}
