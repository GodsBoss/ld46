package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

func NewGame() *engine.Game {
	game := &engine.Game{
		States: map[string]engine.State{
			titleStateID: &title{},
		},
	}
	game.Transition(titleStateID)
	return game
}
