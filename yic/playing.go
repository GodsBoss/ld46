package yic

import (
	"math"

	"github.com/GodsBoss/ld46/pkg/engine"
)

const playingStateID = "playing"

type playing struct {
	levels *levels

	headAnimation float64
}

func (p *playing) Init() {
}

func (p *playing) Tick(ms int) *engine.Transition {
	p.headAnimation = math.Mod(p.headAnimation+float64(ms)/1000.0, 1.0)
	return nil
}

func (p *playing) Objects() map[string][]engine.Object {
	lvl := p.levels.ChosenLevel()
	headX, headY := lvl.realCoordinate(lvl.headX, lvl.headY)
	objects := map[string][]engine.Object{
		"background": []engine.Object{
			engine.Object{
				Key: "bg_playing",
				X:   0,
				Y:   0,
			},
		},
		"fields": []engine.Object{},
		"entities": []engine.Object{
			engine.Object{
				Key:       "head_toddler",
				X:         headX,
				Y:         headY,
				Animation: p.headAnimation,
			},
		},
	}

	for row := range lvl.fields {
		for col := range lvl.fields[row] {
			field := lvl.fields[row][col]
			key := fieldTypeSpriteKeyMapping[field.typ]
			if key == "" {
				continue
			}
			rx, ry := lvl.realCoordinate(col, row)
			objects["fields"] = append(
				objects["fields"],
				engine.Object{
					Key: key,
					X:   rx,
					Y:   ry,
				},
			)
		}
	}
	return objects
}

var fieldTypeSpriteKeyMapping = map[int]string{
	fieldWay:       "field_way",
	fieldObstacle:  "field_obstacle",
	fieldBuildSpot: "field_buildspot",
}

var fieldSize = vector2D{
	X: 18,
	Y: 18,
}
