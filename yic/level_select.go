package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const levelSelectStateID = "level_select"

type levelSelect struct {
	levels *levels
}

func (ls *levelSelect) Init() {}

func (ls *levelSelect) Tick(ms int) *engine.Transition {
	return nil
}

func (ls *levelSelect) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type != engine.KeyUp {
		return nil
	}
	switch event.Key {
	case "b":
		return engine.NewTransition(titleStateID)
	}
	if _, ok := ls.levels.byKey[event.Key]; ok {
		ls.levels.chosen = event.Key
		return engine.NewTransition(playingStateID)
	}
	return nil
}

func (ls *levelSelect) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_level_select",
				X:   0,
				Y:   0,
			},
		},
	}
}
