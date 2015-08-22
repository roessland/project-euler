// Bruteforce
package main

import "fmt"
import _ "math/big"
import "github.com/roessland/gopkg/mathutil"

func main() {
    var min_value float64 = 2.0/5.0
    var max_value float64 = 3.0/7.0
    var D int64 = 1000000

    // Closest value
    var max_f float64 = 0.0

    for d := int64(1); d <= D; d++ {
        min_n := int64(float64(d) * min_value) - 1
        max_n := int64(float64(d) * max_value) + 1
        for n := min_n; n < max_n; n++ {
            f := float64(n) / float64(d)
            if min_value < f && f < max_value && f > max_f{
                gcd := mathutil.GCD(n, d)
                max_f = f
            }
        }
    }
    fmt.Printf("%v / %v = %v\n", n/gcd, d/gcd, float64(n)/float64(d))
    fmt.Println("habeeb it")
}
