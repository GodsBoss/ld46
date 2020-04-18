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
	return engine.NewTransition(playingStateID)
}
