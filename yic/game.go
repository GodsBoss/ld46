package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

func NewGame() *engine.Game {
	lvls := createLevels()
	game := &engine.Game{
		States: map[string]engine.State{
			titleStateID: &title{},
			playingStateID: &playing{
				levels: lvls,
			},
			levelSelectStateID: &levelSelect{
				levels: lvls,
			},
			hiscoreStateID:  &hiscore{},
			gameOverStateID: &gameOver{},
		},
	}
	game.Transition(titleStateID)
	return game
}
