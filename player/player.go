package player

import "github.com/Joey92/poker-player-deathstar/leanpoker"

const VERSION = "Default Go folding player"

func BetRequest(state *leanpoker.Game) int {
	var bet int
	
	if (state.IsCheckable(bet)) {
		return 0;
	}
	
	return bet
}

func Showdown(state *leanpoker.Game) {

}

func Version() string {
	return VERSION
}
