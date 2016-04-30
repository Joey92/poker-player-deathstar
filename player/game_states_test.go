package player

import (
    "testing"
    "github.com/lean-poker/poker-player-go/leanpoker"
)

func TestIsTwoPair(test *testing.T) {
    
    cards := []*leanpoker.Card{
        &leanpoker.Card{Rank: "5", Suit:"Spade"},
        &leanpoker.Card{Rank: "5", Suit:"Heart"},
    }
    
    if !isTwoPair(cards...) {
        
        return
    }
}

func TestIsThreeOfAKind(test *testing.T) {
    
    cards := []*leanpoker.Card{
        &leanpoker.Card{Rank: "5", Suit:"Spade"},
        &leanpoker.Card{Rank: "5", Suit:"Heart"},
        &leanpoker.Card{Rank: "5", Suit:"Heart"},
    }
    
    if !isThreeOfAKind(cards...) {
        
        return
    }
}

func TestIsFourOfAKind(test *testing.T) {
    
    cards := []*leanpoker.Card{
        &leanpoker.Card{Rank: "5", Suit:"Spade"},
        &leanpoker.Card{Rank: "5", Suit:"Heart"},
        &leanpoker.Card{Rank: "5", Suit:"Spade"},
        &leanpoker.Card{Rank: "5", Suit:"Heart"},
    }
    
    if !isFourOfAKind(cards...) {
        
        return
    }
}