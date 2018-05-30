package logic

import "testing"

func TestVerifyCards(t *testing.T) {
	cards := []int8{1, 2}
	t.Log(VerifyCards(cards))
	cards2 := []int8{3, 3, 3}
	t.Log(VerifyCards(cards2))
}

func TestVerifyStraight(t *testing.T) {
	card1 := []int8{1, 2, 3, 4}
	card2 := []int8{1, 10, 11, 12, 13}
	card3 := []int8{6, 7, 8, 9, 10, 11}
	t.Log(VerifyStraight(card1), VerifyStraight(card2), VerifyStraight(card3))
}

func TestVerifySuccessivePair(t *testing.T) {
	card1 := []int8{1, 1, 13, 13}
	card2 := []int8{1, 1, 2, 2}
	card3 := []int8{7, 7, 8, 8, 9, 9}
	card4 := []int8{1, 1, 10, 10, 11, 11, 12, 12, 13, 13}
	t.Log(VerifySuccessivePair(card1),VerifySuccessivePair(card2),VerifySuccessivePair(card3),VerifySuccessivePair(card4))
}
