package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const titleStateID = "title"

type title struct{}

func (t *title) Init() {}

func (t *title) Tick(ms int) *engine.Transition {
	return nil
}

func (t *title) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type != engine.KeyUp {
		return nil
	}
	switch event.Key {
	case "f":
		return engine.NewTransition(hiscoreStateID)
	case "l":
		return engine.NewTransition(levelSelectStateID)
	}
	return nil
}

func (t *title) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_title",
				X:   0,
				Y:   0,
			},
		},
	}
}
