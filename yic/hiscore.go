package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const hiscoreStateID = "hiscore"

type hiscore struct{}

func (h *hiscore) Init() {}

func (h *hiscore) Tick(ms int) *engine.Transition {
	return nil
}

func (h *hiscore) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type != engine.KeyUp {
		return nil
	}
	return engine.NewTransition(titleStateID)
}

func (h *hiscore) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{}
}
