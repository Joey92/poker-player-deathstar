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

func TestRankCard(test *testing.T) {
    cardRanks := map[int]string{
        1: "1", 2: "2", 3: "3",
        14: "K", 15: "A",
    }
    
    for rank, cardRank := range cardRanks {
        value := rankCard(cardRank)
        if value != rank {
            test.Fail()
            return
        }
    }
}

func TestRankHoleCards(test *testing.T)  {
    cards := []*leanpoker.Card{
        &leanpoker.Card{Rank: "1", Suit:"Spade"},
        &leanpoker.Card{Rank: "5", Suit:"Heart"},
    }
    
    if 6 != rankHoleCards(cards) {
        test.Fail()
    }
}