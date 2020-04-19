package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const titleStateID = "title"

type title struct {
	textManager *textManager
}

func (t *title) Init() {
	t.textManager = newTextManager()
	t.textManager.New("level_select_hint", 10, 284).SetContent("Press L for level selection")
	t.textManager.New("title_1", 20, 20).SetScale(2).SetContent("Your Inner Child")
	t.textManager.New("title_2", 40, 40).SetScale(2).SetContent("Keep It Alive!")
	t.textManager.New("prolog", 60, 80).SetContent(prolog)
}

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
		"ui": t.textManager.Objects(),
	}
}

const prolog = `
Do not let adult responsibilites
destroy your inner child! Avoid
thinking too much about your
money, your job, your car, your
house, or your childhood dreams
may leave you forever!
`
