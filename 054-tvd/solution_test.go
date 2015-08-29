package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestString(t *testing.T) {
	assert.Equal(t, "TS", Card{10, 'S'}.String(), "")
}

func TestCardFromStr(t *testing.T) {
	assert.Panics(t, func() { CardFromStr("5") }, "invalid length")
	assert.Panics(t, func() { CardFromStr("1D") }, "invalid value")
	assert.Panics(t, func() { CardFromStr("3A") }, "invalid suit")
	assert.Equal(t, Card{2, 'S'}, CardFromStr("2S"), "")
	assert.Equal(t, Card{2, 'C'}, CardFromStr("2C"), "")
	assert.Equal(t, Card{2, 'D'}, CardFromStr("2D"), "")
	assert.Equal(t, Card{10, 'H'}, CardFromStr("TH"), "")
}

func TestValueInHand(t *testing.T) {
	hand := []Card{
		Card{3, 'H'},
		Card{5, 'S'},
	}
	assert.True(t, ValueInHand(3, hand), "")
	assert.False(t, ValueInHand(4, hand), "")
	assert.True(t, ValueInHand(5, hand), "")
}

func TestCardInHand(t *testing.T) {
	hand := []Card{
		Card{3, 'H'},
		Card{5, 'S'},
	}
	assert.True(t, CardInHand(Card{3, 'H'}, hand), "")
	assert.False(t, CardInHand(Card{3, 'S'}, hand), "")
	assert.True(t, CardInHand(Card{5, 'S'}, hand), "")
}

func TestValuesInHand(t *testing.T) {
	hand := []Card{
		Card{3, 'H'},
		Card{5, 'S'},
		Card{7, 'S'},
	}
	assert.True(t, ValuesInHand([]int{3, 5, 7}, hand), "")
	assert.False(t, ValuesInHand([]int{3, 4, 5}, hand), "")
	assert.True(t, ValuesInHand([]int{7}, hand), "")
}

func TestCardsInHand(t *testing.T) {
	hand := []Card{
		Card{3, 'H'},
		Card{5, 'S'},
		Card{7, 'S'},
	}
	assert.True(t, CardsInHand([]Card{Card{3, 'H'}}, hand), "")
	assert.False(t, CardsInHand([]Card{Card{3, 'S'}, Card{3, 'H'}}, hand), "")
	assert.True(t, CardsInHand([]Card{Card{5, 'S'}, Card{7, 'S'}}, hand), "")
}

func TestSameSuit(t *testing.T) {
	assert.True(t, SameSuit([]Card{Card{5, 'H'}, Card{4, 'H'}}), "")
	assert.False(t, SameSuit([]Card{Card{5, 'S'}, Card{4, 'H'}}), "")
}

func TestStraightFlush(t *testing.T) {
	straightflush := []Card{
		Card{2, 'S'},
		Card{3, 'S'},
		Card{4, 'S'},
		Card{5, 'S'},
		Card{6, 'S'},
	}
	notstraightflush := []Card{
		Card{2, 'H'},
		Card{7, 'H'},
		Card{3, 'H'},
		Card{5, 'H'},
		Card{6, 'H'},
	}
	assert.Equal(t, 6, StraightFlush(straightflush), "")
	assert.Equal(t, 0, StraightFlush(notstraightflush), "")
}
