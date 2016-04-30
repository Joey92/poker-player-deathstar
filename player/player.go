package player

import (
	"github.com/Joey92/poker-player-deathstar/leanpoker"
)

const VERSION = "Default Go folding player"

func BetRequest(state *leanpoker.Game) int {
	
	p := state.GetPlayer()
	
	if p == nil {
		return 1000
	}
	
	var bet int
	
	if (state.IsCheckable(bet)) {
		return bet
	}
	
	if isPair(p.HoleCards) {
		bet += 1000	
	}
	
	return bet
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}

func isPair(cards []*leanpoker.Card) bool  {
	if len(cards) < 2 {
		return false
	}
	
	return cards[0].Rank == cards[1].Rank
}

