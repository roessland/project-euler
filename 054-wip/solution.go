package main

import "fmt"
import "io"
import "os"
import "bufio"

var values_int map[byte]int = map[byte]int{
	'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
	'8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
}
var values_byte map[int]byte = map[int]byte{
	2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7',
	8: '8', 9: '9', 10: 'T', 11: 'J', 12: 'Q', 13: 'K', 14: 'A',
}

type Card struct {
	V int  // Value
	S byte // Suit
}

func (c Card) String() string {
	return fmt.Sprintf("%c%c", values_byte[c.V], c.S)
}

func CardFromStr(c string) Card {
	if len(c) != 2 {
		panic("String length is wrong")
	}
	val, ok := values_int[c[0]]
	if !ok {
		panic("Invalid card value")
	}
	switch c[1] {
	case 'H':
	case 'D':
	case 'C':
	case 'S':
	default:
		panic("Invalid card suit")
	}
	return Card{val, c[1]}
}

func ValueInHand(value int, h []Card) bool {
	for _, hc := range h {
		if value == hc.V {
			return true
		}
	}
	return false
}

func ValuesInHand(values []int, h []Card) bool {
	for _, value := range values {
		foundValue := false
		for _, hc := range h {
			if value == hc.V {
				foundValue = true
				break
			}
		}
		if !foundValue {
			return false
		}
	}
	return true
}

func CardInHand(c Card, h []Card) bool {
	for _, hc := range h {
		if c == hc {
			return true
		}
	}
	return false
}

func CardsInHand(cs []Card, h []Card) bool {
	for _, c := range cs {
		foundCard := false
		for _, hc := range h {
			if c == hc {
				foundCard = true
				break
			}
		}
		if !foundCard {
			return false
		}
	}
	return true
}

func SameSuit(cs []Card) bool {
	s := cs[0].S
	for _, c := range cs {
		if c.S != s {
			return false
		}
	}
	return true
}

// Returns 0 if it's not a straight flush. The highest value otherwise.
func StraightFlush(h []Card) int {
	if !SameSuit(h) {
		return 0
	}
	if ValuesInHand([]int{14, 2, 3, 4, 5}, h) {
		return 5
	}
	if ValuesInHand([]int{2, 3, 4, 5, 6}, h) {
		return 6
	}
	if ValuesInHand([]int{3, 4, 5, 6, 7}, h) {
		return 7
	}
	if ValuesInHand([]int{4, 5, 6, 7, 8}, h) {
		return 8
	}
	if ValuesInHand([]int{5, 6, 7, 8, 9}, h) {
		return 9
	}
	if ValuesInHand([]int{6, 7, 8, 9, 10}, h) {
		return 10
	}
	if ValuesInHand([]int{7, 8, 9, 10, 11}, h) {
		return 11
	}
	if ValuesInHand([]int{8, 9, 10, 11, 12}, h) {
		return 12
	}
	if ValuesInHand([]int{9, 10, 11, 12, 13}, h) {
		return 13
	}
	if ValuesInHand([]int{10, 11, 12, 13, 14}, h) {
		return 14
	}
	return 0
}

// Returns 0 if there are not four of a kind.
// Returns the value of the four cards if they exist.
// Also returns the value of the remaining card if it was a four of a kind.
func FourOfAKind(h []Card) (int, int) {

}

// Returns 1 if player 1 wins, and 2 if player 2 wins
func Winner(A, B []Card) int {
	return 1
}

func main() {
	wins := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		c, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		A := []Card{
			CardFromStr(c[0:2]),
			CardFromStr(c[3:5]),
			CardFromStr(c[6:8]),
			CardFromStr(c[9:11]),
			CardFromStr(c[12:14]),
		}
		B := []Card{
			CardFromStr(c[15:17]),
			CardFromStr(c[18:20]),
			CardFromStr(c[21:23]),
			CardFromStr(c[24:26]),
			CardFromStr(c[27:29]),
		}

		if Winner(A, B) == 1 {
			wins++
		}
	}

	fmt.Printf("Player 1 won %v times\n", wins)
}
