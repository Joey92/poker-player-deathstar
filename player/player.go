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
		bet += state.CurrentBuyIn * 2	
	}
	
	/*if (len(state.CommunityCards) > 0) {
		bet += calcComunityCards(state, bet)
	}*/
	
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

/*func calcComunityCards(g *leanpoker.Game, currentBet) {
	
}*/

