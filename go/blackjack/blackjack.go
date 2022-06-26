package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	const (
		stand = "S"
		hit   = "H"
		split = "P"
		win   = "W"
	)

	card1Val := ParseCard(card1)
	card2Val := ParseCard(card2)
	dealerCardVal := ParseCard(dealerCard)
	cardsSum := card1Val + card2Val

	switch {
	case cardsSum == 22:
		return split
	case cardsSum == 21:
		if dealerCardVal < 10 {
			return win
		}
		return stand
	case cardsSum >= 17 && cardsSum <= 20:
		return stand
	case cardsSum >= 12 && cardsSum <= 16:
		if dealerCardVal < 7 {
			return stand
		}
		return hit
	default:
		return hit
	}
}
