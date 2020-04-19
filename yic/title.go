package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const titleStateID = "title"

type title struct {
	textManager *textManager

	animation float64
}

func (t *title) Init() {
	t.textManager = newTextManager()
	t.textManager.New("level_select_hint", 10, 284).SetContent("Press L for level selection")
	t.textManager.New("title_1", 20, 20).SetScale(2).SetContent("Your Inner Child")
	t.textManager.New("title_2", 40, 40).SetScale(2).SetContent("Keep It Alive!")
	t.textManager.New("prolog", 60, 80).SetContent(prolog)

	t.animation = 0.0
}

func (t *title) Tick(ms int) *engine.Transition {
	t.animation += float64(ms) / 1000.0
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
	objects := map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_title",
				X:   0,
				Y:   0,
			},
		},
		"ui": t.textManager.Objects(),
		"entities": []engine.Object{
			engine.Object{
				Key:       responsibilityType1,
				X:         50,
				Y:         160,
				Animation: t.animation,
			},
			engine.Object{
				Key:       responsibilityType2,
				X:         100,
				Y:         160,
				Animation: t.animation,
			},
			engine.Object{
				Key:       responsibilityType3,
				X:         150,
				Y:         160,
				Animation: t.animation,
			},
			engine.Object{
				Key:       responsibilityType4,
				X:         200,
				Y:         160,
				Animation: t.animation,
			},
			engine.Object{
				Key:       "title_no",
				X:         60,
				Y:         170,
				Z:         1,
				Animation: t.animation,
			},
			engine.Object{
				Key:       "title_no",
				X:         110,
				Y:         170,
				Z:         1,
				Animation: t.animation,
			},
			engine.Object{
				Key:       "title_no",
				X:         160,
				Y:         170,
				Z:         1,
				Animation: t.animation,
			},
			engine.Object{
				Key:       "title_no",
				X:         210,
				Y:         170,
				Z:         1,
				Animation: t.animation,
			},
			engine.Object{
				Key:       "head_toddler",
				X:         300,
				Y:         142,
				Z:         1,
				Animation: t.animation,
			},
			engine.Object{
				Key:       "title_yes",
				X:         320,
				Y:         170,
				Z:         1,
				Animation: t.animation,
			},
		},
	}
	return objects
}

const prolog = `
Do not let adult responsibilites
destroy your inner child! Avoid
thinking too much about your
money, your job, your car, your
house, or your childhood dreams
may leave you forever!
`
