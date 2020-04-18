package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const levelSelectStateID = "level_select"

type levelSelect struct{}

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
	return nil
}

func (ls *levelSelect) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{}
}
