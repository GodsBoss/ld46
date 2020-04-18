package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const playingStateID = "playing"

type playing struct {
	levels        *levels
	remainingTime int
}

func (p *playing) Init() {
	p.remainingTime = 2000
}

func (p *playing) Tick(ms int) *engine.Transition {
	p.remainingTime -= ms
	if p.remainingTime <= 0 {
		return engine.NewTransition(titleStateID)
	}
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
