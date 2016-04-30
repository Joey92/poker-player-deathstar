package player

import (
	"github.com/lean-poker/poker-player-go/leanpoker"
	"strings"
)

const VERSION = "Default Go folding player"
const MAX_RANK = 30
const JQKA_RANK = 22
const LOW_RANK = 5

func BetRequest(state *leanpoker.Game) int {
	
	p := state.GetPlayer()
	
	if p == nil {
		return 1000
	}
	
	if (state.IsCheckable()) {
		return 0
	}
	
	HoleRank := rankHoleCards(p.HoleCards)
	
	if len(state.CommunityCards) == 0 {
		if (state.CurrentBuyIn > state.Pot + state.CurrentBuyIn * HoleRank || HoleRank <= LOW_RANK) {
			if (HoleRank >= 18) {
				// follow game with good hand
				return state.CurrentBuyIn
			}
			return 0
		}
	} else {
		
		switch true {
			case isFourOfAKind(append(state.CommunityCards, p.HoleCards...)):
			case isThreeOfAKind(append(state.CommunityCards, p.HoleCards...)):
				return p.Stack
			break;
			case isTwoPair(append(state.CommunityCards, p.HoleCards...)):
				if (HoleRank >= JQKA_RANK) {
					// all in
					return p.Stack
				}
			
				return state.CurrentBuyIn + HoleRank
			break;
		}
			
	}
	
	return HoleRank
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

