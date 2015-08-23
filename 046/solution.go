package main

import "fmt"
import "github.com/roessland/gopkg/primegen"

func main() {
    // Tactic: Generate all numbers of this form up to a certain size. Filter
    // out primes. The remainder are odd composities that cannot be written as
    // the sum of a prime and twice a square.
    MAX := int64(10000)

    isPrime := primegen.Map(MAX)
    primes := primegen.SliceFromMap(isPrime)
 
    _ = isPrime
    _ = primes

    // Generate lots of numbers
    wasGenerated := make([]bool, MAX+1)
    for _, p := range primes {

        // If the prime is this big, the number is bigger than max
        if p >= MAX {
            break
        }

        // Starting from zero removes all primes
        for i := int64(0); ; i++ {
            num := p + 2*i*i
            if num > MAX {
                break
            }
            wasGenerated[num] = true
        }


    }

    // Remove all even numbers
    for i := int64(0); i <= MAX; i += 2 {
        wasGenerated[i] = true
    }

    // Print some numbers
    for num, val := range wasGenerated {
        if !val {
            fmt.Printf("Possible candidate: %v\n", num)
        }
    }


}
