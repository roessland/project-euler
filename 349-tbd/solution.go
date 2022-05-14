package main

import "fmt"
import "log"

type Pos struct {
	X, Y int
}

func (p Pos) Add(v Pos) Pos {
	return Pos{p.X + v.X, p.Y + v.Y}
}

var N Pos = Pos{0, 1}
var W Pos = Pos{1, 0}
var S Pos = Pos{0, -1}
var E Pos = Pos{-1, 0}

func CW(v Pos) Pos {
	switch v {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	default:
		log.Fatal("Unknown direction", v)
	}
	return Pos{}
}

func CCW(v Pos) Pos {
	switch v {
	case N:
		return W
	case E:
		return N
	case S:
		return E
	case W:
		return S
	default:
		log.Fatal("Unknown direction", v)
	}
	return Pos{}
}

func DrawGrid(grid map[Pos]int) {
	fmt.Println()
	for y := -30; y <= 30; y++ {
		for x := -30; x <= 60; x++ {
			if grid[Pos{x, y}] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	grid := map[Pos]int{}
	currPos := Pos{0, 0}
	currDir := Pos{0, 1}
	numBlack := 0
	for i := 1; i <= 140000; i++ {
		if grid[currPos] == 0 {
			currDir = CW(currDir)
			grid[currPos] = 1
			numBlack++
		} else {
			currDir = CCW(currDir)
			grid[currPos] = 0
			numBlack--
		}
		currPos = currPos.Add(currDir)
		if i%104 == 0 && i%10000 == 0 {
			fmt.Println("After", i, "moves there are", numBlack, "black squares")
			DrawGrid(grid)
		}
	}
}
