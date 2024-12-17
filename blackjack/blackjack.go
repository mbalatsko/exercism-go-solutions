package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "ten", "jack", "queen", "king":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handPoints := ParseCard(card1) + ParseCard(card2)
	dealerHandPoints := ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case handPoints == 21 && dealerHandPoints < 10:
		return "W"
	case handPoints == 21 && dealerHandPoints >= 10:
		return "S"
	case (handPoints >= 17 && handPoints <= 20) || (handPoints >= 12 && handPoints <= 16 && dealerHandPoints < 7):
		return "S"
	case (handPoints >= 12 && handPoints <= 16 && dealerHandPoints >= 7):
		return "H"
	case handPoints <= 11:
		return "H"
	default:
		return "H"
	}
}
