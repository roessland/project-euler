package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

var idealFrequencyAZ = []float64{
    0.081, // A
    0.014, // B
    0.027, // C
    0.043, // D
    0.127, // E
    0.022, // F
    0.020, // G
    0.061, // H
    0.070, // I
    0.002, // J
    0.008, // K
    0.040, // L
    0.024, // M
    0.067, // N
    0.075, // O
    0.019, // P
    0.001, // Q
    0.060, // R
    0.063, // S
    0.091, // T
    0.028, // U
    0.010, // V
    0.024, // W
    0.001, // X
    0.020, // Y
    0.001, // Z
    0.000, // special characters
}

func ReadCiphertext(filename string) []uint8 {
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Errorf("Oops!")
    }
    s := strings.TrimSpace(string(buf))
    cipherStr := strings.Split(s, ",")

    cipher := make([]uint8, len(cipherStr))
    for i, numStr := range cipherStr {
        num, err := strconv.ParseUint(numStr, 10, 8)
        if err != nil {
            panic(err)
        }
        cipher[i] = uint8(num)
    }
    return cipher
}

func Decrypt(cipher []uint8, key []uint8) []uint8 {
    plain := make([]uint8, len(cipher))

    k := 0
    for i, c := range cipher {
        plain[i] = c^key[k]
        k = (k+1) % len(key)
    }
    return plain
}

func Frequencies(plain []uint8) []float64 {
    frequencies := make([]float64, 27)

    for _, p := range plain {
        // Change lowercase p into uppercase
        if 97 <= p && p <= 122 {
            p -= 32
        }

        if 65 <= p && p <= 90 {
            frequencies[p-65] += 1.0 / float64(len(plain))
        } else {
            frequencies[26] += 1.0 / float64(len(plain))
        }
    }
    return frequencies
}

func Score(frequencies []float64) float64 {
    var score float64 = 0.0
    for i := 0; i < 27; i++ {
        score += (idealFrequencyAZ[i] - frequencies[i])*(idealFrequencyAZ[i] - frequencies[i])
    }
    return score
}

func Sum(plain []uint8) int64 {
    var sum int64 = 0
    for _, p := range plain {
        sum += int64(p)
    }
    return sum
}


func main() {
    cipher := ReadCiphertext("cipher.txt")

    for k1 := 97; k1 <= 122; k1++ {
        for k2 := 97; k2 <= 122; k2++ {
            for k3 := 97; k3 <= 122; k3++ {
                key := []uint8{uint8(k1),uint8(k2),uint8(k3)}
                plain := Decrypt(cipher, key)
                frequencies := Frequencies(plain)
                score := Score(frequencies)
                if score < 0.07 { // Found by trial and error
                    fmt.Printf("%v\n", string(plain))
                    fmt.Printf("Key: %v = %v\n", key, string(key))
                    fmt.Printf("Sum of ASCII: %v\n", Sum(plain))
                }
            }
        }
    }
}
