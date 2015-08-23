package main

import "github.com/roessland/gopkg/mathutil"
import "fmt"

func ConcatenatedProduct(integer, n int64) int64 {
    digits := []int64{}
    for i := int64(1); i <= n; i++ {
        digits = append(digits, mathutil.ToDigits(integer*i, 10)...)
    }
    return mathutil.FromDigits(digits, 10)
}

func TestConcatenatedProduct() {
    if ConcatenatedProduct(192, 3) != 192384576 { panic("babana") }
}

func main() {
    TestConcatenatedProduct()

    // All possible integers
    max := int64(0)
    for integer := int64(1); integer < 15000; integer++ {
        // Increase product length until result is > 987654321
        // Store the biggest number
        for n := int64(1); n <= 9; n++ {
            prod := ConcatenatedProduct(integer, n)
            if prod > 987654321 {
                break
            }
            if mathutil.IsPandigital(prod, 9) {
                if prod > max {
                    fmt.Printf("integer=%v, n=%v, %v\n", integer, n, prod)
                    max = prod
                }
            }
        }
    }
    fmt.Printf("%v\n", max)
}
