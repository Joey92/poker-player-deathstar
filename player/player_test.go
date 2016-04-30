package player

import (
    "testing"
    "github.com/Joey92/poker-player-deathstar/leanpoker"
)

func TestShouldBePair(test *testing.T)  {
    cards := []*leanpoker.Card{
        &leanpoker.Card{Rank: "1", Suit:"Spade"},
        &leanpoker.Card{Rank: "1", Suit:"Heart"},
    }
    
    if !isPair(cards) {
        test.Fail()
        return
    }
}

func TestShouldNotBePair(test *testing.T)  {
    cards := []*leanpoker.Card{
        &leanpoker.Card{Rank: "1", Suit:"Spade"},
        &leanpoker.Card{Rank: "5", Suit:"Heart"},
    }
    
    if isPair(cards) {
        test.Fail()
        return
    }
}