package yic

import (
	"math"

	"github.com/GodsBoss/ld46/pkg/engine"

	"math/rand"
)

const gameOverStateID = "game_over"

type gameOver struct {
	textManager *textManager

	adultHeadAnimation float64
	rotation           float64
	types              []string
}

func (g *gameOver) Init() {
	g.textManager = newTextManager()
	g.textManager.New("id", 30, 20).SetContent(epilog)
	g.textManager.New("press_t_for_title", 10, 284).SetContent("Press 't' to return to title")

	allTypes := []string{
		responsibilityType1,
		responsibilityType2,
		responsibilityType3,
		responsibilityType4,
	}
	g.types = make([]string, gameOverItems)
	for i := 0; i < gameOverItems; i++ {
		g.types[i] = allTypes[rand.Intn(len(allTypes))]
	}
}

func (g *gameOver) Tick(ms int) *engine.Transition {
	g.adultHeadAnimation += float64(ms) / 1000.0
	g.rotation += gameOverItemsRotationSpeed * float64(ms) / 1000.0
	return nil
}

func (g *gameOver) HandleKeyEvent(event engine.KeyEvent) *engine.Transition {
	if event.Type == engine.KeyUp && event.Key == "t" {
		return engine.NewTransition(titleStateID)
	}
	return nil
}

func (g *gameOver) Objects() map[string][]engine.Object {
	objects := map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_game_over",
				X:   0,
				Y:   0,
			},
		},
		"entities": []engine.Object{
			engine.Object{
				Key:       "head_adult",
				X:         adultHeadX,
				Y:         adultHeadY,
				Animation: g.adultHeadAnimation,
			},
		},
		"ui": g.textManager.Objects(),
	}
	for i := range g.types {
		angle := 2.0*math.Pi*float64(i)/float64(gameOverItems) + g.rotation
		objects["entities"] = append(
			objects["entities"],
			engine.Object{
				Key:       g.types[i],
				X:         adultHeadX + 8 + int(gameOverItemsDistance*math.Sin(angle)),
				Y:         adultHeadY + 8 + int(gameOverItemsDistance*math.Cos(angle)),
				Animation: g.adultHeadAnimation,
			},
		)
	}
	return objects
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

const (
	adultHeadX = 182
	adultHeadY = 180
)

const (
	gameOverItems              = 10
	gameOverItemsDistance      = 50.0
	gameOverItemsRotationSpeed = 0.25
)
