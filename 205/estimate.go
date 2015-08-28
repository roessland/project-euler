package main

import "fmt"
import "math/rand"

func ColinWin() bool {
	d94 := 9 + rand.Intn(4) + rand.Intn(4) + rand.Intn(4) + rand.Intn(4) +
		rand.Intn(4) + rand.Intn(4) + rand.Intn(4) + rand.Intn(4) + rand.Intn(4)

	d66 := 6 + rand.Intn(6) + rand.Intn(6) + rand.Intn(6) + rand.Intn(6) + rand.Intn(6) + rand.Intn(6)
	if d94 > d66 {
		return true
	}
	return false
}

func main() {
	colinWins := int64(0)
	totalRolls := int64(0)
	for {
		if ColinWin() {
			colinWins++
		}
		totalRolls++

		if totalRolls%100000 == 0 {
			fmt.Printf("%v\n", float64(colinWins)/float64(totalRolls))
		}
	}
}
