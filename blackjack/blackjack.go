package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value = 0
    switch  {
        case card == "ace":
            value = 11
        case card == "ten" || card == "jack" || card == "queen" || card == "king":
    		value = 10
        case card == "nine":
            value = 9
        case card == "eight":
            value = 8
        case card == "seven":
            value = 7
        case card == "six":
            value = 6
        case card == "five":
            value = 5
        case card == "four":
            value = 4
        case card == "three":
            value = 3
        case card == "two":
            value = 2
        default:
            value = 0
        } 
    return value
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return (ParseCard(card1) + ParseCard(card2)) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
        if dealerScore < 10 {
            return "W"
        } else {
            return "S"
        }
    } else {
        return "P"
    }
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
    if handScore < 11 {
        return "H"
    } else if handScore < 17 {
        if dealerScore < 7 {
            return "S"
        } else {
            return "H"
        } 
    }
    return "S"
}
