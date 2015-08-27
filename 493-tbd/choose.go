package main

import "fmt"
import "math/rand"

var balls []int = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6}

func Trial(res chan<- int64) {
	for {
		total := int64(0)

		for i := 0; i < 100000; i++ {
			colorPicked := make([]bool, 7)
			ballPicked := make([]bool, 70)

			ballsLeft := 20
			for ballsLeft > 0 {
				ballNumber := rand.Intn(70)
				if !ballPicked[ballNumber] {
					// Remove that ball from circulation
					ballsLeft--
					ballPicked[ballNumber] = true

					// Note that a ball of this color was found
					colorPicked[balls[ballNumber]] = true
				}
			}

			distinctColors := 0
			for _, yes := range colorPicked {
				if yes {
					distinctColors++
				}
			}
			total += int64(distinctColors)
		}
		res <- total
	}
}

func main() {
	res := make(chan int64, 100)
	total := int64(0)
	numTrials := int64(0)

	go Trial(res)
	go Trial(res)
	go Trial(res)
	go Trial(res)

	for colors := range res {
		total += colors
		numTrials += 100000

		if numTrials%100000 == 0 {
			fmt.Printf("%v\n", float64(total)/float64(numTrials))
		}
	}
}
