package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const playingStateID = "playing"

type playing struct {
	levels *levels
}

func (p *playing) Init() {
}

func (p *playing) Tick(ms int) *engine.Transition {
	return nil
}

func (p *playing) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_playing",
				X:   0,
				Y:   0,
			},
		},
	}
}
