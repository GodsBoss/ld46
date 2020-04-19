package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const hiscoreStateID = "hiscore"

type hiscore struct {
	textManager *textManager
}

func (h *hiscore) Init() {
	h.textManager = newTextManager()
	h.textManager.New("press_t_for_title", 10, 284).SetContent("Press 't' to return to title")
}

func (h *hiscore) Tick(ms int) *engine.Transition {
	return nil
}

func (h *hiscore) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type != engine.KeyUp {
		return nil
	}
	if event.Key == "t" {
		return engine.NewTransition(titleStateID)
	}
	return nil
}

func (h *hiscore) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_hiscore",
				X:   0,
				Y:   0,
			},
		},
		"ui": h.textManager.Objects(),
	}
}
