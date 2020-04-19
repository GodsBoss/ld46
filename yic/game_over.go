package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

const gameOverStateID = "game_over"

type gameOver struct {
	textManager *textManager
}

func (g *gameOver) Init() {
	g.textManager = newTextManager()
	g.textManager.New("id", 30, 20).SetContent(epilog)
	g.textManager.New("press_t_for_title", 10, 284).SetContent("Press 't' to return to title")
}

func (g *gameOver) Tick(ms int) *engine.Transition {
	return nil
}

func (g *gameOver) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type == engine.KeyUp && event.Key == "t" {
		return engine.NewTransition(titleStateID)
	}
	return nil
}

func (g *gameOver) Objects() map[string][]engine.Object {
	return map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_game_over",
				X:   0,
				Y:   0,
			},
		},
		"ui": g.textManager.Objects(),
	}
}

const epilog = `GAME OVER

Congratulations, you are an adult now!

Overwhelmed by responsibilites, dreams and hopes
faded into oblivion. Your inspiration and creativity
drowned by an endless stream of mundane chores.

Devoid of everything you once was, only an empty
shell remained, mindlessly living the same day every
day. A boring existence without ups or downs.

Basically dead inside, you will never experience
the wonders you had in your childhood again, the
fantasy gone forever with no chance of recovery.

Have a good day!
`
