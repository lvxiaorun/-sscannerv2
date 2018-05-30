package logic

import "github.com/lvxiaorun/andlords/tool"

//this func will get a args of number 1-15 and it have sorted small to big
func VerifyCards(cards []int8) bool {
	length := len(cards)
	switch length {
	case 1:
		return true
	case 2:
		//两者相等或者是大小王
		if (cards[0] == 14 && cards[1] == 15) || cards[0] == cards[1] {
			return true
		}
	case 3:
		//只能是三个相等,因为已经排序所以可以直接比较两端
		if cards[0] == cards[2] {
			return true
		}
	case 4:
		//可能的情况
		//1.炸弹 2.三带一 3.顺子 4.连对
		if cards[0] == cards[3] {
			return true //炸弹
		}
	}
	return false
}

//校验是不是三带一 已经排序
func VerifyAAAB(cards []int8) bool {
	//长度在之前已经校验过
	if cards[0] == cards[2] || cards[1] == cards[3] {
		return true
	}
	return false
}

//校验是不是顺子
func VerifyStraight(cards []int8) bool {
	mcards := tool.Int8SliceToMap(cards)
	if mcards[2] > 0 || mcards[14] > 0 || mcards[15] > 0 {
		return false
	}
	var cardsVerify []int8
	if cards[0] == 1 {
		cardsVerify = append(cards[1:], 14)
	} else {
		cardsVerify = cards
	}
	length := len(cardsVerify)
	if cardsVerify[length-1]-cardsVerify[0] != int8(length)-1 {
		return false
	}
	return true
}

func VerifySuccessivePair(cards []int8) bool {
	mcards := tool.Int8SliceToMap(cards)
	for key := range mcards {
		if key == 2 {
			return false
		}
		if mcards[key] != 2 {
			return false
		}
		var cardsVerify []int8
		if cards[0] == 1 {
			cardsVerify = append(cards[2:], 14, 14)
		} else {
			cardsVerify = cards
		}
		length := len(cardsVerify)
		if cardsVerify[length-1]-cardsVerify[0] != int8(length)/2-1 {
			return false
		}
	}
	return true
}
