package basic

//黑红梅方
const (
	CARD_KING        = "card_king"
	CARD_DEPUTY_KING = "card_deputy_king"
)
const (
	CARD_SPADE_ONE   = "card_spade_one"
	CARD_SPADE_TWO   = "card_spade_two"
	CARD_SPADE_THREE = "card_spade_three"
	CARD_SPADE_FOUR  = "card_spade_four"
	CARD_SPADE_FIVE  = "card_spade_five"
	CARD_SPADE_SIX   = "card_spade_six"
	CARD_SPADE_SEVEN = "card_spade_seven"
	CARD_SPADE_EIGHT = "card_spade_eight"
	CARD_SPADE_NINE  = "card_spade_nine"
	CARD_SPADE_TEN   = "card_spade_ten"
	CARD_SPADE_J     = "card_spade_j"
	CARD_SPADE_Q     = "card_spade_q"
	CARD_SPADE_K     = "card_spade_k"
)
const (
	CARD_HEART_ONE   = "card_heart_one"
	CARD_HEART_TWO   = "card_heart_two"
	CARD_HEART_THREE = "card_heart_three"
	CARD_HEART_FOUR  = "card_heart_four"
	CARD_HEART_FIVE  = "card_heart_five"
	CARD_HEART_SIX   = "card_heart_six"
	CARD_HEART_SEVEN = "card_heart_seven"
	CARD_HEART_EIGHT = "card_heart_eight"
	CARD_HEART_NINE  = "card_heart_nine"
	CARD_HEART_TEN   = "card_heart_ten"
	CARD_HEART_J     = "card_heart_j"
	CARD_HEART_Q     = "card_heart_q"
	CARD_HEART_K     = "card_heart_k"
)

const (
	CARD_CLUB_ONE   = "card_club_one"
	CARD_CLUB_TWO   = "card_club_two"
	CARD_CLUB_THREE = "card_club_three"
	CARD_CLUB_FOUR  = "card_club_four"
	CARD_CLUB_FIVE  = "card_club_five"
	CARD_CLUB_SIX   = "card_club_six"
	CARD_CLUB_SEVEN = "card_club_seven"
	CARD_CLUB_EIGHT = "card_club_eight"
	CARD_CLUB_NINE  = "card_club_nine"
	CARD_CLUB_TEN   = "card_club_ten"
	CARD_CLUB_J     = "card_club_j"
	CARD_CLUB_Q     = "card_club_q"
	CARD_CLUB_K     = "card_club_k"
)

const (
	CARD_DIAMOND_ONE   = "card_diamond_one"
	CARD_DIAMOND_TWO   = "card_diamond_two"
	CARD_DIAMOND_THREE = "card_diamond_three"
	CARD_DIAMOND_FOUR  = "card_diamond_four"
	CARD_DIAMOND_FIVE  = "card_diamond_five"
	CARD_DIAMOND_SIX   = "card_diamond_six"
	CARD_DIAMOND_SEVEN = "card_diamond_seven"
	CARD_DIAMOND_EIGHT = "card_diamond_eight"
	CARD_DIAMOND_NINE  = "card_diamond_nine"
	CARD_DIAMOND_TEN   = "card_diamond_ten"
	CARD_DIAMOND_J     = "card_diamond_j"
	CARD_DIAMOND_Q     = "card_diamond_q"
	CARD_DIAMOND_K     = "card_diamond_k"
)

var AllCard = []string{CARD_KING, CARD_DEPUTY_KING,
	CARD_SPADE_ONE, CARD_SPADE_TWO, CARD_SPADE_THREE, CARD_SPADE_FOUR,
	CARD_SPADE_FIVE, CARD_SPADE_SIX, CARD_SPADE_SEVEN, CARD_SPADE_EIGHT,
	CARD_SPADE_NINE, CARD_SPADE_TEN, CARD_SPADE_J, CARD_SPADE_Q,
	CARD_SPADE_K,
	CARD_HEART_ONE, CARD_HEART_TWO, CARD_HEART_THREE, CARD_HEART_FOUR,
	CARD_HEART_FIVE, CARD_HEART_SIX, CARD_HEART_SEVEN, CARD_HEART_EIGHT,
	CARD_HEART_NINE, CARD_HEART_TEN, CARD_HEART_J, CARD_HEART_Q,
	CARD_HEART_K,
	CARD_CLUB_ONE, CARD_CLUB_TWO, CARD_CLUB_THREE, CARD_CLUB_FOUR,
	CARD_CLUB_FIVE, CARD_CLUB_SIX, CARD_CLUB_SEVEN, CARD_CLUB_EIGHT,
	CARD_CLUB_NINE, CARD_CLUB_TEN, CARD_CLUB_J, CARD_CLUB_Q,
	CARD_CLUB_K,
	CARD_DIAMOND_ONE, CARD_DIAMOND_TWO, CARD_DIAMOND_THREE, CARD_DIAMOND_FOUR,
	CARD_DIAMOND_FIVE, CARD_DIAMOND_SIX, CARD_DIAMOND_SEVEN, CARD_DIAMOND_EIGHT,
	CARD_DIAMOND_NINE, CARD_DIAMOND_TEN, CARD_DIAMOND_J, CARD_DIAMOND_Q,
	CARD_DIAMOND_K,}
var AllCardMap map[string]int8

func init() {
	AllCardMap = make(map[string]int8)
	AllCardMap[CARD_KING] = 15
	AllCardMap[CARD_DEPUTY_KING] = 14
	noKingCard := AllCard[2:]
	for i := range noKingCard {
		value := i + 1
		if value%13 == 0 {
			AllCardMap[noKingCard[i]] = 13
		} else {
			AllCardMap[noKingCard[i]] = int8(value) % 13
		}
	}
}
