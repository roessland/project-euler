package main

import "fmt"
import "github.com/roessland/gopkg/primegen"

var isPrime []bool

type UlamSpiral struct {
    SideLength int64
    NumbersOnDiagonals int64
    LastNumber int64
    PrimesOnDiagonals int64
}

func (s UlamSpiral) AddLayer() UlamSpiral {
    var t UlamSpiral = s
    t.SideLength = s.SideLength + 2
    t.NumbersOnDiagonals = s.NumbersOnDiagonals + 4
    t.LastNumber = s.LastNumber + 4*(s.SideLength + 1)

    // Count new primes on the diagonals
    A := s.LastNumber + s.SideLength + 1
    B := A + s.SideLength + 1
    C := B + s.SideLength + 1
    D := C + s.SideLength + 1
    if isPrime[A] { t.PrimesOnDiagonals++ }
    if isPrime[B] { t.PrimesOnDiagonals++ }
    if isPrime[C] { t.PrimesOnDiagonals++ }
    if isPrime[D] { t.PrimesOnDiagonals++ }

    return t
}

func (s UlamSpiral) PrimeRatio() float64 {
    return float64(s.PrimesOnDiagonals) / float64(s.NumbersOnDiagonals)
}

func main() {
    isPrime  = primegen.Map(240000000)
    spiral := UlamSpiral{1, 1, 1, 0}

    // Add a new layer, since the initial prime ratio is 0, since there are no
    // primes.
    spiral = spiral.AddLayer()

    for spiral.PrimeRatio() > 0.1 {
        fmt.Printf("Last number: %v, Primeratio: %v\n", spiral.LastNumber, spiral.PrimeRatio())
        spiral = spiral.AddLayer()
    }
    fmt.Printf("%v\n", spiral)
}
