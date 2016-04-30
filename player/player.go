package player

import (
	"github.com/lean-poker/poker-player-go/leanpoker"
	"strings"
)

const VERSION = "Default Go folding player"
const MAX_RANK = 30

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
		bet += state.CurrentBuyIn * rankHoleCards(p.HoleCards)
	}
	
	if (state.CurrentBuyIn > state.CurrentBuyIn * rankHoleCards(p.HoleCards)) {
		return 0
	}
	
	if (rankHoleCards(p.HoleCards) == MAX_RANK) {
		return 1000
	}
	
	bet += rankHoleCards(p.HoleCards) * 2
	
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

func rankHoleCards(cards []*leanpoker.Card) int {
	if len(cards) < 2 {
		return 0
	}
	
	card1 := rankCard(cards[0].Rank)
	card2 := rankCard(cards[1].Rank)
	
	return card1 + card2
}

func rankCard(rank string) int {
	ranks := "12345678910JQKA"
	return strings.Index(ranks, rank) + 1
}

/*func calcComunityCards(g *leanpoker.Game, currentBet) {
	
}*/

