package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const gameOverStateID = "game_over"

type gameOver struct{}

func (g *gameOver) Init() {}

func (g *gameOver) Tick(ms int) *engine.Transition {
	return nil
}

func (g *gameOver) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_game_over",
				X:   0,
				Y:   0,
			},
		},
	}
}
