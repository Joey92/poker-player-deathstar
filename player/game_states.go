package player

import "github.com/lean-poker/poker-player-go/leanpoker"

func isTwoPair(cards ...*leanpoker.Card) bool {
    ranks := cardRankCountInStack(cards)
    
    for _, amount := range ranks {
        if amount == 2 {
            return true
        } 
    }
    
    return false
}

func isFourOfAKind(cards ...*leanpoker.Card) bool {
    ranks := cardRankCountInStack(cards)
    
    for _, amount := range ranks {
        if amount == 4 {
            return true
        } 
    }
    
    return false
}

func isThreeOfAKind(cards ...*leanpoker.Card) bool {
    
    ranks := cardRankCountInStack(cards)
    
    for _, amount := range ranks {
        if amount == 3 {
            return true
        } 
    }
    
    return false
}

func isFullHouse(cards ...*leanpoker.Card) bool {
    return isTwoPair(cards...) && isThreeOfAKind(cards...)
}

func cardRankCountInStack(stack []*leanpoker.Card) map[string]int {
	ranks := map[string]int{}
	for _, sCard := range stack {
        _, exists := ranks[sCard.Rank]
        
        if !exists {
            ranks[sCard.Rank] = 0
        }
        ranks[sCard.Rank]++
	}
    
    return ranks
}